CREATE TABLE IF NOT EXISTS livestream (
    id_livestream BIGSERIAL PRIMARY KEY,
    fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL,
    title varchar NOT NULL,
    description varchar,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS livestream_schedule (
    id_livestream_schedule SERIAL PRIMARY KEY,
    fk_livestream BIGSERIAL REFERENCES livestream(id_livestream) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP
);

CREATE TABLE IF NOT EXISTS livestream_product (
    id_livestream_product BIGSERIAL PRIMARY KEY,
    fk_livestream BIGSERIAL REFERENCES livestream(id_livestream) NOT NULL,
    fk_product BIGSERIAL REFERENCES product(id_product) NOT NULL,
    priority INTEGER 
);

CREATE TABLE IF NOT EXISTS livestream_external_variant (
    id_livestream_external_variant BIGSERIAL PRIMARY KEY,
    fk_livestream_product BIGSERIAL REFERENCES livestream_product(id_livestream_product) NOT NULL,
    fk_external_variant BIGSERIAL NOT NULL,
    quantity INTEGER NOT NULL
);