// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"testing"
)

func TestRapidaEvent_Get(t *testing.T) {
	tests := []struct {
		name     string
		event    RapidaEvent
		expected string
	}{
		{
			name:     "talk pause event",
			event:    TalkPause,
			expected: "talk.onPause",
		},
		{
			name:     "talk interruption event",
			event:    TalkInterruption,
			expected: "talk.onInterrupt",
		},
		{
			name:     "talk transcript event",
			event:    TalkTranscript,
			expected: "talk.onTranscript",
		},
		{
			name:     "talk start event",
			event:    TalkStart,
			expected: "talk.onStart",
		},
		{
			name:     "talk complete event",
			event:    TalkComplete,
			expected: "talk.onComplete",
		},
		{
			name:     "talk generation event",
			event:    TalkGeneration,
			expected: "talk.onGeneration",
		},
		{
			name:     "talk complete generation event",
			event:    TalkCompleteGeneration,
			expected: "talk.onCompleteGeneration",
		},
		{
			name:     "talk start conversation event",
			event:    TalkStartConversation,
			expected: "talk.onStartConversation",
		},
		{
			name:     "talk complete conversation event",
			event:    TalkCompleteConversation,
			expected: "talk.onCompleteConversation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.event.Get()
			if result != tt.expected {
				t.Errorf("RapidaEvent.Get() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestAssistantServerEvent_Get(t *testing.T) {
	// Test the Get method for AssistantServerEvent type
	event := AssistantServerEvent("test.event")
	result := event.Get()
	if result != "test.event" {
		t.Errorf("AssistantServerEvent.Get() = %q, want %q", result, "test.event")
	}
}

func TestAssistantWebhookEvent_Get(t *testing.T) {
	tests := []struct {
		name     string
		event    AssistantWebhookEvent
		expected string
	}{
		{
			name:     "message received event",
			event:    MessageReceived,
			expected: "message.received",
		},
		{
			name:     "message sent event",
			event:    MessageSent,
			expected: "message.sent",
		},
		{
			name:     "conversation begin event",
			event:    ConversationBegin,
			expected: "conversation.begin",
		},
		{
			name:     "conversation resume event",
			event:    ConversationResume,
			expected: "conversation.resume",
		},
		{
			name:     "conversation completed event",
			event:    ConversationCompleted,
			expected: "conversation.completed",
		},
		{
			name:     "conversation failed event",
			event:    ConversationFailed,
			expected: "conversation.failed",
		},
		{
			name:     "assistant initiated event",
			event:    AssistantInitiated,
			expected: "conversation.initiated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.event.Get()
			if result != tt.expected {
				t.Errorf("AssistantWebhookEvent.Get() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestRapidaEventConstants(t *testing.T) {
	// Verify all RapidaEvent constants have the expected values
	events := map[RapidaEvent]string{
		TalkPause:                "talk.onPause",
		TalkInterruption:         "talk.onInterrupt",
		TalkTranscript:           "talk.onTranscript",
		TalkStart:                "talk.onStart",
		TalkComplete:             "talk.onComplete",
		TalkGeneration:           "talk.onGeneration",
		TalkCompleteGeneration:   "talk.onCompleteGeneration",
		TalkStartConversation:    "talk.onStartConversation",
		TalkCompleteConversation: "talk.onCompleteConversation",
	}

	for event, expected := range events {
		if string(event) != expected {
			t.Errorf("RapidaEvent constant %v = %q, want %q", event, string(event), expected)
		}
	}
}

func TestAssistantWebhookEventConstants(t *testing.T) {
	// Verify all AssistantWebhookEvent constants have the expected values
	events := map[AssistantWebhookEvent]string{
		MessageReceived:       "message.received",
		MessageSent:           "message.sent",
		ConversationBegin:     "conversation.begin",
		ConversationResume:    "conversation.resume",
		ConversationCompleted: "conversation.completed",
		ConversationFailed:    "conversation.failed",
		AssistantInitiated:    "conversation.initiated",
	}

	for event, expected := range events {
		if string(event) != expected {
			t.Errorf("AssistantWebhookEvent constant %v = %q, want %q", event, string(event), expected)
		}
	}
}

func TestRapidaEvent_TypeConversion(t *testing.T) {
	// Test that we can convert string to RapidaEvent and back
	events := []RapidaEvent{
		TalkPause,
		TalkInterruption,
		TalkStart,
		TalkComplete,
	}

	for _, event := range events {
		str := string(event)
		backToEvent := RapidaEvent(str)
		if backToEvent != event {
			t.Errorf("Type conversion failed for %v: got %v after round trip", event, backToEvent)
		}
	}
}
