package repository

import (
	"context"
	"fmt"
	"github.com/JohnnyJa/AdServer/PackageService/internal/model"
	"github.com/JohnnyJa/AdServer/PackageService/service"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	service.Service
	ReadPackages(ctx context.Context, packageIds []uuid.UUID) ([]model.Package, error)
}

type repository struct {
	config *Config
	logger *logrus.Logger
	db     *pgxpool.Pool
}

func New(config *Config, logger *logrus.Logger) Repository {
	return &repository{
		config: config,
		logger: logger,
	}
}

func (r *repository) ReadPackages(ctx context.Context, packageIds []uuid.UUID) ([]model.Package, error) {
	query := `
		SELECT p.id, p.name, array_agg(pz.zone_id) as zone_ids
		FROM package p
		LEFT JOIN package_zone pz ON p.id = pz.package_id
		WHERE p.id = ANY($1)
		GROUP BY p.id
	`
	rows, err := r.db.Query(ctx, query, packageIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Package
	for rows.Next() {
		var r model.Package
		err := rows.Scan(
			&r.Id,
			&r.Name,
			&r.ZoneIds,
		)
		if err != nil {
			return nil, fmt.Errorf("scan active_profile_view: %w", err)
		}
		result = append(result, r)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("iteration error: %w", rows.Err())
	}

	return result, nil
}

func (r *repository) Start() error {
	conn, err := pgxpool.New(context.Background(), r.config.ConnectionString)
	if err != nil {
		return err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return err
	}

	r.db = conn
	return nil
}

func (r *repository) Stop() error {
	r.db.Close()
	return nil
}
