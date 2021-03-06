package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/appengine/v2"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type config struct {
	AllowedHostname string `json:"allowed_hostname"`
}

func getConfig() (*config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	appID := appengine.AppID(ctx)

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretmanager client: %w", err)
	}
	defer client.Close()

	// for now require a secret whose name matches the app name
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", appID, appID),
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %w", err)
	}

	out := &config{}
	decoder := json.NewDecoder(bytes.NewReader(result.Payload.Data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(out); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return out, nil
}
