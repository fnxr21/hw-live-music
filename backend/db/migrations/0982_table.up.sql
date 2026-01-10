CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE DATABASE live_music;


CREATE TABLE IF NOT EXISTS live_music.ref_song_status (
	status_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	status_name text NOT NULL,
	requested_at timestamp(6) DEFAULT now() NOT NULL,
	approved_at timestamp(6) NULL,
	approved_by uuid NULL,
	is_active bool DEFAULT true NOT NULL,
	created_at timestamp(6) DEFAULT now() NOT NULL,
	updated_at timestamp(6) DEFAULT now() NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT ref_song_status_pkey PRIMARY KEY (status_id)
);


CREATE TABLE IF NOT EXISTS live_music.ref_songs (
	song_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	title text NOT NULL,
	artist text NOT NULL,
	duration int4 NULL,
	header_image_url text NULL,
	url text NULL,
	release_song_date date NULL,
	is_active bool DEFAULT true NOT NULL,
	created_at timestamp(6) DEFAULT now() NOT NULL,
	updated_at timestamp(6) DEFAULT now() NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT ref_songs_pkey PRIMARY KEY (song_id)
);




CREATE TABLE IF NOT EXISTS live_music.ref_tables (
	table_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	table_number int4 NOT NULL,
	is_active bool DEFAULT true NOT NULL,
	created_at timestamp(6) DEFAULT now() NOT NULL,
	updated_at timestamp(6) DEFAULT now() NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT ref_tables_pkey PRIMARY KEY (table_id)
);





CREATE TABLE IF NOT EXISTS live_music.ref_users (
	user_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" text NOT NULL,
	"password" text NOT NULL,
	"token" text NULL,
	"role" text NOT NULL,
	is_active bool DEFAULT true NOT NULL,
	created_at timestamp(6) DEFAULT now() NOT NULL,
	updated_at timestamp(6) DEFAULT now() NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT ref_users_pkey PRIMARY KEY (user_id)
);




CREATE TABLE IF NOT EXISTS live_music.trx_live_playlists (
	live_playlist_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	song_request_id uuid NOT NULL,
	order_number int4 NULL,
	is_current bool DEFAULT false NOT NULL,
	table_id uuid NULL,
	is_active bool DEFAULT true NOT NULL,
	created_at timestamp(6) DEFAULT now() NOT NULL,
	updated_at timestamp(6) DEFAULT now() NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT trx_live_playlists_pkey PRIMARY KEY (live_playlist_id)
);


CREATE TABLE IF NOT EXISTS live_music.trx_song_requests (
	song_request_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	table_id uuid NULL,
	song_id uuid NOT NULL,
	status uuid NOT NULL,
	requested_at timestamp(6) DEFAULT now() NOT NULL,
	approved_at timestamp(6) NULL,
	approved_by uuid NULL,
	is_active bool DEFAULT true NOT NULL,
	created_at timestamp(6) DEFAULT now() NOT NULL,
	updated_at timestamp(6) DEFAULT now() NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT trx_song_requests_pkey PRIMARY KEY (song_request_id)
);
