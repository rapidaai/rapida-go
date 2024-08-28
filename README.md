# Rapida SDK

The Rapida SDK provides a powerful interface for interacting with Rapida AI services. This SDK simplifies the process of making API calls, handling authentication, and managing responses from Rapida endpoints.

## Installation

To install the Rapida SDK, use the following command:

```
go get github.com/rapidaai/rapida-sdk/rapida
```

## Quick Start

Here's a basic example to get you started with the Rapida SDK:

```go
package main

import (
    "context"
    "fmt"
    "github.com/rapidaai/rapida-sdk/rapida"
)

func main() {
    options := rapida.NewRapidaClientOption()
    options.SetRapidaApiKey("your-api-key-here")

    client, err := rapida.GetClient(options)
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }

    // Use the client to make API calls
    // ...
}
```

## Key Components

### RapidaClientOption

The `RapidaClientOption` struct allows you to configure the SDK client. Key methods include:

- `NewRapidaClientOption()`: Creates a new option struct with default values.
- `NewRapidaClientOptionWithParams()`: Creates a new option struct with custom parameters.
- `SetRapidaApiKey()`: Sets the API key for authentication.
- `SetRapidaEndpointUrl()`: Sets the endpoint URL.
- `SetRapidaAssistantUrl()`: Sets the assistant URL.
- `SetRapidaEnvironment()`: Sets the environment (e.g., production, staging).
- `SetRapidaRegion()`: Sets the region for API calls.
- `SetSecure()`: Sets whether to use a secure connection.

### Initialize Rapida Client

The `Client` interface defines the core functionality for making API calls:

```go
client, _ := rapida.GetClient(rapida_builders.ClientOptionBuilder().WithApiKey("{RAPIDA_KEY}").Build())
```

### Calling Rapida endpoint

The `Client` interface defines the core functionality for making API calls:

```go

endpoint, _ := rapida_definitions.NewEndpoint(2084859551571509248, "vrsn_2084859551600869376")
res, _ := client.Invoke(rapida_builders.InvokeRequestBuilder(endpoint).Build())
data, err := res.GetData()
if err == nil {
    for _, c := range data {
        print(c.ToText())
    }
}

```
