-- Account

CREATE TYPE account_admin_role AS ENUM ('admin', 'staff');
CREATE TYPE account_admin_status AS ENUM ('active', 'locked', 'deactivated');

CREATE TABLE IF NOT EXISTS account_admins (
	account_admin_id     UUID NOT NULL DEFAULT gen_random_uuid(),
	mail                 TEXT NOT NULL,
	password             TEXT NOT NULL,
	account_admin_role   account_admin_role NOT NULL,
	account_admin_status account_admin_status NOT NULL,
	created_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_admins PRIMARY KEY (account_admin_id),
	CONSTRAINT unique_account_admins_email UNIQUE (mail)
);

CREATE TYPE login_type AS ENUM ('email', 'phone', 'google', 'facebook');
CREATE TYPE account_status AS ENUM ('pending', 'active', 'locked', 'deactivated');

CREATE TABLE IF NOT EXISTS accounts (
	account_id     UUID NOT NULL DEFAULT gen_random_uuid(),
	login_type     login_type NOT NULL,
	login_id       TEXT,
	mail           TEXT,
	password       TEXT,
	phone          TEXT,
	account_status account_status NOT NULL,
	created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_accounts PRIMARY KEY (account_id)
);

CREATE TYPE gender AS ENUM ('male', 'female', 'unspecified');

CREATE TABLE IF NOT EXISTS account_profiles (
	account_id UUID NOT NULL,
	first_name TEXT NOT NULL,
	last_name  TEXT NOT NULL,
	birthday   DATE,
	gender     gender,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_account_profiles PRIMARY KEY (account_id),
	CONSTRAINT fk_account_profiles_account_id FOREIGN KEY(account_id) REFERENCES accounts(account_id)
);