function getUserNotifications() {
    if (getCurrentUserID() == -1) {
        window.location.href = "index.html"
    }
    fetch('https://https://' + getCurrentEnv() + '.herokuapp.com/get_user_notifications', {
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

function displayNotifications(notification) {
    console.log(notification);
    var html = '<h3>Recent Notifications</h3>';
    if (notification.length == 0) {
        html += '<p>No notifications yet.</p>'
    }
    for (var i = 0; i < notification.length; i++) {
        html += '<div class="row p-1 align-items-center h-100">';
        switch (notification[i].notification_type) {
            //like
            case 0:
                html +=
                    '<div class="col-12"><a href = viewProfile.html?profileID=' + notification[i].user_id + '>@' + notification[i].user_name + '</a> liked <a href = viewListing.html?itemID=' + notification[i].item_id + '>' + notification[i].item_name + '</a>. <font style="opacity:.6" size="2px"><i>' +
                    getTimeStamp(notification[i].ctime) + '</i></font></div>'
                break;
                //comment
            case 1:
                html +=
                    '<div class="col-12"><a href = viewProfile.html?profileID=' + notification[i].user_id + '>@' + notification[i].user_name + '</a> commented on <a href = viewListing.html?itemID=' + notification[i].item_id + '>' + notification[i].item_name + '</a>: "' + notification[i].notification_string + '". <font style="opacity:.6" size="2px"><i>' +
                    getTimeStamp(notification[i].ctime) + '</i></font></div>'
                break;
                //sold
            case 2:
                html +=
                    '<div class="col-10"><a href = viewProfile.html?profileID=' + notification[i].user_id + '>@' + notification[i].user_name + '</a> purchased <a href = viewListing.html?itemID=' + notification[i].item_id + '>' + notification[i].item_name + '</a>. <font style="opacity:.6" size="2px"><i>' +
                    getTimeStamp(notification[i].ctime) + '</i></font></div>'
                break;
                //review
            case 3:
                html +=
                    '<div class="col-10"><a href = viewProfile.html?profileID=' + notification[i].user_id + '>@' + notification[i].user_name + '</a> left a review : "' + notification[i].notification_string + '". <font style="opacity:.6" size="2px"><i>' +
                    getTimeStamp(notification[i].ctime) + '</i></font></div>'
                break;
                //owned listing sold out
            case 4:
                html +=
                    '<div class="col-10">Your listing <a href = viewListing.html?itemID=' + notification[i].item_id + '>' + notification[i].item_name + '</a> just sold out. <font style="opacity:.6" size="2px"><i>' +
                    getTimeStamp(notification[i].ctime) + '</i></font></div>'
                break;
                //liked listing sold out
            case 5:
                html +=
                    '<div class="col-10">Your favourite listing <a href = viewListing.html?itemID=' + notification[i].item_id + '>' + notification[i].item_name + '</a> just sold out. <font style="opacity:.6" size="2px"><i>' +
                    getTimeStamp(notification[i].ctime) + '</i></font></div>'
                break;
        }
        html += '</div>';
    }
    document.getElementById("notifications").innerHTML = html;
}