-- adding not null will fail otherwise
UPDATE users
SET password_hash = '\\xDEADBEEF'
WHERE password_hash IS NULL;

ALTER TABLE users
ALTER COLUMN password_hash SET NOT NULL;
