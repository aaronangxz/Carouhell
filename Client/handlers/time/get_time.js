function getLastSeen(loginDate) {
    var todayDate = (Date.now() / 1000)
    var dateDifference = todayDate - loginDate

    if (dateDifference > 86400) {
        date = 'Last seen on ' + convertUnixToTimeStamp(loginDate) + ', ' + convertUnixToTimeStampDetailTime(loginDate)
    } else {
        date = 'Last seen ' + getTimeStamp(loginDate)
    }
    return date
}

function getTimeStamp(listingDate) {
    var postedDate = listingDate
    var todayDate = (Date.now() / 1000)
    var dateDifference = todayDate - postedDate

    if (dateDifference > 5259486) {
        date = Math.floor((todayDate - postedDate) / 2629743) + ' months ago'
    } else if (dateDifference > 2629743) {
        date = "1 month ago"
    } else if (dateDifference > 1209600) {
        date = Math.floor((todayDate - postedDate) / 604800) + ' weeks ago'
    } else if (dateDifference > 604800) {
        date = "1 week ago"
    } else if (dateDifference > 172800) {
        date = Math.floor((todayDate - postedDate) / 86400) + ' days ago'
    } else if (dateDifference > 86400) {
        date = "1 day ago"
    } else if (dateDifference > 7200) {
        date = Math.floor((todayDate - postedDate) / 3600) + ' hours ago'
    } else if (dateDifference > 3600) {
        date = "an hour ago"
    } else if (dateDifference > 3000) {
        date = "50 minutes ago"
    } else if (dateDifference > 2400) {
        date = "40 minutes ago"
    } else if (dateDifference > 1800) {
        date = "30 minutes ago"
    } else if (dateDifference > 1200) {
        date = "20 minutes ago"
    } else if (dateDifference > 600) {
        date = "10 minutes ago"
    } else if (dateDifference > 300) {
        date = "5 minutes ago"
    } else if (dateDifference < 300) {
        date = "moments ago"
    }
    return date
}