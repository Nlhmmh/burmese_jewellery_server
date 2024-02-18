-- Jewellery

CREATE TABLE IF NOT EXISTS categories (
	category_id UUID NOT NULL DEFAULT gen_random_uuid(),
	name        TEXT NOT NULL,
	description TEXT NOT NULL,
	image_url   TEXT NOT NULL,
	created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_categories_category_id PRIMARY KEY (category_id),
	CONSTRAINT unique_categories_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS gems (
	gem_id     UUID NOT NULL DEFAULT gen_random_uuid(),
	name       TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_gems_gem_id PRIMARY KEY (gem_id),
	CONSTRAINT unique_gems_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS materials (
	material_id UUID NOT NULL DEFAULT gen_random_uuid(),
	name        TEXT NOT NULL,
	created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	UNIQUE (name),
	CONSTRAINT pk_materials_material_id PRIMARY KEY (material_id),
	CONSTRAINT unique_materials_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS jewelleries (
	jewellery_id UUID NOT NULL DEFAULT gen_random_uuid(),
	category_id  UUID NOT NULL,
	gem_id       UUID NOT NULL,
	material_id  UUID NOT NULL,
	name         TEXT NOT NULL,
	description  TEXT NOT NULL,
	price        INTEGER NOT NULL,
	image_url    TEXT NOT NULL,
	is_published BOOL NOT NULL,
	quantity     INTEGER NOT NULL DEFAULT 0,
	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_jewelleries PRIMARY KEY (jewellery_id),
	CONSTRAINT fk_jewelleries_category_id FOREIGN KEY(category_id) REFERENCES categories(category_id),
	CONSTRAINT fk_jewelleries_gem_id FOREIGN KEY(gem_id) REFERENCES gems(gem_id),
	CONSTRAINT fk_jewelleries_material_id FOREIGN KEY(material_id) REFERENCES materials(material_id),
	CONSTRAINT unique_jewelleries_name UNIQUE (name)
);
