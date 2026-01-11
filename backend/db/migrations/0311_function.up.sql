

-- Function to notify real-time changes for inserts and updates only
CREATE OR REPLACE FUNCTION live_music.notify_realtime_simple()
RETURNS trigger AS $$
DECLARE
  row_id uuid;
BEGIN
  -- determine primary key depending on table
  IF TG_TABLE_NAME = 'trx_song_requests' THEN
    row_id := NEW.song_request_id;
  ELSIF TG_TABLE_NAME = 'trx_live_playlists' THEN
    row_id := NEW.live_playlist_id;
  END IF;

  -- send notification with table name and row id
  PERFORM pg_notify(
    'realtime_channel',
    json_build_object(
      'table', TG_TABLE_NAME,
      'operation', TG_OP,
      'id', row_id
    )::text
  );

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Trigger for trx_song_requests
CREATE TRIGGER trg_trx_song_requests
AFTER INSERT OR UPDATE
ON live_music.trx_song_requests
FOR EACH ROW
EXECUTE FUNCTION live_music.notify_realtime_simple();

-- Trigger for trx_live_playlists
CREATE TRIGGER trg_trx_live_playlists
AFTER INSERT OR UPDATE
ON live_music.trx_live_playlists
FOR EACH ROW
EXECUTE FUNCTION live_music.notify_realtime_simple();






--

-- Function: insert into live playlist when song request is approved
CREATE OR REPLACE FUNCTION live_music.trg_insert_live_playlist()
RETURNS trigger AS $$
DECLARE
  approved_status uuid;
  max_order int;
BEGIN
  -- Get UUID for APPROVED status
  SELECT status_id INTO approved_status
  FROM live_music.ref_song_status
  WHERE status_name = 'APPROVED';

  -- Only run if status changed to APPROVED
  IF NEW.status = approved_status AND OLD.status IS DISTINCT FROM NEW.status THEN
    -- Get current max order_number for this table_id
SELECT COALESCE(MAX(order_number), 0) INTO max_order
FROM live_music.trx_live_playlists;


    -- Insert new row into live playlist
    INSERT INTO live_music.trx_live_playlists(
      song_request_id,
      order_number,
      is_current,
      table_id,
      is_active,
      created_at,
      updated_at
    )
    VALUES (
      NEW.song_request_id,
      max_order + 1,
      FALSE,
      NEW.table_id,
      TRUE,
      NOW(),
      NOW()
    );
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger for approval
CREATE TRIGGER trg_song_request_approved
AFTER UPDATE ON live_music.trx_song_requests
FOR EACH ROW
EXECUTE FUNCTION live_music.trg_insert_live_playlist();





-- admin insert with approved

CREATE OR REPLACE FUNCTION live_music.trg_insert_live_playlist_admin()
RETURNS trigger AS $$
DECLARE
    approved_status uuid;
    max_order int;
BEGIN
    -- Get UUID for APPROVED status
    SELECT status_id INTO approved_status
    FROM live_music.ref_song_status
    WHERE status_name = 'APPROVED';

    -- Only run if status = APPROVED
    IF NEW.status = approved_status THEN

        -- Get current max order_number ignoring table_id
        SELECT COALESCE(MAX(order_number), 0) INTO max_order
        FROM live_music.trx_live_playlists;

        -- Insert into live playlist
        INSERT INTO live_music.trx_live_playlists(
            song_request_id,
            order_number,
            is_current,
            table_id,   -- keep NULL since it's admin
            is_active,
            created_at,
            updated_at
        )
        VALUES (
            NEW.song_request_id,
            max_order + 1,
            FALSE,
            NULL,
            TRUE,
            NOW(),
            NOW()
        );

    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;




-- Drop trigger if exists
DROP TRIGGER IF EXISTS trg_song_request_admin_insert
ON live_music.trx_song_requests;

-- Create trigger for INSERT only
CREATE TRIGGER trg_song_request_admin_insert
AFTER INSERT ON live_music.trx_song_requests
FOR EACH ROW
EXECUTE FUNCTION live_music.trg_insert_live_playlist_admin();





--
-- DROP FUNCTION live_music.trg_live_playlist_after_update();

CREATE OR REPLACE FUNCTION live_music.trg_live_playlist_after_update()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
    ----------------------------------
    -- Case 1: Playlist item marked inactive
    ----------------------------------
    IF OLD.is_active = TRUE AND NEW.is_active = FALSE THEN
        
        -- Mark corresponding song_request as inactive
        UPDATE live_music.trx_song_requests
        SET is_active = FALSE,
            updated_at = NOW()
        WHERE song_request_id = NEW.song_request_id;

        -- Reorder remaining active playlist items
        WITH ranked AS (
            SELECT 
                live_playlist_id,
                ROW_NUMBER() OVER (ORDER BY order_number) as new_order
            FROM live_music.trx_live_playlists
            WHERE is_active = TRUE
              AND live_playlist_id != NEW.live_playlist_id
        )
        UPDATE live_music.trx_live_playlists p
        SET order_number = r.new_order,
            is_current = (r.new_order = 1),
            updated_at = NOW()
        FROM ranked r
        WHERE p.live_playlist_id = r.live_playlist_id;
        
    END IF;

    ----------------------------------
    -- Case 2: Playlist item moved to order_number = 1
    ----------------------------------
    IF NEW.order_number = 1 
        AND NEW.is_active = TRUE
        AND COALESCE(OLD.order_number, 0) != 1 THEN
        
        -- Set is_current = FALSE for all other items
        UPDATE live_music.trx_live_playlists
        SET is_current = FALSE,
            updated_at = NOW()
        WHERE is_active = TRUE
          AND live_playlist_id != NEW.live_playlist_id
          AND is_current = TRUE;
        
        -- Update song request timestamp
        UPDATE live_music.trx_song_requests
        SET updated_at = NOW()
        WHERE song_request_id = NEW.song_request_id;
    END IF;

    ----------------------------------
    -- Case 3: Drag and Drop Reordering
    ----------------------------------
    IF NEW.is_active = TRUE 
        AND OLD.is_active = TRUE
        AND OLD.order_number IS DISTINCT FROM NEW.order_number 
        AND NEW.order_number IS NOT NULL THEN
        
        WITH shifted AS (
            SELECT 
                live_playlist_id,
                order_number as old_order,
                CASE
                    WHEN OLD.order_number < NEW.order_number THEN
                        CASE 
                            WHEN order_number > OLD.order_number 
                                AND order_number <= NEW.order_number 
                            THEN order_number - 1
                            ELSE order_number
                        END
                    WHEN OLD.order_number > NEW.order_number THEN
                        CASE 
                            WHEN order_number >= NEW.order_number 
                                AND order_number < OLD.order_number 
                            THEN order_number + 1
                            ELSE order_number
                        END
                    ELSE order_number
                END as new_order
            FROM live_music.trx_live_playlists
            WHERE is_active = TRUE
              AND live_playlist_id != NEW.live_playlist_id
        )
        UPDATE live_music.trx_live_playlists p
        SET order_number = s.new_order,
            is_current = (s.new_order = 1),
            updated_at = NOW()
        FROM shifted s
        WHERE p.live_playlist_id = s.live_playlist_id
          AND p.order_number != s.new_order;
          
    END IF;

    RETURN NEW;
END;
$function$
;

-- Recreate the trigger
DROP TRIGGER IF EXISTS trg_live_playlist_after_update ON live_music.trx_live_playlists;

CREATE TRIGGER trg_live_playlist_after_update
AFTER UPDATE ON live_music.trx_live_playlists
FOR EACH ROW
EXECUTE FUNCTION live_music.trg_live_playlist_after_update();