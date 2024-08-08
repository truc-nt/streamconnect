create table account
(
    id                     serial primary key,
    username               varchar(100) not null unique,
    hashed_password        varchar(200) not null,
    full_name              varchar(100),
    email                  varchar(100) not null unique,
    is_enabled             boolean,
    created_date_time      TIMESTAMP,
    last_updated_date_time TIMESTAMP
);

create table account_roles
(
    user_role_id SERIAL PRIMARY KEY,
    username     varchar(100) NOT NULL,
    role         varchar(20)  NOT NULL,
    UNIQUE (username, role),
    FOREIGN KEY (username) REFERENCES account (username)
);