CREATE TABLE IF NOT EXISTS "user" (
  id_user BIGSERIAL PRIMARY KEY,
  username VARCHAR(100) NOT NULL UNIQUE,
  hashed_password VARCHAR(200) NOT NULL,
  email VARCHAR(100) NOT NULL unique,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS acl_role (
  id_acl_role SMALLSERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS acl_role_user (
  id_acl_role_user SMALLSERIAL PRIMARY KEY,
  fk_acl_role SMALLSERIAL REFERENCES acl_role(id_acl_role) NOT NULL,
  fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL,
  fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_role (
  id_user_role SERIAL PRIMARY KEY,
  username varchar(100) NOT NULL,
  role varchar(20) NOT NULL,
  UNIQUE (username, role),
  FOREIGN KEY (username) REFERENCES "user"(username)
);

CREATE TABLE IF NOT EXISTS shop (
  id_shop BIGSERIAL PRIMARY KEY, 
  fk_user BIGSERIAL REFERENCES "user"(id_user) NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ecommerce (
  id_ecommerce SMALLSERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS ext_shop (
  id_ext_shop BIGSERIAL PRIMARY KEY,
  fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL,
  fk_ecommerce SMALLSERIAL REFERENCES ecommerce(id_ecommerce) NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  constraint check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS ext_shop_shopify_auth (
  id_ext_shop_shopify_auth BIGSERIAL PRIMARY KEY,
  fk_ext_shop BIGSERIAL REFERENCES ext_shop(id_ext_shop) NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE,
  access_token TEXT DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);