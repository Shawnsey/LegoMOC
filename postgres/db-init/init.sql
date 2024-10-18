-- create tables
CREATE TABLE public.users
(
    "id" uuid UNIQUE,
    "first_name" Text,
    "last_name" Text,
    "email" varchar(100) UNIQUE,
    "address" varchar(200),
    "phone_number" numeric(10),
    "creator_flag" boolean DEFAULT false,
    PRIMARY KEY ("id")
);
ALTER TABLE IF EXISTS public.users
    OWNER TO shawn;

CREATE TABLE public.orders
(
    "order_id" uuid UNIQUE,
    "customer_id" uuid,
    "line_items" uuid[],
    "created_at" timestamp with time zone,
    "shipped_at" timestamp with time zone,
    "completed_at" timestamp with time zone,
    PRIMARY KEY ("order_id")
);

ALTER TABLE IF EXISTS public.orders
    OWNER to shawn;

Create table public.creations
(
	"id" uuid UNIQUE,
	"name" varchar(30),
	"creator_id" uuid REFERENCES users (id),
	"instructions_link" varchar(100),
	"image_id" uuid,
	"price" numeric(5,2),
	Primary key ("id")
);
alter table if exists public.creations
	owner to shawn;

create table public.images
(
    "id" uuid UNIQUE,
    "imageLink" varchar(100),
    PRIMARY key ("id")
);
ALTER TABLE IF EXISTS public.images
    OWNER to shawn;

CREATE TABLE public.payment_info
(
    "id" uuid UNIQUE,
    "user_id" uuid REFERENCES users (id),
    "creditcard_type" varchar(15),
    "encrypted_card_number" numeric(16),
    "expiration" varchar(5),
    "security_code" numeric(3),
    PRIMARY key ("id")
);
ALTER TABLE IF EXISTS public.payment_info
    OWNER TO shawn;

CREATE TABLE public.pieces
(
    "id" integer UNIQUE,
    "description" varchar(200),
    "image_link" varchar(100),
    PRIMARY KEY ("id")
);
ALTER TABLE IF EXISTS public.pieces
    OWNER TO shawn;

-- Load data
INSERT INTO public.users 
(
    id,
    first_name,
    last_name,
    email,
    address,
    phone_number,
    creator_flag
)
values 
(
    '98311391-88ca-48c7-ad1d-5ccb7fcb4e19',
    'john',
    'doe',
    'johndoe@funmail.com',
    '519 lexington way macon, GA. 12345',
    5552689462,
    TRUE
);

INSERT INTO public.orders
(
    order_id,
    customer_id,
    line_items,
    created_at
)
values
(
    '40a13f82-210d-44db-b1de-78de042c6f75',
    '98311391-88ca-48c7-ad1d-5ccb7fcb4e19',
    ARRAY['98311391-88ca-48c7-ad1d-5ccb7fcb4e19'::uuid],
    now()
);