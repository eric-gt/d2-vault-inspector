-- create "users" table
CREATE TABLE users (
    id bigserial,
    bungie_id integer NOT NULL,
    bungie_name varchar(100) NOT NULL,
    PRIMARY KEY (id)
)
