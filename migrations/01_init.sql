CREATE DATABASE custom_queue;

CREATE TABLE public.events (
    id SERIAL,
    status_id BIGINT,
    topic_id BIGINT,
    created_at TIMESTAMP WITHOUT TIMEZONE,
    message_data jsonb
);

CREATE INDEX public.status_id_idx ON public.events (status_id);
CREATE INDEX public.topic_id_idx ON public.events (topic_id);

CREATE TABLE public.statuses (
    id SERIAL,
    title NVARCHAR(100)
);

CREATE TABLE public.topic (
    id SERIAL,
    title NVARCHAR(300)
);