function displayWalletTransactions(transactions) {
    console.log(transactions);
    var html = '<h3>Recent Transactions</h3>';
    for (var i = 0; i < transactions.length; i++) {
        html += '<div class="row">';
        switch (transactions[i].transaction_type) {
            //top up
            case 0:
                html += '<div class="col-2.5"><font size="3px">' + convertUnixToTimeStamp(transactions[i].transaction_ctime) + ' ' + convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime) + '</font></div>' +
                    '<div class="col-2 text-center"><span><i class="fas fa-arrow-right fa-2x"></i></span> <span><i class="fas fa-wallet fa-2x"></i></span></div>' +
                    '<div class="col-5">Top up Wallet </div>' +
                    '<div class="col-1.5 text-left" id = "walletTrxAmount"><h6> + ' + (parseInt(transactions[i].transaction_amount) / 100).toLocaleString('en-US', {
                        style: 'currency',
                        currency: 'USD',
                    }) + '</h6></div>';
                break;
                //purchase
            case 1:
                html += '<div class="col-2.5"><font size="3px">' + convertUnixToTimeStamp(transactions[i].transaction_ctime) + ' ' + convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime) + '</font></div>' +
                    '<div class="col-2 text-center" id="walletImgContainer">' +
                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_' + transactions[i].item_id + '.jpg">' +
                    // '<img src="'+transactions[i].item_image+'" />'+
                    '</div>' +
                    '<div class="col-6"><a href="viewListing.html?itemID=' + transactions[i].item_id + '">PURCHASED: ' + transactions[i].item_name + '</a></div>' +
                    '<div class="col-1.5 text-right" id = "walletTrxAmount"><h6> - ' + (parseInt(transactions[i].transaction_amount) / 100).toLocaleString('en-US', {
                        style: 'currency',
                        currency: 'USD',
                    }) + '</h6></div>';
                break;
                //sales
            case 2:
                html += '<div class="col-2.5"><font size="3px">' + convertUnixToTimeStamp(transactions[i].transaction_ctime) + ' ' + convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime) + '</font></div>' +
                    '<div class="col-2 text-center" id="walletImgContainer">' +
                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_' + transactions[i].item_id + '.jpg">' +
                    '</div>' +
                    '<div class="col-6"><a href="viewListing.html?itemID=' + transactions[i].item_id + '">SOLD: ' + transactions[i].item_name + '</a></div>' +
                    '<div class="col-1.5 text-left" id = "walletTrxAmount"><h6> + ' + (parseInt(transactions[i].transaction_amount) / 100).toLocaleString('en-US', {
                        style: 'currency',
                        currency: 'USD',
                    }) + '</h6></div>';
                break;
        }
        html += '</div>';
    }
    document.getElementById("walletTransactions").innerHTML = html;
}