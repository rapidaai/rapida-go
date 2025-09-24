# Rapida SDK

[![Build and Publish Golang SDK](https://github.com/rapidaai/rapida-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/rapidaai/rapida-go/actions/workflows/build.yml)

The Rapida SDK provides a powerful interface for interacting with Rapida AI services. This SDK simplifies the process of making API calls, handling authentication, and managing responses from Rapida endpoints.

## Installation

To install the Rapida SDK, use the following command:

```
go get github.com/rapidaai/rapida-go@v0.0.7
```

## Quick Start

Here's how to get started with the Rapida SDK:

```go
package main

import (
    "context"
    "fmt"
    "github.com/rapidaai/rapida-go/rapida/connections"
)

func main() {
    connectionConfig := connections.DefaultConnectionConfig(
        connections.WithSDK("your-api-key-here"), // Replace with your SDK API Key
    )

    client := connectionConfig.CreateClient(context.Background())

    // Use the client to make API calls
    fmt.Println("Client initialized successfully!", client)
}
```

## Authentication

You can configure the Rapida SDK to authenticate using your **API Key** or **Personal Token**:

### Authenticating with API Key

```go
import (
    "os"
    "github.com/rapidaai/rapida-go/rapida/connections"
)

connectionConfig := connections.DefaultConnectionConfig(
    connections.WithSDK(
        os.Getenv("RAPIDA_API_KEY"), // API Key from environment variables
    ),
)
```

### Authenticating with Personal Token

```go
import (
    "os"
    "github.com/rapidaai/rapida-go/rapida/connections"
)

connectionConfig := connections.DefaultConnectionConfig(
    connections.WithPersonalToken(
        os.Getenv("RAPIDA_AUTHORIZATION_TOKEN"), // Authorization Token
        os.Getenv("RAPIDA_AUTH_ID"),            // Authentication ID
        os.Getenv("RAPIDA_PROJECT_ID"),         // Project ID
    ),
)
```

## Initializing a Client

Once the configuration is set, you can initialize the Rapida client. Example:

```go
package main

import (
    "context"
    "github.com/rapidaai/rapida-go/rapida/connections"
)

func main() {
    connectionConfig := connections.DefaultConnectionConfig(
        connections.WithSDK("your-api-key-here"),
    )

    client := connectionConfig.CreateClient(context.Background())

    // Client ready for API calls
    fmt.Println("Client initialized successfully!")
}
```

## Invoking API Endpoints

Here’s an example of how to call a specific API endpoint with the initialized client:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/rapidaai/rapida-go/rapida"
    "github.com/rapidaai/rapida-go/rapida/builders"
    "github.com/rapidaai/rapida-go/rapida/definitions"
)

func main() {
    connectionConfig := connections.DefaultConnectionConfig(
        connections.WithSDK("your-api-key-here"),
    )

    client := connectionConfig.CreateClient(context.Background())

    endpoint, err := definitions.NewEndpoint(2084859551571509248, "vrsn_2084859551600869376")
    if err != nil {
        log.Fatalf("Failed to create endpoint: %v", err)
    }

    request := builders.InvokeRequestBuilder(endpoint).Build()
    response, err := client.Invoke(request)
    if err != nil {
        log.Fatalf("API call failed: %v", err)
    }

    data, err := response.GetData()
    if err != nil {
        log.Fatalf("Failed to process response: %v", err)
    }

    for _, record := range data {
        fmt.Println(record.ToText())
    }
}
```

## Configuration Options

The `DefaultConnectionConfig` accepts multiple options for configuring the SDK. Key options include:

- `WithSDK(apiKey string)`: Sets the API key for authentication.
- `WithPersonalToken(authToken, authId, projectId string)`: Configures the connection for personal tokens.
- `WithEndpointURL(url string)`: Overrides the default API endpoint URL.
- `WithTimeout(timeout time.Duration)`: Sets the timeout for API requests.

## Compatibility

This SDK requires **Go 1.18 or later**. Ensure your system meets this requirement:

```bash
go version
```

To upgrade or specify a version, use the following command:

```bash
go get github.com/rapidaai/rapida-go@v0.0.7
```

---

## Conclusion

The Rapida SDK provides everything necessary to integrate seamlessly with Rapida AI services, offering flexible configuration and authentication options. With the examples provided, you should be able to get started quickly and make advanced API calls as needed.