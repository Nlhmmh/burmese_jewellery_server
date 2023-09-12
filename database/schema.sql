-- Account

CREATE TYPE login_type AS ENUM ('email', 'phone', 'google', 'facebook');
CREATE TYPE account_status AS ENUM ('pending', 'active', 'locked', 'deactivated');

CREATE TABLE IF NOT EXISTS accounts (
	account_id     UUID NOT NULL DEFAULT gen_random_uuid(),
	login_type     login_type NOT NULL,
	login_id       VARCHAR(255),
	mail           VARCHAR(255),
	password       VARCHAR(255),
	phone          VARCHAR(255),
	account_status account_status NOT NULL,
	created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_accounts_account_id PRIMARY KEY (account_id)
);

CREATE TYPE gender AS ENUM ('male', 'female', 'unspecified');

CREATE TABLE IF NOT EXISTS account_profiles (
	account_id UUID NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name  VARCHAR(255) NOT NULL,
	birthday   DATE,
	gender     gender,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_profiles_account_id PRIMARY KEY (account_id),
	CONSTRAINT fk_account_profiles_account_id FOREIGN KEY(account_id) REFERENCES accounts(account_id)
);

CREATE TYPE account_admin_role AS ENUM ('admin', 'staff');
CREATE TYPE account_admin_status AS ENUM ('active', 'locked', 'deactivated');

CREATE TABLE IF NOT EXISTS account_admins (
	account_admins_id    UUID NOT NULL DEFAULT gen_random_uuid(),
	mail                 VARCHAR(255) NOT NULL,
	password             VARCHAR(255) NOT NULL,
	account_admin_role   account_admin_role NOT NULL,
	account_admin_status account_admin_status NOT NULL,
	created_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_admins_account_admins_id PRIMARY KEY (account_admins_id)
);

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
	CONSTRAINT pk_jewelleries_jewellery_id PRIMARY KEY (jewellery_id),
	CONSTRAINT fk_jewelleries_category_id FOREIGN KEY(category_id) REFERENCES categories(category_id),
	CONSTRAINT fk_jewelleries_gem_id FOREIGN KEY(gem_id) REFERENCES gems(gem_id),
	CONSTRAINT fk_jewelleries_material_id FOREIGN KEY(material_id) REFERENCES materials(material_id)
);
