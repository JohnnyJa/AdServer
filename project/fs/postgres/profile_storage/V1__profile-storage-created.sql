CREATE TABLE profiles
(
    id          UUID PRIMARY KEY,
    name        TEXT,
    bid_price   FLOAT       NOT NULL,
    views_limit INT         NOT NULL,
    start_date  TIMESTAMPTZ NOT NULL,
    end_date    TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE creatives
(
    id            UUID PRIMARY KEY,
    profile_id    UUID NOT NULL REFERENCES profiles (id),
    media_url     TEXT NOT NULL,
    width         INT  NOT NULL,
    height        INT  NOT NULL,
    creative_type TEXT NOT NULL, -- e.g. banner, native, video
    created_at    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE profile_targeting
(
    id         UUID PRIMARY KEY,
    profile_id UUID NOT NULL REFERENCES profiles (id) ON DELETE CASCADE,
    key        TEXT NOT NULL, -- e.g. "geo", "os", "language"
    value      TEXT NOT NULL  -- e.g. "PL", "Android", "en"
);

CREATE TABLE profile_package
(
    profile_id UUID REFERENCES profiles (id) ON DELETE CASCADE,
    package_id UUID,
    CONSTRAINT profile_package_pkey PRIMARY KEY (profile_id, package_id)
)

