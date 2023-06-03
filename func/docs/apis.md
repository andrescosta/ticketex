# Resources

## Reservations

```HTTP
POST/GET /reservation
{
    "id":"aaaaaaa"
    "adventure":{
        "id":"aaaaa"
    },
    "capacity"[
        {
            "type":"vip",
            "current":100,
            "max":1000
        }
    ]
}
```

PATCH /reservation/{id}
{
    "capacity"[
        {
            "type":"vip",
            "max":1000
        }
    ]
}

POST /reservation/{id}/users/{id}
{} // EMPTY BODY
(create with status pending_confirmation)

{
    "id":"aaaaaaaaaaa",
    "human_id":"aaaaaaaaaa",
    "status":pending_confirmation
    "expiration": date
    "user":{
        "id":"aaa"
    }
}

GET /reservation/{id}/users/{id}
{
    "id":"aaaaaaaaaaa",
    "status":[reserved,pending_confirmation,pending_payment]
    "expiration": [date|empty for reserved]
    "user":{
        "id":"aaa"
    }
}

PATCH /reservation/{id}/users/{id}/action/confirmed 
(only if the status is pending_confirmation)
(move to status pending_payment)
PATCH /reservation/{id}/users/{id}/action/paid 
(only if the status is pending_payment)
(move to status reserved)
PATCH /reservation/{id}/users/{id}/action/returned 
(will be marked as deleted. get will not return it)
(will increase available capacity to 1)

## Tickets

POST /tickets/reservation/{id}
(will validate with partner if the cc was charged using proc_id )
{
    "proc_id"
}

{
    "id":"11111111111",
    "human_readable_id":"111111111"
}
(will create a record in the database ("reservation_id","human_readable_id") with expiration_time)
GET /tickets/human/{human_readable_id}
{
    "id":"11111111111",
    "human_readable_id":"111111111"
}

## preferences

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
GET /users/{user_id}/preferences/channels?type=[tickets|profile]
{
    {
        "id":"aaaaaaaa"
        "channel":"[email,phone_sms,phone_call]",
        "value":"email_address|phone_number",
        "type":"[tickets|profile]"
    }
}


## Messaging

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

GET /messages/{user_id}
{
    [
        {
            "id":"",
            "status":"[pending,sent,error]
        }
    ]
}