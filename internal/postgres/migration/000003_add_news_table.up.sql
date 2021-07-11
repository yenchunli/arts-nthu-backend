CREATE TABLE "news" (
    "id" bigserial PRIMARY KEY,
    "username" varchar(20) NOT NULL,
    "author" varchar(20) NOT NULL,
    "title" varchar(60) NOT NULL UNIQUE,
    "title_en" varchar(120) DEFAULT '', 
    "start_date" varchar(10) NOT NULL,
    "type" varchar(20) NOT NULL,
    "draft" boolean NOT NULL DEFAULT 'true',
    "content" text NOT NULL ,
    "content_en" text DEFAULT '',
    "create_at" bigint NOT NULL,
    "update_at" bigint NOT NULL
);