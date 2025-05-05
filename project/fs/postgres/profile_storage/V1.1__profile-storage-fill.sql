-- Профайли
INSERT INTO profiles (id, name, start_date, end_date)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'Profile A', now() - interval '1 day', now() + interval '5 days'),
    ('22222222-2222-2222-2222-222222222222', 'Profile B', now() - interval '2 days', now() + interval '3 days'),
    ('33333333-3333-3333-3333-333333333333', 'Expired Profile', now() - interval '10 days', now() - interval '5 days');

-- Креативи
INSERT INTO creatives (id, profile_id, media_url, width, height, creative_type)
VALUES
    ('aaa11111-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', 'https://cdn.example.com/banner1.jpg', 300, 250, 'banner'),
    ('aaa22222-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', 'https://cdn.example.com/banner2.jpg', 728, 90, 'banner'),
    ('bbb11111-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '22222222-2222-2222-2222-222222222222', 'https://cdn.example.com/native1.jpg', 600, 400, 'native'),
    ('ccc11111-cccc-cccc-cccc-cccccccccccc', '33333333-3333-3333-3333-333333333333', 'https://cdn.example.com/expired.jpg', 300, 250, 'banner');

-- Таргетинги
INSERT INTO creative_targeting (id, creative_id, key, value)
VALUES
    ('ddd11111-dddd-dddd-dddd-dddddddddddd', 'aaa11111-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'geo', 'PL'),
    ('ddd22222-dddd-dddd-dddd-dddddddddddd', 'aaa22222-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'geo', 'UA'),
    ('ddd33333-dddd-dddd-dddd-dddddddddddd', 'bbb11111-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'os', 'Android'),
    ('ddd44444-dddd-dddd-dddd-dddddddddddd', 'ccc11111-cccc-cccc-cccc-cccccccccccc', 'geo', 'DE');

-- Профайл-Пакет звʼязки
INSERT INTO profile_package (profile_id, package_id)
VALUES
    ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111'),
    ('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111'),
    ('33333333-3333-3333-3333-333333333333', '22222222-2222-2222-2222-222222222222');
