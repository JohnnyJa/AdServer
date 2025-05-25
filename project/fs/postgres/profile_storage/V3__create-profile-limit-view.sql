CREATE VIEW profile_limit_view AS
SELECT
    p.id as profile_id,
    p.views_limit as profile_limit
FROM profiles p
WHERE now() BETWEEN p.start_date AND p.end_date


