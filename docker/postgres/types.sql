--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1 (Debian 16.1-1.pgdg120+1)
-- Dumped by pg_dump version 16.1 (Debian 16.1-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: types; Type: SCHEMA; Schema: -; Owner: s2usr
--

CREATE SCHEMA types;


ALTER SCHEMA types OWNER TO s2usr;

--
-- Name: flag; Type: TYPE; Schema: types; Owner: s2usr
--

CREATE TYPE types.flag AS ENUM (
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


ALTER TYPE types.flag OWNER TO s2usr;

--
-- PostgreSQL database dump complete
--

