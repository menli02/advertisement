-- Users
CREATE TABLE IF NOT EXISTS users (
    id         BIGSERIAL PRIMARY KEY,
    phone      VARCHAR(20) UNIQUE NOT NULL,
    first_name VARCHAR(100) NOT NULL DEFAULT '',
    last_name  VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Categories
CREATE TABLE IF NOT EXISTS categories (
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    slug       VARCHAR(120) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Advertisements
CREATE TABLE IF NOT EXISTS advertisements (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL REFERENCES users(id),
    category_id BIGINT NOT NULL REFERENCES categories(id),
    title       VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    slug        VARCHAR(300) UNIQUE NOT NULL,
    price       NUMERIC(12, 2) NOT NULL DEFAULT 0,
    currency    VARCHAR(10) NOT NULL DEFAULT 'USD',
    status      VARCHAR(20) NOT NULL DEFAULT 'active'
                CHECK (status IN ('active', 'inactive', 'sold', 'expired')),
    view_count  BIGINT NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_ads_user_id      ON advertisements(user_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_ads_category_id  ON advertisements(category_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_ads_slug         ON advertisements(slug) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_ads_status       ON advertisements(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_ads_created_at   ON advertisements(created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_ads_price        ON advertisements(price) WHERE deleted_at IS NULL;

-- Full-text search index
CREATE INDEX IF NOT EXISTS idx_ads_fts ON advertisements
    USING GIN(to_tsvector('english', title || ' ' || description))
    WHERE deleted_at IS NULL;

-- Advertisement Images
CREATE TABLE IF NOT EXISTS advertisement_images (
    id         BIGSERIAL PRIMARY KEY,
    ad_id      BIGINT NOT NULL REFERENCES advertisements(id) ON DELETE CASCADE,
    url        TEXT NOT NULL,
    position   INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ad_images_ad_id ON advertisement_images(ad_id);
