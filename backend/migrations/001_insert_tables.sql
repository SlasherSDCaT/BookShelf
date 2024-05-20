
INSERT INTO users (username, user_role, password_hash) VALUES
('Slasher', 'ADMIN', '1234'),
('Jane', 'user', '1234'),
('Alice', 'user', '1234');

INSERT INTO collections (user_id, name, description, rating) VALUES
(1, 'Fiction Favorites', 'A collection of my favorite fiction books', 4),
(2, 'Science Reads', 'Interesting books about science', 4),
(3, 'Fantasy Collection', 'A collection of great fantasy novels', 5),
(1, 'Favorites of 2023', 'My favorite books read in 2023', 4);

INSERT INTO books (user_id, title, author, genre, description, body) VALUES
(1, '1984', 'George Orwell', 'Dystopian', 'A dystopian novel set in Airstrip One', 'Full text of 1984...'),
(2, 'A Brief History of Time', 'Stephen Hawking', 'Science', 'An overview of cosmology', 'Full text of A Brief History of Time...'),
(3, 'The Hobbit', 'J.R.R. Tolkien', 'Fantasy', 'A fantasy novel about the adventures of Bilbo Baggins', 'Full text of The Hobbit...'),
(2, 'Sapiens: A Brief History of Humankind', 'Yuval Noah Harari', 'History', 'An exploration of the history of Homo sapiens', 'Full text of Sapiens...'),
(3, 'The Lord of the Rings', 'J.R.R. Tolkien', 'Fantasy', 'Epic fantasy novel', 'Full text of The Lord of the Rings...');

INSERT INTO comments (user_id, book_id, rating, text) VALUES
(1, 1, 5, 'An eye-opening book!'),
(2, 2, 4, 'Very informative and well-written'),
(3, 3, 5, 'An absolute classic in fantasy literature'),
(1, 2, 5, 'A fascinating perspective on human history'),
(2, 3, 5, 'One of the best fantasy series ever written');

INSERT INTO collections_books (collection_id, book_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(1, 2),
(1, 3);

INSERT INTO collection_ratings (collection_id, rating) VALUES
(1, 5),
(2, 4),
(3, 5),
(4, 1);