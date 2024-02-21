CREATE TABLE IF NOT EXISTS account_favourites (
	account_id   UUID NOT NULL,
	jewellery_id UUID NOT NULL,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_favourites PRIMARY KEY (account_id, jewellery_id),
	CONSTRAINT fk_account_favourites_account_id FOREIGN KEY(account_id) REFERENCES accounts(account_id),
	CONSTRAINT fk_account_favourites_jewellery_id FOREIGN KEY(jewellery_id) REFERENCES jewelleries(jewellery_id)
);
