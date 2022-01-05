function displayItemContent(data) {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const isPromoted = urlParams.get('promoted')

    console.log(isPromoted)
    var topBread =
        '<li class="breadcrumb-item"><a href=categories.html?category=' + data.item_category + '>' + categories_Arr[data.item_category] + '</a></li>' +
        '<li class="breadcrumb-item active" aria-current="page">' + data.item_name + '</li>'
    document.getElementById("topBread").innerHTML += topBread;

    var status = ""
    if (data.item_status == 2) {
        status += '<span class="badge badge-pill badge-secondary">Sold Out</span>'
    } else if (data.item_stock - data.item_quantity == 0) {
        status += '<span class="badge badge-pill badge-success">Available</span>'
    } else if (((data.item_quantity / data.item_stock) * 100) <= 10) {
        status += '<span class="badge badge-pill badge-danger">Low in stock</span>'
    } else if (((data.item_quantity / data.item_stock) * 100) <= 25) {
        status += '<span class="badge badge-pill badge-warning">Selling Fast</span>'
    } else if (data.item_status == 1) {
        status += '<span class="badge badge-pill badge-success">Available</span>'
    }

    if (isPromoted == 1) {
        status += '<span> </span><span class="badge badge-pill badge-secondary">Promoted</span>'
    }

    var progress = ""
    if (((data.item_quantity / data.item_stock) * 100) <= 10 && ((data.item_quantity / data.item_stock) * 100) != 0) {
        progress += '<div class="progress" style="width: 25%">' +
            '<div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" role="progressbar" style="width: 90%" aria-valuenow="90" aria-valuemin="0" aria-valuemax="100">' + (data.item_stock - data.item_quantity) + ' Bought</div>' +
            '</div>'
    } else if (((data.item_quantity / data.item_stock) * 100) <= 25 && ((data.item_quantity / data.item_stock) * 100) != 0) {
        progress += '<div class="progress" style="width: 25%;">' +
            '<div class="progress-bar progress-bar-striped bg-warning" role="progressbar" style="width: 75%" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100">' + (data.item_stock - data.item_quantity) + ' Bought</div>' +
            '</div>'
    }

    var sellerType = ''
    if (data.seller_type == 1) {
        sellerType = '<span title = "Official Store"><i class="fas fa-check-circle" style="color:Dodgerblue"></i></span>'
    } else if (data.seller_type == 2) {
        sellerType = '<span title = "Preferred Seller"><i class="fas fa-medal" style="color:Coral"></i></span>'
    }

    var content = "";
    content += '<div class="row">' +
        '<div class="col-4"><img src="https://' + getCurrentEnv() + '.s3.ap-southeast-1.amazonaws.com/listing_' + data.item_id + '.jpg" class="img-fluid" id="singleItemImage"/></div>' +
        '<div class="col-8">' +
        '<div class="row"><div class="col"><h1>' + data.item_name + ' </h1></div></div>' +
        '<div class="row">' +
        '<div class="col"><h2>' + '  S' + (parseInt(data.item_price) / 100).toLocaleString('en-US', {
            style: 'currency',
            currency: 'USD',
        }) + '</h2></div>' +
        '</div>' +
        '<div class="row">' +
        '<div class="col"><h3>' + status + '</h3></div>' +
        '</div>' +
        progress +

        '<div class="row">' +
        '<div class="col"><span><i class="fas fa-clock"></i></span> Posted ' + getTimeStamp(data.listing_ctime) + '</div>' +
        '</div>' +
        '<div class="row">' +
        '<div class="col">' +
        '<a href="javascript:void(0);" onclick="addListingLikes(' + data.item_id + ');" id="like_' + data.item_id + '">' + checkIfUserLikedListing(data.is_liked) + '</a><span id="likecount_' + data.item_id + '"> ' + data.listing_likes + '</span>' +
        '</div>' +
        '</div>' +
        '<div class="row mt-3">' +
        '<div class="col"><span><i class="fas fa-layer-group"></i></span> ' + categories_Arr[data.item_category] + '</div>' +
        '</div>' +
        '<div class="row">' +
        '<div class="col"><span><i class="fas fa-map-marker-alt"></i></span> ' + location_Arr[data.item_location] + '</div>' +
        '</div>' +
        '<div class="row">' +
        '<div class="col"><span><i class="fas fa-user"></i> </span><a href="viewProfile.html?profileID=' + data.seller_id + '"style="color: black">@' + data.seller_name + '</a> ' + sellerType + '</div>' +
        '</div>' +
        '<div class="row mt-3">' +
        '<div class="col"><h3>Description <h3></div>' +
        '</div>' +
        '<div class="row">' +
        '<div class="col">' + data.item_description + '</div>' +
        '</div>' +
        '</div>' +
        '</div>';
    if (getCurrentUserID() == -1) {
        content +=
            '<form id="purchaseForm" method="post">' +
            '<div class="row mt-3" id="buyerBtns">' +
            '<div class="col-4"></div>' +
            // '<div class="col-3 text-center">'+
            //     '<h5> Quantity Available: '+ data.item_quantity + '</h5>' +
            // '</div>'+
            // '<div class="col-1">'+
            //     '<input type="button" class="btn btn-secondary disabled" value="Sold Out"/>'+
            //     // '<input type="button" id="btnBuyNow" class="btn btn-primary" value="Buy Now"/>'+
            // '</div>'+
            '<div class="col-4"></div>' +
            '</div>' +
            '</form>';
    } else if (data.seller_id == getCurrentUserID()) // user is listing's owner
    {
        // display seller's option
        content +=
            '<div class="row mt-5" id="sellerBtns">' +
            '<div class="col-2">' +
            // '<input type="button" id="btnEdit" class="btn btn-primary" value="Edit Listing"/>'+
            '<a href="editListing.html?itemID=' + data.item_id + '" id="btnEdit" class="btn btn-secondary">Edit Listing</a> ' +
            '</div>' +
            '<div class="col-1">' +
            '<a href="javascript:void(0);" onclick="deleteListing(' + data.item_id + ')" id="deleteBtn" class="btn btn-danger">Delete Listing</a>' +
            '</div>' +
            '</div>';
    } else // is a buyer
    {
        if (data.item_quantity == 0) {
            content +=
                '<form id="purchaseForm" method="post">' +
                '<div class="row mt-3" id="buyerBtns">' +
                '<div class="col-4"></div>' +
                // '<div class="col-3 text-center">'+
                //     '<h5> Quantity Available: '+ data.item_quantity + '</h5>' +
                // '</div>'+
                // '<div class="col-1">'+
                //     '<input type="button" class="btn btn-secondary disabled" value="Sold Out"/>'+
                //     // '<input type="button" id="btnBuyNow" class="btn btn-primary" value="Buy Now"/>'+
                // '</div>'+
                '<div class="col-4"></div>' +
                '</div>' +
                '</form>';
        } else {
            content +=
                '<form id="purchaseForm" method="post">' +
                '<div class="row mt-3" id="buyerBtns">' +
                '<div class="col-4"></div>' +
                '<div class="col-2">' +
                '<div class="form-group">' +
                '<div class="input-group">' +
                '<span class="input-group-prepend">' +
                '<button type="button" class="btn btn-default btn-number" disabled="disabled" data-type="minus" data-field="quant[1]">' +
                '<span><i class="bi bi-dash-circle"></i></span>' +
                '</button>' +
                '</span>' +
                '<input type="text" name="quant[1]" id="qtyToPurchase" class="form-control input-number" value="1" min="1" max="10">' +
                '<span class="input-group-append">' +
                '<button type="button" class="btn btn-default btn-number" data-type="plus" data-field="quant[1]">' +
                '<span><i class="bi bi-plus-circle"></i></span>' +
                '</button>' +
                '</span>' +
                '</div>' +
                // '<input type="number" class="form-control" id="qtyToPurchase" placeholder="Enter Quantity">'+
                '</div>' +
                '</div>' +
                '<div class="col-3">' +
                '<span>Available: ' + data.item_quantity + '</span>' +
                '<span id="buyButtonSpan"><input type="button" id="buyButton" onclick="buyNow(' + data.item_id + ',' + data.seller_id + ')" class="btn btn-primary" value="Buy Now"/></span>' +
                '<div class="col-4"></div>' +
                '</div>' +
                '</form>';
        }
    }

    document.getElementById("listingContent").innerHTML = content;

    const aTag = document.getElementById('btnEdit');
    if (aTag) {
        console.log("aTag")
        aTag.addEventListener('click', function() { sessionStorage.setItem(data.item_id, JSON.stringify(data)); });
    }

    // const button = document.getElementById('btnBuyNow');
    // if(button)
    // {
    //     button.addEventListener('click', function(){ buyNow(data); });
    // }

    // const editButton = document.getElementById('btnEdit');
    // if(editButton)
    // {
    //     editButton.addEventListener('click', function(){ editListing(data); });
    // }

}