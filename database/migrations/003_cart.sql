CREATE TABLE IF NOT EXISTS account_cart_jewelleries (
	account_id   UUID NOT NULL,
	jewellery_id UUID NOT NULL,
	quantity     INTEGER NOT NULL,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_cart_jewelleries PRIMARY KEY (account_id, jewellery_id),
	CONSTRAINT fk_account_cart_jewelleries_account_id FOREIGN KEY(account_id) REFERENCES accounts(account_id),
	CONSTRAINT fk_account_cart_jewelleries_jewellery_id FOREIGN KEY(jewellery_id) REFERENCES jewelleries(jewellery_id)
);
