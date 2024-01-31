currentURL = window.location.href;
currentBaseURL = currentURL.split('/').slice(0, 3).toString().replace(',,', '//');
currentCollection = currentURL.split('/').slice(4);

// General conditional search API
async function fetchEventData(collectionInUse, searchRequest, searchCondition) {
    // Replace 'http://localhost:8080/api' with your actual API endpoint
    const apiUrl = `${currentBaseURL}/api/search/${encodeURI(collectionInUse)}/${encodeURI(searchRequest)}/${encodeURI(searchCondition)}`;

    return fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Network response was not ok: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            return data["result"];
        })
        .catch(error => {
            // Handle errors here
            console.error('Error:', error);
            throw error;
        });
}

// Document counting in the current collection API
async function fetchCollectionDocumentQty(collectionInUse) {
    const apiUrl = `${currentBaseURL}/api/search/${encodeURI(collectionInUse)}/documentCount`;

    return fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Network response was not ok: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            return data["documentQty"];
        })
        .catch(error => {
            console.error('Error:', error);
            throw error;
        });
}

// Get the number of documents in the current collection
// Using animation to show the number of documents (from 0 to the actual number, for pretty display)
async function collectionDocumentQtyGetData() {
    try {
        const data = await fetchCollectionDocumentQty(currentCollection);
        const numberOfDocumentsElement = document.getElementById("number-of-document-number");
        const startValue = 0;
        const endValue = data;
        const duration = 1000; // Animation duration in milliseconds
        const frameRate = 30; // Number of frames per second

        const increment = (endValue - startValue) / (duration / 1000 * frameRate);
        let currentValue = startValue;

        const counterInterval = setInterval(() => {
            currentValue += increment;
            numberOfDocumentsElement.textContent = Math.round(currentValue).toLocaleString();

            if (currentValue >= endValue) {
                numberOfDocumentsElement.textContent = endValue.toLocaleString();
                clearInterval(counterInterval);
            }
        }, 1000 / frameRate);
    } catch (error) {
        // Handle errors if needed
        console.error('Error:', error);
    }
}

collectionDocumentQtyGetData();

// Get basic information of the current EVTX file
async function collectionDocumentOverallMetadata() {
    try {
        // Search condition for the API (using MongoDB aggregation); start with the first document
        const searchConditionStart = [
            { $sort: { "event.system.timecreated.systemtime": 1 } },
            { $limit: 1 }
        ];

        const dataStart = await fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionStart));
        const timestampStartElementString = document.getElementById("event-starting-timestamp");
        const dataJSONStart = JSON.parse(dataStart);
        // Convert the timestamp to ISO format
        const startingTimestamp = new Date(dataJSONStart[0]["event"]["system"]["timecreated"]["systemtime"]);
        const formattedStartingTimestamp = startingTimestamp.toISOString().replace(/T/, ' ').replace(/\..+/, '');
        timestampStartElementString.textContent = formattedStartingTimestamp;

        // Search condition for the API (using MongoDB aggregation); end with the last document
        const searchConditionEnd = [
            { $sort: { "event.system.timecreated.systemtime": -1 } },
            { $limit: 1 }
        ];

        const dataEnd = await fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionEnd));
        const timestampEndElementString = document.getElementById("event-ending-timestamp");
        const dataJSONEnd = JSON.parse(dataEnd);
        // Convert the timestamp to ISO format
        const endingTimestamp = new Date(dataJSONEnd[0]["event"]["system"]["timecreated"]["systemtime"]);
        const formattedEndingTimestamp = endingTimestamp.toISOString().replace(/T/, ' ').replace(/\..+/, '');
        timestampEndElementString.textContent = formattedEndingTimestamp;

        // Search condition for the API (using MongoDB aggregation); get the computer name of the current EVTX file
        const searchConditionComputerName = [
            {
                $group: {
                    _id: "$event.system.computer",
                    uniqueComputers: { $addToSet: "$event.system.computer" }
                }
            },
            {
                $project: {
                    _id: 0,
                    uniqueComputers: 1
                }
            }
        ];

        const dataComputerName = JSON.parse(await fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionComputerName)));
        let dataComputerNameString = "";
        dataComputerName.forEach(row => {
            dataComputerNameString += row["uniqueComputers"].toString();
            dataComputerNameString += "\n";
        });
        document.getElementById("event-captured-computer-name").textContent = dataComputerNameString;

        // Search condition for the API (using MongoDB aggregation); get channel names of the current EVTX file
        const searchConditionChannelName = [
            {
                $group: {
                    _id: "$event.system.channel",
                    uniqueChannels: { $addToSet: "$event.system.channel" }
                }
            },
            {
                $project: {
                    _id: 0,
                    uniqueChannels: 1
                }
            }
        ];

        const dataChannelName = JSON.parse(await fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionChannelName)));
        document.getElementById("event-channel-name").textContent = dataChannelName[0]["uniqueChannels"].toString();

        // Search condition for the API (using MongoDB aggregation); get the information of provider
        const searchConditionProvider = [
            {
                $group: {
                    _id: "$event.system.provider",
                    uniqueProviders: { $addToSet: "$event.system.provider" }
                }
            },
            {
                $project: {
                    _id: 0,
                    uniqueProviders: 1
                }
            }
        ];

        const dataProvider = JSON.parse(await fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionProvider)));
        let dataProviderString = "";
        dataProviderString += dataProvider[0]["uniqueProviders"][0]["name"] + '<br>';
        dataProviderString += " (GUID: <code>" + dataProvider[0]["uniqueProviders"][0]["guid"] + "</code>)";
        document.getElementById("event-tool-provider").innerHTML = dataProviderString;

    } catch (error) {
        // Handle errors if needed
        console.error('Error:', error);
    }
}

collectionDocumentOverallMetadata();

// Get the number of documents per event type and draw the chart with chart.js
async function eventDocumentQtyDrawChart() {
    try {
        // Search condition for the API (using MongoDB aggregation)
        const searchCondition = [
            { "$group": { "_id": "$event.eventdata.EventName", "count": { "$sum": 1 } } },
            { "$sort": { "count": -1 } }
        ];

        const data = await fetchEventData(currentCollection, "aggregate", JSON.stringify(searchCondition));
        const eventDocumentQtyElement = document.getElementById("event-id-distribution-chart");

        // Check if data is not null or undefined
        if (data) {
            const dataJson = JSON.parse(data);

            // Extract _id and count arrays
            const ids = dataJson.map(item => item._id);
            const counts = dataJson.map(item => item.count);

            // Draw the chart using Chart.js

            const ctx = eventDocumentQtyElement.getContext('2d');
            new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: ids,
                    datasets: [{
                        label: 'Event ID Distribution',
                        data: counts,
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        borderColor: 'rgba(75, 192, 192, 1)',
                        borderWidth: 1,
                        barThickness: 45,
                        label: '# of events'
                    }]
                },
                options: {
                    scales: {
                        y: {
                            beginAtZero: true,
                            stepSize: 1,
                        }
                    },
                    plugins: {
                        legend: {
                            display: false,
                            labels: {
                                font: {
                                    family: 'IBMPlexSans',
                                }
                            }
                        }
                    },
                    elements: {
                        bar: {
                            borderWidth: 1,
                            borderColor: 'rgba(0, 0, 0, 1)'
                        }
                    }
                }
            });
        }
    } catch (error) {
        // Handle errors if needed
        console.error('Error while drawing chart for eventDocuemtnQtyGetData():', error);
    }
}

eventDocumentQtyDrawChart();

// Get source image distribution per event ID and draw the chart with chart.js
async function imageDistributionDrawChart(searchConditionFieldName, tableElementId) {
    try {
        const searchCondition = [
            {
                $match: {
                    [searchConditionFieldName]: { $ne: null }
                }
            },
            {
                $group: {
                    _id: "$" + [searchConditionFieldName],
                    count: { $sum: 1 }
                }
            },
            {
                $sort: { count: -1 }
            },
            {
                $project: {
                    _id: 0,
                    sourceImage: "$_id",
                    count: 1
                }
            },
        ];

        const data = await fetchEventData(currentCollection, "aggregate", JSON.stringify(searchCondition));
        const dataJSON = JSON.parse(data);

        // Initialize bootstrap tooltip
        const tooltipTriggerList = document.querySelectorAll('[data-bs-toggle="tooltip"]')
const tooltipList = [...tooltipTriggerList].map(tooltipTriggerEl => new bootstrap.Tooltip(tooltipTriggerEl))

        // Create bootstrap table based on the information(labels and counts)
        const target = document.getElementById(tableElementId);
        const table = document.createElement('table');
        table.classList.add('table', 'table-hover'); // Add Bootstrap table classes

        // Create table header
        const thead = document.createElement('thead');
        const headerRow = document.createElement('tr');
        headerRow.classList.add('table-info');
        const headerCell1 = document.createElement('th');
        headerCell1.textContent = 'Executable Name';
        const headerCell2 = document.createElement('th');
        headerCell2.textContent = 'Count';
        headerRow.appendChild(headerCell1);
        headerRow.appendChild(headerCell2);
        thead.appendChild(headerRow);
        table.appendChild(thead);

        // Create table body
        const tbody = document.createElement('tbody');

        dataJSON.forEach(row => {
            const sourceImageFullPath = row["sourceImage"];
            const sourceImageExecutableName = sourceImageFullPath.split('\\').pop();
            const count = row["count"];

            const dataRow = document.createElement('tr');
            const dataCell1 = document.createElement('td');
            const tooltip = document.createElement('span');
            tooltip.setAttribute('data-bs-toggle', 'tooltip');
            tooltip.setAttribute('data-bs-placement', 'top');
            tooltip.setAttribute('title', sourceImageFullPath);
            tooltip.innerHTML = "<code>" + sourceImageExecutableName + "</code>";
            dataCell1.appendChild(tooltip);

            const dataCell2 = document.createElement('td');
            dataCell2.textContent = count.toLocaleString('en-US', { style: 'decimal' });

            dataCell2.style.textAlign = 'right'; // Right align the cell

            dataRow.appendChild(dataCell1);
            dataRow.appendChild(dataCell2);

            tbody.appendChild(dataRow);
        });

        table.appendChild(tbody);

        // Append the table to the target element
        target.appendChild(table);



    } catch (error) {
        // Handle errors if needed
        console.error('Error:', error);
    }
}

imageDistributionDrawChart("event.eventdata.SourceImage", "source-image-distribution-table");
imageDistributionDrawChart("event.eventdata.TargetImage", "target-image-distribution-table");