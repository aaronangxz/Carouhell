<h1> Golang Backend </h1>

<h2>Available API endpoints</h2>

<h3>GET - GetAllListings()</h3>

```json
/get_all_listings

Request
{
    ""
}

Response

    "Data": [
        {
            "item_id": 65,
            "item_name": "Balenciaga Shirt",
            "item_price": 1,
            "item_quantity": 99,
            "item_purchasedquantity": 0,
            "item_description": "Buy me",
            "item_shippinginfo": 0,
            "item_paymentinfo": 0,
            "item_location": "Jurong Island",
            "item_status": 1,
            "item_category": 1,
            "item_image": "https://tic2601test.s3.ap-southeast-1.amazonaws.com/balenciagashirt.jpg",
            "seller_id": 35,
            "seller_name": "im_a_potato",
            "listing_ctime": 1633965691,
            "listing_mtime": 1633965691,
            "listing_likes": 69
        }
    ],
    "Respmeta": {
        "DebugMsg": "GetAllListings success. results: 1",
        "ErrorCode": 0
    }
}
```
	
<h3>POST - GetLatestListings()</h3>

```json
/get_latest_listings

Request
{
    "item_category": null,
    "item_status": null,
    "limit": null
}

Response

    "Data": [
        {
            "item_id": 125,
            "item_name": "Chicken Rice",
            "item_price": 1000,
            "item_quantity": 5,
            "item_purchasedquantity": 0,
            "item_description": "Add Rice",
            "item_shippinginfo": 0,
            "item_paymentinfo": 0,
            "item_location": "Orchard",
            "item_status": 1,
            "item_category": 0,
            "item_image": "https://tic2601test.s3.ap-southeast-1.amazonaws.com/chickenrice.jpg",
            "seller_id": 25,
            "seller_name": "noname",
            "listing_ctime": 1634488836,
            "listing_mtime": 1634488836,
            "listing_likes": 0
        },
    ],
    "Respmeta": {
        "DebugMsg": "GetLatestListings success. results: 1",
        "ErrorCode": 0
    }
}
```

<h3>POST - CreateListing()</h3>

```json
/create_listing

Request
{
    "item_name": null,
    "item_price": null,
    "item_quantity": null,
    "item_description": null,
    "item_shipping_info": null,
    "item_payment_info": null,
    "item_location": null,
    "item_category": null,
    "item_image": null,
    "seller_id": null
}

Response

    "Respmeta": {
        "DebugMsg": "Successfully create listing.",
        "ErrorCode": 0
    }
}
```

<h3>PATCH - UpdateSingleListing()</h3>

```json
/update_single_listing

Request
{
    "item_name": null,
    "item_price": null,
    "item_quantity": null,
    "item_description": null,
    "item_shipping_info": null,
    "item_payment_info": null,
    "item_location": null,
    "item_category": null,
    "item_image": null,
    "seller_id": null
}

Response

    "Respmeta": {
        "DebugMsg": "Successfully updated listing.",
        "ErrorCode": 0
    }
}
```
<h3>DELETE - DeleteSingleListing()</h3>

```json
/delete_single_listing

Request
{
    "item_id": null
}

Response

    "Respmeta": {
        "DebugMsg": "Successfully delete listing.",
        "ErrorCode": 0
    }
}
```

<h3>GET - GetSingleListingByItemID()</h3>

```json
/get_single_listing_by_itemid

Request
{
    "item_id": 95
}

Response
{
    "Data": {
        "item_id": 95,
        "item_name": "Nokia 3310",
        "item_price": 9999,
        "item_quantity": 1,
        "item_purchasedquantity": 0,
        "item_description": "Best phone ever.",
        "item_shippinginfo": 0,
        "item_paymentinfo": 0,
        "item_location": "JB",
        "item_status": 1,
        "item_category": 0,
        "item_image": "https://tic2601test.s3.ap-southeast-1.amazonaws.com/AzIye1m.jpg",
        "seller_id": 56,
        "seller_name": "JT",
        "listing_ctime": 1634054558,
        "listing_mtime": 1634054558,
        "listing_likes": 42
    },
    "Respmeta": {
        "DebugMsg": "GetListingByItemID success.",
        "ErrorCode": 0
    }
}
```

WIP