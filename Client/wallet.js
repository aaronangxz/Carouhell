function getWalletDetails()
{
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
            document.getElementById("walletBalance").innerHTML = "SGD $" + (parseInt(data.Data.wallet_info.wallet_balance)/100).toLocaleString('en-US', {
                style: 'currency',
                currency: 'USD',
              }) ;
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
            html +='<div class="col-2">'+convertUnixToTimeStamp(transactions[i].transaction_ctime)+'<br>'+convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime)+'</div>'+
            '<div class="col-3 text-center"><span><i class="fas fa-arrow-right fa-2x"></i></span> <span><i class="fas fa-wallet fa-2x"></i></span></div>'+
                '<div class="col-5">Top up Wallet </div>'+
                '<div class="col-2 text-center"><h4> + ' +(parseInt(transactions[i].transaction_amount)/100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                  })+'</h4></div>';
                break;
            //purchase
            case 1:
                html +='<div class="col-2">'+convertUnixToTimeStamp(transactions[i].transaction_ctime)+'<br>'+convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime)+'</div>'+
                '<div class="col-3 text-center" id="imgContainer">'+
                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+transactions[i].item_id+'.jpg">'+
                    // '<img src="'+transactions[i].item_image+'" />'+
                '</div>'+
                '<div class="col-5"><a href="viewListing.html?itemID='+transactions[i].item_id+'">BOUGHT: '+transactions[i].item_name+'</a></div>'+
                '<div class="col-2 text-center"><h4> - ' + (parseInt(transactions[i].transaction_amount)/100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                  }) +'</h4></div>';
                break;
            //sales
            case 2:
                html +='<div class="col-2">'+convertUnixToTimeStamp(transactions[i].transaction_ctime)+'<br>'+convertUnixToTimeStampDetailTime(transactions[i].transaction_ctime)+'</div>'+
                '<div class="col-3 text-center" id="imgContainer">'+
                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+transactions[i].item_id+'.jpg">'+
                '</div>'+
                '<div class="col-5"><a href="viewListing.html?itemID='+transactions[i].item_id+'">SOLD: '+transactions[i].item_name+'</a></div>'+
                '<div class="col-2 text-center"><h4> + ' + (parseInt(transactions[i].transaction_amount)/100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                  }) +'</h4></div>';
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