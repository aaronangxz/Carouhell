function displayItemComments(data, itemID) {
    var commentFooter = ''
    var showCommentButton =
        '<form id="listingComment" method="post">' +
        '<div class="row mt-3">' +
        '<div class="col-5">' +
        '<div class="form-group">' +
        '<input type="text" class="form-control" id="comment" placeholder="Add a comment...">' +
        '</div>' +
        '</div>' +
        '<div class="col-2">' +
        '<input type="button" class="btn btn-secondary" onclick="addListingComment(getCurrentUserID(),' + itemID + ')" value="Submit"/>' +
        '</div>' +
        '</div>' +
        '</form>'
    if (getCurrentUserID() == -1) {
        showCommentButton = ""
    }
    var comments =
        '<div class="row mt-3">' +
        '<div class="col"><h3><span><i class="far fa-comment-dots" style="color:grey"></i></span> Comments <h3></div>' +
        '</div>';
    if (data.length == 0) {
        if (getCurrentUserID() == -1) {
            comments += '<font style="opacity:.6" size="3px"><i><a href="loginForm.html">Login</a> to start a discussion.</i></font>'
        } else {
            comments += '<font style="opacity:.6" size="3px"><i>Start a discussion.</i></font>'
        }
    } else {
        if (getCurrentUserID() == -1) {
            commentFooter += '<font style="opacity:.6" size="3px"><i><a href="loginForm.html">Login</a> to join the discussion.</i></font>'
        } else {
            commentFooter += '<font style="opacity:.6" size="3px"><i>Join the discussion.</i></font>'
        }
    }
    for (var i = 0; i < data.length; i++) {
        console.log("comments: " + data[i]);
        comments +=
            '<div class="row text-wrap" >' +
            '<div class="col-3"><a href="viewProfile.html?profileID=' + data[i].user_id + '">' + data[i].user_name + '</a> | ' + convertUnixToTimeStamp(data[i].ctime) + '</div>' +
            '<div class="col-9">' + data[i].comment + '</div>' +
            '</div>';
    }
    comments += showCommentButton + commentFooter;

    document.getElementById("commentsSection").innerHTML = comments;
}