SET CHARSET UTF8;

DROP DATABASE IF EXISTS go_sample;

DROP DATABASE IF EXISTS development;
CREATE DATABASE IF NOT EXISTS development DEFAULT CHARACTER SET utf8;

DROP TABLE IF EXISTS lists;

CREATE TABLE lists (
  id   integer primary key auto_increment,
  text text,
  created_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp not null DEFAULT CURRENT_TIMESTAMP
);
