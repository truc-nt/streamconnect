CREATE TABLE IF NOT EXISTS cart (
    id_cart BIGSERIAL PRIMARY KEY,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL
);

CREATE TABLE IF NOT EXISTS cart_item (
    id_cart_item BIGSERIAL PRIMARY KEY,
    fk_cart BIGSERIAL REFERENCES cart(id_cart) NOT NULL,
    fk_ext_variant BIGSERIAL REFERENCES ext_variant(id_ext_variant) NOT NULL,
    quantity INT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    CONSTRAINT check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS cart_item_livestream_ext_variant (
    id BIGSERIAL PRIMARY KEY,
    fk_cart_item BIGSERIAL REFERENCES cart_item(id_cart_item) NOT NULL,
    fk_livestream_ext_variant BIGSERIAL REFERENCES livestream_ext_variant(id_livestream_ext_variant) NOT NULL
);
