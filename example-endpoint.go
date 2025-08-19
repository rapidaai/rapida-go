package main

import (
	"fmt"

	"github.com/rapidaai/rapida-go/rapida"
	rapida_builders "github.com/rapidaai/rapida-go/rapida/builders"
)

func main() {
	text()
	audio()
}

func assistant() {
	fmt.Println("Running an text endpoint example")
	client, _ := rapida.GetClient(rapida_builders.
		ClientOptionBuilder().
		WithApiKey("rpd-prj-0360813184560c71cf11ecd7945e1aeadc9894d0912c9fa4edb001cf4319e404").
		WithEndpointUrl("localhost:9005").
		Build())

	assistant, _ := rapida_builders.NewAssistantDefinitionBuilder().
		WithAssistantId(2091457615258189824).
		WithAssistantVersion("vrsn_2091457615270772736").
		Build()

	client.SendMessage(rapida_builders.
		NewMessagingRequestBuilder(assistant).
		WithMessage(rapida_builders.NewMessageBuilder().AddContent(
			rapida_builders.NewContentBuilder().WithString("Hello, how are you doing?").Build(),
		).Build()).
		Build())
}

func text() {
	fmt.Println("Running an text endpoint example")
	client, _ := rapida.GetClient(rapida_builders.
		ClientOptionBuilder().
		WithApiKey("rpd-prj-0360813184560c71cf11ecd7945e1aeadc9894d0912c9fa4edb001cf4319e404").
		WithEndpointUrl("localhost:9005").
		Build())
	endpoint, err := rapida_builders.NewEndpointDefinitionBuilder().
		WithEndpointId(2084270574271463424).
		Build()
	fmt.Println("Calling with %+v", endpoint)
	requestBuilder := rapida_builders.NewInvokeRequestBuilder(endpoint)
	requestBuilder.AddStringInput("doing", "hello")
	res, err := client.Invoke(requestBuilder.
		Build())
	if err == nil {
		if res.IsSuccess() {
			data, _ := res.GetData()
			for _, c := range data {
				cnt, _ := c.ToText()
				println(cnt)
			}
		}
	}
}

func audio() {
	fmt.Println("Running an audio endpoint example")
	client, err := rapida.GetClient(rapida_builders.
		ClientOptionBuilder().
		WithApiKey("0360813184560c71cf11ecd7945e1aeadc9894d0912c9fa4edb001cf4319e404").
		WithEndpointUrl("localhost:9005").
		Build())

	if err != nil {
		fmt.Println("Getclient error with %+v", err)
		return
	}

	endpoint, err := rapida_builders.NewEndpointDefinitionBuilder().
		WithEndpointId(2091069247890391040).
		WithEndpointVersion("vrsn_2091069247986860032").
		Build()
	fmt.Println("Calling with %+v", endpoint)
	requestBuilder := rapida_builders.NewInvokeRequestBuilder(endpoint)
	err = requestBuilder.AddStringOption("language", "en")
	fmt.Println("Calling with err%+v", err)
	err = requestBuilder.AddBooleanOption("sentiment", true)

	err = requestBuilder.AddStringMetadata("sentiment", "chahiye")
	fmt.Println("Calling with err%+v", err)
	requestBuilder.AddFileInput("file", "/Users/prashant.srivastav/Documents/codes/lexatic/cookbook/endpoints/audio-recording-sentiment-intent-detection/agent-call-recording-audio.wav")
	res, err := client.Invoke(requestBuilder.
		Build())
	if err != nil {
		print(err)

	}
	if res.IsSuccess() {
		data, _ := res.GetData()
		for _, c := range data {
			cnt, _ := c.ToText()
			println(cnt)
		}
	}
	print(res.GetHumanErrorMessage())
}
