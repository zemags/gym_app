create table users (
  id serial not null unique,
  username varchar(50) not null,
  email varchar(100) not null unique,
  password_hash varchar(255) not null,
  admin bool,
  blocked bool,
  created_at timestamp not null default now(),
  updated_at timestamp
);