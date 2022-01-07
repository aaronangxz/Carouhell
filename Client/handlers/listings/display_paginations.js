function displayPaginations(pages, current, next) {
    paginationHtml =
        '<p></p>' +
        '<nav aria-label="...">' +
        '<ul class="pagination justify-content-center">'

    //disable previous button on page 1
    if (current == 1) {
        paginationHtml +=
            '<li class="page-item disabled">' +
            '<a class="page-link" href="#" tabindex="-1" aria-disabled="true">Previous</a>' +
            '</li>'
    } else {
        paginationHtml +=
            '<li class="page-item">' +
            '<a class="page-link" href="index.html?page=' + (current - 1) + '" onclick="getLatestListingPaginated(' + (current - 1) + ')">Previous</a>' +
            '</li>'
    }

    //active
    for (let i = 1; i <= pages; i++) {
        //active page
        if (i == current) {
            paginationHtml +=
                '<li class="page-item active">' +
                '<a class="page-link" href="#">' + i + '</a>' +
                '</li>'
        } else {
            //other pages
            paginationHtml +=
                '<li class="page-item">' +
                '<a class="page-link" href="index.html?page=' + i + '" onclick="getLatestListingPaginated(' + i + ')">' + i + '</a>' +
                '</li>'
        }
    }

    if (next != 0) {
        //next button
        paginationHtml +=
            '<li class="page-item">' +
            '<a class="page-link" href="index.html?page=' + (current + 1) + '" onclick="getLatestListingPaginated(' + (current + 1) + ')">Next</a>' +
            '</li>'
    } else {
        paginationHtml +=
            '<li class="page-item disabled">' +
            '<a class="page-link" >Next</a>' +
            '</li>'
    }
    paginationHtml +=
        '</ul>' +
        '</nav>'
    document.getElementById("paginationDisplay").innerHTML = paginationHtml
}