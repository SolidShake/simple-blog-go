CREATE TABLE IF NOT EXISTS post (
    post_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title varchar(255) NOT NULL,
    content text,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS tag (
    tag_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS post_tag (
    post_id int NOT NULL,
    tag_id int NOT NULL,
    FOREIGN KEY (post_id) REFERENCES post(post_id),
    FOREIGN KEY (tag_id) REFERENCES tag(tag_id),
    UNIQUE (post_id, tag_id)
);