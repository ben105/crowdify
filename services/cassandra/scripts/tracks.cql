USE crowdify;

-- Create unprocessed_tracks table
CREATE TABLE unprocessed_tracks (
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
  artists frozen<list<map<text, text>>>,
);

-- Create a secondary index on album_id for querying
CREATE INDEX ON tracks (album_id);
