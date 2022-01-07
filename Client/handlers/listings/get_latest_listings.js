function getLatestListing() {
    document.getElementById("footer").innerHTML += '<div class="loader-wrapper">' +
        '<span class="loader"><span class="loader-inner"></span></span>' +
        '</div>'
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/v2/get_latest_listings', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "user_id": parseInt(getCurrentUserID()),
                "pagination": parseInt(1)
            })
        })
        .then(response => response.json())
        .then(result => {
            for (const d of result.Data.Listings) {
                displayListing(d);
            }
            displayPaginations(result.Data.Paginations.total_page, result.Data.Paginations.current_page, result.Data.Paginations.next_page);
            $(".loader-wrapper").fadeOut("slow");
        })
        .catch(error => { console.log('NO:', JSON.stringify(error)); });
    sessionStorage.setItem('prevLocation', window.location)
}

function getLatestListingPaginated(page) {
    if (page == null) {
        page = 1
    }
    document.getElementById("footer").innerHTML += '<div class="loader-wrapper">' +
        '<span class="loader"><span class="loader-inner"></span></span>' +
        '</div>'
    document.getElementById("cards").innerHTML = '<p></p>'
    document.getElementById("paginationDisplay").innerHTML = '<p></p>'

    fetch('https://' + getCurrentEnv() + '.herokuapp.com/v2/get_latest_listings', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "user_id": parseInt(getCurrentUserID()),
                "pagination": parseInt(page)
            })
        })
        .then(response => response.json())
        .then(result => {
            for (const d of result.Data.Listings) {
                displayListing(d);
            }
            displayPaginations(result.Data.Paginations.total_page, result.Data.Paginations.current_page, result.Data.Paginations.next_page);
            $(".loader-wrapper").fadeOut("slow");
        })
        .catch(error => { console.log('NO:', JSON.stringify(error)); });
    sessionStorage.setItem('prevLocation', window.location)
}