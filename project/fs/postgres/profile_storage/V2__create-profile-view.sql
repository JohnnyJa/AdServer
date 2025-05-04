CREATE VIEW active_profile_view AS
        SELECT
            p.id as profile_id,
            p.name as profile_name,
            c.id as creative_id,
            c.media_url,
            c.width,
            c.height,
            c.creative_type,
            ct.key,
            ct.value,
            ARRAY_AGG(DISTINCT pp.package_id) AS package_ids
        FROM profiles p
            JOIN profile_package pp ON pp.profile_id = p.id
            JOIN creatives c on p.id = c.profile_id
            LEFT JOIN creative_targeting ct on ct.creative_id = c.id
        WHERE now() BETWEEN p.start_date AND p.end_date
        GROUP BY
            p.id, p.name,
            c.id, c.media_url, c.width, c.height, c.creative_type,
            ct.key, ct.value;
