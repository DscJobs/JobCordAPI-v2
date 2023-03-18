--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Ubuntu 15.2-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.2 (Ubuntu 15.2-1.pgdg22.04+1)

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
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cv; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.cv (
    itag uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    userid text NOT NULL,
    overview text NOT NULL,
    hire text NOT NULL,
    birthday timestamp with time zone DEFAULT now() NOT NULL,
    link text,
    email text,
    job text,
    vanity text,
    private boolean DEFAULT false NOT NULL,
    developer boolean DEFAULT false NOT NULL,
    current boolean DEFAULT false NOT NULL,
    exptoggle boolean DEFAULT false NOT NULL,
    nitro boolean DEFAULT false NOT NULL,
    views integer DEFAULT 0 NOT NULL,
    likes integer DEFAULT 0 NOT NULL,
    date timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: details; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.details (
    itag uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    userid text NOT NULL,
    expservers integer NOT NULL,
    exp integer NOT NULL,
    active integer NOT NULL,
    salary integer NOT NULL
);


--
-- Name: internal_user_cache; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.internal_user_cache (
    id text NOT NULL,
    username text NOT NULL,
    discriminator text NOT NULL,
    avatar text NOT NULL,
    bot boolean NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    last_updated timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: requests; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.requests (
    itag uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    userid text NOT NULL,
    cv text NOT NULL,
    content text NOT NULL,
    tag text NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    itag uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    userid text NOT NULL,
    token text NOT NULL,
    votes text[] DEFAULT '{}'::text[] NOT NULL,
    banned boolean DEFAULT false NOT NULL,
    staff boolean DEFAULT false NOT NULL,
    premium boolean DEFAULT false NOT NULL,
    lifetime_premium boolean DEFAULT false NOT NULL,
    premium_duration timestamp with time zone DEFAULT now() NOT NULL,
    notifications jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: cv cv_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cv
    ADD CONSTRAINT cv_pkey PRIMARY KEY (itag);


--
-- Name: cv cv_userid_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cv
    ADD CONSTRAINT cv_userid_key UNIQUE (userid);


--
-- Name: details details_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.details
    ADD CONSTRAINT details_pkey PRIMARY KEY (itag);


--
-- Name: details details_userid_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.details
    ADD CONSTRAINT details_userid_key UNIQUE (userid);


--
-- Name: internal_user_cache internal_user_cache_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.internal_user_cache
    ADD CONSTRAINT internal_user_cache_pkey PRIMARY KEY (id);


--
-- Name: requests requests_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.requests
    ADD CONSTRAINT requests_pkey PRIMARY KEY (itag);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (itag);


--
-- Name: users users_userid_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_userid_key UNIQUE (userid);


--
-- Name: users userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT userid_fkey FOREIGN KEY (userid) REFERENCES public.users(userid) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: details userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.details
    ADD CONSTRAINT userid_fkey FOREIGN KEY (userid) REFERENCES public.users(userid) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: requests userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.requests
    ADD CONSTRAINT userid_fkey FOREIGN KEY (userid) REFERENCES public.users(userid) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--
