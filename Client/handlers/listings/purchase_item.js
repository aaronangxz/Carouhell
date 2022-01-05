function buyNow(itemID, sellerID) {
    setPrevSecureLocation();
    console.log("buy now ", itemID);
    var qtyToPurchase = document.getElementById("qtyToPurchase").value;
    console.log("qtyToPurchase: " + qtyToPurchase);
    if (qtyToPurchase != "" || !qtyToPurchase) {
        if (Number(qtyToPurchase) && Number(qtyToPurchase) != 0) {
            fetch('https://' + getCurrentEnv() + '.herokuapp.com/purchase_single_item', {
                    method: 'POST',
                    headers: {
                        'Authorization': getToken(),
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        "item_id": parseInt(itemID),
                        "user_id": parseInt(getCurrentUserID()),
                        "purchase_quantity": parseInt(qtyToPurchase)
                    })
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                    if (data.Respmeta.ErrorCode != 0) {
                        if (confirm("Failed Purchasing due to " + data.Respmeta.DebugMsg)) {
                            location.reload();
                        }
                    } else // successful
                    {
                        if (confirm("Successfully Purchased!")) {
                            if (data.Data.is_reviewed) {
                                window.location.href = "wallet.html";
                            } else {
                                window.location.href = "review.html?sellerID=" + sellerID;
                            }
                        }
                    }
                })
                .catch(error => console.log(error));
        }
    } else {
        alert("Please enter quantity");
    }

}