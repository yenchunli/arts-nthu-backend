CREATE TABLE "users" (
  "username" varchar(20) PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar(20) NOT NULL,
  "email" varchar(50) UNIQUE NOT NULL,
  "password_change_at" bigint NOT NULL,  
  "create_at" bigint NOT NULL
);