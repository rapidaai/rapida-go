// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package connections

import (
	"context"
	"testing"

	"github.com/rapidaai/rapida-go/rapida/configs"
)

func TestNewConnectionConfig(t *testing.T) {
	t.Run("default endpoints", func(t *testing.T) {
		auth := RapidaCredential{"api_key": "test-key"}
		cc := NewConnectionConfig(auth, nil, false)

		if cc == nil {
			t.Fatal("NewConnectionConfig() returned nil")
		}

		// Verify it implements ConnectionConfig interface
		var _ ConnectionConfig = cc
	})

	t.Run("custom endpoints", func(t *testing.T) {
		auth := RapidaCredential{"api_key": "test-key"}
		customEndpoints := map[string]string{
			"assistant": "custom-assistant:443",
			"web":       "custom-web:443",
			"endpoint":  "custom-endpoint:443",
		}
		cc := NewConnectionConfig(auth, customEndpoints, false)

		if cc == nil {
			t.Fatal("NewConnectionConfig() returned nil")
		}
	})

	t.Run("insecure connection", func(t *testing.T) {
		auth := RapidaCredential{"api_key": "test-key"}
		cc := NewConnectionConfig(auth, nil, true)

		if cc == nil {
			t.Fatal("NewConnectionConfig() returned nil")
		}
	})

	t.Run("partial custom endpoints", func(t *testing.T) {
		auth := RapidaCredential{"api_key": "test-key"}
		customEndpoints := map[string]string{
			"assistant": "custom-assistant:443",
		}
		cc := NewConnectionConfig(auth, customEndpoints, false)

		if cc == nil {
			t.Fatal("NewConnectionConfig() returned nil")
		}
	})

	t.Run("nil auth", func(t *testing.T) {
		cc := NewConnectionConfig(nil, nil, false)

		if cc == nil {
			t.Fatal("NewConnectionConfig() returned nil with nil auth")
		}
	})
}

func TestConnectionConfig_WithCustomEndpoint(t *testing.T) {
	auth := RapidaCredential{"api_key": "test-key"}
	cc := NewConnectionConfig(auth, nil, false)

	t.Run("update all endpoints", func(t *testing.T) {
		customEndpoints := map[string]string{
			"assistant": "new-assistant:443",
			"web":       "new-web:443",
			"endpoint":  "new-endpoint:443",
		}
		result := cc.WithCustomEndpoint(customEndpoints)

		if result == nil {
			t.Error("WithCustomEndpoint() returned nil")
		}
	})

	t.Run("update single endpoint", func(t *testing.T) {
		customEndpoints := map[string]string{
			"assistant": "single-assistant:443",
		}
		result := cc.WithCustomEndpoint(customEndpoints)

		if result == nil {
			t.Error("WithCustomEndpoint() returned nil")
		}
	})

	t.Run("nil endpoints", func(t *testing.T) {
		result := cc.WithCustomEndpoint(nil)

		if result == nil {
			t.Error("WithCustomEndpoint(nil) returned nil")
		}
	})

	t.Run("empty endpoints map", func(t *testing.T) {
		result := cc.WithCustomEndpoint(map[string]string{})

		if result == nil {
			t.Error("WithCustomEndpoint({}) returned nil")
		}
	})
}

func TestConnectionConfig_WithInsecureConnection(t *testing.T) {
	auth := RapidaCredential{"api_key": "test-key"}
	cc := NewConnectionConfig(auth, nil, false)

	result := cc.WithInsecureConnection()

	if result == nil {
		t.Error("WithInsecureConnection() returned nil")
	}

	// Verify it returns the same interface for chaining
	var _ ConnectionConfig = result
}

func TestConnectionConfig_WithLocal(t *testing.T) {
	auth := RapidaCredential{"api_key": "test-key"}
	cc := NewConnectionConfig(auth, nil, false)

	result := cc.WithLocal()

	if result == nil {
		t.Error("WithLocal() returned nil")
	}

	// Verify it returns the same interface for chaining
	var _ ConnectionConfig = result
}

func TestConnectionConfig_WithAuth(t *testing.T) {
	auth := RapidaCredential{
		"api_key":    "test-key",
		"project_id": "test-project",
	}
	cc := NewConnectionConfig(auth, nil, false)

	ctx := context.Background()
	result := cc.WithAuth(ctx)

	if result == nil {
		t.Error("WithAuth() returned nil context")
	}
}

func TestConnectionConfig_MethodChaining(t *testing.T) {
	auth := RapidaCredential{"api_key": "test-key"}

	// Test method chaining
	cc := NewConnectionConfig(auth, nil, false).
		WithInsecureConnection().
		WithLocal().
		WithCustomEndpoint(map[string]string{"assistant": "custom:443"})

	if cc == nil {
		t.Error("Method chaining returned nil")
	}
}

func TestRapidaCredential(t *testing.T) {
	t.Run("create credential", func(t *testing.T) {
		cred := RapidaCredential{
			"api_key":    "test-api-key",
			"project_id": "test-project-id",
		}

		if cred["api_key"] != "test-api-key" {
			t.Errorf("Credential api_key = %q, want %q", cred["api_key"], "test-api-key")
		}
		if cred["project_id"] != "test-project-id" {
			t.Errorf("Credential project_id = %q, want %q", cred["project_id"], "test-project-id")
		}
	})

	t.Run("empty credential", func(t *testing.T) {
		cred := RapidaCredential{}

		if len(cred) != 0 {
			t.Errorf("Empty credential should have length 0, got %d", len(cred))
		}
	})
}

func TestDefaultEndpoints(t *testing.T) {
	// Verify the default endpoints are used from configs
	auth := RapidaCredential{"api_key": "test-key"}
	_ = NewConnectionConfig(auth, nil, false)

	// These should be the default values from configs package
	if configs.ASSISTANT_API == "" {
		t.Error("configs.ASSISTANT_API should not be empty")
	}
	if configs.WEB_API == "" {
		t.Error("configs.WEB_API should not be empty")
	}
	if configs.ENDPOINT_API == "" {
		t.Error("configs.ENDPOINT_API should not be empty")
	}
}

func TestConnectionConfig_GetGrpcOption(t *testing.T) {
	t.Run("secure connection options", func(t *testing.T) {
		auth := RapidaCredential{"api_key": "test-key"}
		cc := NewConnectionConfig(auth, nil, false).(*connectionConfig)

		opts := cc.GetGrpcOption()

		if len(opts) == 0 {
			t.Error("GetGrpcOption() returned empty slice")
		}
	})

	t.Run("insecure connection options", func(t *testing.T) {
		auth := RapidaCredential{"api_key": "test-key"}
		cc := NewConnectionConfig(auth, nil, true).(*connectionConfig)

		opts := cc.GetGrpcOption()

		if len(opts) == 0 {
			t.Error("GetGrpcOption() returned empty slice for insecure connection")
		}
	})
}

// Note: The actual client creation tests would require a running server
// or mocking the gRPC connections. These tests verify the configuration setup.
func TestConnectionConfig_ClientCreation(t *testing.T) {
	auth := RapidaCredential{"api_key": "test-key"}
	cc := NewConnectionConfig(auth, nil, true)

	// These tests verify that the client creation methods exist and can be called
	// They may return errors since there's no server to connect to
	t.Run("TalkServiceClient creation", func(t *testing.T) {
		_, err := cc.TalkServiceClient()
		// We expect either success or a connection-related error
		// The important thing is that the method doesn't panic
		_ = err // Accept any result
	})

	t.Run("AuthenticationServiceClient creation", func(t *testing.T) {
		_, err := cc.AuthenticationServiceClient()
		_ = err
	})

	t.Run("AssistantServiceClient creation", func(t *testing.T) {
		_, err := cc.AssistantServiceClient()
		_ = err
	})

	t.Run("AssistantDeploymentServiceClient creation", func(t *testing.T) {
		_, err := cc.AssistantDeploymentServiceClient()
		_ = err
	})

	t.Run("DeploymentClient creation", func(t *testing.T) {
		_, err := cc.DeploymentClient()
		_ = err
	})

	t.Run("EndpointServiceClient creation", func(t *testing.T) {
		_, err := cc.EndpointServiceClient()
		_ = err
	})

	t.Run("VaultServiceClient creation", func(t *testing.T) {
		_, err := cc.VaultServiceClient()
		_ = err
	})
}
