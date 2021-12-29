/* constantly checking if need to reset Nav Bar links according to user's status(logged in?) */
function resetNavBar() {
    if (localStorage.getItem('userID') != null) // user logged in
    {
        document.getElementById("register").style.display = "none";
        document.getElementById("login").style.display = "none";
        document.getElementById("sell").style.display = "block";
        document.getElementById("profile").style.display = "block";
        document.getElementById("favourites").style.display = "block";
        document.getElementById("wallet").style.display = "block";
        document.getElementById("logout").style.display = "block";
        document.getElementById("notification").style.display = "block";

        return true;
    } else {
        ""
        document.getElementById("register").style.display = "block";
        document.getElementById("login").style.display = "block";
        document.getElementById("sell").style.display = "none";
        document.getElementById("profile").style.display = "none";
        document.getElementById("favourites").style.display = "none";
        document.getElementById("wallet").style.display = "none";
        document.getElementById("logout").style.display = "none";
        document.getElementById("notification").style.display = "none";

        return false;
    }
}