DROP FUNCTION IF EXISTS public."fn_get_directory_real_path"(int4, varchar);

CREATE FUNCTION public."fn_get_directory_real_path"(p_id int4, p_parent_path varchar DEFAULT '')
RETURNS varchar
AS
$FUNCTION$
DECLARE
    v_parent_id int4;
    v_real_name varchar;
    v_current_path varchar;
BEGIN
    SELECT
        d.parent_id,
        d.real_name
    INTO
        v_parent_id,
        v_real_name
    FROM public."directories" as d
    WHERE d.id = p_id;

    IF p_parent_path = '' THEN
        v_current_path := v_real_name;
    ELSE
        v_current_path := CONCAT(v_real_name, '/', p_parent_path);
    END IF;

    IF v_parent_id IS NULL THEN
        RETURN  v_current_path;
    ELSE
        RETURN public."fn_get_directory_real_path"(v_parent_id, v_current_path);
    END IF;
END;
$FUNCTION$
LANGUAGE plpgsql;

COMMENT ON FUNCTION public."fn_get_directory_real_path"(int4, varchar)
IS 'Función recursiva que obtiene la ruta real de un directorio en base a su id y el id de su directorio padre.';