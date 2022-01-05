function addListingLikes(itemID) {
    if (getCurrentUserID() < 0) {
        $('#promptLogIn').modal('show')
            //$('.modal-backdrop').remove();
        $(document).on('click', '#toLogin', function() {
            window.location.href = 'loginForm.html';
        })
    } else // is logged in
    {
        var userID = sessionStorage.getItem('status');

        fetch('https://' + getCurrentEnv() + '.herokuapp.com/add_listing_likes', {
                method: 'POST',
                headers: {
                    'Authorization': getToken(),
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "item_id": itemID,
                    "user_id": parseInt(getCurrentUserID())
                })
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);

                if (data.Data.is_liked == true) {
                    document.getElementById("like_" + itemID).innerHTML = '<span><i class="fas fa-heart" style="color:red"> </i></span>'
                    document.getElementById("likecount_" + itemID).innerHTML = '<span> ' + data.Data.latest_likes_count + '</span>'
                } else {
                    document.getElementById("like_" + itemID).innerHTML = '<span><i class="far fa-heart" style="color:black"></i></span>'
                    document.getElementById("likecount_" + itemID).innerHTML = '<span> ' + data.Data.latest_likes_count + '</span>'
                }
                // document.location.reload();
            })
            .catch(error => console.log(error));
    }
}