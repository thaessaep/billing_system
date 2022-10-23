CREATE TABLE IF NOT EXISTS "users"(

	user_id integer primary key,
	balance integer NOT NULL
	
);

CREATE TABLE IF NOT EXISTS "reserve_bills"(
	
	success boolean,
	order_id integer NOT NULL,
	service_id integer NOT NULL,
	cost integer NOT NULL,
	user_id integer REFERENCES "users"

)