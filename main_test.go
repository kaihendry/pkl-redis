package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRedisBinds(t *testing.T) {
	tests := []struct {
		name                string
		pklConfig           string
		expectedRedisBinds  []string
		expectedErrContains string
	}{
		{
			name:      "normal config with both toggles",
			pklConfig: "config.pkl",
			expectedRedisBinds: []string{
				"docker exec redis redis-cli -n 6 SET RL:toggle SHADOW",
				"docker exec redis redis-cli -n 6 SET CB:toggle OFF",
			},
		},
		{
			name:      "normal config with both toggles",
			pklConfig: "config2.pkl",
			expectedRedisBinds: []string{
				"docker exec redis redis-cli -n 6 SET RL:toggle ON",
				"docker exec redis redis-cli -n 6 SET CB:toggle ON",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genRedisBinds(tt.pklConfig)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tt.expectedRedisBinds, got); diff != "" {
				t.Errorf("redis binds mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
