CREATE TABLE IF NOT EXISTS user_address (
    id_user_address BIGSERIAL PRIMARY KEY,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    is_default BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE TABLE IF NOT EXISTS "order" (
    id_order BIGSERIAL PRIMARY KEY,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL,
    fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL,
    fk_shipping_method SMALLSERIAL REFERENCES shipping_method(id_shipping_method) NOT NULL,
    fk_payment_method SMALLSERIAL REFERENCES payment_method(id_payment_method) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ext_order (
    id_ext_order BIGSERIAL PRIMARY KEY,
    fk_order BIGSERIAL REFERENCES "order"(id_order) NOT NULL,
    fk_ext_shop BIGSERIAL REFERENCES ext_shop(id_ext_shop) NOT NULL,
    ext_order_id_mapping TEXT NOT NULL,
    shipping_fee DECIMAL(10, 2) NOT NULL,
    shipping_fee_discount DECIMAL(10, 2) NOT NULL,
    internal_discount DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    external_discount DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ext_order_voucher (
    id_order_voucher BIGSERIAL PRIMARY KEY,
    fk_ext_order BIGSERIAL REFERENCES ext_order(id_ext_order) NOT NULL,
    fk_voucher BIGSERIAL REFERENCES voucher(id_voucher) NOT NULL
);

CREATE TABLE IF NOT EXISTS order_item (
    id_order_item BIGSERIAL PRIMARY KEY,
    fk_order BIGSERIAL REFERENCES "order"(id_order) NOT NULL,
    fk_ext_variant BIGSERIAL REFERENCES ext_variant(id_ext_variant) NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    paid_price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS order_item_livestream_ext_variant (
    id BIGSERIAL PRIMARY KEY,
    fk_order_item BIGSERIAL REFERENCES order_item(id_order_item) NOT NULL,
    fk_livestream_ext_variant BIGSERIAL REFERENCES livestream_ext_variant(id_livestream_ext_variant) NOT NULL
);

CREATE TABLE IF NOT EXISTS voucher (
    id_voucher BIGSERIAL PRIMARY KEY,
    fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL,
    code TEXT UNIQUE NOT NULL,
    discount DECIMAL(10, 2) NOT NULL,
    max_discount DECIMAL(10, 2),
    method TEXT NOT NULL,
    type TEXT NOT NULL,
    target TEXT NOT NULL,
    quantity INT NOT NULL,
    min_purchase DECIMAL(10, 2) DEFAULT 0.00 NOT NULL,
    start_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_method check (method in ('each', 'across')),
    CONSTRAINT check_type check (type in ('percentage', 'fixed')),
    CONSTRAINT check_target check (target in ('item', 'shipping'))
);

CREATE TABLE IF NOT EXISTS voucher_user (
    id_voucher_user BIGSERIAL PRIMARY KEY,
    fk_voucher BIGSERIAL REFERENCES voucher(id_voucher) NOT NULL,
    fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL,
    CONSTRAINT unique_voucher_user UNIQUE(fk_voucher, fk_user)
);

CREATE TABLE IF NOT EXISTS shipping_method (
    id_shipping_method SMALLSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    CONSTRAINT unique_shipping_method UNIQUE(name)
)

CREATE TABLE IF NOT EXISTS payment_method (
    id_payment_method SMALLSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    CONSTRAINT unique_payment_method UNIQUE(name)
)

CREATE TABLE IF NOT EXISTS order_address (
    id_order_address BIGSERIAL PRIMARY KEY,
    fk_order BIGSERIAL REFERENCES "order"(id_order) NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL
)

INSERT INTO shipping_method (name) VALUES ('standard')
INSERT INTO payment_method (name) VALUES ('cash on delivery')

