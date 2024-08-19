CREATE TABLE IF NOT EXISTS cart (
    id_cart BIGSERIAL PRIMARY KEY,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL
);

CREATE TABLE IF NOT EXISTS cart_livestream_external_variant (
    id_cart_livestream_external_variant BIGSERIAL PRIMARY KEY,
    fk_cart BIGSERIAL REFERENCES cart(id_cart) NOT NULL,
    fk_livestream_external_variant BIGSERIAL REFERENCES livestream_external_variant(id_livestream_external_variant) NOT NULL,
    quantity INT NOT NULL
);