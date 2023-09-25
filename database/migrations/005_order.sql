CREATE TYPE order_status AS ENUM ('proceeding', 'shipped', 'delivered', 'returned', 'cancelled');

CREATE TABLE IF NOT EXISTS account_orders (
	order_id     UUID NOT NULL DEFAULT gen_random_uuid(),
	account_id   UUID NOT NULL,
	order_status order_status NOT NULL,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_orders PRIMARY KEY (order_id),
	CONSTRAINT fk_account_orders_account_id FOREIGN KEY(account_id) REFERENCES accounts(account_id)
);

CREATE TABLE IF NOT EXISTS account_order_jewelleries (
	order_id     UUID NOT NULL,
	jewellery_id UUID NOT NULL,
	quantity     INTEGER NOT NULL,
	price        INTEGER NOT NULL,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_order_jewelleries PRIMARY KEY (order_id, jewellery_id),
	CONSTRAINT fk_account_order_jewelleries_order_id FOREIGN KEY(order_id) REFERENCES account_orders(order_id),
	CONSTRAINT fk_account_order_jewelleries_jewellery_id FOREIGN KEY(jewellery_id) REFERENCES jewelleries(jewellery_id)
);

CREATE TABLE IF NOT EXISTS account_order_addresses (
	order_id     UUID NOT NULL,
	country_code CHAR(2) NOT NULL,
	post_code    TEXT,
	state        TEXT,
	city         TEXT NOT NULL,
	address      TEXT NOT NULL,
	remarks      TEXT,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_order_addresses PRIMARY KEY (order_id),
	CONSTRAINT fk_account_order_addresses_order_id FOREIGN KEY(order_id) REFERENCES account_orders(order_id)
);
COMMENT ON COLUMN account_order_addresses.state is 'State, Province, Division';