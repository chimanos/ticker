CREATE TABLE Message
(
  id serial NOT NULL,
  pair text,
	market text,
	price double precision,
	best_ask double precision,
	best_bid double precision,
	time bigint,
  PRIMARY KEY (id)
)


