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
                    '<div class="col-12"><b>' + notification[i].user_name + '</b> liked <b>' + notification[i].item_name + '</b>. <font style="opacity:.6" size="2px"><i>' +
                    Math.ceil((Date.now() / 1000 - notification[i].ctime) / 86400) + ' days ago</i></font></div>'
                break;
                //comment
            case 1:
                html +=
                    '<div class="col-12"><b>' + notification[i].user_name + '</b> commented on <b>' + notification[i].item_name + '</b> : "' + notification[i].notification_string + '". <font style="opacity:.6" size="2px"><i>' +
                    Math.ceil((Date.now() / 1000 - notification[i].ctime) / 86400) + ' days ago</i></font></div>'
                break;
                //sold
            case 2:
                html +=
                    '<div class="col-10"><b>' + notification[i].user_name + '</b> purchased <b>' + notification[i].item_name + '</b>. <font style="opacity:.6" size="2px"><i>' +
                    Math.ceil((Date.now() / 1000 - notification[i].ctime) / 86400) + ' days ago</i></font></div>'
                break;
                //review
            case 3:
                html +=
                    '<div class="col-10"><b>' + notification[i].user_name + '</b> left a review : "' + notification[i].notification_string + '". <font style="opacity:.6" size="2px"><i>' +
                    Math.ceil((Date.now() / 1000 - notification[i].ctime) / 86400) + ' days ago</i></font></div>'
                break;
        }
        html += '</div>';
    }
    document.getElementById("notifications").innerHTML = html;
}