package service

import (
	"errors"
	"fmt"
	"strings"

	objectsv3 "buf.build/gen/go/agntcy/oasf/protocolbuffers/go/objects/v3"
)

func buildVSCodeCopilotMCPConfig(record *objectsv3.Record) (*VSCodeCopilotMCPConfig, error) {
	var mcpExt *objectsv3.Extension
	for _, ext := range record.Extensions {
		if ext.Name == "schema.oasf.agntcy.org/features/runtime/mcp" {
			mcpExt = ext
			break
		}
	}

	if mcpExt == nil {
		return nil, errors.New("MCP extension not found in record")
	}

	serversVal, ok := mcpExt.Data.Fields["servers"]
	if !ok {
		return nil, errors.New("invalid or missing 'servers' in MCP extension data")
	}

	serversStruct := serversVal.GetStructValue()
	if serversStruct == nil {
		return nil, errors.New("'servers' is not a struct")
	}

	servers := make(map[string]Server)
	inputs := []Input{}

	for serverName, serverVal := range serversStruct.Fields {
		serverMap := serverVal.GetStructValue()
		if serverMap == nil {
			continue
		}

		command, ok := serverMap.Fields["command"]
		if !ok {
			return nil, fmt.Errorf("missing 'command' for server '%s'", serverName)
		}

		args := []string{}
		if argsVal, ok := serverMap.Fields["args"]; ok {
			for _, arg := range argsVal.GetListValue().Values {
				args = append(args, arg.GetStringValue())
			}
		}

		env := map[string]string{}
		if envVal, ok := serverMap.Fields["env"]; ok {
			envStruct := envVal.GetStructValue()
			if envStruct != nil {
				for key, val := range envStruct.Fields {
					env[key] = val.GetStringValue()

					if after, ok0 := strings.CutPrefix(val.GetStringValue(), "${input:"); ok0 {
						id := strings.TrimSuffix(after, "}")
						inputs = append(inputs, Input{
							ID:          id,
							Type:        "promptString",
							Password:    true,
							Description: fmt.Sprintf("Secret value for %s", id),
						})
					}
				}
			}
		}

		servers[serverName] = Server{
			Command: command.GetStringValue(),
			Args:    args,
			Env:     env,
		}
	}

	return &VSCodeCopilotMCPConfig{
		Servers: servers,
		Inputs:  inputs,
	}, nil
}

func buildGHCopilotRecord() (*objectsv3.Record, error) {
	// TODO

	return nil, nil
}
