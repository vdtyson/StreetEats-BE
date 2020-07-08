begin;

create type user_type as enum ('WORKER','CUSTOMER');
create table se_users
(
    uid         bigserial,
    user_type   user_type   not null,
    firstName   varchar(50) not null,
    lastName    varchar(50) not null,
    phone       varchar(50),
    email       varchar(255),
    pw_hash     varchar,
    created_at  timestamp default (now()),
    last_update timestamp default (now()),
    primary key (uid)
);

create table customers
(
    username varchar(64) unique not null,
    primary key (uid)
) inherits (se_users);

end;