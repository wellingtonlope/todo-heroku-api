CREATE TABLE IF NOT EXISTS todo
(
    id          uuid DEFAULT gen_random_uuid() primary key,
    title       text,
    description text
);