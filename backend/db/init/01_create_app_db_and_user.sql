CREATE DATABASE notenest_db;
CREATE USER notenest_user WITH PASSWORD 'notenest_user_password';

grant all privileges
on database notenest_db
to notenest_user
;

grant usage
on schema pg_catalog
to notenest_user
;

ALTER DATABASE notenest_db OWNER TO notenest_user;

ALTER SCHEMA public OWNER TO notenest_user;

