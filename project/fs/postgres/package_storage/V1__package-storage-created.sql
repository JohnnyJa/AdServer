CREATE TABLE package (
                         id UUID PRIMARY KEY,
                         name TEXT,
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE package_zone (
            package_id UUID REFERENCES package(id) ON DELETE CASCADE ,
            zone_id UUID,
            CONSTRAINT package_zone_pkey PRIMARY KEY (package_id, zone_id)
)


