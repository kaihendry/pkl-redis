package main

import "testing"

func TestRedisBinds(t *testing.T) {
	// input is a configuration file
	// output are redis cli commands
	// assert against expected output

	testCases := []struct {
		pklConfig             string
		expectedRedisCommands []string
	}{
		{
			pklConfig: "config.pkl",
			expectedRedisCommands: []string{
				"docker exec redis redis-cli -n 6 SET RL:toggle SHADOW",
				"docker exec redis redis-cli -n 6 SET CB:toggle OFF",
			},
		},
	}
	for _, tc := range testCases {
		redisCommands, err := genRedisBinds(tc.pklConfig)
		if err != nil {
			t.Errorf("genRedisBinds(%s) failed: %v", tc.pklConfig, err)
		}
		if len(redisCommands) != len(tc.expectedRedisCommands) {
			t.Errorf("genRedisBinds(%s) = %v, want %v", tc.pklConfig, redisCommands, tc.expectedRedisCommands)
		}
		for i, redisCommand := range redisCommands {
			if redisCommand != tc.expectedRedisCommands[i] {
				t.Errorf("ORDER genRedisBinds(%s) = %v, want %v", tc.pklConfig, redisCommands, tc.expectedRedisCommands)
			}
		}
	}
}
