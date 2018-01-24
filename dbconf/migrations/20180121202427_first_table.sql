
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE ngi_daily (
	name	text NOT NULL,
	time    text NOT NULL,
	price   text NOT NULL,
	percentage text NOT NULL,
	PRIMARY KEY(name, time)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE ngi_daily;
