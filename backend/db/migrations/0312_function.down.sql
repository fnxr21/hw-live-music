    -- Realtime notification triggers
    DROP TRIGGER IF EXISTS trg_trx_song_requests
    ON live_music.trx_song_requests;

    DROP TRIGGER IF EXISTS trg_trx_live_playlists
    ON live_music.trx_live_playlists;

    -- Song request approval trigger
    DROP TRIGGER IF EXISTS trg_song_request_approved
    ON live_music.trx_song_requests;

    -- Live playlist update trigger
    DROP TRIGGER IF EXISTS trg_live_playlist_update
    ON live_music.trx_live_playlists;

    -- Drop trigger if exists
DROP TRIGGER IF EXISTS trg_song_request_admin_insert
ON live_music.trx_song_requests;


    DROP FUNCTION IF EXISTS live_music.notify_realtime_simple();
    DROP FUNCTION IF EXISTS live_music.trg_insert_live_playlist();
    DROP FUNCTION IF EXISTS live_music.trg_live_playlist_update();
