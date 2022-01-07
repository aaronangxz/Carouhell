function topUpWallet() {
    var amt = document.getElementById("amt").value;
    if (Number(amt) && amt % 1 === 0) {
        amt = parseInt(amt);
    } else {
        alert("only whole number allowed");
        return;
    }
    amt *= 100;

    fetch('https://' + getCurrentEnv() + '.herokuapp.com/top_up_user_wallet', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                "user_id": parseInt(getCurrentUserID()),
                "amount": amt
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.Respmeta.ErrorCode != 0) {
                alert("unable to top up wallet due to: " + data.Respmeta.DebugMsg);
            } else // successful
            {
                alert("Top up successful!");
                location.reload();
            }
        })
        .catch(error => console.log(error));
}