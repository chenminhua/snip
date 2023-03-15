## CREATE db and table
```sql
CREATE DATABASE dbname;

CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    UNIQUE (name)
);

DROP TABLE IF EXISTS users
```

## size
```sql
SELECT pg_database_size('forumdb');

SELECT pg_size_pretty(pg_database_size('forumdb'));
```

## db/table info

```sql
SELECT * FROM pg_database WHERE datname='forumdb';

SELECT * FROM pg_class WHERE relname='users';
```

## manage temp tables

CREATE TABLE only visible in the session where it was created.

```sql
CREATE TEMP TABLE IF NOT EXISTS temp_users (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    UNIQUE (name) 
);
```

## unlogged tables

```sql
CREATE UNLOGGED TABLE IF NOT EXISTS temp_users (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    UNIQUE (name) 
);
```