// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"testing"
)

func TestRapidaStage_Get(t *testing.T) {
	tests := []struct {
		name     string
		stage    RapidaStage
		expected string
	}{
		{
			name:     "authentication stage",
			stage:    AuthenticationStage,
			expected: "user-authentication",
		},
		{
			name:     "transcription stage",
			stage:    TranscriptionStage,
			expected: "audio-transcription",
		},
		{
			name:     "assistant identification stage",
			stage:    AssistantIdentificaionStage,
			expected: "assistant-identificaion",
		},
		{
			name:     "undefined stage",
			stage:    UndefinedStage,
			expected: "undefined",
		},
		{
			name:     "query formulation stage",
			stage:    QueryFormulationStage,
			expected: "query-formulation",
		},
		{
			name:     "information retrieval stage",
			stage:    InformationRetrievalStage,
			expected: "information-retrieval",
		},
		{
			name:     "document retrieval stage",
			stage:    DocumentRetrievalStage,
			expected: "document-retrieval",
		},
		{
			name:     "context augmentation stage",
			stage:    ContextAugmentationStage,
			expected: "context-augmentation",
		},
		{
			name:     "text generation stage",
			stage:    TextGenerationStage,
			expected: "text-generation",
		},
		{
			name:     "tool call stage",
			stage:    ToolCallStage,
			expected: "tool-call",
		},
		{
			name:     "tool execute stage",
			stage:    ToolExecuteStage,
			expected: "tool-execute",
		},
		{
			name:     "output evaluation stage",
			stage:    OutputEvaluationStage,
			expected: "output-evaluation",
		},
		{
			name:     "get tool stage",
			stage:    GetToolStage,
			expected: "get-tool",
		},
		{
			name:     "provider model identification stage",
			stage:    ProviderModelIdentificationStage,
			expected: "model-identificaion",
		},
		{
			name:     "create message stage",
			stage:    CreateMessageStage,
			expected: "create-message",
		},
		{
			name:     "update message stage",
			stage:    UpdateMessageStage,
			expected: "update-message",
		},
		{
			name:     "create message metric stage",
			stage:    CreateMessageMetricStage,
			expected: "create-message-metric",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.stage.Get()
			if result != tt.expected {
				t.Errorf("RapidaStage.Get() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestRapidaStageConstants(t *testing.T) {
	// Verify all constants have the expected values
	constants := map[RapidaStage]string{
		AuthenticationStage:              "user-authentication",
		TranscriptionStage:               "audio-transcription",
		AssistantIdentificaionStage:      "assistant-identificaion",
		UndefinedStage:                   "undefined",
		QueryFormulationStage:            "query-formulation",
		InformationRetrievalStage:        "information-retrieval",
		DocumentRetrievalStage:           "document-retrieval",
		ContextAugmentationStage:         "context-augmentation",
		TextGenerationStage:              "text-generation",
		ToolCallStage:                    "tool-call",
		ToolExecuteStage:                 "tool-execute",
		OutputEvaluationStage:            "output-evaluation",
		GetToolStage:                     "get-tool",
		ProviderModelIdentificationStage: "model-identificaion",
		CreateMessageStage:               "create-message",
		UpdateMessageStage:               "update-message",
		CreateMessageMetricStage:         "create-message-metric",
	}

	for stage, expected := range constants {
		if string(stage) != expected {
			t.Errorf("RapidaStage constant %v = %q, want %q", stage, string(stage), expected)
		}
	}
}

func TestRapidaStage_TypeConversion(t *testing.T) {
	// Test that we can convert string to RapidaStage and back
	stages := []RapidaStage{
		AuthenticationStage,
		TranscriptionStage,
		TextGenerationStage,
		ToolCallStage,
	}

	for _, stage := range stages {
		// Convert to string and back
		str := string(stage)
		backToStage := RapidaStage(str)
		if backToStage != stage {
			t.Errorf("Type conversion failed for %v: got %v after round trip", stage, backToStage)
		}
	}
}
