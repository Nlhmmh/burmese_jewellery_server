CREATE TABLE IF NOT EXISTS faqs (
	faq_id     UUID NOT NULL DEFAULT gen_random_uuid(),
	question   TEXT NOT NULL,
	answer     TEXT NOT NULL,
	is_active  BOOL NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_faqs PRIMARY KEY (faq_id)
);
