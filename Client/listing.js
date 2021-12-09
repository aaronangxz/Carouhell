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
        console.log("Converted Base64 version is " + document.getElementById("imgTest").innerHTML);
      }
      fileReader.readAsDataURL(fileToLoad);
    }
}

function createListing(userID)
{
    if (getCurrentUserID() == -1){
        window.location.href = "index.html"
    }
    setPrevSecureLocation();
    var itemName = document.getElementById('itemName').value;
    var itemPrice = document.getElementById("itemPrice").value;
    var itemQty = document.getElementById("itemQty").value;
    var itemDesc = document.getElementById("itemDesc").value;
    var itemCat = document.getElementById("itemCat");
    var itemCatValue = itemCat.options[itemCat.selectedIndex].value;
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
                "item_price": parseFloat(itemPrice)*100,
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
                    window.location.href = "viewListing.html?itemID=" + data.Data.item_id;
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
                    window.location.href = "index.html";
                }
            }
    })
    .catch(error => console.log(error)); 
}

function addListingComment(userID, itemID)
{
    var comment = document.getElementById('comment').value;
    if (comment == ""){
        confirm("Comment cannot be empty.")
        return
    }
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
                location.reload();
            }
    })
    .catch(error => console.log(error)); 

}

function viewListingByItemId(itemID)
{
    setPrevLocation();
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
            getRecommendedListingsByItemId(data.Data.item_id);
            document.title = 'Carouhell - '+data.Data.item_name;
        })
        .catch(error => console.log(error));   
}

function getRecommendedListingsByItemId(itemID)
{
    document.getElementById("recommendedListingsSection").innerHTML +='<div class="row mt-3">' +
    '<div class="col"><h3><i class="fas fa-fire" style="color:red"></i></span> You may also like <h3></div>'+
'</div>'
    fetch('https://tic2601-t11.herokuapp.com/v2/get_recommended_listings_by_itemid', {
            method: 'POST',
            headers:{
                'Authorization': getToken(),
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                "item_id": parseInt(itemID),
                "user_id": parseInt(getCurrentUserID())
            })
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
           
            for(const d of data.Data){
                displayListing(d,'true');
             }
        })
        .catch(error => console.log(error)); 
}

function displayItemContent(data)
{
    
    var status = ""
    if (data.item_status == 2){
        status +='<span class="badge badge-pill badge-secondary">Sold Out</span>'
    }else if (data.item_stock - data.item_quantity == 0){
        status +='<span class="badge badge-pill badge-success">Available</span>'
    }else if (((data.item_quantity/data.item_stock) * 100) <= 10){
        status +='<span class="badge badge-pill badge-danger">Low in stock</span>'
    }else if (((data.item_quantity/data.item_stock) * 100) <= 25){
        status +='<span class="badge badge-pill badge-warning">Selling Fast</span>'
    }else if (data.item_status == 1){
        status +='<span class="badge badge-pill badge-success">Available</span>'
    }

    var progress = ""
    if (((data.item_quantity/data.item_stock) * 100) <= 10 && ((data.item_quantity/data.item_stock) * 100) != 0){
        progress += '<div class="progress" style="width: 25%">'+
        '<div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" role="progressbar" style="width: 90%" aria-valuenow="90" aria-valuemin="0" aria-valuemax="100">'+ (data.item_stock-data.item_quantity) + ' Bought</div>'+
      '</div>'
    }else if (((data.item_quantity/data.item_stock) * 100) <= 25 && ((data.item_quantity/data.item_stock) * 100) != 0){
        progress += '<div class="progress" style="width: 25%;">'+
        '<div class="progress-bar progress-bar-striped bg-warning" role="progressbar" style="width: 75%" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100">'+ (data.item_stock-data.item_quantity) + ' Bought</div>'+
      '</div>'
    }

    var sellerType = ''
    if (data.seller_type == 1){
        sellerType = '<span title = "Official Store"><i class="fas fa-check-circle" style="color:Dodgerblue"></i></span>'
    }else if (data.seller_type == 2){
        sellerType = '<span title = "Preferred Seller"><i class="fas fa-medal" style="color:Coral"></i></span>'
    }

    var content = "";
    content += '<div class="row">' +
        '<div class="col-4"><img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+data.item_id+'.jpg" class="img-fluid" /></div>'+
        '<div class="col-8">'+
            '<div class="row"><div class="col"><h1>'+data.item_name+' </h1></div></div>'+
              '<div class="row">' +
                '<div class="col"><h2>'+ '  S' +(parseInt(data.item_price)/100).toLocaleString('en-US', {
                    style: 'currency',
                    currency: 'USD',
                  })+'</h2></div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><h3>'+status+'</h3></div>'+
            '</div>'+
            progress+

            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-clock"></i></span> Posted '+getTimeStamp(data.listing_ctime)+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col">'+
                    '<a href="javascript:void(0);" onclick="addListingLikes('+data.item_id+');" id="like_'+ data.item_id+'">'+ checkIfUserLikedListing(data.is_liked)+'</a><span id="likecount_'+ data.item_id+'"> '+ data.listing_likes+'</span>'+
                '</div>'+
            '</div>'+
            '<div class="row mt-3">' +
                '<div class="col"><span><i class="fas fa-layer-group"></i></span> '+categories_Arr[data.item_category]+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-map-marker-alt"></i></span> '+location_Arr[data.item_location]+'</div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col"><span><i class="fas fa-user"></i> </span><a href="viewProfile.html?profileID='+data.seller_id+'"style="color: black">@'+data.seller_name+'</a> '+sellerType+'</div>'+
            '</div>'+
            '<div class="row mt-3">' +
                '<div class="col"><h3>Description <h3></div>'+
            '</div>'+
            '<div class="row">' +
                '<div class="col">'+data.item_description+'</div>'+
            '</div>'+
        '</div>'+
    '</div>';
    if (getCurrentUserID() == -1){
        content += 
        '<form id="purchaseForm" method="post">'+
            '<div class="row mt-3" id="buyerBtns">'+
                '<div class="col-4"></div>'+
                // '<div class="col-3 text-center">'+
                //     '<h5> Quantity Available: '+ data.item_quantity + '</h5>' +
                // '</div>'+
                // '<div class="col-1">'+
                //     '<input type="button" class="btn btn-secondary disabled" value="Sold Out"/>'+
                //     // '<input type="button" id="btnBuyNow" class="btn btn-primary" value="Buy Now"/>'+
                // '</div>'+
                '<div class="col-4"></div>'+
            '</div>'+
        '</form>';
    }
    else if(data.seller_id == getCurrentUserID()) // user is listing's owner
    {
        // display seller's option
        content += 
        '<div class="row mt-5" id="sellerBtns">'+
            '<div class="col-2">'+
                // '<input type="button" id="btnEdit" class="btn btn-primary" value="Edit Listing"/>'+
                 '<a href="editListing.html?itemID='+data.item_id+'" id="btnEdit" class="btn btn-secondary">Edit Listing</a> '+
            '</div>'+
            '<div class="col-1">'+
                '<a href="javascript:void(0);" onclick="deleteListing('+data.item_id+')" id="deleteBtn" class="btn btn-danger">Delete Listing</a>'+
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
                    // '<div class="col-3 text-center">'+
                    //     '<h5> Quantity Available: '+ data.item_quantity + '</h5>' +
                    // '</div>'+
                    // '<div class="col-1">'+
                    //     '<input type="button" class="btn btn-secondary disabled" value="Sold Out"/>'+
                    //     // '<input type="button" id="btnBuyNow" class="btn btn-primary" value="Buy Now"/>'+
                    // '</div>'+
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
                        '<div class="col-3">'+
                            '<div class="form-group">'+
                            '<input type="number" class="form-control" id="qtyToPurchase" placeholder="Enter Quantity">'+
                            '<p>Quantity Available: '+ data.item_quantity + '</p>'+
                            '</div>'+
                    '</div>'+
                    '<div class="col-1">'+
                        '<input type="button" onclick="buyNow('+data.item_id+','+data.seller_id+')" class="btn btn-primary" background-color value="Buy Now"/>'+
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
    var commentFooter = ''
    var showCommentButton = 
    '<form id="listingComment" method="post">'+
        '<div class="row mt-3">' +
            '<div class="col-10">'+
                '<div class="form-group">'+
                    '<input type="text" class="form-control" id="comment" placeholder="Add a comment...">'+
                '</div>'+
            '</div>'+
            '<div class="col-2">'+
                '<input type="button" class="btn btn-secondary" onclick="addListingComment(getCurrentUserID(),'+itemID+')" value="Submit"/>'+
            '</div>'+
        '</div>'+
    '</form>'
    if (getCurrentUserID()== -1){
        showCommentButton = ""
    }
    var comments =
    '<div class="row mt-3">' +
        '<div class="col"><h3><span><i class="far fa-comment-dots" style="color:grey"></i></span> Comments <h3></div>'+
    '</div>';
    if (data.length == 0){
        if (getCurrentUserID()== -1){
            comments += '<font style="opacity:.6" size="3px"><i><a href="loginForm.html">Login</a> to start a discussion.</i></font>'
        }else{
            comments += '<font style="opacity:.6" size="3px"><i>Start a discussion.</i></font>'
        }        
    }else{
        if (getCurrentUserID()== -1){
             commentFooter += '<font style="opacity:.6" size="3px"><i><a href="loginForm.html">Login</a> to join the discussion.</i></font>'
        }else{
            commentFooter += '<font style="opacity:.6" size="3px"><i>Join the discussion.</i></font>'
        }
    }
    for(var i = 0; i < data.length; i++)
    {
        console.log("comments: " + data[i]);
        comments +=
        '<div class="row text-wrap" >' +
            '<div class="col-3"><a href="viewProfile.html?profileID='+data[i].user_id+'">'+data[i].user_name +'</a> | '+convertUnixToTimeStamp(data[i].ctime)+ '</div>'+
            '<div class="col-9">'+data[i].comment+'</div>'+
        '</div>';
    }
    comments += showCommentButton + commentFooter;

    document.getElementById("commentsSection").innerHTML = comments;
}

function addListingLikes(itemID)
{
    if (getCurrentUserID() < 0)
    {
        $('#promptLogIn').modal('show')
        //$('.modal-backdrop').remove();
        $(document).on('click','#toLogin',function(){
            window.location.href = 'loginForm.html';
       })
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

            if(data.Data.is_liked == true){
                document.getElementById("like_"+ itemID).innerHTML = '<span><i class="fas fa-heart" style="color:red"> </i></span>'
                document.getElementById("likecount_"+ itemID).innerHTML = '<span> '+ data.Data.latest_likes_count +'</span>'
            }else{
                document.getElementById("like_"+ itemID).innerHTML = '<span><i class="far fa-heart" style="color:black"></i></span>'
                document.getElementById("likecount_"+ itemID).innerHTML = '<span> '+ data.Data.latest_likes_count +'</span>'
            }
            // document.location.reload();
        })
        .catch(error => console.log(error)); 
    }
}

function checkIfUserLikedListing(is_liked)
{
    if(is_liked && getCurrentUserID() != -1)
        return '<span><i class="fas fa-heart" style="color:red"> </i></span>';
    else
        return '<span><i class="far fa-heart" style="color:black"></i></span>';
}

function toViewListing(itemID)
{
    window.location.href = 'viewListing.html?itemID='+ itemID;

    // if (getCurrentUserID() < 0)
    // {
    //    if(confirm("Please Log In first"))
    //    {
    //     window.location.href = 'loginForm.html';
    //    }
    // }
    // else
    // {
    //     window.location.href = 'viewListing.html?itemID='+ itemID;
    // }
}

function displayListing(d, isRecommend)
{
    //console.log(d);

    var bool_value = isRecommend == 'true';
    var element = "cards"
    var rcmdHeader = ''
    if (bool_value === true){
        element = "recommendedListings"
        rcmdHeader += '<div class="row mt-3">' +
        '<div class="col"><h3>💬 Comments <h3></div>'+
    '</div>';
    }

    var status = ""
    if (d.item_status == 2){
        status +='<span class="badge badge-pill badge-secondary">Sold Out</span>'
    }else if (d.item_stock - d.item_quantity == 0){
        status +='<span class="badge badge-pill badge-success">Available</span>'
    }else if (((d.item_quantity/d.item_stock) * 100) <= 10){
        status +='<span class="badge badge-pill badge-danger">Low in stock</span>'
    }else if (((d.item_quantity/d.item_stock) * 100) <= 25 ){
        status +='<span class="badge badge-pill badge-warning">Selling Fast</span>'
    }else if (d.item_status == 1){
        status +='<span class="badge badge-pill badge-success">Available</span>'
    }

    var likes = ""
    if (d.listing_likes >= 10){
        likes += '<span><i class="fas fa-chart-line" style="color:red"></i></span>'
    }

    var sellerType = ''
    if (d.seller_type == 1){
        sellerType = '<span title = "Official Store"><i class="fas fa-check-circle" style="color:Dodgerblue"></i></span>'
    }else if (d.seller_type == 2){
        sellerType = '<span title = "Preferred Seller"><i class="fas fa-medal" style="color:Coral"></i></span>'
    }

    document.getElementById(element).innerHTML +=
    '<div class="col-lg-3 col-md-4 col-sm-6 col-xs-6 mt-4">'+
        '<div class="card border-0" id="'+d.item_id+'">'+
            '<div class="card-header card-header-color border-0 pb-0"><a href="viewProfile.html?profileID='+d.seller_id+'" style="color: black">@'+d.seller_name+'</a> '+ sellerType +
            '</div>'+
            '<div class="row">'+
                '<div class="col">'+
                    '<p class = "pl-4 pt-0" id="itemDay"><font style="opacity:.8" size="2px">'+ getTimeStamp(d.listing_ctime) +' </font></p>'+
                '</div>'+
            '</div>'+
                '<div class="card-body pb-3">'+
                    '<div class="container">'+
                        '<div class="row ">'+
                            '<a href="'+'viewListing.html?itemID='+d.item_id+'"class="btn btn-link stretched-link"></a> '+
                                '<div class="col-lg-12 col-md-12 col-sm-12 col-xs-6 text-center" id="imgContainer">'+
                                    '<img src="https://tic2601-t11.s3.ap-southeast-1.amazonaws.com/listing_'+d.item_id+'.jpg" class="img-cover img-fluid" />'+
                                '</div>'+
                        '</div>'+
                        '<div class="row">'+
                        '<h6 class="card-title pl-0" id="itemName">'+ d.item_name +'</h6>'+
                        '</div>'+ 
                    '</div>'+
                '</div>'+
                '<div class="card-footer d-flex border-0">'+
                    '<div class="row">'+
                        '<div class="col">'+
                            '<h5 class="card-title">S'+ (parseInt(d.item_price)/100).toLocaleString('en-US', {
                                style: 'currency',
                                currency: 'USD',
                                }) +
                            '</h5>'+
                            '<h6 class="card-title" id="likebutton">'+
                                '<a href="javascript:void(0);" onclick="addListingLikes('+d.item_id+');" id="like_'+ d.item_id+'">'+ checkIfUserLikedListing(d.is_liked)+'</a><span id="likecount_'+ d.item_id+'"> '+ d.listing_likes+'</span>'+ 
                            '</h6>'+
                            '<h6 class="card-title">'+ status + '  ' + likes +
                            '</h6>'+
                        '</div>'+
                    '</div>'+
                '</div>'+
        '</div>'+
    '</div>'; 
}

function getSearchItem()
{
    document.getElementById("landingHomePage").innerHTML =   '<div class="loader-wrapper">'+
    '<span class="loader"><span class="loader-inner"></span></span>'+
    '</div>';
    var searchItem = document.getElementById('searchItem').value;
    if (searchItem == ''){
        location.reload();
    }
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
            document.getElementById("searchResultText").innerHTML = 
            '<div class="col">'+
                '<h1>No result found</h1>'+
            '</div>';
            $(".loader-wrapper").fadeOut("slow");
        document.getElementById("landingHomePage").innerHTML = ''
            document.getElementById("cards").innerHTML = "";
            // if(confirm("0 search results for " +searchItem))
            // {
            //     location.reload();
            // }
        }
        else // successful
        {
            console.log(data);
            document.getElementById("landingHomePage").innerHTML = '';
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
    sessionStorage.setItem('prevLocation',window.location)
}

function getFilterResults()
{
    document.getElementById("footer").innerHTML = ''
    document.getElementById("cards").innerHTML =   '<div class="loader-wrapper">'+
    '<span class="loader"><span class="loader-inner"></span></span>'+
    '</div>'   
   // filterOptions    
    var sortByCat = document.getElementById("sortByCat");
    var selectedCategory = sortByCat.options[sortByCat.selectedIndex].value;
    var sortByLocation = document.getElementById("sortByLocation");
    var selectedLocation = sortByLocation.options[sortByLocation.selectedIndex].value;
    var minPrice = document.getElementById("minPrice").value;
    var maxPrice = document.getElementById("maxPrice").value;
    var sortbyFlag = document.getElementById("sortFlag");
    var selectedSortFlag = sortbyFlag.options[sortbyFlag.selectedIndex].value;
    

    if(selectedCategory)
        selectedCategory = parseInt(selectedCategory);

    if(selectedLocation)
        selectedLocation = parseInt(selectedLocation);
    
    if(parseInt(minPrice) <= 0)
        minPrice = null;        

    if(minPrice || minPrice != "")
        minPrice = parseFloat(minPrice)*100;
    else
        minPrice = null;

    if(parseInt(maxPrice) <= 0)
        maxPrice = null;         
    
    if(maxPrice || maxPrice != "")
        maxPrice = parseFloat(maxPrice)*100;
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
                "location": selectedLocation
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
            $('#noResults').modal('show')
            //$('.modal-backdrop').remove();
            $('#noResults').on('hidden.bs.modal', function (e) {
                location.reload();
              })
        }
        else // successful
        {
            document.getElementById("landingHomePage").innerHTML = '';
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
    sessionStorage.setItem('prevLocation',window.location)
}

function buyNow(itemID, sellerID)
{
    setPrevSecureLocation();
    console.log("buy now ", itemID);
    var qtyToPurchase = document.getElementById("qtyToPurchase").value;
    console.log("qtyToPurchase: " + qtyToPurchase);
    if(qtyToPurchase != "" || !qtyToPurchase)
    {
        if(Number(qtyToPurchase) && Number(qtyToPurchase) != 0)
        {
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
                            if(data.Data.is_reviewed)
                            {
                                window.location.href = "wallet.html";
                            }
                            else
                            {
                                window.location.href = "review.html?sellerID="+sellerID;
                            }                            
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
    if (getCurrentUserID() == -1){
        window.location.href = "index.html"
    }
    setPrevSecureLocation();
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
        document.getElementById("title").innerHTML = '<h1>My Favourites('+data.Data.length+')</h1>';
        if(data.Respmeta.ErrorCode != 0)
        {
            if(confirm("Unable to get user liked listings due to the following reason: " + data.Respmeta.DebugMsg))
            {
                //location.reload();
            }
        }
        else // successful
        {
            document.getElementById("cards").innerHTML = "";
            // document.getElementById("title").innerHTML = "<h1>Your Favourite Listings</h1>";
            for(const d of data.Data){
                displayListing(d);
            }
        }
    })
    .catch(error => console.log(error)); 
}

function viewProfileByUserID(profileID)
{
    setPrevLocation();
    var currentUser = getCurrentUserID();
    if(!profileID)
    {
        // viewing own listings
        profileID = currentUser;
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
            document.getElementById("title").innerHTML = "<h2>"+data.Data.user_listings.length+" listings</h2>";

            for(const d of data.Data.user_listings){
                displayListing(d);
            }
            displayUserReviews(data.Data);
            document.title = 'Carouhell - @'+data.Data.account_info.user_name + "'s profile";
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
    document.getElementById("footer").innerHTML +=   '<div class="loader-wrapper">'+
        '<span class="loader"><span class="loader-inner"></span></span>'+
        '</div>'   
    fetch('https://tic2601-t11.herokuapp.com/v2/get_latest_listings', {
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
    .then(result => {/*result.Data*/  
        $(".loader-wrapper").fadeOut("slow");
        console.log(result);
        for(const d of result.Data){
         displayListing(d);
      }
    })
    .catch(error => {console.log('NO:', JSON.stringify(error));});
    sessionStorage.setItem('prevLocation',window.location)
}

function displayUserReviews(data)
{
    var sellerType = ''
    if (data.account_info.seller_type == 1){
        sellerType = '<span title = "Official Store"><i class="fas fa-check-circle" style="color:Dodgerblue"></i></span>'
    }else if (data.account_info.seller_type == 2){
        sellerType = '<span title = "Preferred Seller"><i class="fas fa-medal" style="color:Coral"></i></span>'
    }

    var reviews =
    '<div class="row mt-3">' +
        '<div class="col"><h1>@'+data.account_info.user_name + ' ' + sellerType +'</h1></div>'+
    '</div>' +
    '<div class="row mt-1">' +
        '<div class="col"><small style="color:grey;">Last seen '+ getLastSeen(data.account_info.user_last_login) +'</small></div>'+
    '</div>'+
    '<div class="row mt-3">' +
        '<div class="col"><h2>'+data.ratings.user_ratings+'<span><i class="fas fa-star" style="color:gold"></i></span> ('+data.review_count+' Reviews)</h2></div>'+
    '</div>';
    


    for(var i = 0; i < data.user_reviews.length; i++)
    {
        if (i == 10){
            break
        }

        rate = ""
        if (data.user_reviews[i].ratings == 1){
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        }else if (data.user_reviews[i].ratings == 2){
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        }else if (data.user_reviews[i].ratings == 3){
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        }else if (data.user_reviews[i].ratings == 4){
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="far fa-star" style="color:gold"></i></span>'
        }else if (data.user_reviews[i].ratings == 5){
            rate += '<span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span><span><i class="fas fa-star" style="color:gold"></i></span>'
        }

        reviews +=
        '<div class="row text-wrap">' +
            '<div class="col-3">'+convertUnixToTimeStamp(data.user_reviews[i].ctime)+' | <a href="viewProfile.html?profileID='+data.user_reviews[i].user_id+'">'+data.user_reviews[i].user_name+'</a></div>'+
            '<div class="col-2">'+data.user_reviews[i].ratings +'/5'+rate+'</div>'+
            '<div class="col-7">'+data.user_reviews[i].review_text+'</div>'+
        '</div>';
    }
    
    reviews += '<div class="mb-4"><hr class="solid"></div>'
    document.getElementById("reviewSection").innerHTML = reviews;
}

function addUserReview(sellerID)
{
    setPrevSecureLocation();
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
                   window.location.href="wallet.html";
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
    if (getCurrentUserID() == -1){
        window.location.href = "index.html"
    }
    setPrevSecureLocation();
    var itemName = document.getElementById('itemName').value;
    var itemPrice = document.getElementById("itemPrice").value;
    var itemQty = document.getElementById("itemQty").value;
    var itemDesc = document.getElementById("itemDesc").value;
    var itemCat = document.getElementById("itemCat");
    var itemCatValue = itemCat.options[itemCat.selectedIndex].value;
    var location = document.getElementById("itemLocation");
    var locationValue = location.options[location.selectedIndex].value;

    console.log("Price: " + itemPrice);

    if(itemPrice || itemPrice != "")
        itemPrice = parseFloat(itemPrice);
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
    const file1 = document.getElementById("img").files[0];
    const file2 = document.getElementById("img").files[1];
    console.log("file1: " + file1);
    console.log("file2: " + file2);
    if(file1)
    {
        var reader = new FileReader();
      
        reader.onload = function () {
            base64String = reader.result.replace("data:", "")
                .replace(/^.+,/, "");
      
            imageBase64Stringsep = base64String;
        }
        reader.readAsDataURL(file1);
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
                            location.reload();
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
                            location.reload();
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
