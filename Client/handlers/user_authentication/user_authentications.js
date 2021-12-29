function loginUser() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    if (!username | !password) {
        $('#loginCredentials').modal('show')
    } else {
        document.getElementById("footer").innerHTML += '<div class="loader-wrapper">' +
            '<span class="loader"><span class="loader-inner"></span></span>' +
            '</div>'
        fetch('https://tic2601-t11.herokuapp.com/authenticate_user', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    "user_name": username,
                    "user_password": password
                })
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.Respmeta.ErrorCode != 0) {
                    var myModal = document.getElementById('loginFail');
                    var registerFailedModal = bootstrap.Modal.getOrCreateInstance(myModal)

                    myModal.addEventListener('show.bs.modal', function() {
                        var modalTitle = myModal.querySelector('.modal-title')
                        var modalBodyInput = myModal.querySelector('.modal-body')
                        modalTitle.textContent = 'Login Fail: Error ' + data.Respmeta.ErrorCode
                        modalBodyInput.textContent = data.Respmeta.DebugMsg
                    })
                    registerFailedModal.show();
                    myModal.addEventListener('hide.bs.modal', function() {
                        location.reload();
                    })
                } else // successful
                {
                    setLocalStorage('userID', data.Data.user_id, 604800);
                    localStorage.setItem('token', 'Bearer ' + data.Token.access_token);
                    setCurrentUserName(username);
                    if (getPrevLocation() != null) {
                        window.location.href = getPrevLocation();
                    } else {
                        window.location.href = "index.html"
                    }
                }
            })
            .catch(error => console.log(error));
    }
}

function signOut() {
    var curr = window.location
    if (curr == getPrevLocation()) {
        window.location.href = getPrevLocation();
    } else {
        window.location.href = "index.html"
    }
    localStorage.clear();
    sessionStorage.clear();
}