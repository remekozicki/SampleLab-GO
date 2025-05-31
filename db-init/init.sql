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

-- assortment table
CREATE SEQUENCE IF NOT EXISTS assortment_id_seq;
ALTER TABLE assortment ALTER COLUMN id SET DEFAULT nextval('assortment_id_seq');
SELECT setval('assortment_id_seq', COALESCE((SELECT MAX(id) FROM assortment), 1));

ALTER TABLE assortment_indications
DROP CONSTRAINT fk7aiabke38hkwv0ai7bmvfqtcu;

ALTER TABLE assortment_indications
    ADD CONSTRAINT fk7aiabke38hkwv0ai7bmvfqtcu FOREIGN KEY (assortment_id)
        REFERENCES assortment(id) ON DELETE CASCADE;


