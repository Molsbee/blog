CREATE TABlE IF NOT EXISTS articles(
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL,
    published TINYINT NOT NULL DEFAULT 0,
    title VARCHAR(255),
    content VARCHAR(4000),
    author VARCHAR(50)
) ENGINE=INNODB;