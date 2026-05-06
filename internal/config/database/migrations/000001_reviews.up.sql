CREATE TABLE IF NOT EXISTS reviews(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    review TEXT,
    author_name VARCHAR(255),
    url_photo VARCHAR(255)
)