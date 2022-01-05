var CurrentEnv = ''

function setCurrentEnv(env) {
    CurrentEnv = env
    console.log("CurrentEnv : " + CurrentEnv)
}

function getCurrentEnv() {
    if (CurrentEnv == "local" || "test") {
        return "carouhell-sg-api-test"
    } else {
        return "tic2601-t11"
    }
}