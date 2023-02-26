package mynet

import "testing"

func TestCheckIPString(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"192.168.0.1", true},
		{"::1", true},
		{"localhost", false},
		{"not-an-ip-address", false},
	}

	for _, testCase := range testCases {
		if output := CheckIPString(testCase.input); output != testCase.expected {
			t.Errorf("Test failed: input=%v, expected=%v, output=%v", testCase.input, testCase.expected, output)
		}
	}
}