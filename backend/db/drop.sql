-- Drop triggers first
DROP TRIGGER IF EXISTS trg_trx_song_requests ON live_music.trx_song_requests;
DROP TRIGGER IF EXISTS trg_trx_live_playlists ON live_music.trx_live_playlists;
DROP TRIGGER IF EXISTS trg_song_request_approved ON live_music.trx_song_requests;
DROP TRIGGER IF EXISTS trg_live_playlist_update ON live_music.trx_live_playlists;

-- Drop functions
DROP FUNCTION IF EXISTS live_music.notify_realtime_simple() CASCADE;
DROP FUNCTION IF EXISTS live_music.trg_insert_live_playlist() CASCADE;
DROP FUNCTION IF EXISTS live_music.trg_live_playlist_update() CASCADE;

-- Drop tables
DROP TABLE IF EXISTS live_music.trx_live_playlists CASCADE;
DROP TABLE IF EXISTS live_music.trx_song_requests CASCADE;
DROP TABLE IF EXISTS live_music.ref_song_status CASCADE;
DROP TABLE IF EXISTS live_music.ref_songs CASCADE;
DROP TABLE IF EXISTS live_music.ref_tables CASCADE;
DROP TABLE IF EXISTS live_music.ref_users CASCADE;

-- Finally, drop the schema itself (optional)
DROP SCHEMA IF EXISTS live_music CASCADE;

-- Drop extension (optional, if you don’t need uuid generation anymore)
DROP EXTENSION IF EXISTS "uuid-ossp";
