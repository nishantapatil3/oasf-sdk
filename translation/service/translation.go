// Copyright AGNTCY Contributors (https://github.com/agntcy)
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"encoding/json"
	"fmt"

	translationv1 "buf.build/gen/go/agntcy/oasf-sdk/protocolbuffers/go/translation/v1"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func NewTranslationService() *TranslationService {
	return &TranslationService{}
}

func (t TranslationService) RecordToVSCodeCopilot(req *translationv1.RecordToVSCodeCopilotRequest) (*structpb.Struct, error) {
	vsCodeCopilotMCPConfig, err := buildVSCodeCopilotMCPConfig(req.Record)
	if err != nil {
		return nil, fmt.Errorf("failed to build VSCode MCP config: %w", err)
	}

	return toStruct(map[string]any{
		"mcpConfig": *vsCodeCopilotMCPConfig,
	})
}

func (t TranslationService) GHCopilotToRecord() (*structpb.Struct, error) {
	GHCopilotRecord, err := buildGHCopilotRecord()
	if err != nil {
		return nil, fmt.Errorf("failed to build MCP - Record: %w", err)
	}

	return toStruct(map[string]any{
		"record": GHCopilotRecord,
	})

}

func (t TranslationService) RecordToA2A(req *translationv1.RecordToA2ARequest) (*structpb.Struct, error) {
	a2aCard, err := buildA2ACard(req.Record)
	if err != nil {
		return nil, fmt.Errorf("failed to build A2A card: %w", err)
	}

	return toStruct(map[string]any{
		"a2aCard": *a2aCard,
	})
}

func (t TranslationService) A2AToRecord() (*structpb.Struct, error) {
	A2ARecord, err := buildA2ARecord()
	if err != nil {
		return nil, fmt.Errorf("failed to build A2A - Record: %w", err)
	}

	return toStruct(map[string]any{
		"record": A2ARecord,
	})
}

func toStruct(v any) (*structpb.Struct, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}
	return structpb.NewStruct(data)
}
