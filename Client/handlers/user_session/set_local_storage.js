function setLocalStorage(key, value, ttl) {
    const item = {
        value: value,
        expiry: (Date.now() / 1000) + ttl,
    }
    localStorage.setItem(key, JSON.stringify(item))
}

function setCurrentUserName(name) {
    localStorage.setItem('userName', name)
}

function setPrevLocation() {
    sessionStorage.setItem('prevLocation', window.location)
}

function setPrevSecureLocation() {
    sessionStorage.setItem('prevSecureLocation', window.location)
}