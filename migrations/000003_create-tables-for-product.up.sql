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

CREATE TABLE IF NOT EXISTS external_variant (
  id_external_variant BIGSERIAL PRIMARY KEY,
  fk_external_shop BIGSERIAL REFERENCES external_shop(id_external_shop) NOT NULL,
  fk_variant BIGINT REFERENCES variant(id_variant),
  external_product_id_mapping VARCHAR(255),
  external_id_mapping VARCHAR(255) NOT NULL,
  sku TEXT,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  option JSON,
  price DECIMAL(10, 2),
  image_url TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT unique_external_product_variant UNIQUE (external_product_id_mapping, external_id_mapping),
  CONSTRAINT check_status check (status in ('active', 'inactive'))
);