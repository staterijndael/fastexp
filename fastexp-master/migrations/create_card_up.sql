CREATE TABLE cards (
  id bigserial not null primary key,
  name varchar not null unique,
  shortdesc varchar not null,
  fulldesc varchar not null
);