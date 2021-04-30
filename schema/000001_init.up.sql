create table users (
  id integer unique primary key autoincrement,
  name varchar(50) not null,
  email varchar(100) not null unique,
  password_hash varchar(255) not null
);