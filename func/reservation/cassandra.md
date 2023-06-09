# Keyspace
```
CREATE KEYSPACE resv
  WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
  };
```

## Tables
```
CREATE TABLE reservations (
    adventure_id varchar, 
    type varchar,
    status varchar,
    max int,
    current int,
    PRIMARY KEY(adventure_id, type)
)

CREATE TABLE reservation_user (
    user_id uuid (PK),
    adventure_id varchar,
    type varchar,
    quantity shortint
    status varchar,
    expiration duration,
    PRIMARY KEY(user_id, type)
)
```