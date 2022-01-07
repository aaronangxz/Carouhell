function getUserLikedListing() {
    document.getElementById("footer").innerHTML += '<div class="loader-wrapper">' +
        '<span class="loader"><span class="loader-inner"></span></span>' +
        '</div>'
    if (getCurrentUserID() == -1) {
        window.location.href = "index.html"
    }
    setPrevSecureLocation();
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/get_user_liked_listings', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "user_id": parseInt(getCurrentUserID())
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            document.getElementById("title").innerHTML = '<h1>My Favourites(' + data.Data.length + ')</h1>';
            if (data.Respmeta.ErrorCode != 0) {
                if (confirm("Unable to get user liked listings due to the following reason: " + data.Respmeta.DebugMsg)) {
                    //location.reload();
                }
            } else // successful
            {
                document.getElementById("cards").innerHTML = "";
                $(".loader-wrapper").fadeOut("fast");
                for (const d of data.Data) {
                    displayListing(d);
                }
            }
        })
        .catch(error => console.log(error));
}