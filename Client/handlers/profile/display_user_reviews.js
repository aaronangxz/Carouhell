function displayUserReviews(data) {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const isPromoted = urlParams.get('promoted')

    var sellerType = ''
    if (data.account_info.seller_type == 1) {
        sellerType = '<span title = "Official Store"><i class="fas fa-check-circle" style="color:Dodgerblue"></i></span>'
    } else if (data.account_info.seller_type == 2) {
        sellerType = '<span title = "Preferred Seller"><i class="fas fa-medal" style="color:Coral"></i></span>'
    }

    if (isPromoted == 1) {
        sellerType += '<span> </span><span class="badge badge-pill badge-secondary">Promoted</span>'
    }

    var reviews =
        '<div class="row mt-3">' +
        '<div class="col"><h1>@' + data.account_info.user_name + ' ' + sellerType + '</h1></div>' +
        '</div>' +
        '<div class="row mt-1">' +
        '<div class="col"><small style="color:grey;"> ' + getLastSeen(data.account_info.user_last_login) + '</small></div>' +
        '</div>' +
        '<div class="row mt-3">' +
        '<div class="col"><h2>' + data.ratings.user_ratings + '<span><i class="fas fa-star" style="color:gold"></i></span> (' + data.review_count + ' Reviews)</h2></div>' +
        '</div>';

    for (var i = 0; i < data.user_reviews.length; i++) {
        if (i == 10) {
            break
        }

        rate = ""
        if (data.user_reviews[i].ratings == 1) {
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        } else if (data.user_reviews[i].ratings == 2) {
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        } else if (data.user_reviews[i].ratings == 3) {
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        } else if (data.user_reviews[i].ratings == 4) {
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        } else if (data.user_reviews[i].ratings == 5) {
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span>'
        }

        reviews +=
            '<div class="row text-wrap">' +
            '<div class="col-3">' + convertUnixToTimeStamp(data.user_reviews[i].ctime) + ' | <a href="viewProfile.html?profileID=' + data.user_reviews[i].user_id + '">' + data.user_reviews[i].user_name + '</a></div>' +
            '<div class="col-2">' + data.user_reviews[i].ratings + '/5' + rate + '</div>' +
            '<div class="col-7">' + data.user_reviews[i].review_text + '</div>' +
            '</div>';
    }

    reviews += '<div class="mb-4"><hr class="solid"></div>'
    document.getElementById("reviewSection").innerHTML = reviews;
}