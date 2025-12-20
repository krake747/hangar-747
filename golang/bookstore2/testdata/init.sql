CREATE TABLE IF NOT EXISTS books (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    copies INTEGER NOT NULL DEFAULT 0
);

INSERT OR IGNORE INTO books (id, title, author, copies) VALUES
    ('abc', 'In the Company of Cheerful Ladies', 'Alexander McCall Smith', 0),
    ('def', 'White Heat', 'Dominic Sandbrook', 2);