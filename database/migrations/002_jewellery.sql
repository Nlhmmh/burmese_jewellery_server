-- Jewellery

CREATE TABLE IF NOT EXISTS categories (
	category_id UUID NOT NULL DEFAULT gen_random_uuid(),
	name        VARCHAR(255) NOT NULL,
	description VARCHAR(255) NOT NULL,
	image_url   VARCHAR(255) NOT NULL,
	created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_categories_category_id PRIMARY KEY (category_id)
);

CREATE TABLE IF NOT EXISTS gems (
	gem_id     UUID NOT NULL DEFAULT gen_random_uuid(),
	name       VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_gems_gem_id PRIMARY KEY (gem_id)
);

CREATE TABLE IF NOT EXISTS materials (
	material_id UUID NOT NULL DEFAULT gen_random_uuid(),
	name        VARCHAR(255) NOT NULL,
	created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_materials_material_id PRIMARY KEY (material_id)
);

CREATE TABLE IF NOT EXISTS jewelleries (
	jewellery_id UUID NOT NULL DEFAULT gen_random_uuid(),
	category_id  UUID NOT NULL,
	gem_id       UUID NOT NULL,
	material_id  UUID NOT NULL,
	name         VARCHAR(255) NOT NULL,
	description  VARCHAR(255) NOT NULL,
	quantity     INTEGER NOT NULL,
	price        FLOAT NOT NULL,
	image_url    VARCHAR(255) NOT NULL,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_jewelleries PRIMARY KEY (jewellery_id),
	CONSTRAINT fk_jewelleries_category_id FOREIGN KEY(category_id) REFERENCES categories(category_id),
	CONSTRAINT fk_jewelleries_gem_id FOREIGN KEY(gem_id) REFERENCES gems(gem_id),
	CONSTRAINT fk_jewelleries_material_id FOREIGN KEY(material_id) REFERENCES materials(material_id)
);