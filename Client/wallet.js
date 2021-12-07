function getWalletDetails()
{
    if (getCurrentUserID() == -1){
        window.location.href = "index.html"
    }
    setPrevSecureLocation();
    fetch('https://tic2601-t11.herokuapp.com/get_user_wallet_details', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'}, 
        body: JSON.stringify({
            "user_id": parseInt(getCurrentUserID())
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
        {
           alert("unable to retrieve wallet due to: " + data.Respmeta.DebugMsg);
        }
        else // successful
        {
            document.getElementById("walletBalance").innerHTML = (parseInt(data.Data.wallet_info.wallet_balance)/100).toLocaleString('en-US', {
                style: 'currency',
                currency: 'USD',
              }) ;
              var lastTopUp = data.Data.wallet_info.last_top_up

              if (lastTopUp != null){
                document.getElementById("walletLastTopUp").innerHTML ='<font style="opacity:.6">Last top up on '+ convertUnixToTimeStamp(lastTopUp)+ ', ' + convertUnixToTimeStampDetailTime(lastTopUp)+'</font>'
              }

            displayWalletTransactions(data.Data.transactions);
        }
    })
    .catch(error => console.log(error)); 
}

function displayWalletTransactions(transactions)
{
    console.log(transactions);
    var html = '<h3>Recent Transactions</h3>';
    for(var i = 0 ; i<transactions.length; i++)
    {
        html += '<div class="row p-3 align-items-center h-100">';
        switch (transactions[i].transaction_type)
        {
            //top up
            case 0: 
            html +='<div class="col-2.5"><font size="3px">'+convertUnixToTimeStamp(transactions[i].transaction_ctime)+' '+convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime)+'</font></div>'+
            '<div class="col-2 text-center"><span><i class="fas fa-arrow-right fa-2x"></i></span> <span><i class="fas fa-wallet fa-2x"></i></span></div>'+
                '<div class="col-6">Top up Wallet </div>'+
                '<div class="col-1.5 text-left"><h6> + ' +(parseInt(transactions[i].transaction_amount)/100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                  })+'</h6></div>';
                break;
            //purchase
            case 1:
                html +='<div class="col-2.5"><font size="3px">'+convertUnixToTimeStamp(transactions[i].transaction_ctime)+' ' +convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime)+'</font></div>'+
                '<div class="col-2 text-center" id="imgContainer">'+
                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+transactions[i].item_id+'.jpg">'+
                    // '<img src="'+transactions[i].item_image+'" />'+
                '</div>'+
                '<div class="col-6"><a href="viewListing.html?itemID='+transactions[i].item_id+'">PURCHASED: '+transactions[i].item_name+'</a></div>'+
                '<div class="col-1.5 text-left"><h6> - ' + (parseInt(transactions[i].transaction_amount)/100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                  }) +'</h6></div>';
                break;
            //sales
            case 2:
                html +='<div class="col-2.5"><font size="3px">'+convertUnixToTimeStamp(transactions[i].transaction_ctime)+' '+convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime)+'</font></div>'+
                '<div class="col-2 text-center" id="imgContainer">'+
                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+transactions[i].item_id+'.jpg">'+
                '</div>'+
                '<div class="col-6"><a href="viewListing.html?itemID='+transactions[i].item_id+'">SOLD: '+transactions[i].item_name+'</a></div>'+
                '<div class="col-1.5 text-left"><h6> + ' + (parseInt(transactions[i].transaction_amount)/100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                  }) +'</h6></div>';
                break;
        }
        html += '</div>';
    }
    document.getElementById("walletTransactions").innerHTML = html;
}

function topUpWallet()
{
    var amt = document.getElementById("amt").value;
    if(Number(amt) && amt % 1 === 0)
    {
        amt = parseInt(amt);
    }
    else
    {
        alert("only whole number allowed");
        return;
    }
    amt*= 100;

    fetch('https://tic2601-t11.herokuapp.com/top_up_user_wallet', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'}, 
        body: JSON.stringify({
            "user_id": parseInt(getCurrentUserID()),
            "amount":amt
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
        {
           alert("unable to top up wallet due to: " + data.Respmeta.DebugMsg);
        }
        else // successful
        {
            alert("Top up successful!");
            location.reload();
        }
    })
    .catch(error => console.log(error)); 
}