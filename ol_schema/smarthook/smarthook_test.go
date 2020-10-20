package smarthooksschema

import (
	"fmt"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/smarthooks"
	"github.com/stretchr/testify/assert"
)

func TestRulesSchema(t *testing.T) {
	t.Run("creates and returns a map of a Smarthooks Schema", func(t *testing.T) {
		provSchema := Schema()
		assert.NotNil(t, provSchema["type"])
		assert.NotNil(t, provSchema["status"])
		assert.NotNil(t, provSchema["disabled"])
		assert.NotNil(t, provSchema["risk_enabled"])
		assert.NotNil(t, provSchema["location_enabled"])
		assert.NotNil(t, provSchema["retries"])
		assert.NotNil(t, provSchema["timeout"])
		assert.NotNil(t, provSchema["env_vars"])
	})
}

func TestInflate(t *testing.T) {
	tests := map[string]struct {
		ResourceData   map[string]interface{}
		ExpectedOutput smarthooks.SmartHook
	}{
		"creates and returns the address of a SmartHook": {
			ResourceData: map[string]interface{}{
				"id":               "32f9dfee-a02c-4932-98ec-37838ce62ba0",
				"type":             "pre-authentication",
				"function":         "function myFunc(){...}",
				"retries":          0,
				"timeout":          2,
				"disabled":         false,
				"env_vars":         []interface{}{"API_KEY"},
				"risk_enabled":     false,
				"location_enabled": false,
			},
			ExpectedOutput: smarthooks.SmartHook{
				ID:              oltypes.String("32f9dfee-a02c-4932-98ec-37838ce62ba0"),
				Type:            oltypes.String("pre-authentication"),
				Function:        oltypes.String("function myFunc(){...}"),
				Retries:         oltypes.Int32(int32(0)),
				Timeout:         oltypes.Int32(int32(2)),
				Disabled:        oltypes.Bool(false),
				EnvVars:         []string{"API_KEY"},
				RiskEnabled:     oltypes.Bool(false),
				LocationEnabled: oltypes.Bool(false),
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			subj := Inflate(test.ResourceData)
			assert.Equal(t, test.ExpectedOutput, subj)
		})
	}
}

func TestValidTypes(t *testing.T) {
	tests := map[string]struct {
		InputKey       string
		InputValue     string
		ExpectedOutput []error
	}{
		"no errors on valid input": {
			InputKey:       "type",
			InputValue:     "pre-authentication",
			ExpectedOutput: nil,
		},
		"errors on invalid input": {
			InputKey:       "type",
			InputValue:     "asdf",
			ExpectedOutput: []error{fmt.Errorf("type must be one of [pre-authentication], got: asdf")},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, errs := validTypes(test.InputValue, test.InputKey)
			assert.Equal(t, test.ExpectedOutput, errs)
		})
	}
}
