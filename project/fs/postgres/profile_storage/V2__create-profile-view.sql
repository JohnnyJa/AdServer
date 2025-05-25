CREATE VIEW active_profile_view AS
        SELECT
            p.id as profile_id,
            p.name as profile_name,
            p.bid_price,
            c.id as creative_id,
            c.media_url,
            c.width,
            c.height,
            c.creative_type,
            pt.key,
            pt.value,
            ARRAY_AGG(DISTINCT pp.package_id) AS package_ids
        FROM profiles p
            JOIN profile_package pp ON pp.profile_id = p.id
            JOIN creatives c on p.id = c.profile_id
            LEFT JOIN profile_targeting pt on pt.profile_id = p.id
        WHERE now() BETWEEN p.start_date AND p.end_date
        GROUP BY
            p.id, p.name, p.bid_price,
            c.id, c.media_url, c.width, c.height, c.creative_type,
            pt.key, pt.value;


