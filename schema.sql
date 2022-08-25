CREATE TABLE products(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    guid VARCHAR(55) UNIQUE NOT NULL,
    name VARCHAR(255) not NULL UNIQUE,
    price REAL NOT NULL,
    description TEXT,
    createdAt TEXT NOT NULL,
);