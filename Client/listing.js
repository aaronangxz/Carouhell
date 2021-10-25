function encodeImageFileAsURL(inputFileToLoad) {

    var filesSelected = document.getElementById("inputFileToLoad").files[0];
    if (filesSelected.length > 0) {
      var fileToLoad = filesSelected[0];

      var fileReader = new FileReader();

      fileReader.onload = function(fileLoadedEvent) {
        var srcData = fileLoadedEvent.target.result; // <--- data: base64
        console.log("srcData: " + srcData);
        var newImage = document.createElement('img');
        newImage.src = srcData;

        document.getElementById("imgTest").innerHTML = newImage.outerHTML;
        alert("Converted Base64 version is " + document.getElementById("imgTest").innerHTML);
        console.log("Converted Base64 version is " + document.getElementById("imgTest").innerHTML);
      }
      fileReader.readAsDataURL(fileToLoad);
    }
}

function createListing(userID)
{
    var itemName = document.getElementById('itemName').value;
    var itemPrice = document.getElementById("itemPrice").value;
    var itemQty = document.getElementById("itemQty").value;
    var itemDesc = document.getElementById("itemDesc").value;
    var itemCat = document.getElementById("itemCat");
    var itemCatValue = itemCat.options[itemCat.selectedIndex].value;
    // var itemImg = document.getElementById("img").value;
    var collection = document.getElementById("collection");
    var collectionValue = collection.options[collection.selectedIndex].value;
    var location = document.getElementById("location");
    var locationValue = location.options[location.selectedIndex].value;
    


    var base64String = "";
    const file = document.getElementById("img").files[0];
    var reader = new FileReader();
      
    reader.onload = function () {
        base64String = reader.result.replace("data:", "")
            .replace(/^.+,/, "");
  
        imageBase64Stringsep = base64String;
    }
    reader.readAsDataURL(file);
    reader.addEventListener('load', (event) => {
        console.log('load finish: ' + base64String ); // base 64

        fetch('https://tic2601-t11.herokuapp.com/create_listing', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'}, 
            body: JSON.stringify({
                "item_name": itemName,
                "item_price": parseInt(itemPrice),
                "item_quantity":parseInt(itemQty),
                "item_description":itemDesc,
                "item_shipping_info": 0,
                "item_payment_info": 0,
                "item_location": locationValue,
                "item_category": parseInt(itemCatValue),
                "item_image": base64String,
                "seller_id": parseInt(userID)
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if(data.Respmeta.ErrorCode != 0)
            {
                if(confirm("Failed creating listing"))
                {
                    location.reload();
                }
            }
            else // successful
            {
                if(confirm("Listing Created!"))
                {
                    //window.location.href = "viewListing.html?itemID=" + data.item_id;
                }
            }
        })
        .catch(error => console.log(error)); 

    });
    

    
}   

function deleteListing(itemID)
{
    fetch('https://tic2601-t11.herokuapp.com/delete_single_listing', {
        method: 'DELETE',
        // headers: {'Content-Type': 'application/json'}, 
        body: JSON.stringify({
            "item_id": itemID
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
            {
                if(confirm("Failed deleting listing"))
                {
                    location.reload();
                }
            }
            else // successful
            {
                if(confirm("Listing Deleted!"))
                {
                    window.location.href = "profile.html";
                }
            }
    })
    .catch(error => console.log(error)); 
}

function addListingComment(userID, itemID)
{
    var comment = document.getElementById('comment').value;
    console.log("comment:" + comment);
    fetch('https://tic2601-t11.herokuapp.com/add_listing_comments', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'}, 
        body: JSON.stringify({
            "item_id": 65,
            "user_id":55,
            "comment":comment
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
            {
                if(confirm("Comment failed"))
                {
                   // window.location.reload(false); 
                }
            }
            else // successful
            {
                if(confirm("Successfully commented"))
                {
                   // window.location.reload(false); 
                }
            }
    })
    .catch(error => console.log(error)); 

}

function viewListingByItemId(itemID)
{
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var requestionOptions = {
        method: 'POST',
        headers: myHeaders,
        body: JSON.stringify({
            "item_id": parseInt(itemID)
        }),
        redirect: 'follow'
    };

    fetch("https://tic2601-t11.herokuapp.com/get_single_listing_by_itemid", requestionOptions)
    .then(response => response.json())
    // .then(result => console.log(result))
    .then(data => renderIndividualListing(data.Data))
    .catch(error => console.log('error', error)); 
}

function renderIndividualListing (data)
{
    console.log(data);
    var content = "";
    content += '<div class="row">' +
        '<div class="col-4"><img src="'+data.item_image+'" class="img-fluid"></div>'+
        '<div class="col-8">'+
            '<div class="row"><h1>'+data.item_name+' - $'+data.item_price+' ['+itemStatus_Arr[data.item_status]+'] </h1></div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-clock"></i></span> Posted on '+convertUnixToTimeStamp(data.listing_ctime)+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-heart"></i></span> '+data.listing_likes+'</div>'+
            '</div>'+
            '<div class="row mt-3">' +
                '<div class="col">Category: '+categories_Arr[data.item_category]+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-map-marker-alt"></i></span> '+data.item_location+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-user"></i> </span>'+data.seller_name+'</div>'+
            '</div>'+
            '<div class="row mt-3">' +
                '<div class="col"><h3>Description <h3></div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col">'+data.item_description+'</div>'+
            '</div>'+
        '</div>'+
    '</div>'+
    '<div class="row mt-3">' +
        '<div class="col"><h3>Comments <h3></div>'+
    '</div>'+
    '<form id="listingComment" method="post">'+
        '<div class="row mt-3">' +
            '<div class="col-10">'+
                '<div class="form-group">'+
                    '<input type="text" class="form-control" id="comment" placeholder="Add a comment...">'+
                '</div>'+
            '</div>'+
            '<div class="col-2">'+
                '<input type="button" class="btn btn-primary" onclick="addListingComment('+sessionStorage.getItem('userID')+','+data.item_id+')" value="Submit"/>'+
            '</div>'+
        '</div>'+
    '</form>'+
    '<div class="row mt-3" id="comments">' +
        '<div class="col">'+
        '</div>'+
    '</div>'+
    
    '<div class="row mt-3" id="buyerBtns">'+
        '<div class="col">';
        if(data.user_id != sessionStorage.getItem('userID') && itemStatus_Arr[data.item_status] == "Available")
        {
            content += '<a href="javascript:void(0);" onclick="buyNow('+data+')" class="btn btn-primary">Buy Now</a>';
        }
    content +='</div>'+
    '</div>';
    if(data.user_id == sessionStorage.getItem('userID')) // means it's own listings
    {
        content += '<div class="row" id="sellerBtns">'+
                        '<div class="col">'+
                            '<a href="editListing.html?itemID='+data.item_id+'" id="editBtn" class="btn btn-primary">Edit Listing</a> '+
                        '</div>'+
                        '<div class="col">'+
                            '<a href="javascript:void(0);" onclick="deleteListing('+data.item_id+')" id="deleteBtn" class="btn btn-primary">Delete Listing</a>'+
                        '</div>'+
                    '</div>';
    }
    

    document.getElementById("listingContent").innerHTML = content;
}

function addListingLikes(itemID)
{
    console.log("is user logged in? -" + isLoggedIn());
    if(!isLoggedIn) //not logged in
    {
        window.location.href = "index.html";
    }
    else // is logged in
    {
        var userID = sessionStorage.getItem('status');

        fetch('https://tic2601-t11.herokuapp.com/add_listing_likes', {
            method: 'POST',
            // headers: {'Content-Type': 'application/json'}, 
            body: JSON.stringify({
                "item_id": itemID,
                "user_id": userID
            })
        })
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.log(error)); 
    }
}

function getAllListing() {
    fetch('https://tic2601-t11.herokuapp.com/get_all_listings', {
      method: 'GET'
    })
    .then(response => response.json())
    .then(result => {/*result.Data*/  for(const d of result.Data){
          console.log(d);  
          document.getElementById("cards").innerHTML += 
          '<div class="col-md-4 mt-5">'+
            '<div class="card" id="'+d.item_id+'">'+
              '<div class="card-header">'+
                d.seller_name+
              '</div>'+
              '<div class="card-body pb-5">'+
                '<div class="container">'+
                    '<div class="row">'+
                        '<div class="col text-center" id="imgContainer">'+
                            '<img src="'+d.item_image+'" />'+
                        '</div>'+
                    '</div>'+
                    '<div class="row">'+
                        '<div class="col">'+
                            '<h5 class="card-title">'+ d.item_name +'</h5>'+
                        '</div>'+
                    '</div>'+ 
                    '<div class="row">'+
                        '<div class="col">'+
                            '<h5 class="card-title">$'+ d.item_price +'</h5>'+
                        '</div>'+                       
                    '</div>'+
                    '<div class="row">'+
                        '<div class="col">'+
                            '<h5 class="card-title">'+
                            '<a href="javascript:void(0);" onclick="addListingLikes('+d.item_id+');"><span><i class="fas fa-heart"></i></span></a> ' + d.listing_likes + 
                            '</h5>'+
                        '</div>'+
                    '</div>'+
                    '<div class="row">'+
                        '<div class="col">'+
                            '<h5 class="card-title">'+ itemStatus_Arr[d.item_status] +'</h5>'+
                        '</div>'+
                    '</div>'+
                    '<div class="row mt-3">'+
                        '<div class="col text-center">'+
                            '<a href="ViewListing.html?itemID='+d.item_id+'" class="btn btn-primary">View Listing</a> '+
                        '</div>'+
                    '</div>'+ 
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'; 
      }
    })
    .catch(error => {console.log('NO:', JSON.stringify(error));});
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