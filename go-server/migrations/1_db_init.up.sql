

create table if not exists users (
  id serial primary key,
  username varchar(128) not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  deleted_at timestamp
);
