function loadListingDetails(data) {
    console.log("edit" + JSON.stringify(data));
    //window.location.href ="editListing.html";
    document.getElementById("itemName").value = data.item_name;
    document.getElementById("itemPrice").value = parseInt(data.item_price) / 100;
    document.getElementById("itemQty").value = data.item_quantity;
    document.getElementById("itemDesc").value = data.item_description;
    document.getElementById("itemCat").value = data.item_category;
    document.getElementById("itemLocation").value = data.item_location;
    document.getElementById("myImg").src = 'https://' + getCurrentEnv() + '.s3.ap-southeast-1.amazonaws.com/listing_' + data.item_id + '.jpg';
}