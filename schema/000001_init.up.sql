create table users (
  id serial not null unique,
  name varchar(50) not null,
  email varchar(100) not null unique,
  password_hash varchar(255) not null
);