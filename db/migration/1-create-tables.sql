-- +migrate Up notransaction
CREATE TYPE "status" AS ENUM ('PENDING', 'SENT', 'FAILED');
CREATE TYPE "priority" AS ENUM ('LOW', 'HIGH');

CREATE TABLE IF NOT EXISTS mailing_lists (
    id BIGINT PRIMARY KEY,
    subject TEXT NOT NULL,
    receipient TEXT NOT NULL,
    sender_email TEXT NOT NULL,
    sender_name TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    "status" status NOT NULL,
    "priority" priority NOT NULL,
    sent_at TIMESTAMP DEFAULT NULL,
    last_sending_attempt TIMESTAMP DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS mailing_lists;