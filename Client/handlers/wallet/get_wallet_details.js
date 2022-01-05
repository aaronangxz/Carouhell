function getWalletDetails() {
    if (getCurrentUserID() == -1) {
        window.location.href = "index.html"
    }
    setPrevSecureLocation();
    fetch('https://' + getCurrentEnv() + '.herokuapp.com/get_user_wallet_details', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                "user_id": parseInt(getCurrentUserID())
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                alert("unable to retrieve wallet due to: " + data.Respmeta.DebugMsg);
            } else // successful
            {
                document.getElementById("walletBalance").innerHTML = (parseInt(data.Data.wallet_info.wallet_balance) / 100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                });
                var lastTopUp = data.Data.wallet_info.last_top_up

                if (lastTopUp != null) {
                    document.getElementById("walletLastTopUp").innerHTML = '<font style="opacity:.6">Last top up on ' + convertUnixToTimeStamp(lastTopUp) + ', ' + convertUnixToTimeStampDetailTime(lastTopUp) + '</font>'
                }

                displayWalletTransactions(data.Data.transactions);
            }
        })
        .catch(error => console.log(error));
}