CREATE TABLE IF NOT EXISTS notification (
  id_notification BIGSERIAL primary key,
  fk_user         BIGINT REFERENCES "user" (id_user) NOT NULL,
  title           TEXT                               NOT NULL,
  message         TEXT                               NOT NULL,
  type            TEXT                               NOT NULL,
  status          TEXT                               NOT NULL default 'NEW',
  redirect_url    TEXT,
  created_at      TIMESTAMP                          NOT NULL default now()
);

CREATE TABLE IF NOT EXISTS livestream_product_follower (
  fk_user                      BIGINT REFERENCES "user" (id_user)                                             NOT NULL,
  fk_livestream_product        BIGINT REFERENCES livestream_product (id_livestream_product) on delete cascade NOT NULL,
  created_at                   TIMESTAMP                                                                      NOT NULL default now(),
  primary key (fk_livestream_product, fk_user)
);