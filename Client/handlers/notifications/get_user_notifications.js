function getUserNotifications() {
    if (getCurrentUserID() == -1) {
        window.location.href = "index.html"
    }
    fetch('https://tic2601-t11.herokuapp.com/get_user_notifications', {
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
                alert("unable to retrieve notifications due to: " + data.Respmeta.DebugMsg);
            } else // successful
            {
                displayNotifications(data.Data.notification_list);
            }
        })
        .catch(error => console.log(error));
}