function createListing(userID) {
    var myModal = document.getElementById('failedCreateListing');
    var createListingError = bootstrap.Modal.getOrCreateInstance(myModal)
    var modalTitle = myModal.querySelector('.modal-title')
    var modalBodyInput = myModal.querySelector('.modal-body')

    if (getCurrentUserID() == -1) {
        window.location.href = "index.html"
    }
    setPrevSecureLocation();

    var itemName = document.getElementById('itemName').value;
    var itemPrice = document.getElementById("itemPrice").value;
    var itemQty = document.getElementById("itemQty").value;
    var itemDesc = document.getElementById("itemDesc").value;
    var itemCat = document.getElementById("itemCat");
    var itemCatValue = itemCat.options[itemCat.selectedIndex].value;
    var location = document.getElementById("itemLocation");
    var locationValue = location.options[location.selectedIndex].value;

    var base64String = "";
    const file = document.getElementById("imgBrowser").files[0];
    var reader = new FileReader();

    reader.onload = function() {
        base64String = reader.result.replace("data:", "")
            .replace(/^.+,/, "");

        imageBase64Stringsep = base64String;
    }

    if (file && itemName && itemPrice && itemQty && itemDesc && itemCatValue && locationValue != null) {
        reader.readAsDataURL(file);
    } else {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "All fields and image are required."
        })
        createListingError.show();
    }

    if (itemName.length > 40) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Item Name should not exceed 40 characters."
        })
        createListingError.show();
    } else if (itemPrice > 42949672) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Maximum price allowed is S$42,949,672."
        })
        createListingError.show();
    } else if (itemQty > 42949672 || itemQty <= 0) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Quantity allowed is 1 - 42,949,672."
        })
        createListingError.show();
    } else if (itemDesc.length > 500) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Item Description should not exceed 500 characters."
        })
        createListingError.show();
    } else {
        document.getElementById("footer").innerHTML += '<div class="loader-wrapper">' +
            '<span class="loader"><span class="loader-inner"></span></span>' +
            '</div>'

        reader.addEventListener('load', (event) => {
            console.log('load finish: ' + base64String); // base 64

            fetch('https://' + getCurrentEnv() + '.herokuapp.com/create_listing', {
                    method: 'POST',
                    headers: {
                        'Authorization': getToken(),
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        "item_name": itemName,
                        "item_price": parseFloat(itemPrice) * 100,
                        "item_quantity": parseInt(itemQty),
                        "item_description": itemDesc,
                        "item_location": parseInt(locationValue),
                        "item_category": parseInt(itemCatValue),
                        "item_image": base64String,
                        "seller_id": parseInt(userID)
                    })
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                    if (data.Respmeta.ErrorCode != 0) {
                        // var myModal = document.getElementById('failedCreateListing');
                        // var registerFailedModal = bootstrap.Modal.getOrCreateInstance(myModal)

                        myModal.addEventListener('show.bs.modal', function() {
                            // var modalTitle = myModal.querySelector('.modal-title')
                            // var modalBodyInput = myModal.querySelector('.modal-body')
                            modalTitle.textContent = 'Unable to create listing: Error ' + data.Respmeta.ErrorCode
                            modalBodyInput.textContent = data.Respmeta.DebugMsg
                        })
                        createListingError.show();
                    } else // successful
                    {
                        $(".loader-wrapper").fadeOut("slow");
                        window.location.href = "viewListing.html?itemID=" + data.Data.item_id;
                    }
                })
                .catch(error => console.log(error));
        });
    }


}