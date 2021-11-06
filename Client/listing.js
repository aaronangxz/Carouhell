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
    // var collection = document.getElementById("collection");
    // var collectionValue = collection.options[collection.selectedIndex].value;
    var location = document.getElementById("itemLocation");
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
            headers:{
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                "item_name": itemName,
                "item_price": parseInt(itemPrice),
                "item_quantity":parseInt(itemQty),
                "item_description":itemDesc,
                "item_location": parseInt(locationValue),
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
                    window.history.back();
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
        headers:{
            'Authorization': getToken(),
            'Content-Type': 'application/json'
        }, 
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
                    window.location.href = "viewProfile.html";
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
        headers:{
            'Authorization': getToken(),
            'Content-Type': 'application/json'
        }, 
        body: JSON.stringify({
            "item_id": parseInt(itemID),
            "user_id":parseInt(userID),
            "comment":comment
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
            {
                if(confirm("Comment failed due to the following reason: " + data.Respmeta.DebugMsg))
                {
                    location.reload();
                }
            }
            else // successful
            {
                if(confirm("Successfully commented"))
                {
                   location.reload();
                }
            }
    })
    .catch(error => console.log(error)); 

}

function viewListingByItemId(itemID)
{
    fetch('https://tic2601-t11.herokuapp.com/get_single_listing_by_itemid', {
            method: 'POST',
            headers:{
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                "item_id": parseInt(itemID)
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            displayItemContent(data.Data);
            displayItemComments(data.Data.listing_comments, data.Data.item_id);
        })
        .catch(error => console.log(error)); 
}


function displayItemContent(data)
{
    var content = "";
    content += '<div class="row">' +
        '<div class="col-4"><img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+data.item_id+'.jpg" class="img-fluid" /></div>'+
        '<div class="col-8">'+
            '<div class="row"><h1>'+data.item_name+' - $'+parseInt(data.item_price)/100+' ['+itemStatus_Arr[data.item_status]+'] </h1></div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-clock"></i></span> Posted on '+convertUnixToTimeStamp(data.listing_ctime)+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><a href="javascript:void(0);" onclick="addListingLikes('+data.item_id+');"></a> ' + data.listing_likes+'</div>'+
            '</div>'+
            '<div class="row mt-3">' +
                '<div class="col">Category: '+categories_Arr[data.item_category]+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-map-marker-alt"></i></span> '+location_Arr[data.item_location]+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-user"></i> </span><a href="viewProfile.html?profileID='+data.seller_id+'">@'+data.seller_name+'</a></div>'+
            '</div>'+
            '<div class="row mt-3">' +
                '<div class="col"><h3>Description <h3></div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col">'+data.item_description+'</div>'+
            '</div>'+
        '</div>'+
    '</div>';
    if(data.seller_id == getCurrentUserID()) // user is listing's owner
    {
        // display seller's option
        content += 
        '<div class="row mt-3" id="sellerBtns">'+
            '<div class="col">'+
                // '<input type="button" id="btnEdit" class="btn btn-primary" value="Edit Listing"/>'+
                 '<a href="editListing.html?itemID='+data.item_id+'" id="btnEdit" class="btn btn-primary">Edit Listing</a> '+
            '</div>'+
            '<div class="col">'+
                '<a href="javascript:void(0);" onclick="deleteListing('+data.item_id+')" id="deleteBtn" class="btn btn-primary">Delete Listing</a>'+
            '</div>'+
        '</div>';
    }
    else // is a buyer
    {
        if(data.item_quantity == 0)
        {
            content += 
            '<form id="purchaseForm" method="post">'+
                '<div class="row mt-3" id="buyerBtns">'+
                    '<div class="col-4"></div>'+
                    '<div class="col-3 text-center">'+
                        '<h5> Quantity Available: '+ data.item_quantity + '</h5>' +
                    '</div>'+
                    '<div class="col-1">'+
                        '<input type="button" class="btn btn-secondary disabled" value="Sold Out"/>'+
                        // '<input type="button" id="btnBuyNow" class="btn btn-primary" value="Buy Now"/>'+
                    '</div>'+
                    '<div class="col-4"></div>'+
                '</div>'+
            '</form>';
        }
        else
        {
            content += 
            '<form id="purchaseForm" method="post">'+
                '<div class="row mt-3" id="buyerBtns">'+
                    '<div class="col-4"></div>'+
                    '<div class="col-3 text-center">'+
                        '<div class="form-group">'+
                            '<input type="number" class="form-control" id="qtyToPurchase" placeholder="Enter Quantity">'+
                        '</div>'+
                        '<h5>Quantity Available: '+ data.item_quantity + '</h5>' +
                    '</div>'+
                    '<div class="col-1">'+
                        '<input type="button" onclick="buyNow('+data.item_id+','+data.seller_id+')" class="btn btn-primary" value="Buy Now"/>'+
                        // '<input type="button" id="btnBuyNow" class="btn btn-primary" value="Buy Now"/>'+
                    '</div>'+
                    '<div class="col-4"></div>'+
                '</div>'+
            '</form>';
        }
    }

    document.getElementById("listingContent").innerHTML = content;

    const aTag = document.getElementById('btnEdit');
    if(aTag){
        console.log("aTag")
        aTag.addEventListener('click', function(){sessionStorage.setItem(data.item_id, JSON.stringify(data));});
    }

    // const button = document.getElementById('btnBuyNow');
    // if(button)
    // {
    //     button.addEventListener('click', function(){ buyNow(data); });
    // }
    
    // const editButton = document.getElementById('btnEdit');
    // if(editButton)
    // {
    //     editButton.addEventListener('click', function(){ editListing(data); });
    // }
    
}

function displayItemComments(data, itemID)
{
    var comments =
    '<div class="row mt-3">' +
        '<div class="col"><h3>Comments <h3></div>'+
    '</div>';
    for(var i = 0; i < data.length; i++)
    {
        console.log("comments: " + data[i]);
        comments +=
        '<div class="row text-wrap" >' +
            '<div class="col-3"><a href="viewProfile.html?profileID='+data[i].user_id+'">'+data[i].user_name +'</a> ∙ '+convertUnixToTimeStamp(data[i].ctime)+'</div>'+
            '<div class="col-9">'+data[i].comment+'</div>'+
        '</div>';
    }
    comments +=
    '<form id="listingComment" method="post">'+
        '<div class="row mt-3">' +
            '<div class="col-10">'+
                '<div class="form-group">'+
                    '<input type="text" class="form-control" id="comment" placeholder="Add a comment...">'+
                '</div>'+
            '</div>'+
            '<div class="col-2">'+
                '<input type="button" class="btn btn-primary" onclick="addListingComment(getCurrentUserID(),'+itemID+')" value="Submit"/>'+
            '</div>'+
        '</div>'+
    '</form>';

    document.getElementById("commentsSection").innerHTML = comments;
}

function addListingLikes(itemID)
{
    if(!getCurrentUserID) //not logged in
    {
        window.location.href = "index.html";
    }
    else // is logged in
    {
        var userID = sessionStorage.getItem('status');

        fetch('https://tic2601-t11.herokuapp.com/add_listing_likes', {
            method: 'POST',
            headers:{
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                "item_id": itemID,
                "user_id": parseInt(getCurrentUserID())
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            location.reload();
        })
        .catch(error => console.log(error)); 
    }
}

function checkIfUserLikedListing(is_liked)
{
    if(is_liked)
        return '<span><i class="fas fa-heart"></i></span>';
    else
        return '<span><i class="far fa-heart"></i></span>';
}

function toViewListing(itemID)
{
    if (getCurrentUserID() < 0)
    {
       if(confirm("Please Log In first"))
       {
        window.location.href = 'loginForm.html';
       }
    }
    else
    {
        window.location.href = 'ViewListing.html?itemID='+ itemID;
    }
}


function displayListing(d)
{
    console.log(d);
    document.getElementById("cards").innerHTML += 
    '<div class="col-md-4 mt-5">'+
    '<div class="card" id="'+d.item_id+'">'+
        '<div class="card-header"><a href="viewProfile.html?profileID='+d.seller_id+'">@'+d.seller_name+'</a></div>'+
        '<div class="card-body pb-5">'+
        '<div class="container">'+
            '<div class="row">'+
                '<div class="col text-center" id="imgContainer">'+
                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+d.item_id+'.jpg" />'+
                '</div>'+
            '</div>'+
            '<div class="row">'+
                '<div class="col">'+
                    '<h5 class="card-title">'+ d.item_name +'</h5>'+
                '</div>'+
            '</div>'+ 
            '<div class="row">'+
                '<div class="col">'+
                    '<h5 class="card-title">$'+ parseInt(d.item_price)/100 +'</h5>'+
                '</div>'+                       
            '</div>'+
            '<div class="row">'+
                '<div class="col">'+
                    '<h5 class="card-title">'+
                    '<a href="javascript:void(0);" onclick="addListingLikes('+d.item_id+');">'+checkIfUserLikedListing(d.is_liked)+'</a> ' + d.listing_likes + 
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
                    '<a href="javascript:void(0);" onclick="toViewListing('+d.item_id+');" class="btn btn-primary">View Listing</a> '+
                '</div>'+
            '</div>'+ 
        '</div>'+
        '</div>'+
    '</div>'+
    '</div>'; 
}

function getSearchItem()
{
    var searchItem = document.getElementById('searchItem').value;
    console.log(searchItem + "- searchItem" );
    fetch('https://tic2601-t11.herokuapp.com/get_listings_using_filters', {
        method: 'POST',
        headers:{
            'Authorization': getToken(),
            'Content-Type': 'application/json'
        }, 
        body: JSON.stringify({
            "search_keyword": searchItem,
        })
    })
    .then(response => response.json())
    .then(data => {
        //console.log(data);
        if(data.Respmeta.ErrorCode != 0)
        {
            if(confirm("0 search results for " +searchItem))
            {
                location.reload();
            }
        }
        else // successful
        {
            console.log(data);
            document.getElementById("cards").innerHTML = "";
            document.getElementById("searchResultText").innerHTML = "";
            document.getElementById("searchResultText").innerHTML += 
            '<div class="col">'+
                '<h1>'+data.Data.length+' search result(s) found</h1>'+
            '</div>';
            for(const d of data.Data){
               displayListing(d,);
            }
        }
    })
    .catch(error => console.log(error)); 
}

function getFilterResults()
{
   // filterOptions

   console.log("GO GET FILTER RESULTS");
    
    var sortByCat = document.getElementById("sortByCat");
    var selectedCategory = sortByCat.options[sortByCat.selectedIndex].value;
    var sortByLocation = document.getElementById("sortByLocation");
    var selectedLocation = sortByLocation.options[sortByLocation.selectedIndex].value;
    var minPrice = document.getElementById("minPrice").value;
    var maxPrice = document.getElementById("maxPrice").value;
    var sortbyFlag = document.getElementById("sortFlag");
    var selectedSortFlag = sortbyFlag.options[sortbyFlag.selectedIndex].value;
    
    console.log("selected Cat: " + selectedCategory);
    console.log("selected Location: " + selectedLocation);
    console.log("min: " + minPrice);
    console.log("max: " + maxPrice);
    console.log("sort flag: " + selectedSortFlag);

    if(selectedCategory)
        selectedCategory = parseInt(selectedCategory);

    if(selectedLocation)
        selectedLocation = parseInt(selectedLocation);
    
    if(minPrice || minPrice != "")
        minPrice = parseInt(minPrice);
    else
        minPrice = null;
        
    if(maxPrice || maxPrice != "")
        maxPrice = parseInt(maxPrice);
    else
        maxPrice = null;
        
    
    if(selectedSortFlag)
        selectedSortFlag = parseInt(selectedSortFlag);


    fetch('https://tic2601-t11.herokuapp.com/get_listings_using_filters', {
        method: 'POST',
        headers:{
            'Authorization': getToken(),
            'Content-Type': 'application/json'
        }, 
        body: JSON.stringify({
            "category_filter": {
                "item_category": selectedCategory
            },
            "location_filter": {
                "location": selectedCategory
            },
            "price_filter": {
                "min_price": minPrice,
                "max_price": maxPrice,
            },
            "sort_flag": selectedSortFlag
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
        {
            if(confirm("Unable to filter due to the following reason: " + data.Respmeta.DebugMsg))
            {
                location.reload();
            }
        }
        else // successful
        {
            document.getElementById("cards").innerHTML = "";
            document.getElementById("searchResultText").innerHTML = "";
            document.getElementById("searchResultText").innerHTML += 
            '<div class="col">'+
                '<h1>'+data.Data.length+' search result(s) found</h1>'+
            '</div>';
            for(const d of data.Data){
                displayListing(d,);
            }
        }
    })
    .catch(error => console.log(error)); 
}

function buyNow(itemID, sellerID)
{
    console.log("buy now ", itemID);
    var qtyToPurchase = document.getElementById("qtyToPurchase").value;
    console.log("qtyToPurchase: " + qtyToPurchase);
    if(qtyToPurchase != "" || !qtyToPurchase)
    {
        if(Number(qtyToPurchase) && Number(qtyToPurchase) != 0)
        {
            alert("Proceeding Transactions");
            //do whatever

            fetch('https://tic2601-t11.herokuapp.com/purchase_single_item', {
                method: 'POST',
                headers:{
                    'Authorization': getToken(),
                    'Content-Type': 'application/json'
                }, 
                body: JSON.stringify({
                    "item_id": parseInt(itemID),
                    "user_id": parseInt(getCurrentUserID()),
                    "purchase_quantity":parseInt(qtyToPurchase)
                })
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if(data.Respmeta.ErrorCode != 0)
                    {
                        if(confirm("Failed Purchasing due to " + data.Respmeta.DebugMsg))
                        {
                            location.reload();
                        }
                    }
                    else // successful
                    {
                        if(confirm("Successfully Purchased!"))
                        {
                            console.log("seller ID" + sellerID);
                            window.location.href = "review.html?sellerID="+sellerID;
                        }
                    }
            })
            .catch(error => console.log(error)); 
        }
    }
    else
    {
        alert("Please enter quantity");
    }
    
}

function getUserLikedListing() {
    fetch('https://tic2601-t11.herokuapp.com/get_user_liked_listings', {
        method: 'POST',                
        headers:{
            'Authorization': getToken(),
            'Content-Type': 'application/json'
        },  
        body: JSON.stringify({
            "user_id": parseInt(getCurrentUserID())
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
        {
            if(confirm("Unable to get user liked listings due to the following reason: " + data.Respmeta.DebugMsg))
            {
                location.reload();
            }
        }
        else // successful
        {
            document.getElementById("cards").innerHTML = "";
            document.getElementById("title").innerHTML = "<h1>Your Favourite Listings</h1>";
            for(const d of data.Data){
                displayListing(d);
            }
        }
    })
    .catch(error => console.log(error)); 
}

function viewProfileByUserID(profileID)
{
    var currentUser = getCurrentUserID();
    if(!profileID)
    {
        // viewing own listings
        profileID = currentUser;
    }
    else
    {
        //view other seller's profile
    }

    fetch('https://tic2601-t11.herokuapp.com/get_user_details', {
        method: 'POST',
        headers:{
            'Authorization': getToken(),
            'Content-Type': 'application/json'
        }, 
        body: JSON.stringify({
            "user_id": parseInt(profileID)
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
        {
            if(confirm("Unable to get user's details due to the following reason: " + data.Respmeta.DebugMsg))
            {
                location.reload();
            }
        }
        else // successful
        {
            document.getElementById("cards").innerHTML = "";
            document.getElementById("title").innerHTML = "<h3>Listings ("+data.Data.user_listings.length+" listings)</h3>";
            for(const d of data.Data.user_listings){
                displayListing(d);
            }
            displayUserReviews(data.Data);
            // document.getElementById("cards").innerHTML = "";
            // document.getElementById("title").innerHTML = "<h1>Your Listings</h1>";
            // for(const d of data.Data){
            //     displayListing(d);
            // }
        }
    })
    .catch(error => console.log(error)); 

}

function getAllListing() {
    fetch('https://tic2601-t11.herokuapp.com/get_all_listings', {
      method: 'GET'
    })
    .then(response => response.json())
    .then(result => {/*result.Data*/  for(const d of result.Data){
          //console.log(d);  
         displayListing(d);
      }
    })
    .catch(error => {console.log('NO:', JSON.stringify(error));});
}

function getLatestListing() {
    fetch('https://tic2601-t11.herokuapp.com/get_latest_listings', {
      method: 'GET',
      headers:{
        'Authorization': getToken(),
        'Content-Type': 'application/json'
    }
    })
    .then(response => response.json())
    .then(result => {/*result.Data*/  
        console.log(result);
        for(const d of result.Data){
          //console.log(d);  
         displayListing(d);
      }
    })
    .catch(error => {console.log('NO:', JSON.stringify(error));});
}

function displayUserReviews(data)
{
    var reviews =
    '<div class="row mt-3">' +
        '<div class="col"><h3>@'+data.account_info.user_name+'\'s Profile </h3></div>'+
    '</div>' +
    '<div class="row mt-3">' +
        '<div class="col"><h3>Ratings: '+data.ratings.user_ratings+' out of 5<span><i class="fas fa-star"></i></span> (Total of '+data.review_count+' Reviews)</h3></div>'+
    '</div>';
    
    for(var i = 0; i < data.user_reviews.length; i++)
    {
        reviews +=
        '<div class="row text-wrap">' +
            '<div class="col-3">'+convertUnixToTimeStamp(data.user_reviews[i].ctime)+' ∙ <a href="viewProfile.html?profileID='+data.user_reviews[i].user_id+'">'+data.user_reviews[i].user_name+'</a></div>'+
            '<div class="col-2">'+data.user_reviews[i].ratings +'/5 <span><i class="far fa-star"></i></span></div>'+
            '<div class="col-7">'+data.user_reviews[i].review_text+'</div>'+
        '</div>';
    }

    // if(data.account_info.user_id != getCurrentUserID()) // view other ppl profile
    // {
    //     reviews += 
    //     '<form id="userReview" method="post">'+
    //         '<div class="row mt-3">' +
    //             '<div class="col-2">'+
    //                 '<div class="form-group">'+
    //                     '<input type="number" class="form-control" id="rating" placeholder="/5 star">'+
    //                 '</div>'+
    //             '</div>'+
    //             '<div class="col-8">'+
    //                 '<div class="form-group">'+
    //                     '<input type="text" class="form-control" id="review" placeholder="Add a review...">'+
    //                 '</div>'+
    //             '</div>'+
    //             '<div class="col-2">'+
    //                 '<input type="button" class="btn btn-primary" onclick="addUserReview('+data.account_info.user_id+')" value="Submit"/>'+
    //             '</div>'+
    //         '</div>'+
    //     '</form>';
    // }
    
    document.getElementById("reviewSection").innerHTML = reviews;
}

function addUserReview(sellerID)
{
    var rating = document.getElementById("rating");
    var ratingValue = rating.options[rating.selectedIndex].value;
    var review = document.getElementById('review').value;

    if(!review || !ratingValue || review == "")
    {
        alert("Review incomplete, please try again");
        location.reload();
    }

    fetch('https://tic2601-t11.herokuapp.com/add_user_review', {
        method: 'POST',
        headers:{
            'Authorization': getToken() ,
            'Content-Type': 'application/json'
        }, 
        body: JSON.stringify({
            "user_id": parseInt(getCurrentUserID()),
            "seller_id": parseInt(sellerID),
            "ratings": parseInt(ratingValue),
            "review_text": review
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if(data.Respmeta.ErrorCode != 0)
            {
                if(confirm("Review failed due to the following reason: " + data.Respmeta.DebugMsg))
                {
                    location.reload();
                }
            }
            else // successful
            {
                if(confirm("Successfully added review"))
                {
                   window.location.href="viewProfile.html?profileID="+sellerID;
                }
            }
    })
    .catch(error => console.log(error)); 
}

function loadListingDetails(data)
{
    console.log("edit" + JSON.stringify(data));
    //window.location.href ="editListing.html";
    document.getElementById("itemName").value = data.item_name;
    document.getElementById("itemPrice").value = parseInt(data.item_price)/100;
    document.getElementById("itemQty").value = data.item_quantity;
    document.getElementById("itemDesc").value = data.item_description;
    document.getElementById("itemCat").value = data.item_category;
    document.getElementById("itemLocation").value = data.item_location;
    document.getElementById("myImg").src = 'https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+data.item_id+'.jpg';
}


function editListing(itemID)
{
    var itemName = document.getElementById('itemName').value;
    var itemPrice = document.getElementById("itemPrice").value;
    var itemQty = document.getElementById("itemQty").value;
    var itemDesc = document.getElementById("itemDesc").value;
    var itemCat = document.getElementById("itemCat");
    var itemCatValue = itemCat.options[itemCat.selectedIndex].value;
    var location = document.getElementById("itemLocation");
    var locationValue = location.options[location.selectedIndex].value;

    if(itemPrice || itemPrice != "")
        itemPrice = parseInt(itemPrice);
    else
        itemPrice = null;

    if(itemQty || itemQty != "")
        itemQty = parseInt(itemQty);
    else
        itemQty = null;

    if(!itemName || itemName == "")
        itemName = null;

    if(!itemDesc || itemDesc == "")
        itemDesc = null;
    
    if(itemCatValue)
        itemCatValue = parseInt(itemCatValue);

    if(locationValue)
        locationValue = parseInt(locationValue);


    var base64String = "";
    const file = document.getElementById("img").files[0];
    console.log("file: " + file);
    if(file)
    {
        var reader = new FileReader();
      
        reader.onload = function () {
            base64String = reader.result.replace("data:", "")
                .replace(/^.+,/, "");
      
            imageBase64Stringsep = base64String;
        }
        reader.readAsDataURL(file);
        reader.addEventListener('load', (event) => {
            console.log('load finish: ' + base64String ); // base 64  
            fetch('https://tic2601-t11.herokuapp.com/update_single_listing', {
                method: 'PATCH',
                headers:{
                    'Authorization': getToken(),
                    'Content-Type': 'application/json'
                }, 
                body: JSON.stringify({
                    "item_id": parseInt(itemID),
                    "item_name": itemName,
                    "item_price": itemPrice*100,
                    "item_quantity":itemQty,
                    "item_description":itemDesc,
                    "item_location": locationValue,
                    "item_category": itemCatValue,
                    "item_image": base64String,
                    "seller_id": parseInt(getCurrentUserID())
                })
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if(data.Respmeta.ErrorCode != 0)
                    {
                        if(confirm("Failed updating listing"))
                        {
                            //location.reload();
                        }
                    }
                    else // successful
                    {
                        if(confirm("Listing Updated!"))
                        {
                            window.location.href = "viewListing.html?itemID="+itemID;
                        }
                    }
            })
            .catch(error => console.log(error));
        });
    }
    else
    {
        alert("no photo");
        fetch('https://tic2601-t11.herokuapp.com/update_single_listing', {
                method: 'PATCH',
                headers:{
                'Authorization': getToken(),
                'Content-Type': 'application/json'
                }, 
                body: JSON.stringify({
                    "item_id": parseInt(itemID),
                    "item_name": itemName,
                    "item_price": itemPrice*100,
                    "item_quantity":itemQty,
                    "item_description":itemDesc,
                    "item_location": locationValue,
                    "item_category": itemCatValue,
                    "item_image": null,
                    "seller_id": parseInt(getCurrentUserID())
                })
            })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if(data.Respmeta.ErrorCode != 0)
                    {
                        if(confirm("Failed updating listing"))
                        {
                            //location.reload();
                        }
                    }
                    else // successful
                    {
                        if(confirm("Listing Updated!"))
                        {
                            window.location.href = "viewListing.html?itemID="+itemID;
                        }
                    }
            })
            .catch(error => console.log(error));
    }
}
