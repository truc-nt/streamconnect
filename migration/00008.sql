create table if not exists notification
(
  id_notification bigserial primary key,
  fk_user         bigint REFERENCES "user" (id_user) NOT NULL,
  title           text                               not null,
  message         text                               not null,
  type            text                               not null,
  status          text                               not null default 'NEW',
  redirect_url    text,
  created_at      timestamp                          not null default now()
);

create table if not exists livestream_product_follow
(
  fk_user                      bigint REFERENCES "user" (id_user)                                             NOT NULL,
  fk_livestream_product        bigint REFERENCES livestream_product (id_livestream_product) on delete cascade NOT NULL,
  created_at                   timestamp                                                                      not null default now(),
  primary key (fk_livestream_product, fk_user)
);