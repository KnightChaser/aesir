currentURL = window.location.href;
currentBaseURL = currentURL.split('/').slice(0, 3).toString().replace(",,", "//");
currentCollection = currentURL.split('/').slice(4);

function selectizeTabSourceImages() {
    searchConditionSourceImages = [
        {
            $match: {
                "event.eventdata.SourceImage": { $ne: null }
            }
        },
        {
            $group: {
                _id: "$event.eventdata.SourceImage",
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
        }
    ];

    fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionSourceImages)).then((data) => {
        let sourceImageStatistics = JSON.parse(data);

        sourceImageStatistics.forEach((item, index) => {
            sourceImageStatistics[index].sourceImagesAbbreviated = sourceImageStatistics[index].sourceImage.split('\\').pop();
        });

        $("#source-image-selectize-tab").selectize({
            persist: false,
            maxItems: null,
            valueField: "sourceImage",
            labelField: "sourceImage",
            searchField: "sourceImage",
            options: sourceImageStatistics,
            render: {
                item: function (item, escape) {
                    return (
                        "<div>" +
                        "<span>" + escape(item.sourceImagesAbbreviated) + "</span>" +
                        " <span>(" + escape(item.count) + " matches)</span>" +
                        "</div>"
                    );
                },
                option: function (item, escape) {
                    return (
                        "<div>" +
                        '<span class="label"><b>' +
                        escape(item.sourceImagesAbbreviated) +
                        "</b><br>" +
                        "</span>" +
                        '<span class="caption">' + escape(item.count) + " matches" + "<br>" +
                        "(<code>" + escape(item.sourceImage) + "</code>)" +
                        "</span>" +
                        "</div>"
                    );
                },
            },
            createFilter: function (input) {
                var match, regex;

                regex = new RegExp("^" + input + "$", "i");
                match = sourceImagesAbbreviated.find(item => item.match(regex));
                if (match) return false;

                return true;
            },
            create: function (input) {
                return { sourceImage: input };
            },
        });

    });
}

selectizeTabSourceImages();

function selectizeTabTargetImages() {
    searchConditionTargetImages = [
        {
            $match: {
                "event.eventdata.TargetImage": { $ne: null }
            }
        },
        {
            $group: {
                _id: "$event.eventdata.TargetImage",
                count: { $sum: 1 }
            }
        },
        {
            $sort: { count: -1 }
        },
        {
            $project: {
                _id: 0,
                targetImage: "$_id",
                count: 1
            }
        }
    ];

    fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionTargetImages)).then((data) => {
        let targetImageStatistics = JSON.parse(data);

        targetImageStatistics.forEach((item, index) => {
            targetImageStatistics[index].targetImagesAbbreviated = targetImageStatistics[index].targetImage.split('\\').pop();
        });

        $("#target-image-selectize-tab").selectize({
            persist: false,
            maxItems: null,
            valueField: "targetImage",
            labelField: "targetImage",
            searchField: "targetImage",
            options: targetImageStatistics,
            render: {
                item: function (item, escape) {
                    return (
                        "<div>" +
                        "<span>" + escape(item.targetImagesAbbreviated) + "</span>" +
                        " <span>(" + escape(item.count) + " matches)</span>" +
                        "</div>"
                    );
                },
                option: function (item, escape) {
                    return (
                        "<div>" +
                        '<span class="label"><b>' +
                        escape(item.targetImagesAbbreviated) +
                        "</b><br>" +
                        "</span>" +
                        '<span class="caption">' + escape(item.count) + " matches" + "<br>" +
                        "(<code>" + escape(item.targetImage) + "</code>)" +
                        "</span>" +
                        "</div>"
                    );
                },
            },
            createFilter: function (input) {
                var match, regex;

                regex = new RegExp("^" + input + "$", "i");
                match = sourceImagesAbbreviated.find(item => item.match(regex));
                if (match) return false;

                return true;
            },
            create: function (input) {
                return { sourceImage: input };
            },
        });

    });
}

selectizeTabTargetImages();

function selectizeTabRuleName() {
    searchConditionRuleName = [
        {
            $match: {
                "event.eventdata.RuleName": { $ne: null }
            }
        },
        {
            $group: {
                _id: "$event.eventdata.RuleName",
                count: { $sum: 1 }
            }
        },
        {
            $sort: { count: -1 }
        },
        {
            $project: {
                _id: 0,
                ruleName: "$_id",
                count: 1
            }
        }
    ];

    fetchEventData(currentCollection, "aggregate", JSON.stringify(searchConditionRuleName)).then((data) => {
        let targetImageStatistics = JSON.parse(data);

        $("#rule-name-selectize-tab").selectize({
            persist: false,
            maxItems: null,
            valueField: "ruleName",
            labelField: "ruleName",
            searchField: "ruleName",
            options: targetImageStatistics,
            render: {
                item: function (item, escape) {
                    return (
                        "<div>" +
                        "<span>" + escape(item.ruleName) + "</span>" +
                        " <span>(" + escape(item.count) + " matches)</span>" +
                        "</div>"
                    );
                },
                option: function (item, escape) {
                    return (
                        "<div>" +
                        '<span class="label"><b>' +
                        escape(item.ruleName) +
                        "</b><br>" +
                        "</span>" +
                        '<span class="caption">' + escape(item.count) + " matches" + "<br>" +
                        "</span>" +
                        "</div>"
                    );
                },
            },
            createFilter: function (input) {
                var match, regex;

                regex = new RegExp("^" + input + "$", "i");
                match = sourceImagesAbbreviated.find(item => item.match(regex));
                if (match) return false;

                return true;
            },
            create: function (input) {
                return { sourceImage: input };
            },
        });

    });
}

selectizeTabRuleName();

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

    fetchEventDataMultipleCondition(currentCollection, JSON.stringify(searchCondition)).then((response) => {

        // Clear the table before appending the search result
        $("#search-result-table-row").empty();

        // Showing the number of results.
        $("#searched-document-count").text(response["count"]);

        // Order searchResult JSON data by objectId in descending order
        let searchResult = JSON.parse(response["result"])
        searchResult.sort((a, b) => (a._id < b._id) ? 1 : -1);

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
                '<td>' + ruleName + '</td>' +
                '<td><code>' + sourceImageData.sourceImage + '</code></td>' +
                '<td><code>' + targetImageData.targetImage + '</code></td>' +
                '<td>' + eventOccurredTime + '</td>' +
                '<td>' +
                "<button type='button' class='btn btn-primary' style='margin:auto; display:block;' onclick='showAdditionalInfoAlert(\`" + encodeURI(additionalInformationAlertHTML) + "\`)'>Go</button>" +
                '</td>' +
                '</tr>'
            );
        })
    })
});