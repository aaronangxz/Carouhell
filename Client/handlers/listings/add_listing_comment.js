function addListingComment(userID, itemID) {
    var comment = document.getElementById('comment').value;
    if (comment == "") {
        confirm("Comment cannot be empty.")
        return
    }
    console.log("comment:" + comment);
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/add_listing_comments', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "item_id": parseInt(itemID),
                "user_id": parseInt(userID),
                "comment": comment
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                if (confirm("Comment failed due to the following reason: " + data.Respmeta.DebugMsg)) {
                    location.reload();
                }
            } else // successful
            {
                location.reload();
            }
        })
        .catch(error => console.log(error));

}