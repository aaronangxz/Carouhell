function convertUnixToTimeStamp(unix) {
    let unix_timestamp = unix;
    // Create a new JavaScript Date object based on the timestamp
    // multiplied by 1000 so that the argument is in milliseconds, not seconds.
    var date = new Date(unix_timestamp * 1000);
    // Hours part from the timestamp
    var hours = date.getHours();
    // Minutes part from the timestamp
    var minutes = "0" + date.getMinutes();
    // Seconds part from the timestamp
    var seconds = "0" + date.getSeconds();

    var day = date.getDate();
    var month = date.getMonth() + 1;
    var year = date.getFullYear();

    var formattedDate = day + '-' + month + '-' + year;
    return formattedDate;
}

function convertUnixToTimeStampDetailTime(unix) {
    let unix_timestamp = unix;
    // Create a new JavaScript Date object based on the timestamp
    // multiplied by 1000 so that the argument is in milliseconds, not seconds.
    var date = new Date(unix_timestamp * 1000);
    // Hours part from the timestamp
    var hours = date.getHours();
    // Minutes part from the timestamp
    var minutes = date.getMinutes();
    // Seconds part from the timestamp
    var seconds = date.getSeconds();

    var ampm = ""

    if (parseInt(minutes) < 10) {
        minutes = '0' + minutes;
    }

    if (parseInt(hours) == 0) {
        hours = 12;
        ampm = "AM"
    } else if (parseInt(hours) > 12) {
        hours = hours - 12;
        ampm = "PM"
    } else {
        ampm = "AM"
    }

    var formattedDate = hours + ':' + minutes + ' ' + ampm;
    return formattedDate;
}