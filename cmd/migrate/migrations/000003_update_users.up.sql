-- Drop old columns if they don't match the new data
ALTER TABLE users
DROP COLUMN IF EXISTS email,
DROP COLUMN IF EXISTS password,
DROP COLUMN IF EXISTS created_at;

-- Add new columns if they don't exist
ALTER TABLE users
ADD COLUMN IF NOT EXISTS name TEXT,
ADD COLUMN IF NOT EXISTS age INTEGER,
ADD COLUMN IF NOT EXISTS is_admin BOOLEAN;