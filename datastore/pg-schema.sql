-- pre
drop user if exists s2usr;
create user s2usr with password 's2usr';
alter user s2usr with superuser;
drop database if exists s2;
create database s2;
alter database s2 owner to s2usr;