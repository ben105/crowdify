-- Create unprocessed_tracks table
CREATE TABLE crowdify.unprocessed_tracks (
  id text PRIMARY KEY,
  name text,
  type text,
  duration_ms int,
  popularity int,
  explicit boolean,
  track_number int,
  disc_number int,
  
  -- Album information
  album_id text,
  album_name text,
  album_type text,
  album_total_tracks int,
  album_release_date text,
  album_release_date_precision text,
  
  -- Artists information (could be multiple)
  artists frozen<list<map<text, text>>>
);

-- Create tracks table
CREATE TABLE crowdify.tracks (
  id text PRIMARY KEY,
  name text,
  type text,
  duration_ms int,
  popularity int,
  explicit boolean,
  track_number int,
  disc_number int,

  -- TODO: Add the fleshed out artists details

  -- Album information
  album_id text,
  album_name text,
  album_type text,
  album_total_tracks int,
  album_release_date text,
  album_release_date_precision text
);

-- TODO: Create a secondary index on the artist as well.

-- Create a secondary index on album_id for querying
CREATE INDEX ON crowdify.tracks (album_id);