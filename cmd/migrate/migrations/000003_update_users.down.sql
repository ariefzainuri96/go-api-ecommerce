ALTER TABLE users
ALTER COLUMN id TYPE bigserial USING id::bigserial;

-- Remove added columns
ALTER TABLE users
DROP COLUMN IF EXISTS name,
DROP COLUMN IF EXISTS age,
DROP COLUMN IF EXISTS is_admin;

-- Add old columns back
ALTER TABLE users
ADD COLUMN email varchar(255) UNIQUE NOT NULL,
ADD COLUMN password bytea NOT NULL,
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();