CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    userName TEXT,
    firstName TEXT,
    lastName TEXT,
    createdAt TIMESTAMP NOT NULL DEFAULT now(),
    lastSeenAt TIMESTAMP NOT NULL DEFAULT now()
);