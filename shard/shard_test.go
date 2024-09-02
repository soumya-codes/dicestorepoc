package shard

import (
	"testing"

	"github.com/soumya-codes/AlgoAndDS/generics/ops"
	"github.com/stretchr/testify/assert"
)

func TestShardEvaluateRequest(t *testing.T) {
	// Initialize the shard
	sh := NewShard(1)

	// Define test cases
	tests := []struct {
		description string
		operations  []*ops.Operation
		expected    string
	}{
		{
			description: "Test storing and retrieving a string",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"stringKey", "Hello, World!"}},
				{Cmd: "GET", Args: []string{"stringKey"}},
			},
			expected: "+OK\r\n$13\r\nHello, World!\r\n",
		},
		{
			description: "Test storing and retrieving an int8 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"int8Key", "127"}},
				{Cmd: "GET", Args: []string{"int8Key"}},
			},
			expected: "+OK\r\n$3\r\n127\r\n",
		},
		{
			description: "Test storing and retrieving a negative int8 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"negativeInt8Key", "-128"}},
				{Cmd: "GET", Args: []string{"negativeInt8Key"}},
			},
			expected: "+OK\r\n$4\r\n-128\r\n",
		},
		{
			description: "Test storing and retrieving an int16 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"int16Key", "32767"}},
				{Cmd: "GET", Args: []string{"int16Key"}},
			},
			expected: "+OK\r\n$5\r\n32767\r\n",
		},
		{
			description: "Test storing and retrieving a negative int16 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"negativeInt16Key", "-32768"}},
				{Cmd: "GET", Args: []string{"negativeInt16Key"}},
			},
			expected: "+OK\r\n$6\r\n-32768\r\n",
		},
		{
			description: "Test storing and retrieving an int32 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"int32Key", "2147483647"}},
				{Cmd: "GET", Args: []string{"int32Key"}},
			},
			expected: "+OK\r\n$10\r\n2147483647\r\n",
		},
		{
			description: "Test storing and retrieving an int64 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"int64Key", "9223372036854775807"}},
				{Cmd: "GET", Args: []string{"int64Key"}},
			},
			expected: "+OK\r\n$19\r\n9223372036854775807\r\n",
		},
		{
			description: "Test storing and retrieving a negative int32 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"negativeInt32Key", "-2147483648"}},
				{Cmd: "GET", Args: []string{"negativeInt32Key"}},
			},
			expected: "+OK\r\n$11\r\n-2147483648\r\n",
		},
		{
			description: "Test storing and retrieving a negative int64 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"negativeInt64Key", "-9223372036854775808"}},
				{Cmd: "GET", Args: []string{"negativeInt64Key"}},
			},
			expected: "+OK\r\n$20\r\n-9223372036854775808\r\n",
		},
		{
			description: "Test storing and retrieving a uint8 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"uint8Key", "255"}},
				{Cmd: "GET", Args: []string{"uint8Key"}},
			},
			expected: "+OK\r\n$3\r\n255\r\n",
		},
		{
			description: "Test storing and retrieving a uint16 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"uint16Key", "65535"}},
				{Cmd: "GET", Args: []string{"uint16Key"}},
			},
			expected: "+OK\r\n$5\r\n65535\r\n",
		},
		{
			description: "Test storing and retrieving a uint32 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"uint32Key", "4294967295"}},
				{Cmd: "GET", Args: []string{"uint32Key"}},
			},
			expected: "+OK\r\n$10\r\n4294967295\r\n",
		},
		{
			description: "Test storing and retrieving a uint64 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"uint64Key", "18446744073709551615"}},
				{Cmd: "GET", Args: []string{"uint64Key"}},
			},
			expected: "+OK\r\n$20\r\n18446744073709551615\r\n",
		},
		{
			description: "Test storing and retrieving a float32 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"float32Key", "3.4028235e+38"}},
				{Cmd: "GET", Args: []string{"float32Key"}},
			},
			expected: "+OK\r\n$39\r\n340282350000000000000000000000000000000\r\n",
		},
		{
			description: "Test storing and retrieving a float64 value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"float64Key", "1.7976931348623157e+308"}},
				{Cmd: "GET", Args: []string{"float64Key"}},
			},
			expected: "+OK\r\n$309\r\n179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\r\n",
		},
		{
			description: "Test storing and retrieving a byte array as a string",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"byteArrayKey", string([]byte{0xf0, 0x9f, 0x98, 0x8d})}},
				{Cmd: "GET", Args: []string{"byteArrayKey"}},
			},
			expected: "+OK\r\n$4\r\n😍\r\n",
		},
		{
			description: "Test storing and retrieving an empty string",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"emptyStringKey", ""}},
				{Cmd: "GET", Args: []string{"emptyStringKey"}},
			},
			expected: "+OK\r\n$0\r\n\r\n",
		},
		{
			description: "Test storing and retrieving a special character",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"specialCharKey", "\xf0"}},
				{Cmd: "GET", Args: []string{"specialCharKey"}},
			},
			expected: "+OK\r\n$1\r\n\xf0\r\n",
		},
		{
			description: "Test storing and retrieving a UTF-8 encoded string",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"utf8Key", "こんにちは世界"}},
				{Cmd: "GET", Args: []string{"utf8Key"}},
			},
			expected: "+OK\r\n$21\r\nこんにちは世界\r\n",
		},
		// TODO: Check this test case, should be treated as a string.
		/*		{
					description: "Test storing and retrieving a large integer beyond int64 range",
					operations: []*ops.Operation{
						{Cmd: "SET", Args: []string{"largeIntKey", "18446744073709551616"}},
						{Cmd: "GET", Args: []string{"largeIntKey"}},
					},
					expected: "+OK\r\n$20\r\n18446744000000000000\r\n",
				},
		*/{
			description: "Test storing and retrieving a negative float value",
			operations: []*ops.Operation{
				{Cmd: "SET", Args: []string{"negativeFloatKey", "-1.7976931348623157e+308"}},
				{Cmd: "GET", Args: []string{"negativeFloatKey"}},
			},
			expected: "+OK\r\n$310\r\n-179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\r\n",
		},
	}

	// Run the test cases
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Prepare the request
			req := ops.Request{
				OperationSet: tt.operations,
			}

			// Evaluate the request
			resp := sh.EvaluateRequest(&req)

			// Validate the response
			assert.Equal(t, tt.expected, string(resp))
		})
	}
}
