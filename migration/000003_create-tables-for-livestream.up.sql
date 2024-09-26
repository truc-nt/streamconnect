CREATE TABLE IF NOT EXISTS livestream (
    id_livestream BIGSERIAL PRIMARY KEY,
    fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR,
    status VARCHAR NOT NULL CHECK (status in ('created', 'started', 'played', 'ended')),
    meeting_id VARCHAR NOT NULL,
    hls_url VARCHAR,
    start_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
    priority INTEGER NOT NULL,
    is_livestreamed BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE TABLE IF NOT EXISTS livestream_ext_variant (
    id_livestream_ext_variant BIGSERIAL PRIMARY KEY,
    fk_livestream_product BIGSERIAL REFERENCES livestream_product(id_livestream_product) NOT NULL,
    fk_ext_variant BIGSERIAL REFERENCES ext_variant(id_ext_variant) NOT NULL,
    quantity INTEGER NOT NULL,
    CONSTRAINT unique_livestream_ext_variant UNIQUE (fk_livestream_product, fk_ext_variant)
);