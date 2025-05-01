
CREATE TABLE profiles (
                          id UUID PRIMARY KEY,
                          name TEXT,
                          package_id UUID NOT NULL,
                          is_active BOOLEAN DEFAULT TRUE,
                          created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE creatives (
                           id UUID PRIMARY KEY,
                           profile_id UUID NOT NULL REFERENCES profiles(id),
                           media_url TEXT NOT NULL,
                           width INT NOT NULL,
                           height INT NOT NULL,
                           creative_type TEXT NOT NULL, -- e.g. banner, native, video
                           is_active BOOLEAN DEFAULT TRUE,
                           created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE creative_targetings (
                                     id UUID PRIMARY KEY,
                                     creative_id UUID NOT NULL REFERENCES creatives(id) ON DELETE CASCADE,
                                     key TEXT NOT NULL,           -- e.g. "geo", "os", "language"
                                     value TEXT NOT NULL          -- e.g. "PL", "Android", "en"
);

