CREATE TABLE IF NOT EXISTS cart (
    id_cart BIGSERIAL PRIMARY KEY,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL
);

INSERT INTO cart (fk_user) VALUES (1);

CREATE TABLE IF NOT EXISTS cart_item (
    id_cart_item BIGSERIAL PRIMARY KEY,
    fk_cart BIGSERIAL REFERENCES cart(id_cart) NOT NULL,
    fk_external_variant BIGSERIAL REFERENCES external_variant(id_external_variant) NOT NULL,
    quantity INT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    CONSTRAINT check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS cart_item_livestream_external_variant (
    id BIGSERIAL PRIMARY KEY,
    fk_cart_item BIGSERIAL REFERENCES cart_item(id_cart_item) NOT NULL,
    fk BIGSERIAL REFERENCES livestream_external_variant(id_livestream_external_variant) NOT NULL
);
