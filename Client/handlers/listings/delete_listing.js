function deleteListing(itemID) {
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/delete_single_listing', {
            method: 'DELETE',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "item_id": itemID
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                if (confirm("Failed deleting listing")) {
                    location.reload();
                }
            } else // successful
            {
                if (confirm("Listing Deleted!")) {
                    window.location.href = "index.html";
                }
            }
        })
        .catch(error => console.log(error));
}