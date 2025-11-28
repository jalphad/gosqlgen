--
-- PostgreSQL database dump
--

-- Dumped from database version 16.6 (Debian 16.6-1.pgdg120+1)
-- Dumped by pg_dump version 16.6 (Debian 16.6-1.pgdg120+1)

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
-- Name: atlas_schema_revisions; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA atlas_schema_revisions;


--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: atlas_schema_revisions; Type: TABLE; Schema: atlas_schema_revisions; Owner: -
--

CREATE TABLE atlas_schema_revisions.atlas_schema_revisions (
                                                               version character varying NOT NULL,
                                                               description character varying NOT NULL,
                                                               type bigint DEFAULT 2 NOT NULL,
                                                               applied bigint DEFAULT 0 NOT NULL,
                                                               total bigint DEFAULT 0 NOT NULL,
                                                               executed_at timestamp with time zone NOT NULL,
                                                               execution_time bigint NOT NULL,
                                                               error text,
                                                               error_stmt text,
                                                               hash character varying NOT NULL,
                                                               partial_hashes jsonb,
                                                               operator_version character varying NOT NULL
);


--
-- Name: backends; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.backends (
                                 created_at timestamp with time zone,
                                 updated_at timestamp with time zone,
                                 deleted_at timestamp with time zone,
                                 name text NOT NULL,
                                 type text,
                                 properties bytea
);


--
-- Name: collections; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.collections (
                                    id uuid DEFAULT gen_random_uuid(),
                                    created_at timestamp with time zone,
                                    updated_at timestamp with time zone,
                                    deleted_at timestamp with time zone,
                                    backend text NOT NULL,
                                    backend_id text NOT NULL,
                                    pcode text,
                                    request_id uuid
);


--
-- Name: instance_collections; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.instance_collections (
                                             collection_backend text NOT NULL,
                                             collection_backend_id text NOT NULL,
                                             instance_backend text NOT NULL,
                                             instance_backend_id text NOT NULL
);


--
-- Name: instances; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.instances (
                                  id uuid DEFAULT gen_random_uuid(),
                                  created_at timestamp with time zone,
                                  updated_at timestamp with time zone,
                                  deleted_at timestamp with time zone,
                                  backend text NOT NULL,
                                  backend_id text NOT NULL,
                                  pcode text,
                                  request_id uuid
);


--
-- Name: purposes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.purposes (
                                 created_at timestamp with time zone,
                                 updated_at timestamp with time zone,
                                 deleted_at timestamp with time zone,
                                 code text NOT NULL
);


--
-- Name: requests; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.requests (
                                 id uuid DEFAULT gen_random_uuid() NOT NULL,
                                 created_at timestamp with time zone,
                                 updated_at timestamp with time zone,
                                 deleted_at timestamp with time zone,
                                 backend text,
                                 backend_id text,
                                 pcode text,
                                 payload bytea,
                                 status text,
                                 job_id bigint,
                                 description text
);


--
-- Name: services; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.services (
                                 id uuid DEFAULT gen_random_uuid() NOT NULL,
                                 created_at timestamp with time zone,
                                 updated_at timestamp with time zone,
                                 deleted_at timestamp with time zone,
                                 backend_ref text,
                                 backend_id text,
                                 name text,
                                 config bytea
);


--
-- Name: services_purposes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.services_purposes (
                                          service_id uuid NOT NULL,
                                          purpose_code text NOT NULL
);


--
-- Name: atlas_schema_revisions atlas_schema_revisions_pkey; Type: CONSTRAINT; Schema: atlas_schema_revisions; Owner: -
--

ALTER TABLE ONLY atlas_schema_revisions.atlas_schema_revisions
    ADD CONSTRAINT atlas_schema_revisions_pkey PRIMARY KEY (version);


--
-- Name: backends backends_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.backends
    ADD CONSTRAINT backends_pkey PRIMARY KEY (name);


--
-- Name: collections collections_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.collections
    ADD CONSTRAINT collections_pkey PRIMARY KEY (backend, backend_id);


--
-- Name: instance_collections instance_collections_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.instance_collections
    ADD CONSTRAINT instance_collections_pkey PRIMARY KEY (collection_backend, collection_backend_id, instance_backend, instance_backend_id);


--
-- Name: instances instances_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.instances
    ADD CONSTRAINT instances_pkey PRIMARY KEY (backend, backend_id);


--
-- Name: purposes purposes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.purposes
    ADD CONSTRAINT purposes_pkey PRIMARY KEY (code);


--
-- Name: requests requests_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.requests
    ADD CONSTRAINT requests_pkey PRIMARY KEY (id);


--
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- Name: services_purposes services_purposes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services_purposes
    ADD CONSTRAINT services_purposes_pkey PRIMARY KEY (service_id, purpose_code);


--
-- Name: idx_backend_collection; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_backend_collection ON public.collections USING btree (backend, backend_id);


--
-- Name: idx_backend_instance; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_backend_instance ON public.instances USING btree (backend, backend_id);


--
-- Name: idx_backend_service; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_backend_service ON public.services USING btree (name, backend_ref, backend_id);


--
-- Name: idx_backends_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_backends_created_at ON public.backends USING btree (created_at);


--
-- Name: idx_backends_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_backends_deleted_at ON public.backends USING btree (deleted_at);


--
-- Name: idx_collections_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_collections_created_at ON public.collections USING btree (created_at);


--
-- Name: idx_collections_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_collections_deleted_at ON public.collections USING btree (deleted_at);


--
-- Name: idx_collections_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_collections_id ON public.collections USING btree (id);


--
-- Name: idx_instances_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_instances_created_at ON public.instances USING btree (created_at);


--
-- Name: idx_instances_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_instances_deleted_at ON public.instances USING btree (deleted_at);


--
-- Name: idx_instances_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_instances_id ON public.instances USING btree (id);


--
-- Name: idx_purposes_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_purposes_created_at ON public.purposes USING btree (created_at);


--
-- Name: idx_purposes_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_purposes_deleted_at ON public.purposes USING btree (deleted_at);


--
-- Name: idx_requests_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_requests_created_at ON public.requests USING btree (created_at);


--
-- Name: idx_requests_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_requests_deleted_at ON public.requests USING btree (deleted_at);


--
-- Name: idx_requests_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_requests_id ON public.requests USING btree (id);


--
-- Name: idx_services_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_services_created_at ON public.services USING btree (created_at);


--
-- Name: idx_services_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_services_deleted_at ON public.services USING btree (deleted_at);


--
-- Name: idx_services_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_services_id ON public.services USING btree (id);


--
-- Name: instance_collections fk_instance_collections_collection; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.instance_collections
    ADD CONSTRAINT fk_instance_collections_collection FOREIGN KEY (collection_backend, collection_backend_id) REFERENCES public.collections(backend, backend_id);


--
-- Name: instance_collections fk_instance_collections_instance; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.instance_collections
    ADD CONSTRAINT fk_instance_collections_instance FOREIGN KEY (instance_backend, instance_backend_id) REFERENCES public.instances(backend, backend_id);


--
-- Name: collections fk_requests_collections; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.collections
    ADD CONSTRAINT fk_requests_collections FOREIGN KEY (request_id) REFERENCES public.requests(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: instances fk_requests_instances; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.instances
    ADD CONSTRAINT fk_requests_instances FOREIGN KEY (request_id) REFERENCES public.requests(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: services fk_services_backend; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_services_backend FOREIGN KEY (backend_ref) REFERENCES public.backends(name);


--
-- Name: services_purposes fk_services_purposes_purpose; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services_purposes
    ADD CONSTRAINT fk_services_purposes_purpose FOREIGN KEY (purpose_code) REFERENCES public.purposes(code);


--
-- Name: services_purposes fk_services_purposes_service; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services_purposes
    ADD CONSTRAINT fk_services_purposes_service FOREIGN KEY (service_id) REFERENCES public.services(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: -
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


--
-- PostgreSQL database dump complete
--
