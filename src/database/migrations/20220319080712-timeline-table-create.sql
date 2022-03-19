-- +migrate Up
CREATE TABLE IF NOT EXISTS `timelines` (
  id INT AUTO_INCREMENT NOT NULL,
  user_id INT NOT NULL,
  content VARCHAR(1024),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE IF EXISTS `timelines`;