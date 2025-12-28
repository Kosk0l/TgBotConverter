CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    userName TEXT,
    firstName TEXT,
    lastName TEXT,
    createdAt TIMESTAMP NOT NULL,
    lastSeenAt TIMESTAMP NOT NULL
);