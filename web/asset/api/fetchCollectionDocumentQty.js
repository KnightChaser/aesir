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