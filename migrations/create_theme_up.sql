CREATE TABLE theme (
  id bigserial not null primary key,
  user_id INTEGER not null,
  title varchar not null unique,
  description varchar not null
);