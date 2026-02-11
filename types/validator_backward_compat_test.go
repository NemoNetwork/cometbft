package types

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/stretchr/testify/assert"
)

// TestNewValidatorBackwardCompatibility tests that NewValidator works with both
// 2 parameters (for cosmos-sdk v0.50.11 compatibility) and 3 parameters.
func TestNewValidatorBackwardCompatibility(t *testing.T) {
	pubKey := ed25519.GenPrivKey().PubKey()

	// Test with 2 parameters (backward compatible)
	val1 := NewValidator(pubKey, 100)
	assert.NotNil(t, val1)
	assert.Equal(t, int64(100), val1.VotingPower)
	assert.False(t, val1.ProposeDisabled) // Should default to false

	// Test with 3 parameters (explicit proposeDisabled)
	val2 := NewValidator(pubKey, 200, false)
	assert.NotNil(t, val2)
	assert.Equal(t, int64(200), val2.VotingPower)
	assert.False(t, val2.ProposeDisabled)

	val3 := NewValidator(pubKey, 300, true)
	assert.NotNil(t, val3)
	assert.Equal(t, int64(300), val3.VotingPower)
	assert.True(t, val3.ProposeDisabled)
}
