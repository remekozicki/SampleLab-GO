-- Address table
CREATE SEQUENCE IF NOT EXISTS address_id_seq;
ALTER TABLE address ALTER COLUMN id SET DEFAULT nextval('address_id_seq');
SELECT setval('address_id_seq', (SELECT COALESCE(MAX(id), 0) FROM address));

-- Client table
CREATE SEQUENCE IF NOT EXISTS client_id_seq;
ALTER TABLE client ALTER COLUMN id SET DEFAULT nextval('client_id_seq');
SELECT setval('client_id_seq', (SELECT COALESCE(MAX(id), 0) FROM client));

-- Users table
CREATE SEQUENCE IF NOT EXISTS users_id_seq;
ALTER TABLE users ALTER COLUMN id SET DEFAULT nextval('users_id_seq');
SELECT setval('users_id_seq', (SELECT COALESCE(MAX(id), 0) FROM users));


-- Indication table
CREATE SEQUENCE IF NOT EXISTS indication_id_seq;
ALTER TABLE indication ALTER COLUMN id SET DEFAULT nextval('indication_id_seq');
SELECT setval('indication_id_seq', COALESCE((SELECT MAX(id) FROM indication), 1), false);

-- (Dodaj inne tabele analogicznie...)