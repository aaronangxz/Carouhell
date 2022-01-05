function editListing(itemID) {
    var myModal = document.getElementById('failedEditListing');
    var editListingError = bootstrap.Modal.getOrCreateInstance(myModal)
    var modalTitle = myModal.querySelector('.modal-title')
    var modalBodyInput = myModal.querySelector('.modal-body')

    document.getElementById("footer").innerHTML += '<div class="loader-wrapper">' +
        '<span class="loader"><span class="loader-inner"></span></span>' +
        '</div>'

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

    console.log("Price: " + itemPrice);

    if (itemPrice || itemPrice != "")
        itemPrice = parseFloat(itemPrice);
    else
        itemPrice = null;

    if (itemQty || itemQty != "")
        itemQty = parseInt(itemQty);
    else
        itemQty = null;

    if (!itemName || itemName == "")
        itemName = null;

    if (!itemDesc || itemDesc == "")
        itemDesc = null;

    if (itemCatValue)
        itemCatValue = parseInt(itemCatValue);

    if (locationValue)
        locationValue = parseInt(locationValue);


    var base64String = "";
    const file1 = document.getElementById("imgBrowser").files[0];
    console.log("file1: " + file1);

    if (itemName.length > 40) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Item Name should not exceed 40 characters."
        })
        editListingError.show();
        $(".loader-wrapper").fadeOut("slow");
    } else if (itemPrice > 42949672) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Maximum price allowed is S$42,949,672."
        })
        editListingError.show();
        $(".loader-wrapper").fadeOut("slow");
    } else if (itemQty > 42949672 || itemQty <= 0) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Quantity allowed is 1 - 42,949,672."
        })
        editListingError.show();
        $(".loader-wrapper").fadeOut("slow");
    } else if (itemDesc.length > 500) {
        myModal.addEventListener('show.bs.modal', function() {
            modalTitle.textContent = 'Unable to create listing'
            modalBodyInput.textContent = "Item Description should not exceed 500 characters."
        })
        editListingError.show();
        $(".loader-wrapper").fadeOut("slow");
    } else {
        if (file1) {
            var reader = new FileReader();

            reader.onload = function() {
                base64String = reader.result.replace("data:", "")
                    .replace(/^.+,/, "");

                imageBase64Stringsep = base64String;
            }
            reader.readAsDataURL(file1);
            reader.addEventListener('load', (event) => {
                console.log('load finish: ' + base64String); // base 64  
                fetch('https://' + getCurrentEnv() + '.herokuapp.com/update_single_listing', {
                        method: 'PATCH',
                        headers: {
                            'Authorization': getToken(),
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({
                            "item_id": parseInt(itemID),
                            "item_name": itemName,
                            "item_price": itemPrice * 100,
                            "item_quantity": itemQty,
                            "item_description": itemDesc,
                            "item_location": locationValue,
                            "item_category": itemCatValue,
                            "item_image": base64String,
                            "seller_id": parseInt(getCurrentUserID())
                        })
                    })
                    .then(response => response.json())
                    .then(data => {
                        console.log(data);
                        if (data.Respmeta.ErrorCode != 0) {
                            myModal.addEventListener('show.bs.modal', function() {
                                modalTitle.textContent = 'Unable to edit listing. Error:' + data.Respmeta.ErrorCode
                                modalBodyInput.textContent = data.Respmeta.DebugMsg
                            })
                            editListingError.show();
                            $(".loader-wrapper").fadeOut("slow");
                        } else // successful
                        {
                            $(".loader-wrapper").fadeOut("slow");
                            window.location.href = "viewListing.html?itemID=" + itemID;
                        }
                    })
                    .catch(error => console.log(error));
            });
        } else {
            fetch('https://' + getCurrentEnv() + '.herokuapp.com/update_single_listing', {
                    method: 'PATCH',
                    headers: {
                        'Authorization': getToken(),
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        "item_id": parseInt(itemID),
                        "item_name": itemName,
                        "item_price": itemPrice * 100,
                        "item_quantity": itemQty,
                        "item_description": itemDesc,
                        "item_location": locationValue,
                        "item_category": itemCatValue,
                        "item_image": null,
                        "seller_id": parseInt(getCurrentUserID())
                    })
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                    if (data.Respmeta.ErrorCode != 0) {
                        myModal.addEventListener('show.bs.modal', function() {
                            modalTitle.textContent = 'Unable to edit listing. Error:' + data.Respmeta.ErrorCode
                            modalBodyInput.textContent = data.Respmeta.DebugMsg
                        })
                        editListingError.show();
                        $(".loader-wrapper").fadeOut("slow");
                    } else // successful
                    {
                        $(".loader-wrapper").fadeOut("slow");
                        window.location.href = "viewListing.html?itemID=" + itemID;
                    }
                })
                .catch(error => console.log(error));
        }
    }
}