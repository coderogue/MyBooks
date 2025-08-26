--
-- PostgreSQL database dump
--

\restrict oc9EKgR4fkmjxMlrKgulweYtzeJ4TrQacnEjBuSdKTsXhD9XEhcVfsAX3aaobAj

-- Dumped from database version 17.6
-- Dumped by pg_dump version 17.6

-- Started on 2025-08-26 13:26:05

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 6 (class 2615 OID 16391)
-- Name: mybooks-db; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA "mybooks-db";


ALTER SCHEMA "mybooks-db" OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 16392)
-- Name: books; Type: TABLE; Schema: mybooks-db; Owner: postgres
--

CREATE TABLE "mybooks-db".books (
    "Id" integer NOT NULL,
    "Title" character varying NOT NULL,
    "Author" character varying NOT NULL,
    "Description" text,
    "Date_Bought" date,
    "Status" integer,
    "Date_Read" date
);


ALTER TABLE "mybooks-db".books OWNER TO postgres;

--
-- TOC entry 4889 (class 0 OID 16392)
-- Dependencies: 218
-- Data for Name: books; Type: TABLE DATA; Schema: mybooks-db; Owner: postgres
--

COPY "mybooks-db".books ("Id", "Title", "Author", "Description", "Date_Bought", "Status", "Date_Read") FROM stdin;
\.


--
-- TOC entry 4743 (class 2606 OID 16398)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: mybooks-db; Owner: postgres
--

ALTER TABLE ONLY "mybooks-db".books
    ADD CONSTRAINT books_pkey PRIMARY KEY ("Id");


-- Completed on 2025-08-26 13:26:06

--
-- PostgreSQL database dump complete
--

\unrestrict oc9EKgR4fkmjxMlrKgulweYtzeJ4TrQacnEjBuSdKTsXhD9XEhcVfsAX3aaobAj

