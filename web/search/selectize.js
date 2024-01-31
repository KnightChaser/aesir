currentURL = window.location.href;
currentBaseURL = currentURL.split('/').slice(0, 3).toString().replace(",,", "//");
currentCollection = currentURL.split('/').slice(4);

searchCondition = [
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

fetchEventData(currentCollection, "aggregate", JSON.stringify(searchCondition)).then((data) => {
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
                    (item.sourceImage
                        ? '<span>' + escape(item.sourceImagesAbbreviated) + "</span>"
                        : "") +
                    (item.count
                        ? ' <span>(' + escape(item.count) + " matches)</span>"
                        : "") +
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