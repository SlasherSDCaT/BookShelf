CREATE TABLE IF NOT EXISTS users(
    id                  SERIAL PRIMARY KEY  NOT NULL UNIQUE,
    username            VARCHAR(255)        NOT NULL UNIQUE,
    user_role           VARCHAR(255)        NOT NULL,
    password_hash       VARCHAR(255)        NOT NULL
);

CREATE TABLE  IF NOT EXISTS collections (
    collection_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    rating FLOAT CHECK (rating >= 0 AND rating <= 5),
    public BOOLEAN DEFAULT FALSE
);

CREATE TABLE  IF NOT EXISTS books (
    book_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    title VARCHAR(200) NOT NULL,
    author VARCHAR(200) NOT NULL,
    genre VARCHAR(100),
    image TEXT,
    description TEXT,
    body TEXT
);

CREATE TABLE  IF NOT EXISTS comments (
    comment_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    book_id INT REFERENCES books(book_id),
    rating INT CHECK (rating >= 0 AND rating <= 5),
    text TEXT
);

-- Пример добавления книги в коллекцию (требует промежуточной таблицы для связи многие ко многим)
CREATE TABLE  IF NOT EXISTS collections_books (
    collection_id INT REFERENCES collections(collection_id),
    book_id INT REFERENCES books(book_id),
    PRIMARY KEY (collection_id, book_id)
);

CREATE TABLE IF NOT EXISTS collection_ratings (
    rating_id SERIAL PRIMARY KEY,
    collection_id INT REFERENCES collections(collection_id),
    rating INT CHECK (rating >= 1 AND rating <= 5)
);


