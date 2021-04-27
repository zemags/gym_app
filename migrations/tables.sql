create table users (
  id integer unique primary key autoincrement,
  login varchar(50) not null,
  email varchar(100) not null UNIQUE,
  passwd text
);
create table body (
  id integer unique primary key autoincrement,
  dt datetime default CURRENT_TIMESTAMP,
  height FLOAT,
  weight FLOAT,
  body FLOAT,
  fat_percent FLOAT,
  waist FLOAT,
  arms FLOAT,
  forearm FLOAT,
  chest FLOAT,
  neck FLOAT,
  thighs FLOAT,
  thigh FLOAT,
  calf FLOAT,
  foreign key (user_id) REFERENCES users (id)
);
create table exercise_category (
  id integer unique primary key autoincrement,
  name varchar(50)
);
create table exercises (
  id integer unique primary key autoincrement,
  name varchar(50) not null UNIQUE,
  foreign key (exercise_category_id) REFERENCES exercise_category (id)
)