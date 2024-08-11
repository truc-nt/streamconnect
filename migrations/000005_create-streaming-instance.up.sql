create table livestream (
    id serial primary key,
    title varchar not null,
    description varchar,
    owner_id int not null references account(id),
    status varchar not null,
    meeting_id varchar not null,
    hls_url varchar,
    start_time timestamp,
    end_time timestamp,
    created_at timestamp not null
);