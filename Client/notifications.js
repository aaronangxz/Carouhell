function getUserNotifications()
{
    fetch('https://tic2601-t11.herokuapp.com/get_user_notifications', {
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
           alert("unable to retrieve notifications due to: " + data.Respmeta.DebugMsg);
        }
        else // successful
        {
              displayNotifications(data.Data.notification_list);
        }
    })
    .catch(error => console.log(error)); 
}

function displayNotifications(notification)
{
    console.log(notification);
    var html = '<h3>Recent Notifications</h3>';
    for(var i = 0 ; i<notification.length; i++)
    {
        html += '<div class="row p-1 align-items-center h-100">';
        switch (notification[i].notification_type)
        {
            //like
            case 0: 
            html +='<div class="col-2">'+convertUnixToTimeStamp(notification[i].ctime)+ ' ' + convertUnixToTimeStampDetailTime(notification[i].ctime)+'</div>'+
                '<div class="col-10"><b>'+ notification[i].user_name +'</b> liked your listing <b>' + notification[i].item_name + '</b>.</div>'
                break;
            //comment
            case 1:
                html +='<div class="col-2">'+convertUnixToTimeStamp(notification[i].ctime)+ ' ' + convertUnixToTimeStampDetailTime(notification[i].ctime)+'</div>'+
                '<div class="col-10"><b>'+ notification[i].user_name +'</b> commented on your listing <b>' + notification[i].item_name + '</b> : "' + notification[i].notification_string +'".</div>'
                break;
            //sold
            case 2:
                html +='<div class="col-2">'+convertUnixToTimeStamp(notification[i].ctime)+ ' ' + convertUnixToTimeStampDetailTime(notification[i].ctime)+'</div>'+
                '<div class="col-10"><b>'+ notification[i].user_name +'</b> purchased your listing <b>' + notification[i].item_name + '</b>.</div>'
                break;
            //review
            case 3:
                html +='<div class="col-2">'+convertUnixToTimeStamp(notification[i].ctime)+ ' ' + convertUnixToTimeStampDetailTime(notification[i].ctime)+'</div>'+
                '<div class="col-10"><b>'+ notification[i].user_name +'</b> left a review : "' + notification[i].notification_string +'".</div>'
                break;
        }
        html += '</div>';
    }
    document.getElementById("notifications").innerHTML = html;
}