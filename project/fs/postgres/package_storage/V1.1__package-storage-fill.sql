-- Створимо два пакети
INSERT INTO package (id, name) VALUES
                                   ('11111111-1111-1111-1111-111111111111', 'Test Package A'),
                                   ('22222222-2222-2222-2222-222222222222', 'Test Package B');

-- Прив’яжемо зони до цих пакетів
INSERT INTO package_zone (package_id, zone_id) VALUES
                                                   ('11111111-1111-1111-1111-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
                                                   ('11111111-1111-1111-1111-111111111111', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'),
                                                   ('22222222-2222-2222-2222-222222222222', 'cccccccc-cccc-cccc-cccc-cccccccccccc');
