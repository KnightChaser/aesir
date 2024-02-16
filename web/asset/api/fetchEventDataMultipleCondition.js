// Document counting in the current collection API
async function fetchEventDataMultipleCondition(collectionInUse, condition) {
    const apiUrl = `${currentBaseURL}/api/searchMultipleCondition/${encodeURI(collectionInUse)}/${encodeURI(condition)}`;

    return fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Network response was not ok: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            return {
                "result": data["result"],
                "count": data["count"]
            };
        })
        .catch(error => {
            console.error('Error:', error);
            throw error;
        });
}