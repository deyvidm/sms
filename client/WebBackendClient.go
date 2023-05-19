package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const defaultAddress = "http://localhost:8080"

type WebBackendClient struct {
	BearerToken string
	Address     string
	client      *http.Client
}

func New(BearerToken string) WebBackendClient {
	return WebBackendClient{
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

func (wbc *WebBackendClient) UpdateInvite(invite *UpdateInvite) error {
	url := "/api/internal/invite/" + invite.ID
	fmt.Println("shooting to ", url)
	bod, err := json.Marshal(invite)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, wbc.Address+url, bytes.NewBuffer(bod))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Set the request header (if needed)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer astynq")

	// Send the request
	resp, err := wbc.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Process the response
	fmt.Println("Response Status:", resp.Status)
	// Read the response body, if required
	// responseBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Response Body:", string(responseBody))

	return nil
}
