CREATE TABLE users (
  id BINARY(36),
  name varchar(200) DEFAULT '' NOT NULL,
  email varchar(200) DEFAULT '' NOT NULL,
  password varchar(200) DEFAULT '' NOT NULL,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  deleted_at datetime(3) NULL,
  PRIMARY KEY (id),
  UNIQUE KEY idx_users_email (email),
  KEY idx_users_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;