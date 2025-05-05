package model

import "github.com/google/uuid"

type Package struct {
	Id      uuid.UUID   `db:"package_id"`
	Name    string      `db:"name"`
	ZoneIds []uuid.UUID `db:"zones_id"`
}
