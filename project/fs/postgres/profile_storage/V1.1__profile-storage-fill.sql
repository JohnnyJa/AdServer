-- Insert into profiles
INSERT INTO profiles (id, name, bid_price, views_limit, start_date, end_date)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'Summer Promo', 1.50, 100, '2025-05-01 00:00:00+00', '2025-06-01 00:00:00+00'),
    ('22222222-2222-2222-2222-222222222222', 'Winter Sale', 2.10, 2, '2025-05-01 00:00:00+00', '2026-01-15 00:00:00+00');

-- Insert into creatives
INSERT INTO creatives (id, profile_id, media_url, width, height, creative_type)
VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', 'https://cdn.example.com/banner1.jpg', 300, 250, 'banner'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '11111111-1111-1111-1111-111111111111', 'https://cdn.example.com/banner2.jpg', 728, 90, 'banner'),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', '22222222-2222-2222-2222-222222222222', 'https://cdn.example.com/video1.mp4', 1920, 1080, 'banner');

-- Insert into profile_targeting
INSERT INTO profile_targeting (id, profile_id, key, value)
VALUES
    ('aaaa1111-0000-0000-0000-000000000001', '11111111-1111-1111-1111-111111111111', 'ImpTargeting', 'Android'),


-- Insert into profile_package
INSERT INTO profile_package (profile_id, package_id)
VALUES
    ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111'),
    ('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111'),
    ('22222222-2222-2222-2222-222222222222', '22222222-2222-2222-2222-222222222222');
