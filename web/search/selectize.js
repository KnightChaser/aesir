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

// If button (id = "search-form-submit-button") is clicked, print the selected values of the selectize.js dropdowns
// Also, get the values of date range picker
$("#search-form-submit-button").click(function () {
    console.log("Source Image: " + $("#source-image-selectize-tab").val());
    console.log("Target Image: " + $("#target-image-selectize-tab").val());
    console.log("Rule Name: " + $("#rule-name-selectize-tab").val());
    console.log("Date Range (start): " + $("#starting-datetime").val());
    console.log("Date Range (end): " + $("#ending-datetime").val());
});