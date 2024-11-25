create table if not exists projects
(
    id          bigint generated always as identity
        primary key,
    name        text not null,
    description text
);

create table if not exists ticket_statuses
(
    id     bigint generated always as identity
        primary key,
    status text not null
);

create table if not exists roles
(
    id          bigint generated always as identity
        primary key,
    name        text                                       not null,
    description text default 'Описание не заполнено'::text not null
);

create table if not exists users
(
    id       bigint generated always as identity
        primary key,
    login    text   not null
        unique,
    email    text   not null
        unique,
    password text   not null,
    role_id  bigint not null
        references roles
);

create table if not exists user_projects
(
    id         bigint generated always as identity
        primary key,
    user_id    bigint not null
        references users,
    project_id bigint not null
        references projects
);

create table if not exists support_tickets
(
    id          bigint generated always as identity
        primary key,
    project_id  bigint                                 not null
        references projects,
    status_id   bigint                   default 1     not null
        references ticket_statuses,
    title       text                                   not null,
    description text,
    created_at  timestamp with time zone default now() not null,
    updated_at  timestamp with time zone default now() not null,
    created_by  text                                   not null,
    answer      text,
    answered_by bigint
        constraint support_tickets_users_id_fk
            references users
);

create table if not exists sessions
(
    id         bigint generated always as identity
        primary key,
    user_id    bigint                                                        not null
        references users,
    token      text                                                          not null,
    created_at timestamp with time zone default now()                        not null,
    expires_at timestamp with time zone default (now() + '7 days'::interval) not null
);

