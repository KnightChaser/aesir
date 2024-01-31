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
async function collectionDocumentQtyGetData() {
    try {
        const data = await fetchCollectionDocumentQty(currentCollection);
        const numberOfDocumentsElement = document.getElementById("number-of-document-number");
        numberOfDocumentsElement.textContent = data;
    } catch (error) {
        // Handle errors if needed
        console.error('Error:', error);
    }
}

collectionDocumentQtyGetData();

// Get the number of documents per event type and draw the chart with chart.js
async function eventDocumentQtyGetData() {
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
                            display: false
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



eventDocumentQtyGetData();