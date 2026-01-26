// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

type RapidaEvent string

const (
	// signals to stop speaking
	// what happend is when user unintentionally interrupt the the voice completly stop
	// pause will make sure that it's not unintentionally interrupted
	TalkPause        RapidaEvent = "talk.onPause"
	TalkInterruption RapidaEvent = "talk.onInterrupt"

	//
	TalkTranscript RapidaEvent = "talk.onTranscript"
	// start and complete
	TalkStart    RapidaEvent = "talk.onStart"
	TalkComplete RapidaEvent = "talk.onComplete"

	TalkGeneration         RapidaEvent = "talk.onGeneration"
	TalkCompleteGeneration RapidaEvent = "talk.onCompleteGeneration"

	TalkStartConversation    RapidaEvent = "talk.onStartConversation"
	TalkCompleteConversation RapidaEvent = "talk.onCompleteConversation"
)

// Get returns the string value of the RapidaStage
func (r RapidaEvent) Get() string {
	return string(r)
}

type AssistantServerEvent string

const (
	AssistantInitiated AssistantWebhookEvent = "conversation.initiated"
	// Triggered when a new conversation is started.
)

func (r AssistantServerEvent) Get() string {
	return string(r)
}

type AssistantWebhookEvent string

const (
	MessageReceived AssistantWebhookEvent = "message.received"
	// Triggered when a message is received.

	MessageSent AssistantWebhookEvent = "message.sent"
	// Triggered when a mes

	ConversationBegin     AssistantWebhookEvent = "conversation.begin"
	ConversationResume    AssistantWebhookEvent = "conversation.resume"
	ConversationCompleted AssistantWebhookEvent = "conversation.completed"
	// Triggered when a conversation ends successfully.

	ConversationFailed AssistantWebhookEvent = "conversation.failed"
	// Triggered when a conversation encounters an error.

)

func (r AssistantWebhookEvent) Get() string {
	return string(r)
}
