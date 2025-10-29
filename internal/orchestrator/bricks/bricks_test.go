package bricks

import (
	"testing"

	"github.com/arduino/arduino-app-cli/internal/orchestrator/bricksindex"
	"github.com/stretchr/testify/require"
)

func TestGetBrickInstanceVariableDetails(t *testing.T) {
	tests := []struct {
		name                   string
		brick                  *bricksindex.Brick
		brickInstanceVariables map[string]string
		expected               []BrickInstanceVariable
	}{
		{
			name: "variable is present in the map",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", Description: "desc"},
				},
			},
			brickInstanceVariables: map[string]string{"VAR1": "value1"},
			expected: []BrickInstanceVariable{
				{Name: "VAR1", Value: "value1", Description: "desc", Required: true},
			},
		},
		{
			name: "variable not present in the map",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", Description: "desc"},
				},
			},
			brickInstanceVariables: map[string]string{},
			expected: []BrickInstanceVariable{
				{Name: "VAR1", Value: "", Description: "desc", Required: true},
			},
		},
		{
			name: "variable with default value",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", DefaultValue: "default", Description: "desc"},
				},
			},
			brickInstanceVariables: map[string]string{},
			expected: []BrickInstanceVariable{
				{Name: "VAR1", Value: "", Description: "desc", Required: false},
			},
		},
		{
			name: "multiple variables",
			brick: &bricksindex.Brick{
				Variables: []bricksindex.BrickVariable{
					{Name: "VAR1", Description: "desc1"},
					{Name: "VAR2", DefaultValue: "def2", Description: "desc2"},
				},
			},
			brickInstanceVariables: map[string]string{"VAR1": "v1"},
			expected: []BrickInstanceVariable{
				{Name: "VAR1", Value: "v1", Description: "desc1", Required: true},
				{Name: "VAR2", Value: "", Description: "desc2", Required: false},
			},
		},
		{
			name:                   "no variables",
			brick:                  &bricksindex.Brick{Variables: []bricksindex.BrickVariable{}},
			brickInstanceVariables: map[string]string{},
			expected:               []BrickInstanceVariable{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getBrickInstanceVariableDetails(tt.brick, tt.brickInstanceVariables)
			require.Equal(t, tt.expected, got)
		})
	}
}
