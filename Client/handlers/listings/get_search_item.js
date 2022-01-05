function getSearchItem() {
    document.getElementById("carouselExampleIndicators").innerHTML = '';
    document.getElementById("landingHomePage").innerHTML = '<div class="loader-wrapper">' +
        '<span class="loader"><span class="loader-inner"></span></span>' +
        '</div>';
    var searchItem = document.getElementById('searchItem').value;
    console.log(searchItem)
    if (searchItem == '' || null) {
        location.reload();
    }
    console.log(searchItem + "- searchItem");
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/get_listings_using_filters', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "search_keyword": searchItem,
            })
        })
        .then(response => response.json())
        .then(data => {
            //console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                document.getElementById("searchResultText").innerHTML =
                    '<div class="col">' +
                    '<h1>Hmm..we don\'t have that here.</h1>' +
                    '<h2>Try to search for something else?</h2>' +
                    '</div>';
                $(".loader-wrapper").fadeOut("slow");
                document.getElementById("landingHomePage").innerHTML = ''
                document.getElementById("cards").innerHTML = "";
            } else // successful
            {
                console.log(data);
                document.getElementById("landingHomePage").innerHTML = '';
                document.getElementById("cards").innerHTML = "";
                document.getElementById("searchResultText").innerHTML = "";
                document.getElementById("searchResultText").innerHTML +=
                    '<div class="col">' +
                    '<h1>' + data.Data.length + ' search result(s) found</h1>' +
                    '</div>';
                for (const d of data.Data) {
                    displayListing(d, );
                }
            }
        })
        .catch(error => console.log(error));
    sessionStorage.setItem('prevLocation', window.location)
}