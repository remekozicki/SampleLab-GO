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
-- Assortment table
CREATE SEQUENCE IF NOT EXISTS assortment_id_seq;
ALTER TABLE assortment ALTER COLUMN id SET DEFAULT nextval('assortment_id_seq');
SELECT setval('assortment_id_seq', COALESCE((SELECT MAX(id) FROM assortment), 1));

-- Product Group table
CREATE SEQUENCE product_group_id_seq;
ALTER TABLE product_group ALTER COLUMN id SET DEFAULT nextval('product_group_id_seq');
SELECT setval('product_group_id_seq', COALESCE((SELECT MAX(id) FROM product_group), 1));

-- Sampling Standard table
CREATE SEQUENCE sampling_standard_id_seq;
ALTER TABLE sampling_standard ALTER COLUMN id SET DEFAULT nextval('sampling_standard_id_seq');
SELECT setval('sampling_standard_id_seq', COALESCE((SELECT MAX(id) FROM sampling_standard), 1));

-- Inspection table
CREATE SEQUENCE inspection_id_seq;
ALTER TABLE inspection ALTER COLUMN id SET DEFAULT nextval('inspection_id_seq');
SELECT setval('inspection_id_seq', COALESCE((SELECT MAX(id) FROM inspection), 1));

-- Examination table
CREATE SEQUENCE IF NOT EXISTS examination_id_seq;
ALTER TABLE examination ALTER COLUMN id SET DEFAULT nextval('examination_id_seq');
SELECT setval('examination_id_seq', (SELECT COALESCE(MAX(id), 0) FROM examination));

-- Sample table
CREATE SEQUENCE IF NOT EXISTS sample_id_seq;
ALTER TABLE sample ALTER COLUMN id SET DEFAULT nextval('sample_id_seq');
SELECT setval('sample_id_seq', (SELECT COALESCE(MAX(id), 0) FROM sample));


-- auto delete assortment_indications
ALTER TABLE assortment_indications
DROP CONSTRAINT fk7aiabke38hkwv0ai7bmvfqtcu;

ALTER TABLE assortment_indications
    ADD CONSTRAINT fk7aiabke38hkwv0ai7bmvfqtcu FOREIGN KEY (assortment_id)
        REFERENCES assortment(id) ON DELETE CASCADE;

-- auto delete assortment
ALTER TABLE assortment
DROP CONSTRAINT fkltruhwpio7er8587wpi9tjyd2;

ALTER TABLE assortment
    ADD CONSTRAINT fkltruhwpio7er8587wpi9tjyd2 FOREIGN KEY (group_id)
        REFERENCES product_group(id) ON DELETE CASCADE;

-- auto delete product_group_sampling_standards
ALTER TABLE product_group_sampling_standards DROP CONSTRAINT fk_product_group;

ALTER TABLE product_group_sampling_standards
    ADD CONSTRAINT fk_product_group FOREIGN KEY (groups_id)
        REFERENCES product_group(id) ON DELETE CASCADE;