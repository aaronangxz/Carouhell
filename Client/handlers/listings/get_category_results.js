function getCategoryResults(selectedCategory) {
    document.getElementById("footer").innerHTML = ''
    document.getElementById("cards").innerHTML = '<div class="loader-wrapper">' +
        '<span class="loader"><span class="loader-inner"></span></span>' +
        '</div>'

    fetch('https://' + getCurrentEnv() + '.herokuapp.com/get_listings_using_filters', {
            method: 'POST',
            headers: {
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "category_filter": {
                    "item_category": parseInt(selectedCategory)
                },
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                if (data.Respmeta.ErrorCode == 3) {
                    document.getElementById("searchResultText").innerHTML =
                        '<div class="col">' +
                        '<h1>Hmm..we don\'t have that here.</h1>' +
                        '<h2>Try to search for something else?</h2>' +
                        '</div>';
                    document.getElementById("landingHomePage").innerHTML = ''
                    document.getElementById("cards").innerHTML = "";
                } else {
                    var myModal = document.getElementById('noResults');
                    var registerFailedModal = bootstrap.Modal.getOrCreateInstance(myModal)

                    myModal.addEventListener('show.bs.modal', function() {
                        var modalTitle = myModal.querySelector('.modal-title')
                        var modalBodyInput = myModal.querySelector('.modal-body')
                        modalTitle.textContent = 'Unable to retrieve results: Error ' + data.Respmeta.ErrorCode
                        modalBodyInput.textContent = data.Respmeta.DebugMsg
                    })
                    registerFailedModal.show();
                    myModal.addEventListener('hide.bs.modal', function() {
                        location.reload();
                    })

                    $('#noResults').modal('show')
                    $('#noResults').on('hidden.bs.modal', function(e) {
                        location.reload();
                    })
                }
            } else // successful
            {
                document.getElementById("landingHomePage").innerHTML = '';
                document.getElementById("cards").innerHTML = "";
                document.getElementById("searchResultText").innerHTML = "";
                document.getElementById("searchResultText").innerHTML +=
                    '<div class="col">' +
                    '<p></p>' +
                    '<h1>' + categories_Arr[selectedCategory] + '</h1>' +
                    '<h6>' + data.Data.length + ' listings found in ' + categories_Arr[selectedCategory] + '</h6>'
                '</div>';
                for (const d of data.Data) {
                    displayListing(d, );
                }
                document.title = 'Carouhell - ' + categories_Arr[selectedCategory];
            }
        })
        .catch(error => console.log(error));
    sessionStorage.setItem('prevLocation', window.location)
}