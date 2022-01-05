function checkIfUserLikedListing(is_liked) {
    if (is_liked && getCurrentUserID() != -1)
        return '<span><i class="fas fa-heart" style="color:red"> </i></span>';
    else
        return '<span><i class="far fa-heart" style="color:black"></i></span>';
}