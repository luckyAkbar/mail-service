-- +migrate Up notransaction
ALTER TABLE mailing_lists ADD COLUMN IF NOT EXISTS receipient_name TEXT NOT NULL;
ALTER TABLE mailing_lists RENAME COLUMN receipient TO receipient_email;

-- +migrate Down
ALTER TABLE mailing_lists DROP COLUMN IF EXISTS receipient_name;
ALTER TABLE mailing_lists RENAME COLUMN receipient_email to receipient;