CREATE TABLE "jobs" (
  "id" varchar(100) PRIMARY KEY,
  "name" varchar(100),
  "state" varchar(100),
  "create_date" timestamptz,
  "update_date" timestamptz,
  "delete_date" timestamptz
);

CREATE TABLE "tasks" (
  "id" varchar(100) PRIMARY KEY,
  "job_id" varchar(100),
  "execute_time" integer,
  "status" varchar(100),
  "create_date" timestamptz,
  "update_date" timestamptz,
  "delete_date" timestamptz
);

CREATE TABLE "executors" (
  "id" varchar(100) PRIMARY KEY,
  "name" varchar(100),
  "status" varchar(100),
  "create_date" timestamptz,
  "update_date" timestamptz,
  "delete_date" timestamptz
);