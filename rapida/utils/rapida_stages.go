// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

type RapidaStage string

const (
	AuthenticationStage              RapidaStage = "user-authentication"
	TranscriptionStage               RapidaStage = "audio-transcription"
	AssistantIdentificaionStage      RapidaStage = "assistant-identificaion"
	UndefinedStage                   RapidaStage = "undefined"
	QueryFormulationStage            RapidaStage = "query-formulation"
	InformationRetrievalStage        RapidaStage = "information-retrieval"
	DocumentRetrievalStage           RapidaStage = "document-retrieval"
	ContextAugmentationStage         RapidaStage = "context-augmentation"
	TextGenerationStage              RapidaStage = "text-generation"
	ToolCallStage                    RapidaStage = "tool-call"
	ToolExecuteStage                 RapidaStage = "tool-execute"
	OutputEvaluationStage            RapidaStage = "output-evaluation"
	GetToolStage                     RapidaStage = "get-tool"
	ProviderModelIdentificationStage RapidaStage = "model-identificaion"
	CreateMessageStage               RapidaStage = "create-message"
	UpdateMessageStage               RapidaStage = "update-message"
	CreateMessageMetricStage         RapidaStage = "create-message-metric"
)

// Get returns the string value of the RapidaStage
func (r RapidaStage) Get() string {
	return string(r)
}
