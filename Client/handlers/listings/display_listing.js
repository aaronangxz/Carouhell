function displayListing(d, isRecommend) {
    var bool_value = isRecommend == 'true';
    var element = "cards"
    var rcmdHeader = ''
    if (bool_value === true) {
        element = "recommendedListings"
        rcmdHeader += '<div class="row mt-3">' +
            '<div class="col"><h3>💬 Comments <h3></div>' +
            '</div>';
    }

    var status = ""
    if (d.item_status == 2) {
        status += '<span class="badge badge-pill badge-secondary">Sold Out</span>'
    } else if (d.item_stock - d.item_quantity == 0) {
        status += '<span class="badge badge-pill badge-success">Available</span>'
    } else if (((d.item_quantity / d.item_stock) * 100) <= 10) {
        status += '<span class="badge badge-pill badge-danger">Low in stock</span>'
    } else if (((d.item_quantity / d.item_stock) * 100) <= 25) {
        status += '<span class="badge badge-pill badge-warning">Selling Fast</span>'
    } else if (d.item_status == 1) {
        status += '<span class="badge badge-pill badge-success">Available</span>'
    }

    var likes = ""
    if (d.listing_likes >= 10) {
        likes += '<span><i class="fas fa-chart-line" style="color:red"></i></span>'
    }

    var sellerType = ''
    if (d.seller_type == 1) {
        sellerType = '<span title = "Official Store"><i class="fas fa-check-circle" style="color:Dodgerblue"></i></span>'
    } else if (d.seller_type == 2) {
        sellerType = '<span title = "Preferred Seller"><i class="fas fa-medal" style="color:Coral"></i></span>'
    }

    document.getElementById(element).innerHTML +=
        '<div class="col-lg-3 col-md-4 col-sm-6 col-xs-6 mt-4">' +
        '<div class="card border-0" id="' + d.item_id + '">' +
        '<div class="card-header card-header-color border-0 pb-0"><a href="viewProfile.html?profileID=' + d.seller_id + '" style="color: black">@' + d.seller_name + '</a> ' + sellerType +
        '</div>' +
        '<div class="row">' +
        '<div class="col">' +
        '<p class = "pl-4 pt-0" id="itemDay"><font style="opacity:.8" size="2px">' + getTimeStamp(d.listing_ctime) + ' </font></p>' +
        '</div>' +
        '</div>' +
        '<div class="card-body pb-3">' +
        '<div class="container">' +
        '<div class="row ">' +
        '<a href="' + 'viewListing.html?itemID=' + d.item_id + '"class="btn btn-link stretched-link"></a> ' +
        '<div class="col-lg-12 col-md-12 col-sm-12 col-xs-6 text-center" id="imgContainer">' +
        '<img src="https://' + getCurrentEnv() + '.s3.ap-southeast-1.amazonaws.com/listing_' + d.item_id + '.jpg" class="img-cover img-fluid" />' +
        '</div>' +
        '</div>' +
        '<div class="row">' +
        '<h6 class="card-title pl-0" id="itemName">' + d.item_name + '</h6>' +
        '</div>' +
        '</div>' +
        '</div>' +
        '<div class="card-footer d-flex border-0">' +
        '<div class="row">' +
        '<div class="col">' +
        '<h5 class="card-title">S' + (parseInt(d.item_price) / 100).toLocaleString('en-US', {
            style: 'currency',
            currency: 'USD',
        }) +
        '</h5>' +
        '<h6 class="card-title" id="likebutton">' +
        '<a href="javascript:void(0);" onclick="addListingLikes(' + d.item_id + ');" id="like_' + d.item_id + '">' + checkIfUserLikedListing(d.is_liked) + '</a><span id="likecount_' + d.item_id + '"> ' + d.listing_likes + '</span>' +
        '</h6>' +
        '<h6 class="card-title">' + status + '  ' + likes +
        '</h6>' +
        '</div>' +
        '</div>' +
        '</div>' +
        '</div>' +
        '</div>';
}