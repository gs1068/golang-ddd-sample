-- +migrate Up
CREATE TABLE IF NOT EXISTS `likes` (
  id INT AUTO_INCREMENT NOT NULL,
  user_id INT NOT NULL,
  timeline_id INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (timeline_id) REFERENCES timelines(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE IF EXISTS `likes`;