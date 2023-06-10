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
CREATE TYPE capacity (
    type text,
    max_people int
);

CREATE TYPE availability (
    type text,
    availability int
);


CREATE TABLE reservations (
    adventure_id text,
    status text,
    capacities map<text, frozen<capacity>>,
    availabilities map<text, frozen<availability>>,
    PRIMARY KEY(adventure_id)
);

INSERT INTO reservations (adventure_id,status,capacities) VALUES ('A1','CLOSED',
{'T1':{type:'T1',max_people:100,availability:100},'T2':{type:'T2',max_people:200,availability:200},'T3':{type:'T3',max_people:200,availability:200}});

SELECT * FROM reservations WHERE adventure_id='A1';

UPDATE reservations SET status='OPENED' WHERE adventure_id='A1' AND type=;

CREATE TABLE reservations_user (
    user_id uuid,
    adventure_id varchar,
    type varchar,
    quantity smallint,
    status varchar,
    expiration duration,
    PRIMARY KEY(user_id, adventure_id)
);
```
