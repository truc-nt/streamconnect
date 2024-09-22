CREATE TABLE IF NOT EXISTS category (
  id_category SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product (
  id_product BIGSERIAL PRIMARY KEY,
  fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  CONSTRAINT check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS variant (
  id_variant BIGSERIAL PRIMARY KEY,
  fk_product BIGSERIAL REFERENCES product(id_product) NOT NULL,
  sku TEXT,
  status TEXT NOT NULL DEFAULT 'active',
  option JSON DEFAULT '{}',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS image_variant (
  id_image_variant BIGSERIAL PRIMARY KEY,
  fk_variant BIGSERIAL REFERENCES variant(id_variant) NOT NULL,
  url TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS ext_variant (
  id_ext_variant BIGSERIAL PRIMARY KEY,
  fk_ext_shop BIGSERIAL REFERENCES ext_shop(id_ext_shop) NOT NULL,
  fk_variant BIGINT REFERENCES variant(id_variant),
  ext_product_id_mapping VARCHAR(255),
  ext_id_mapping VARCHAR(255) NOT NULL,
  sku TEXT,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  option JSON,
  price DECIMAL(10, 2) NOT NULL,
  image_url TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT unique_ext_product_variant UNIQUE (ext_product_id_mapping, ext_id_mapping),
  CONSTRAINT check_status check (status in ('active', 'inactive'))
);