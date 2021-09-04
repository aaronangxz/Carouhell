<h1> Golang Backend </h1>

<h2>Mock Database</h2>

1. SQLite - `test.db`

<h2>Available API endpoints</h2>

<h3>GET - GetAllListings()</h3>

```
/listings

Request
{
    ""
}

Response
{
    "data": {
        "item_id": 1,
        "item_name": "ItemA",
        "item_price": 69
    }
}
```
	
<h3>POST - CreateListing()</h3>

```
/listings

Request
{
    "item_name": "ItemA",
    "item_price": 69
}

Response
{
    "data": {
        "item_id": 1,
        "item_name": "ItemA",
        "item_price": 69
    }
}

```
	
<h3>GET - GetListingByItemID()</h3>

```
/listings/:item_id

Request
{
    "item_id": "1",
}

Response
{
    "data": {
        "item_id": 1,
        "item_name": "ItemA",
        "item_price": 69
    }
}
```
	
<h3>PATCH - UpdateListing()</h3>

```
/listings/:item_id

Request
{
    
    "item_name": "ItemB",
    "item_price": 6969
}

Response
{
    "data": {
        "item_id": 1,
        "item_name": "ItemB",
        "item_price": 6969
    }
}
```
	
<h3>DELETE - DeleteListing()</h3>

```
/listings/:item_id

Request
{
    ""
}

Response
{
  "data": true
}
```

<h2>To Test API</h2>

1. Use Postman<br>
2. Request URL: `localhost:port/item_id`
3. Insert the request body as per exmples above.
4. Send

