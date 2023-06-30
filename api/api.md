User web site:

- get user and created
This is the first API after getting a JWT Token from Auth0

JWT + /users
{
    check if the user exists, if successful: {
        return user data
    } if error {

        create a user(email + ...) using data api
        return user data
    }
}

- get all open adventures
JWT + /adventures

start buying an adventure
JWT + /adventures adventures_id + type + quantity
{
    makes a reservation
}

confirm tickets 
{
    front end validates the credit card, if successful: {
        change status of the reservation to pay 
            and get tickets numbers
    } if error {
        change status of the reservation to cancelled
    }
}

cancel tickets {
    call reservation API to cancel the tickets
}

Partner web site:
get user if created {
    the user will be created as partner 
}

get all partners adventures
JWT + /adventures
{
    call data API with user id
}

create a new adventure
JWT + POST /adventures
{
    call data API with user id
}

publish an adventure
JWT + POST /adventures
{
    call data API with user id + adventure id
}


close an adventure
JWT + POST /adventures
{
    call data API with user id + adventure id
}