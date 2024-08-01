-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS savannah_tech;

-- remove strict mode to avoid invalid time format on "0000-00-00 00:00:00"
SET SQL_MODE = "";

-- select savannah_tech as db
USE savannah_tech;

create table i_repository
(
    id                    int auto_increment
        primary key,
    created_at            datetime      default CURRENT_TIMESTAMP     not null,
    repository_created_at datetime      default '0000-00-00 00:00:00' null,
    repository_updated_at datetime      default '0000-00-00 00:00:00' null,
    name                  varchar(500)  default ''                    not null,
    owner                 varchar(100)  default ''                    not null,
    description           varchar(5000) default ''                    null,
    url                   varchar(5000) default ''                    null,
    language              varchar(100)  default ''                    null,
    forks_count           int           default 0                     null,
    stars_count           int           default 0                     null,
    open_issues_count     int           default 0                     null,
    watchers_count        int           default 0                     null,
    constraint i_repository_pk
        unique (name)
);

create table i_commit
(
    id              int auto_increment
        primary key,
    created_at      datetime      default CURRENT_TIMESTAMP     not null,
    i_repository_id int                                         not null,
    author_name     varchar(100)  default ''                    not null,
    author_email    varchar(100)  default ''                    not null,
    author_login    varchar(100)  default ''                    null,
    date            datetime      default '0000-00-00 00:00:00' not null,
    message         varchar(5000) default ''                    not null,
    url             varchar(5000) default ''                    not null,
    constraint i_commit_i_repository_id_fk
        foreign key (i_repository_id) references i_repository (id)
            on update cascade on delete cascade
);
