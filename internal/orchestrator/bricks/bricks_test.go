package bricks

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arduino/arduino-app-cli/internal/orchestrator/bricksindex"
)

func TestGetBrickInstanceVariableDetails(t *testing.T) {
	tests := []struct {
		name                     string
		brick                    *bricksindex.Brick
		userVariables            map[string]string
		expectedInstanceVariable []BrickInstanceVariable
		expectedVariableMap      map[string]string
	}{
		{
			name: "variable is present in the map",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", Description: "desc"},
				},
			},
			userVariables: map[string]string{"VAR1": "value1"},
			expectedInstanceVariable: []BrickInstanceVariable{
				{Name: "VAR1", Value: "value1", Description: "desc", Required: true},
			},
			expectedVariableMap: map[string]string{"VAR1": "value1"},
		},
		{
			name: "variable not present in the map",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", Description: "desc"},
				},
			},
			userVariables: map[string]string{},
			expectedInstanceVariable: []BrickInstanceVariable{
				{Name: "VAR1", Value: "", Description: "desc", Required: true},
			},
			expectedVariableMap: map[string]string{"VAR1": ""},
		},
		{
			name: "variable with default value",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", DefaultValue: "default", Description: "desc"},
				},
			},
			userVariables: map[string]string{},
			expectedInstanceVariable: []BrickInstanceVariable{
				{Name: "VAR1", Value: "default", Description: "desc", Required: false},
			},
			expectedVariableMap: map[string]string{"VAR1": "default"},
		},
		{
			name: "multiple variables",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", Description: "desc1"},
					{Name: "VAR2", DefaultValue: "def2", Description: "desc2"},
				},
			},
			userVariables: map[string]string{"VAR1": "v1"},
			expectedInstanceVariable: []BrickInstanceVariable{
				{Name: "VAR1", Value: "v1", Description: "desc1", Required: true},
				{Name: "VAR2", Value: "def2", Description: "desc2", Required: false},
			},
			expectedVariableMap: map[string]string{"VAR1": "v1", "VAR2": "def2"},
		},
		{
			name:                     "no variables",
			brick:                    &bricksindex.Brick{Variables: []bricksindex.BrickVariable{}},
			userVariables:            map[string]string{},
			expectedInstanceVariable: []BrickInstanceVariable{},
			expectedVariableMap:      map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualVariableMap, actualInstanceVariables := getBrickVariableDetails(tt.brick, tt.userVariables)
			require.Equal(t, tt.expectedVariableMap, actualVariableMap)
			require.Equal(t, tt.expectedInstanceVariable, actualInstanceVariables)
		})
	}
}
