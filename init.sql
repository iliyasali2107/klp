CREATE USER user;
CREATE DATABASE klp;
GRANT ALL PRIVILEGES ON DATABASE klp to user;

CREATE TABLE IF NOT EXISTS sport_cfs  (
    sport varchar(100),
    cfs FLOAT
);