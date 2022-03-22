create table data
(
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title      varchar not null,
    text       text    not null,
    tag        varchar not null,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL
);

