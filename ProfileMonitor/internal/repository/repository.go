package repository

import (
	"context"
	"fmt"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	service.Service
	ReadProfiles(ctx context.Context) ([]ProfileRow, error)
	ReadProfilesLimits(ctx context.Context) ([]ProfileLimitsRow, error)
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

func (r *repository) ReadProfiles(ctx context.Context) ([]ProfileRow, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
    		profile_id, profile_name, bid_price,
		    creative_id, media_url, width, height, creative_type,
    		key, value,
    		package_ids
		FROM active_profile_view p
`)

	if err != nil {
		return nil, fmt.Errorf("query active_profile_view: %w", err)
	}
	defer rows.Close()

	var result []ProfileRow
	for rows.Next() {
		var r ProfileRow
		err := rows.Scan(
			&r.ProfileID,
			&r.ProfileName,
			&r.BidPrice,
			&r.CreativeID,
			&r.MediaURL,
			&r.Width,
			&r.Height,
			&r.CreativeType,
			&r.TargetingKey,
			&r.TargetingValue,
			&r.PackageIDs,
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

func (r *repository) ReadProfilesLimits(ctx context.Context) ([]ProfileLimitsRow, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
    		profile_id, profile_limit
		FROM profile_limit_view p
`)
	if err != nil {
		return nil, fmt.Errorf("query profile_limit_view: %w", err)
	}

	defer rows.Close()
	var result []ProfileLimitsRow
	for rows.Next() {
		var r ProfileLimitsRow
		err := rows.Scan(
			&r.ProfileID,
			&r.ViewsLimit,
		)
		if err != nil {
			return nil, fmt.Errorf("scan profile_limit_view: %w", err)
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
