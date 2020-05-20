CREATE TABlE IF NOT EXISTS articles(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    published BOOLEAN NOT NULL DEFAULT FALSE,
    title VARCHAR(255),
    content VARCHAR(4000),
    author VARCHAR(50)
)