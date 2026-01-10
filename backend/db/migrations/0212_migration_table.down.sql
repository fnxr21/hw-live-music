-- transactional tables first
DROP TABLE IF EXISTS live_music.trx_live_playlists;
DROP TABLE IF EXISTS live_music.trx_song_requests;

-- reference tables
DROP TABLE IF EXISTS live_music.ref_users;
DROP TABLE IF EXISTS live_music.ref_tables;
DROP TABLE IF EXISTS live_music.ref_songs;
DROP TABLE IF EXISTS live_music.ref_song_status;
