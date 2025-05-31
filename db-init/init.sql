CREATE SEQUENCE IF NOT EXISTS users_id_seq;

ALTER TABLE users
    ALTER COLUMN id SET DEFAULT nextval('users_id_seq');

SELECT setval('users_id_seq', COALESCE((SELECT MAX(id) FROM users), 1));


ALTER TABLE assortment_indications
DROP CONSTRAINT fk7aiabke38hkwv0ai7bmvfqtcu;

ALTER TABLE assortment_indications
    ADD CONSTRAINT fk7aiabke38hkwv0ai7bmvfqtcu FOREIGN KEY (assortment_id)
        REFERENCES assortment(id) ON DELETE CASCADE;

CREATE SEQUENCE IF NOT EXISTS assortment_id_seq;

ALTER TABLE assortment ALTER COLUMN id SET DEFAULT nextval('assortment_id_seq');
SELECT setval('assortment_id_seq', COALESCE((SELECT MAX(id) FROM assortment), 1));
