package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deyvidm/sms/cmd/dispatcher/log"
)

const defaultAddress = "http://localhost:8080"

type WebBackendClient struct {
	BearerToken string
	Address     string
	client      *http.Client
}

func NewWebBackendClient(BearerToken string) *WebBackendClient {
	return &WebBackendClient{
		BearerToken: BearerToken,
		Address:     defaultAddress,
		client:      &http.Client{},
	}
}

type UpdateInvite struct {
	ID     string  `json:"-"`
	Status *string `json:"status,omitempty"`
	Paid   *bool   `json:"paid,omitempty"`
}

var logger = log.GetLogger()

func (wbc *WebBackendClient) UpdateInvite(invite UpdateInvite) error {
	url := "/api/internal/invite/" + invite.ID
	logger.Infof("\tupdating invite %s with %s", invite.ID, *invite.Status)
	bod, err := json.Marshal(invite)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, wbc.Address+url, bytes.NewBuffer(bod))
	if err != nil {
		logger.Errorf("Error creating request: %+v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer astynq")

	resp, err := wbc.client.Do(req)
	if err != nil {
		logger.Errorf("Error sending request: %+v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got %s Response trying to update %s", resp.Status, invite.ID)
	}
	return nil
}
