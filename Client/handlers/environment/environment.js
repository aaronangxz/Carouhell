var CurrentEnv = ''

function setCurrentEnv(env) {
    CurrentEnv = env
    console.log("CurrentEnv : " + CurrentEnv)
}

function getCurrentEnv() {
    if (document.URL.startsWith("http://localhost")) {
        CurrentEnv = "local"
    } else if (document.URL.indexOf("127.0.0.1:5500") != -1) {
        CurrentEnv = "local"
    } else if (document.URL.indexOf("test") != -1) {
        CurrentEnv = "test"
    } else {
        CurrentEnv = "live"
    }

    if (CurrentEnv == "local" || CurrentEnv == "test") {
        return "carouhell-sg-api-test"
    } else {
        return "tic2601-t11"
    }
}