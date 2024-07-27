CREATE TABLE "user" (
  id_user SERIAL PRIMARY KEY
);

CREATE TABLE shop (
  id_shop SERIAL PRIMARY KEY, 
  fk_user INTEGER REFERENCES "user"(id_user) NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ecommerce (
  id_ecommerce SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE external_shop (
  id_external_shop SERIAL PRIMARY KEY,
  fk_shop INTEGER REFERENCES shop(id_shop) NULL,
  fk_ecommerce INTEGER REFERENCES ecommerce(id_ecommerce) NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  constraint check_status check (status in ('active', 'inactive'))
);

CREATE TABLE shopify_external_shop_auth (
  id_shopify_external_shop_auth SERIAL PRIMARY KEY,
  fk_external_shop INTEGER REFERENCES external_shop(id_external_shop) NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE,
  access_token TEXT DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);