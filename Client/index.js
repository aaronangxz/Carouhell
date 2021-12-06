function setLocalStorage(key, value, ttl)
{
    const item = {
		value: value,
		expiry: (Date.now()/1000) + ttl,
	}
    localStorage.setItem(key, JSON.stringify(item))
}

function getCurrentUserID()
{
    const itemStr = localStorage.getItem('userID')
    if (!itemStr) {
		return -1
	}

    const item = JSON.parse(itemStr)

	const now = (Date.now()/1000)
	// compare the expiry time of the item with the current time
	if (now > item.expiry) {
		// If the item is expired, delete the item from storage
		localStorage.removeItem(key)
		return -1
	}
	return item.value
}

function setCurrentUserName(name)
{
    localStorage.setItem('userName',name)
}

function getCurrentUserName()
{
    if(localStorage.getItem('userName') != null || localStorage.getItem('userName') != "")
        return localStorage.getItem('userName');
    else
        return null;
}

function getToken()
{
    if(localStorage.getItem('token') != null || localStorage.getItem('token') != "")
        return localStorage.getItem('token');
    else
        return -1;
}

function setPrevLocation()
{
    sessionStorage.setItem('prevLocation',window.location)
}

function getPrevLocation()
{
    if(sessionStorage.getItem('prevLocation') != null || sessionStorage.getItem('prevLocation') != "")
        return sessionStorage.getItem('prevLocation');
    else
        return -1;
}

function setPrevSecureLocation()
{
    sessionStorage.setItem('prevSecureLocation',window.location)
}

function getPrevSecureLocation()
{
    if(sessionStorage.getItem('prevSecureLocation') != null || sessionStorage.getItem('prevSecureLocation') != "")
        return sessionStorage.getItem('prevSecureLocation');
    else
        return -1;
}

/* constantly checking if need to reset Nav Bar links according to user's status(logged in?) */
function resetNavBar()
{
    if(localStorage.getItem('userID') != null) // user logged in
    {
        document.getElementById("register").style.display = "none";
        document.getElementById("login").style.display = "none";
        document.getElementById("sell").style.display = "block";
        document.getElementById("profile").style.display = "block";
        document.getElementById("favourites").style.display = "block";
        document.getElementById("wallet").style.display = "block";
        document.getElementById("logout").style.display = "block";
        document.getElementById("notification").style.display = "block";

        return true;
    }
    else 
    {
        ""
        document.getElementById("register").style.display = "block";
        document.getElementById("login").style.display = "block";
        document.getElementById("sell").style.display = "none";
        document.getElementById("profile").style.display = "none";
        document.getElementById("favourites").style.display = "none";
        document.getElementById("wallet").style.display = "none";
        document.getElementById("logout").style.display = "none";
        document.getElementById("notification").style.display = "none";

        return false;
    }
}

function registerUser()
{
     var username = document.getElementById("username").value;
     var email = document.getElementById("email").value;
     var password = document.getElementById("password").value;
     var cfmPassword = document.getElementById("cfmPassword").value;
    //  var securityQn = document.getElementById("securityQn");
    //  var selectedQn = securityQn.options[securityQn.selectedIndex].value;
    //  var securityAns = document.getElementById("securityAns").value;

     if(!username | !email | !password | !cfmPassword)
     {
        alert("Please fill in the blanks.");
     }
     else if(password !== cfmPassword)
     {
        alert("password not tally.");
        document.getElementById("password").value = "";
        document.getElementById("cfmPassword").value = "";
     }
     else
     {
        //store into user account db
        fetch('https://tic2601-t11.herokuapp.com/create_account', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'}, 
            body: JSON.stringify({
                "user_name": username,
                "user_email": email,
                "user_password": password,
                "user_security_question": 1,
                "user_security_answer": "x"
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (confirm(data.Respmeta.DebugMsg))
            {
                if(data.Respmeta.ErrorCode != 0) //failed
                {
                    if(confirm("Registration failed due to the following: " + data.Respmeta.DebugMsg))
                    {
                        window.location.href = "registerForm.html";
                    }
                }
                else
                {
                    if(confirm("Please Proceed to Login"))
                    {
                        window.location.href = "loginForm.html";
                    }
                }
            }
        })
        .catch(error => console.log(error)); 
     }  
}

function loginUser()
{
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

     if(!username | !password)
     {
        alert("Please fill in the blanks.");
     }
     else
     {
        document.getElementById("loginLoad").innerHTML =   '<div class="loader-wrapper">'+
        '<span class="loader"><span class="loader-inner"></span></span>'+
        '</div>'        //check if account exist
        fetch('https://tic2601-t11.herokuapp.com/authenticate_user', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'}, 
            body: JSON.stringify({
                "user_name": username,
                "user_password":password
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if(data.Respmeta.ErrorCode != 0)
            {
                if(confirm("Login Fail, please try again"))
                {
                    location.reload();
                }
            }
            else // successful
            {
                setLocalStorage('userID',data.Data.user_id,604800);
                // localStorage.setItem('userID',data.Data.user_id);
                localStorage.setItem('token', 'Bearer ' + data.Token.access_token);
                setCurrentUserName(username);
                console.log(username)
                //document.getElementById('navbar-text').innerHTML += username
                

                if (getPrevLocation() != null){
                    window.location.href = getPrevLocation();
                }else{
                    window.location.href = "index.html"
                }
            }
        })
        .catch(error => console.log(error));  
     }    
}

/*heck for now */
/*logic abit wrong cos need to check if username exists first */
function resetPassword()
{
     var username = document.getElementById("username").value;
     var password = document.getElementById("newPassword").value;
     var cfmPassword = document.getElementById("cfmNewPassword").value;
     var securityAns = document.getElementById("securityAns").value;

     if(!username | !password | !securityAns)
     {
        alert("Please fill in the blanks.");
     }
     else if(password !== cfmPassword)
     {
        alert("password not tally.");
        document.getElementById("newPassword").value = "";
        document.getElementById("cfmNewPassword").value = "";
     }
     else if(securityAns) // get user's security Ans to match
     {
     }
     else
     {
        //reset user's password and login
     }
   
}

function signOut()
{
    var curr = window.location
    if (curr == getPrevLocation()){
        window.location.href = getPrevLocation();
    }else{
        window.location.href = "index.html"
    }
    localStorage.clear();
    sessionStorage.clear();
}

function convertUnixToTimeStamp(unix)
{
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

    var formattedDate = day + '-' + month + '-' + year ;
    return formattedDate;
}

function convertUnixToTimeStampDetailTime(unix)
{
    let unix_timestamp = unix;
    // Create a new JavaScript Date object based on the timestamp
    // multiplied by 1000 so that the argument is in milliseconds, not seconds.
    var date = new Date(unix_timestamp * 1000);
    // Hours part from the timestamp
    var hours =  date.getHours();
    // Minutes part from the timestamp
    var minutes =  date.getMinutes();
    // Seconds part from the timestamp
    var seconds =  date.getSeconds();

    var ampm = ""

    if (parseInt(minutes) < 10){
        minutes = '0' + minutes;
    }

    if (parseInt(hours) == 0){
        hours = 12;
        ampm = "AM"
    }else if (parseInt(hours) > 12){
        hours = hours - 12;
        ampm = "PM"
    }else{
        ampm = "AM"
    }

    var formattedDate = hours + ':'+ minutes + ' ' +ampm;
    return formattedDate;
}

function getLastSeen(unix)
{
    if (Date.now()/1000 - unix < 300){
        date = "moments ago"
    }else if (Date.now()/1000  - unix < 300){
        date = "5 minutes ago"
    }else if (Date.now()/1000  - unix < 600){
        date = "10 minutes ago"
    }else if (Date.now()/1000  - unix < 6000){
        date = "an hour ago"
    }else if (Date.now()/1000  - unix < 86400){
        date = "a day ago"
    }else{
        date = 'on '+ convertUnixToTimeStamp(unix) +', '+ convertUnixToTimeStampDetailTime(unix)
    }
    return date
}
