package repository

import "github.com/google/uuid"

type ProfileRow struct {
	ProfileID      uuid.UUID   `db:"profile_id"`
	ProfileName    string      `db:"profile_name"`
	CreativeID     uuid.UUID   `db:"creative_id"`
	MediaURL       string      `db:"media_url"`
	Width          int         `db:"width"`
	Height         int         `db:"height"`
	CreativeType   string      `db:"creative_type"`
	TargetingKey   *string     `db:"key"`
	TargetingValue *string     `db:"value"`
	PackageIDs     []uuid.UUID `db:"package_ids"`
}
