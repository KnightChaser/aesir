function showAdditionalInfoAlert(htmlContent) {
    htmlContent = decodeURI(htmlContent);
    Swal.fire({
        title: 'Additional Information',
        icon: 'info',
        html: htmlContent,
        showCloseButton: true,
        focusConfirm: false,
        width: 1200,
    });
};

// If button (id = "search-form-submit-button") is clicked, print the selected values of the selectize.js dropdowns
// Also, get the values of date range picker
$("#search-form-submit-button").click(function () {
    let sourceImageSelected = $("#source-image-selectize-tab").val();
    let targetImageSelected = $("#target-image-selectize-tab").val();
    let ruleNameSelected = $("#rule-name-selectize-tab").val();
    let dateRangeStart = $("#starting-datetime").val();
    let dateRangeEnd = $("#ending-datetime").val();

    let searchCondition = {
        $and: []
    };

    if (sourceImageSelected.length > 0)
        searchCondition.$and.push({ "event.eventdata.SourceImage": { $in: sourceImageSelected } });

    if (targetImageSelected.length > 0)
        searchCondition.$and.push({ "event.eventdata.TargetImage": { $in: targetImageSelected } });

    if (ruleNameSelected.length > 0)
        searchCondition.$and.push({ "event.eventdata.RuleName": { $in: ruleNameSelected } });

    if (dateRangeStart.length > 0 && dateRangeEnd.length > 0)
        searchCondition.$and.push({ "event.eventdata.TimeGenerated": { $gte: dateRangeStart, $lte: dateRangeEnd } });

    // If no search condition is selected, fire a SweetAlert2 error message
    if (searchCondition.$and.length === 0) {
        Swal.fire({
            icon: 'error',
            title: 'Search Condition Error',
            text: 'Please select at least one search condition!',
        });
        return;
    }

    const sweetAlert2Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        customClass: {
            popup: 'colored-toast',
        },
        showConfirmButton: false,
        timer: 3500,
        timerProgressBar: true,
    });

    sweetAlert2Toast.fire({
        icon: 'info',
        title: 'Searching...',
    });

    fetchEventDataMultipleCondition(currentCollection, JSON.stringify(searchCondition)).then((response) => {

        // Clear the table before appending the search result
        $("#search-result-table-row").empty();

        // Showing the number of results.
        const count = response["count"];
        const formattedCount = count.toLocaleString();
        $("#searched-document-count").text(formattedCount);

        // Order searchResult JSON data by objectId in descending order
        let searchResult = JSON.parse(response["result"]);

        // Showing the number of entire document in the collection automatically
        fetchCollectionDocumentQty(currentCollection)
        .then((entireDataQty) => {
            $("#entire-document-count").text(entireDataQty.toLocaleString());
            
            let documentCountProportion = (count / entireDataQty) * 100;
            $("#searched-document-count-propotion-progress-bar").css("width", documentCountProportion + "%");

            if (documentCountProportion > 0) {
                $("#searched-document-count-percentage").text("(" + documentCountProportion.toFixed(3) + "% of entire document)");
            } else {
                $("#searched-document-count-percentage").text("(0% of entire document)");
            }
            
        })
        .catch((error) => {
            Swal.fire({
                icon: 'error',
                title: 'Error',
                text: 'Error while fetching the number of entire document : ' + error,
            });
            throw new Error('Error while fetching the number of entire document :', error);
        });

        // get element one by one from searchResult object
        // and display it in the table
        searchResult.forEach((item, index) => {
            // overall metadata
            let objectId = item._id || "N/A";
            let eventRecordId = item.eventrecordid || "N/A";

            // system data (overall common)
            let eventId = item.event.system.eventid || "N/A";
            let eventOccurredTime = item.event.system.timecreated.systemtime || "N/A";
            let computerName = item.event.system.computer || "N/A";
            let securityUserId = item.event.system.security.userid || "N/A";
            let provider = {
                name: item.event.system.provider.name || "N/A",
                guid: item.event.system.provider.guid || "N/A",
            }

            // event data (event-specific)
            let eventName = item.event.eventdata.EventName || "N/A";
            let eventCallTrace = item.event.eventdata.CallTrace ? item.event.eventdata.CallTrace.split("|").map(trace => `<code>${trace}</code>`).join("<br>") : "N/A";
            let ruleName = item.event.eventdata.RuleName || "N/A";
            let grantedAccess = item.event.eventdata.GrantedAccess || "N/A";
            let hashes = item.event.eventdata.Hashes || "N/A";
            let processData = {
                "parentCommandLine": item.event.eventdata.ParentCommandLine || "N/A",
                "parentImage": item.event.eventdata.ParentImage || "N/A",
                "parentProcessId": item.event.eventdata.ParentProcessId || "N/A",
                "processGuid": item.event.eventdata.ProcessGuid || "N/A",
                "processId": item.event.eventdata.ProcessId || "N/A",
                "commandLine": item.event.eventdata.CommandLine || "N/A",
                "description": item.event.eventdata.Description || "N/A",
                "logonGuid": item.event.eventdata.LogonGuid || "N/A",
            }
            let sourceImageData = {
                "sourceImage": item.event.eventdata.SourceImage || "N/A",
                "sourceProcessGUID": item.event.eventdata.SourceProcessGUID || "N/A",
                "sourceProcessId": item.event.eventdata.SourceProcessId || "N/A",
                "sourceThreadId": item.event.eventdata.SourceThreadId || "N/A",
                "soruceUser": item.event.eventdata.SourceUser || "N/A",
            }
            let targetImageData = {
                "targetImage": item.event.eventdata.TargetImage || "N/A",
                "targetProcessGUID": item.event.eventdata.TargetProcessGUID || "N/A",
                "targetProcessId": item.event.eventdata.TargetProcessId || "N/A",
                "targetThreadId": item.event.eventdata.TargetThreadId || "N/A",
                "targetUser": item.event.eventdata.TargetUser || "N/A",
            }

            let additionalInformationAlertHTML = `
            <table class="table fixed">
                <col scope="col" width="220px" />
                <col scope="col" width="600px" />
                <thead>
                    <tr>
                        <th scope="col">metadata</th>
                        <th scope="col">value</th>
                    </tr>
                    <tbody>
                        <tr>
                            <th scope="row"><code>objectId</code></th>
                            <td>${objectId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>eventRecordId</code></th>
                            <td>${eventRecordId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>eventId</code></th>
                            <td>${eventId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>eventOccurredTime</code></th>
                            <td>${eventOccurredTime}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>computerName</code></th>
                            <td>${computerName}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>securityUserId</code></th>
                            <td>${securityUserId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>provider.name</code></th>
                            <td>${provider.name}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>provider.guid</code></th>
                            <td>${provider.guid}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>eventName</code></th>
                            <td>${eventName}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>eventCallTrace</code></th>
                            <td>${eventCallTrace}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>ruleName</code></th>
                            <td>${ruleName}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>grantedAccess</code></th>
                            <td>${grantedAccess}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>sourceImageData.sourceImage</code></th>
                            <td>${sourceImageData.sourceImage}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>sourceImageData.sourceProcessGUID</code></th>
                            <td>${sourceImageData.sourceProcessGUID}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>sourceImageData.sourceProcessId</code></th>
                            <td>${sourceImageData.sourceProcessId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>sourceImageData.sourceThreadId</code></th>
                            <td>${sourceImageData.sourceThreadId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>sourceImageData.soruceUser</code></th>
                            <td>${sourceImageData.soruceUser}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>targetImageData.targetImage</code></th>
                            <td>${targetImageData.targetImage}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>targetImageData.targetProcessGUID</code></th>
                            <td>${targetImageData.targetProcessGUID}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>targetImageData.targetProcessId</code></th>
                            <td>${targetImageData.targetProcessId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>targetImageData.targetThreadId</code></th>
                            <td>${targetImageData.targetThreadId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>targetImageData.targetUser</code></th>
                            <td>${targetImageData.targetUser}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.parentCommandLine</code></th>
                            <td><code>${processData.parentCommandLine}</code></td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.parentImage</code></th>
                            <td>${processData.parentImage}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.parentProcessId</code></th>
                            <td>${processData.parentProcessId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.processGuid</code></th>
                            <td>${processData.processGuid}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.processId</code></th>
                            <td>${processData.processId}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.commandLine</code></th>
                            <td><code>${processData.commandLine}</code></td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.description</code></th>
                            <td>${processData.description}</td>
                        </tr>
                        <tr>
                            <th scope="row"><code>processData.logonGuid</code></th>
                            <td>${processData.logonGuid}</td>
                        </tr>
                    </tbody>
                </thead>
            </table>
            `;

            // Append the search result to the table
            // Append HTML object to the <tbody> element whose ID is "search-result-table-row"
            $("#search-result-table-row").append(
                '<tr>' +
                '<th scope="row">' + eventRecordId + '</th>' +
                '<td style="word-break: break-all;">' + ruleName + '</td>' +
                '<td style="word-break: break-all;"><code>' + sourceImageData.sourceImage + '</code></td>' +
                '<td style="word-break: break-all;"><code>' + targetImageData.targetImage + '</code></td>' +
                '<td>' + eventOccurredTime + '</td>' +
                '<td>' +
                "<button type='button' class='btn btn-primary' style='margin:auto; display:block;' onclick='showAdditionalInfoAlert(\`" + encodeURI(additionalInformationAlertHTML) + "\`)'>Go</button>" +
                '</td>' +
                '</tr>'
            );

        })

        Swal.close();

        sweetAlert2Toast.fire({
            icon: 'success',
            title: 'Search Completed!',
        });
    })
});

// Showing the number of entire document in the collection automatically
$(document).ready(function() {
    fetchCollectionDocumentQty(currentCollection)
        .then((data) => {
            $("#entire-document-count").text(data.toLocaleString());
        })
        .catch((error) => {
            Swal.Fire({
                icon: 'error',
                title: 'Error',
                text: 'Error while fetching the number of entire document : ' + error,
            
            })
            throw new Error('Error while fetching the number of entire document :', error);
        });
});