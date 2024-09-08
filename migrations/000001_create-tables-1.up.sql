CREATE TABLE IF NOT EXISTS "user" (
  id_user                BIGSERIAL PRIMARY KEY,
  username               varchar(100) not null unique,
  hashed_password        varchar(200) not null,
  full_name              varchar(100),
  email                  varchar(100) not null unique,
  is_enabled             boolean,
  created_date_time      TIMESTAMP,
  last_updated_date_time TIMESTAMP
);

create table if not exists user_roles
(
  user_role_id SERIAL PRIMARY KEY,
  username     varchar(100) NOT NULL,
  role         varchar(20)  NOT NULL,
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

CREATE TABLE IF NOT EXISTS external_shop (
  id_external_shop BIGSERIAL PRIMARY KEY,
  fk_shop BIGSERIAL REFERENCES shop(id_shop) NOT NULL,
  fk_ecommerce SMALLSERIAL REFERENCES ecommerce(id_ecommerce) NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  constraint check_status check (status in ('active', 'inactive'))
);

CREATE TABLE IF NOT EXISTS external_shop_shopify_auth (
  id_external_shop_shopify_auth BIGSERIAL PRIMARY KEY,
  fk_external_shop BIGSERIAL REFERENCES external_shop(id_external_shop) NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE,
  access_token TEXT DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);