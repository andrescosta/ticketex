# Resources

## Reservations

```HTTP
POST /reservation/metadata
{
    "adventure_id":"aaaaa",
    "status":"closed"
    "capacity":[
        {
            "type":"vip",
            "max":1000
        }
    ]
}

GET /reservation/metadata/{adventure_id}
{
    "adventure_id":"aaaaa",
    "status":open
    "capacity":[
        {
            "type":"vip",
            "current":0,
            "max":1000
        }
    ]
}
```

```HTTP
PATCH /reservation/metadata/{adventure_id}/capacity/{type}
{
    "max":1000 
    (max must be bigger than current max. if not, error)
}
```

```HTTP
POST /reservation/metadata/{adventure_id}/capacity
{
    "type":aaaaa,
    "max":1000 
    (max must be bigger than current max. if not, error)
}
```


```HTTP
POST /reservation/{adventure_id}/{type}/users/{id}/
{
    "quantity":1
} 
(creates with status pending_confirmation)

{
    "id":"aaaaaaaaaa",
    "status":pending_confirmation,
    "expiration": date,
}
````

```HTTP

GET /reservation/{adventure_id}/{type}/users/{id}/
{
    "adventure_id":"aaaaaaaaaaa",
    "status":[reserved,pending_confirmation,pending_payment]
    "expiration": [date|empty for reserved]
    "user_id":"aaa"
}
```
```HTTP
PATCH /reservation/{adventure_id}/{type}/users/{id}/status/confirmed 
(only if the status is pending_confirmation)
(moves to status pending_payment)

PATCH /reservation/{adventure_id}/{type}/users/{id}/status/paid 
(only if the status is pending_payment)
(moves to status reserved)

PATCH /reservation/{adventure_id}/{type}/users/{id}/status/cancelled 
(masks as deleted. "get" will not return it)
(increases available capacity to 1)
```

## Tickets

```HTTP
POST /tickets/{adventure_id}/{type}/{user_id}
(calls payment processing partner to validate if the cc was charged using proc_id )
{
    "proc_id"
}

{
    "id":"11111111111",
    "human_readable_id":"111111111",
    "status"[paid|error_partner]
}
(creates a record in the database ("reservation_id","human_readable_id") with expiration_time)
```

```HTTP
GET /tickets/human/{human_readable_id}
{
    "id":"11111111111",
    "human_readable_id":"111111111"
}
```

## preferences

```HTTP
POST /users/{user_id}/preferences
{
    "channels":[
        {
            "id":"aaaaaaaa"
            "channel":"[email,phone_sms,phone_call]",
            "value":"email_address|phone_number",
            "type":"[tickets|profile]"
        }
    ]
}
```

```HTTP
GET /users/{user_id}/preferences/channels?type=[tickets|profile]
{
    {
        "id":"aaaaaaaa"
        "channel":"[email,phone_sms,phone_call]",
        "value":"email_address|phone_number",
        "type":"[tickets|profile]"
    }
}
```

## Messaging

```HTTP
POST /messages/{user_id}
{
    "type":["reservation","devolution","expiration"],
    "channel":{
        "value":"email_address|phone_number"
    },
    "data":{
        "ticket_id":"aaaaa",
        "recipient":{
            "fullname":"aaaaaa"
        }
    }
}

{
    "id":""
}
```
```HTTP
GET /messages/{user_id}
{
    [
        {
            "id":"",
            "status":"[pending,sent,error]
        }
    ]
}
```