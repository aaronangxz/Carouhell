function getRecommendedListingsByItemId(itemID) {
    document.getElementById("recommendedListingsSection").innerHTML +=
        '<div class="mb-4"><hr class="solid"></div>' +
        '<div class="row mt-3">' +
        '<div class="col"><h3><i class="fas fa-fire" style="color:red"></i></span> You may also like <h3></div>' +
        '</div>'
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/v2/get_recommended_listings_by_itemid', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "item_id": parseInt(itemID),
                "user_id": parseInt(getCurrentUserID())
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);

            for (const d of data.Data) {
                displayListing(d, 'true');
            }
        })
        .catch(error => console.log(error));
}