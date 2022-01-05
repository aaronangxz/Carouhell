function viewProfileByUserID(profileID) {
    document.getElementById("footer").innerHTML += '<div class="loader-wrapper">' +
        '<span class="loader"><span class="loader-inner"></span></span>' +
        '</div>'
    var currentUser = getCurrentUserID();
    if (!profileID) {
        // viewing own listings
        profileID = currentUser;
    }

    fetch('https://' + getCurrentEnv() + '.herokuapp.com/get_user_details', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "user_id": parseInt(profileID)
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                var myModal = document.getElementById('loadProfileFail');
                var profileLoadFailedModal = bootstrap.Modal.getOrCreateInstance(myModal)

                myModal.addEventListener('show.bs.modal', function() {
                    var modalTitle = myModal.querySelector('.modal-title')
                    var modalBodyInput = myModal.querySelector('.modal-body')
                    modalTitle.textContent = 'Unable to retrieve profile: Error ' + data.Respmeta.ErrorCode
                    modalBodyInput.textContent = data.Respmeta.DebugMsg
                })
                profileLoadFailedModal.show();
                myModal.addEventListener('hide.bs.modal', function() {
                    window.location.href = getPrevLocation()
                })
            } else // successful
            {
                setPrevLocation();
                document.getElementById("cards").innerHTML = "";
                document.getElementById("title").innerHTML = "<h2>" + data.Data.user_listings.length + " listings</h2>";
                for (const d of data.Data.user_listings) {
                    displayListing(d);
                }
                $(".loader-wrapper").fadeOut("fast");
                displayUserReviews(data.Data);
                document.title = 'Carouhell - @' + data.Data.account_info.user_name + "'s profile";
            }
        })
        .catch(error => console.log(error));

}