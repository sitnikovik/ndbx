CREATE KEYSPACE IF NOT EXISTS iot 
WITH REPLICATION = { 
   'class' : 'NetworkTopologyStrategy', 
   'dc1' : 3
};

CREATE TABLE IF NOT EXISTS iot.sensor_data (
    sensor_id text,
    created_at timestamp,
    value float,
    PRIMARY KEY ((sensor_id), created_at)
) WITH CLUSTERING ORDER BY (created_at DESC);