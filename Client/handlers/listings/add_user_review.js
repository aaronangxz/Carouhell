function addUserReview(sellerID) {
    setPrevSecureLocation();
    var rating = document.getElementById("rating");
    var ratingValue = rating.options[rating.selectedIndex].value;
    var review = document.getElementById('review').value;

    if (!review || !ratingValue || review == "") {
        alert("Review incomplete, please try again");
        location.reload();
    }

    fetch('https://' + getCurrentEnv() + '.herokuapp.com/add_user_review', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "user_id": parseInt(getCurrentUserID()),
                "seller_id": parseInt(sellerID),
                "ratings": parseInt(ratingValue),
                "review_text": review
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                if (confirm("Review failed due to the following reason: " + data.Respmeta.DebugMsg)) {
                    location.reload();
                }
            } else // successful
            {
                if (confirm("Successfully added review")) {
                    window.location.href = "wallet.html";
                }
            }
        })
        .catch(error => console.log(error));
}