-- (already exist) user: user, password: postgres

-- create user crud_user with encrypted password 'postgres';
create database crud_test;
grant all privileges on database crud_test to crud_user;