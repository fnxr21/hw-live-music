
SELECT * FROM live_music.trx_song_requests;
SELECT * FROM live_music.trx_live_playlists tlp ;

SELECT * FROM live_music.ref_song_status ;
SELECT * FROM live_music.ref_songs ;
SELECT * FROM live_music.ref_tables;
SELECT * FROM live_music.ref_users ;



-- Users
INSERT INTO live_music.ref_users (name, password, role)
VALUES
('Alice', 'password123', 'ADMIN'),
('Bob', 'password456', 'STAFF');

-- Tables
INSERT INTO live_music.ref_tables (table_number)
VALUES (1), (2);

-- Songs
INSERT INTO live_music.ref_songs (title, artist, duration, url)
VALUES
('Song A', 'Artist 1', 180, 'https://example.com/songA.mp3'),
('Song B', 'Artist 2', 200, 'https://example.com/songB.mp3'),
('Song C', 'Artist 3', 240, 'https://example.com/songC.mp3');

-- Statuses
INSERT INTO live_music.ref_song_status (status_name)
VALUES ('PENDING'), ('APPROVED'), ('REJECTED');


-- First, get IDs
WITH song_ids AS (
  SELECT song_id, title FROM live_music.ref_songs
), table_ids AS (
  SELECT table_id, table_number FROM live_music.ref_tables
), status_ids AS (
  SELECT status_id, status_name FROM live_music.ref_song_status
  
)

INSERT INTO live_music.trx_song_requests (table_id, song_id, status)
SELECT t.table_id, s.song_id, st.status_id
FROM table_ids t
JOIN song_ids s ON s.title = 'Song A'
JOIN status_ids st ON st.status_name = 'PENDING';


-- First, get IDs
WITH song_ids AS (
  SELECT song_id, title FROM live_music.ref_songs
), table_ids AS (
  SELECT table_id, table_number FROM live_music.ref_tables
), status_ids AS (
  SELECT status_id, status_name FROM live_music.ref_song_status
  
)
INSERT INTO live_music.trx_song_requests (table_id, song_id, status)
SELECT t.table_id, s.song_id, st.status_id
FROM table_ids t
JOIN song_ids s ON s.title = 'Song B'
JOIN status_ids st ON st.status_name = 'PENDING';

-- First, get IDs
WITH song_ids AS (
  SELECT song_id, title FROM live_music.ref_songs
), table_ids AS (
  SELECT table_id, table_number FROM live_music.ref_tables
), status_ids AS (
  SELECT status_id, status_name FROM live_music.ref_song_status
  
)

INSERT INTO live_music.trx_song_requests (table_id, song_id, status)
SELECT t.table_id, s.song_id, st.status_id
FROM table_ids t
JOIN song_ids s ON s.title = 'Song C'
JOIN status_ids st ON st.status_name = 'PENDING';





-- Approve Song 
UPDATE live_music.trx_song_requests
SET status = (SELECT status_id FROM live_music.ref_song_status WHERE status_name='APPROVED')
WHERE song_request_id = 'c73fba27-2fcf-44dc-a427-8555f4f0c164';


-- Mark the first playlist item as inactive (finished playing)
UPDATE live_music.trx_live_playlists
SET is_active = FALSE
WHERE order_number = 1;


-- move order number 2 to  position 1
UPDATE live_music.trx_live_playlists
SET order_number = 3
WHERE live_playlist_id = '63ee0720-6929-43c9-ac22-dd2c0fef3c21';



	SELECT tsr.*,rss.status_name ,rt.table_number 
		FROM live_music.trx_song_requests tsr 
		left join live_music.ref_song_status rss on rss.status_id  =tsr.status 
		left join live_music.ref_tables rt on rt.table_id  =tsr.table_id 
		WHERE tsr.is_active = TRUE
		ORDER BY tsr.requested_at ASC


