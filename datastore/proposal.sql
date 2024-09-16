-- connect database first and dump sql
DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF not EXISTS "uuid-ossp";

ALTER DATABASE s2 SET timezone TO 'Asia/Jakarta';

-- create schema type
drop schema if exists types cascade;
create schema types;
alter schema types owner to s2usr;
set search_path = types,public;

create type flag as enum (
    'Valid',
    'Invalid',
    'Active',
    'Passive',
    'Enable',
    'Disable',
    'Up',
    'Down',
    'Insert',
    'Update',
    'Delete',
    'Inactive'
);

-- create schema webintelligence
drop schema if exists webintelligence cascade;
create schema webintelligence;
alter schema webintelligence owner to s2usr;
set search_path = webintelligence,public;

create table crawlpage (
    pagesource text null,
    link text null,
    task text null,
    documenttype text null,
    mimetype text null,
    document bytea default null,
    created timestamp without time zone default now() not null,
    createdby text default 'POSTGRES',
    updated timestamp without time zone default now() not null,
    updatedby text default 'POSTGRES',
    flag types.flag default 'Insert'::types.flag
);

create table crawlhref (
    link text null,
    href text null,
    task text null,
    created timestamp without time zone default now() not null,
    createdby text default 'POSTGRES',
    updated timestamp without time zone default now() not null,
    updatedby text default 'POSTGRES',
    flag types.flag default 'Insert'::types.flag
);