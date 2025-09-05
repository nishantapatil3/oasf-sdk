package service

import (
	"encoding/json"
	"errors"
	"fmt"

	objectsv3 "buf.build/gen/go/agntcy/oasf/protocolbuffers/go/objects/v3"
)

func buildA2ACard(record *objectsv3.Record) (*A2ACard, error) {
	var a2aExt *objectsv3.Extension
	for _, ext := range record.Extensions {
		if ext.Name == "schema.oasf.agntcy.org/features/runtime/a2a" {
			a2aExt = ext
			break
		}
	}

	if a2aExt == nil {
		return nil, errors.New("A2A extension not found in record")
	}

	jsonBytes, err := json.Marshal(a2aExt.Data.AsMap())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal A2A data to JSON: %w", err)
	}

	var card A2ACard
	if err := json.Unmarshal(jsonBytes, &card); err != nil {
		return nil, fmt.Errorf("failed to unmarshal A2A data into A2ACard: %w", err)
	}

	return &card, nil
}

func buildA2ARecord() (*objectsv3.Record, error) {
	// TODO

	return nil, nil
}
