-- +migrate Up
CREATE TABLE IF NOT EXISTS `follows` (
  id INT AUTO_INCREMENT NOT NULL,
  follower_id INT NOT NULL,
  followed_id INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (followed_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE IF EXISTS `tasks`;