function registerUser() {
    var username = document.getElementById("username").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
    var cfmPassword = document.getElementById("cfmPassword").value;
    //  var securityQn = document.getElementById("securityQn");
    //  var selectedQn = securityQn.options[securityQn.selectedIndex].value;
    //  var securityAns = document.getElementById("securityAns").value;

    if (!username | !email | !password | !cfmPassword) {
        $('#registerFieldsEmpty').modal('show')
    } else if (password !== cfmPassword) {
        $('#passwordMismatch').modal('show')
        document.getElementById("password").value = "";
        document.getElementById("cfmPassword").value = "";
    } else {
        //store into user account db
        fetch('https://' + getCurrentEnv() + '.herokuapp.com/create_account', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    "user_name": username,
                    "user_email": email,
                    "user_password": password,
                })
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.Respmeta.ErrorCode != 0) {
                    var myModal = document.getElementById('registerFailed');
                    var registerFailedModal = bootstrap.Modal.getOrCreateInstance(myModal)

                    myModal.addEventListener('show.bs.modal', function() {
                        var modalTitle = myModal.querySelector('.modal-title')
                        var modalBodyInput = myModal.querySelector('.modal-body')
                        modalTitle.textContent = 'Registration Fail: Error ' + data.Respmeta.ErrorCode
                        modalBodyInput.textContent = data.Respmeta.DebugMsg
                    })
                    registerFailedModal.show();
                } else {
                    $('#registerSuccess').modal('show')
                    $(document).on('click', '#registerSuccess', function() {
                        window.location.href = 'loginForm.html';
                    })
                }
            })
            .catch(error => console.log(error));
    }
}