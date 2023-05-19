package client

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type InviteResponseStore struct {
	client *redis.Client
}

func NewInviteResponseStore(client *redis.Client) *InviteResponseStore {
	return &InviteResponseStore{
		client: client,
	}
}

func (irs *InviteResponseStore) SaveNewInviteEntry(phone, inviteID string) error {
	return irs.addValToOrderedSet(phone, inviteID)
}

func (irs *InviteResponseStore) FetchAllInvites(phone string) (map[float64]string, error) {
	return irs.getMapFromOrderedSet(phone)
}

func (irs *InviteResponseStore) addValToOrderedSet(key string, val string) error {
	// Retrieve the highest weight from the ordered set
	highestWeight, err := irs.client.ZRevRangeWithScores(context.Background(), key, 0, 0).Result()
	if err != nil {
		return err
	}

	weight := 1.0
	if len(highestWeight) > 0 {
		weight = highestWeight[0].Score + 1
	}

	// Add the new inviteID to the ordered set with the incremented weight
	err = irs.client.ZAdd(context.Background(), key, redis.Z{
		Score:  weight,
		Member: val,
	}).Err()
	return err
}

func (irs *InviteResponseStore) getAllKeysFromOrderedSet(key string) ([]redis.Z, error) {
	// Retrieve all values from the ordered set with their weights
	results, err := irs.client.ZRangeWithScores(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (irs *InviteResponseStore) getMapFromOrderedSet(key string) (map[float64]string, error) {
	results, err := irs.getAllKeysFromOrderedSet(key)
	if err != nil {
		return nil, err
	}

	// Convert the results to a map with weights as keys
	resultsMap := make(map[float64]string)
	for _, result := range results {
		resultsMap[result.Score] = result.Member.(string)
	}
	return resultsMap, nil
}
