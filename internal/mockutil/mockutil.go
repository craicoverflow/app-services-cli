package mockutil

import (
	"context"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
)

func NewConfigMock(cfg *config.Config) config.IConfig {
	return &config.IConfigMock{
		LocationFunc: func() (string, error) {
			return ":mock_location:", nil
		},
		LoadFunc: func() (*config.Config, error) {
			return cfg, nil
		},
		SaveFunc: func(c *config.Config) error {
			cfg = c
			return nil
		},
		RemoveFunc: func() error {
			cfg = nil
			return nil
		},
	}
}

func NewConnectionMock(conn *connection.Connection, apiClient *managedservices.APIClient) connection.IConnection {
	return &connection.IConnectionMock{
		RefreshTokensFunc: func(ctx context.Context) (string, string, error) {
			if conn.Token.AccessToken == "" && conn.Token.RefreshToken == "" {
				return "", "", connection.ErrNotLoggedIn
			}
			if conn.Token.RefreshToken == "expired" {
				return "", "", connection.ErrSessionExpired
			}

			return "valid", "valid", nil
		},
		LogoutFunc: func(ctx context.Context) error {
			if conn.Token.AccessToken == "" && conn.Token.RefreshToken == "" {
				return connection.ErrNotLoggedIn
			}
			if conn.Token.AccessToken == "expired" && conn.Token.RefreshToken == "expired" {
				return connection.ErrSessionExpired
			}

			return nil
		},
		NewMASClientFunc: func() *managedservices.APIClient {
			return apiClient
		},
	}
}

func NewKafkaRequestTypeMock(name string) managedservices.KafkaRequest {
	var kafkaReq managedservices.KafkaRequest
	kafkaReq.SetId("1")
	kafkaReq.SetName(name)

	return kafkaReq
}