DROP FUNCTION IF EXISTS public."fn_get_media_real_path"(int4);

CREATE FUNCTION public."fn_get_media_real_path"(p_id int4)
RETURNS varchar
AS
$FUNCTION$
BEGIN
    RETURN (
        SELECT
            CONCAT(public."fn_get_directory_real_path"(d.id), '/', m.real_name)
        FROM  public."medias" m
        INNER JOIN public."directories" d ON m.directory_id = d.id
        WHERE m.id = p_id
    );
END;
$FUNCTION$
LANGUAGE plpgsql;


COMMENT ON FUNCTION public."fn_get_media_real_path"(int4) 
IS 'Obtiene la ruta real de un archivo multimedia en base a su id.';