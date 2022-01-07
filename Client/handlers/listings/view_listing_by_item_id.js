function viewListingByItemId(itemID) {
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/get_single_listing_by_itemid', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "item_id": parseInt(itemID)
            })
        })
        .then(response => response.json())
        .then(data => {
            if (data.Respmeta.ErrorCode != 0) {
                var myModal = document.getElementById('loadListingFail');
                var ListingLoadFailedModal = bootstrap.Modal.getOrCreateInstance(myModal)

                myModal.addEventListener('show.bs.modal', function() {
                    var modalTitle = myModal.querySelector('.modal-title')
                    var modalBodyInput = myModal.querySelector('.modal-body')
                    modalTitle.textContent = 'Unable to retrieve listing: Error ' + data.Respmeta.ErrorCode
                    modalBodyInput.textContent = data.Respmeta.DebugMsg
                })
                ListingLoadFailedModal.show();
                myModal.addEventListener('hide.bs.modal', function() {
                    window.location.href = getPrevLocation()
                })

            } else {
                setPrevLocation();
                $(".loader-wrapper").fadeOut("fast");
                displayItemContent(data.Data);
                displayItemComments(data.Data.listing_comments, data.Data.item_id);
                getRecommendedListingsByItemId(data.Data.item_id);
                document.title = 'Carouhell - ' + data.Data.item_name;
            }

        })
        .catch(error => console.log(error));
}