currentURL = window.location.href;
currentBaseURL = currentURL.split('/').slice(0, 3).toString().replace(',,', '//');
currentCollection = currentURL.split('/').slice(4);

// General conditional search API
async function fetchEventData(collectionInUse, searchCondition) {
    // Replace 'http://localhost:8080/api' with your actual API endpoint
    const apiUrl = `${currentBaseURL}/api/search/${encodeURI(collectionInUse)}/${encodeURI(searchCondition)}`;
  
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