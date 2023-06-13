# 

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
    
);

CREATE TYPE availability (
    type text,
    
);


CREATE TABLE reservations (
    adventure_id varchar(100),
    status varchar(100),
    PRIMARY KEY(adventure_id)
);

CREATE TABLE reservation_capacities{
    adventure_id varchar(100),
    type varchar(100),
    availability int,
    max_people int
    PRIMARY KEY(adventure_id, type),
    FOREIGN KEY (adventure_id) REFERENCES reservations(adventure_id)
}

INSERT INTO reservations (adventure_id,status,capacities) VALUES ('A1','CLOSED',
{'T1':{type:'T1',max_people:100,availability:100},'T2':{type:'T2',max_people:200,availability:200},'T3':{type:'T3',max_people:200,availability:200}});

SELECT * FROM reservations WHERE adventure_id='A1';

UPDATE reservations SET status='OPENED' WHERE adventure_id='A1' AND type=;

```
