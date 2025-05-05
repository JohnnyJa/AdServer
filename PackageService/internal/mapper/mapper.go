package mapper

import (
	"github.com/google/uuid"
)

func StringsToUUIDs(input []string) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, len(input))
	for i, item := range input {
		uuid, err := uuid.Parse(item)
		if err != nil {
			return nil, err
		}
		result[i] = uuid
	}
	return result, nil
}

func UUIDsToStrings(input []uuid.UUID) []string {
	result := make([]string, len(input))
	for i, item := range input {
		result[i] = item.String()
	}
	return result
}
