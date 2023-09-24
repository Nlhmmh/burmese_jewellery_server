CREATE TABLE IF NOT EXISTS account_carts (
	cart_id    UUID NOT NULL DEFAULT gen_random_uuid(),
	account_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_carts PRIMARY KEY (cart_id),
	CONSTRAINT fk_account_favourites_account_id FOREIGN KEY(account_id) REFERENCES accounts(account_id)
);

CREATE TABLE IF NOT EXISTS account_cart_jewelleries (
	cart_id      UUID NOT NULL,
	jewellery_id UUID NOT NULL,
	quantity     INTEGER NOT NULL,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_cart_jewelleries PRIMARY KEY (cart_id, jewellery_id),
	CONSTRAINT fk_account_cart_jewelleries_cart_id FOREIGN KEY(cart_id) REFERENCES account_carts(cart_id),
	CONSTRAINT fk_account_cart_jewelleries_jewellery_id FOREIGN KEY(jewellery_id) REFERENCES jewelleries(jewellery_id)
);