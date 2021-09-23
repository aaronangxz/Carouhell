<h1> Golang Backend </h1>

<h2>Mock Database</h2>

1. SQLite - `test.db`

- Tables: 

    - listings
    - notifications

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
/listings/item_id

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
/listings/item_id

Request
{
    ""
}

Response
{
  "data": true
}
```

<h3>GET - GetNotificationsByUserID()</h3>

```
/notifications/user_id

Request
{
    ""
}

Response
{
    "data": [
        {
            "notification_id": 2,
            "user_id": 696969,
            "notification_text": "Buy buy buy buy!"
        },
        {
            "notification_id": 1,
            "user_id": 696969,
            "notification_text": "Welcome to Carouhell!"
        }
    ],
    "resp": "success"
}
```

<h3>GET - CreateMockNotifications()</h3>

```
/create_mock_notifications

Request
{
    "user_id": 696969,
    "notification_text": "New notification text"
}

Response
{
    "resp": "success"
}
```

<h2>To Test API</h2>

1. Use Postman<br>
2. Request URL: `localhost:8080/...`
3. Insert the request body as per exmples above.
4. Send

<h2>Docker - Note to Self</h2>
1. Write `Dockerfile`
2. To build: `docker build --tag tic2601 .`
3. `docker run tic2601` will run container isolated from network.
3. Use `docker run --publish 8080:8080 tic2601` to expose container to network and port. ([host_port]:[container_port])