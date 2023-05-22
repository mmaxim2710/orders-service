CREATE TABLE users (
    uuid uuid not null primary key default uuid_generate_v4(),
    login varchar(255) not null unique,
    email varchar(320) not null unique,
    first_name varchar(128) not null,
    last_name varchar(128) not null,
    encrypted_password varchar not null
)