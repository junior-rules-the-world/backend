create table users
(
    id           serial
        primary key,
    display_name varchar,
    username     varchar               not null
        unique,
    email        varchar               not null
        unique,
    password     varchar               not null,
    role         varchar default user,
    team_id      int                   not null,
    created_at   date    default now() not null,
    updated_at   date    default now() not null
);

