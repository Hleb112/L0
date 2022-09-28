CREATE TABLE IF NOT EXISTS public.new_order
(
    order_uuid uuid NOT NULL,
    track_number character varying COLLATE pg_catalog."default",
    entry character varying COLLATE pg_catalog."default",
    internal_signature "char",
    shardkey "char",
    oof_shard text COLLATE pg_catalog."default",
    delivery json,
    items json,
    date_created text COLLATE pg_catalog."default",
    locale character varying COLLATE pg_catalog."default",
    customer_id character varying COLLATE pg_catalog."default",
    delivery_service character varying COLLATE pg_catalog."default",
    sm_id integer,
    payment json,
    CONSTRAINT order_uuid_pkey PRIMARY KEY (order_uuid)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.new_order
    OWNER to postgres;
