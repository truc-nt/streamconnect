CREATE TABLE IF NOT EXISTS user_address (
    id_user_address BIGSERIAL PRIMARY KEY,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT,
    is_default BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "order" (
    id_order BIGSERIAL PRIMARY KEY,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS external_order (
    id_external_order BIGSERIAL PRIMARY KEY,
    fk_order BIGSERIAL REFERENCES "order"(id_order) NOT NULL,
    external_order_id_mapping TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_item (
    id_order_item BIGSERIAL PRIMARY KEY,
    fk_order BIGSERIAL REFERENCES "order"(id_order) NOT NULL,
    fk_external_variant BIGSERIAL REFERENCES external_variant(id_external_variant) NOT NULL,
    quantity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS order_item_livestream_external_variant (
    id BIGSERIAL PRIMARY KEY,
    fk_order_item BIGSERIAL REFERENCES order_item(id_order_item) NOT NULL,
    fk BIGSERIAL REFERENCES livestream_external_variant(id_livestream_external_variant) NOT NULL
);