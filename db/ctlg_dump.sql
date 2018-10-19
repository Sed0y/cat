--
-- PostgreSQL database dump
--

-- Dumped from database version 10.2
-- Dumped by pg_dump version 10.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE category (
    id integer NOT NULL,
    parent_id integer,
    level integer,
    weight integer,
    name character varying(128),
    active boolean NOT NULL,
    url character varying(256)
);


ALTER TABLE category OWNER TO postgres;

--
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY category (id, parent_id, level, weight, name, active, url) FROM stdin;
5	1	2	17	1.3. Пятавя - второй уровень	t	/
1	0	1	15	1. Первый уровень	t	/
3	0	1	16	2. Третья, первый уровень	t	/
4	1	2	1	Четвёртая перешла в первую	t	/
6	5	3	5	Ещё одна	t	ff
2	1	2	15	1.2. вторая, второй уровень	t	http://localhost:8080/admin/categories
\.


--
-- Name: category category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- Name: category level; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE category
    ADD CONSTRAINT level CHECK (((level > 0) AND (level <= 3))) NOT VALID;


--
-- PostgreSQL database dump complete
--

