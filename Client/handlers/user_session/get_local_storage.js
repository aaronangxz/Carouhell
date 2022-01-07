function getCurrentUserID() {
    const itemStr = localStorage.getItem('userID')
    if (!itemStr) {
        return -1
    }

    const item = JSON.parse(itemStr)

    const now = (Date.now() / 1000)
        // compare the expiry time of the item with the current time
    if (now > item.expiry) {
        // If the item is expired, delete the item from storage
        localStorage.removeItem(key)
        return -1
    }
    return item.value
}

function getCurrentUserName() {
    if (localStorage.getItem('userName') != null || localStorage.getItem('userName') != "")
        return localStorage.getItem('userName');
    else
        return null;
}

function getToken() {
    if (localStorage.getItem('token') != null || localStorage.getItem('token') != "")
        return localStorage.getItem('token');
    else
        return -1;
}

function getPrevLocation() {
    if (sessionStorage.getItem('prevLocation') != null || sessionStorage.getItem('prevLocation') != "")
        return sessionStorage.getItem('prevLocation');
    else
        return -1;
}

function getPrevSecureLocation() {
    if (sessionStorage.getItem('prevSecureLocation') != null || sessionStorage.getItem('prevSecureLocation') != "")
        return sessionStorage.getItem('prevSecureLocation');
    else
        return -1;
}