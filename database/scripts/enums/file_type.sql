DROP TYPE IF EXISTS public."file_type";

CREATE TYPE public."file_type" AS ENUM (
    'image',
    'video',
    'gif'
);

COMMENT ON TYPE public."file_type" IS 'The type of the file';