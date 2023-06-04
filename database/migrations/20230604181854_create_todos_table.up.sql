CREATE TABLE todos (
  id BINARY(36),
  
  title VARCHAR(200) NOT NULL DEFAULT '',
  completed_at datetime(3) NULL,
  user_id BINARY(36) NOT NULL,

  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  deleted_at datetime(3) NULL,
  
  PRIMARY KEY (id),
  KEY idx_users_deleted_at (deleted_at),
  KEY idx_users_completed_at (completed_at),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;