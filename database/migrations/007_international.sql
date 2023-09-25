CREATE TABLE IF NOT EXISTS currencies (
	currency_code VARCHAR(3) NOT NULL,
	name          TEXT NOT NULL,
	decimals      TEXT NOT NULL,
	is_active     BOOL NOT NULL,
	created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_currencies PRIMARY KEY (currency_code)
);
COMMENT ON COLUMN currencies.currency_code is 'ISO 4217';

CREATE TABLE IF NOT EXISTS countries (
	country_code   VARCHAR(2) NOT NULL,
	currency_code  VARCHAR(3) NOT NULL,
	name           TEXT NOT NULL,
	is_active      BOOL NOT NULL,
	flag_image_url TEXT NOT NULL,
	created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_countries PRIMARY KEY (country_code),
	CONSTRAINT fk_countries_currency_code FOREIGN KEY(currency_code) REFERENCES currencies(currency_code)
);
COMMENT ON COLUMN countries.country_code is 'ISO Alpha-2';