CREATE TABLE "user" (
  id_user SERIAL PRIMARY KEY
);

CREATE TABLE user_shopify_auth (
  id_user_shopify_auth SERIAL PRIMARY KEY,
  fk_user INTEGER REFERENCES "user"(id_user) NOT NULL UNIQUE,
  shop_name TEXT NOT NULL,
  client_id TEXT NOT NULL,
  client_secret TEXT NOT NULL,
  access_token TEXT DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);