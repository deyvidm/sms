package client

import (
	"context"

	"github.com/deyvidm/sms/common/types"
	"github.com/redis/go-redis/v9"
)

// This Store keeps track of recipient invites and their current status.
// We log user invites to an ordered queue in Redis (1 for each user)
type InviteResponseStore struct {
	client *redis.Client
}

func NewInviteResponseStore(client *redis.Client) *InviteResponseStore {
	return &InviteResponseStore{
		client: client,
	}
}

// Remove an invitation from the contact's list,
// and decrement all the following weights so that weight order stays continuous
// ex.   (1 2 3 4 5).Pop(3) => {1-->1, 2-->2, 3-->rm, 4-->3, 5-->4} => {1,2,3,4}
func (irs *InviteResponseStore) PopInvite(phone string, inviteID string) error {
	// Retrieve the popped invite's weight
	poppedWeight, err := irs.client.ZScore(context.Background(), phone, inviteID).Result()
	if err != nil {
		return err
	}

	// Delete the specific member from the sorted set
	err = irs.client.ZRem(context.Background(), phone, inviteID).Err()
	if err != nil {
		return err
	}

	// Adjust the scores of the remaining elements with higher weights
	elements, err := irs.client.ZRangeWithScores(context.Background(), phone, 0, -1).Result()
	if err != nil {
		return err
	}

	for _, element := range elements {
		if element.Score > poppedWeight {
			newScore := element.Score - 1

			err = irs.client.ZAdd(context.Background(), phone, redis.Z{
				Score:  newScore,
				Member: element.Member,
			}).Err()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Save a new x:inviteID pair under contact's pending invites
func (irs *InviteResponseStore) SaveNewInviteEntry(phone, inviteID string) error {
	return irs.addValToOrderedSet(phone, inviteID)
}

func (irs *InviteResponseStore) FetchTargetInviteID(contactID string, parsedResponse types.ResponseInfo) (string, error) {
	invites, err := irs.fetchAllInvites(contactID)
	if err != nil {
		return "", err
	}
	if len(invites) == 1 {
		for _, inv := range invites {
			return inv, nil
		}
	}

	if parsedResponse.TargetInviteKey == nil {
		return "", types.MissingKeyError{PendingInvites: len(invites)}
	}
	return invites[*parsedResponse.TargetInviteKey], nil
}

func (irs *InviteResponseStore) fetchAllInvites(phone string) (map[float64]string, error) {
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
