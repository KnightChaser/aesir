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