function getFilterResults() {
    // filterOptions    
    document.getElementById("carouselExampleIndicators").innerHTML = '';
    document.getElementById("paginationDisplay").innerHTML = '<p></p>'
    var sortByCat = document.getElementById("sortByCat");
    var selectedCategory = sortByCat.options[sortByCat.selectedIndex].value;
    var sortByLocation = document.getElementById("sortByLocation");
    var selectedLocation = sortByLocation.options[sortByLocation.selectedIndex].value;
    var minPrice = document.getElementById("minPrice").value;
    var maxPrice = document.getElementById("maxPrice").value;
    var sortbyFlag = document.getElementById("sortFlag");
    var selectedSortFlag = sortbyFlag.options[sortbyFlag.selectedIndex].value;


    if (selectedCategory)
        selectedCategory = parseInt(selectedCategory);

    if (selectedLocation)
        selectedLocation = parseInt(selectedLocation);

    if (parseInt(minPrice) <= 0) {
        minPrice = null;
        document.getElementById('minPrice').value = ''
    }

    if (minPrice || minPrice != "")
        minPrice = parseFloat(minPrice) * 100;
    else
        minPrice = null;

    if (parseInt(maxPrice) <= 0) {
        maxPrice = null;
        document.getElementById('maxPrice').value = ''
    }

    if (maxPrice || maxPrice != "")
        maxPrice = parseFloat(maxPrice) * 100;
    else
        maxPrice = null;

    if (parseInt(minPrice) >= 4294967295 || parseInt(maxPrice) >= 4294967295) {
        var myModal = document.getElementById('noResults');
        var registerFailedModal = bootstrap.Modal.getOrCreateInstance(myModal)

        myModal.addEventListener('show.bs.modal', function() {
            var modalTitle = myModal.querySelector('.modal-title')
            var modalBodyInput = myModal.querySelector('.modal-body')
            modalTitle.textContent = 'Incorrect input'
            modalBodyInput.textContent = 'Maximum price allowed is S$42,949,672'
        })
        registerFailedModal.show();
        myModal.addEventListener('hide.bs.modal', function() {
            location.reload();
        })
    }

    if (selectedSortFlag)
        selectedSortFlag = parseInt(selectedSortFlag);

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
                    "item_category": selectedCategory
                },
                "location_filter": {
                    "location": selectedLocation
                },
                "price_filter": {
                    "min_price": minPrice,
                    "max_price": maxPrice,
                },
                "sort_flag": selectedSortFlag
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
                        //$('.modal-backdrop').remove();
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