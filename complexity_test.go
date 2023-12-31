package gometric

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAMC(t *testing.T) {
	t.Run("valid: input 1", func(t *testing.T) {
		got, err := CalcAMC(2, 2)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 2", func(t *testing.T) {
		got, err := CalcAMC(16, 5)

		assert.Equal(t, 3.2, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 3", func(t *testing.T) {
		got, err := CalcAMC(12, 3)

		assert.Equal(t, 4.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 4", func(t *testing.T) {
		got, err := CalcAMC(5, 3)

		assert.Equal(t, 1.6666666666666667, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 5", func(t *testing.T) {
		got, err := CalcAMC(17, 3)

		assert.Equal(t, 5.666666666666667, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero method complexity", func(t *testing.T) {
		got, err := CalcAMC(0, 20)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: both values are 0", func(t *testing.T) {
		got, err := CalcAMC(0, 0)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: both values are math.MaxInt32", func(t *testing.T) {
		got, err := CalcAMC(math.MaxInt32, math.MaxInt32)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("invalid: zero total number of methods", func(t *testing.T) {
		got, err := CalcAMC(1, 0)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative method complexity", func(t *testing.T) {
		got, err := CalcAMC(-1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative total number of methods", func(t *testing.T) {
		got, err := CalcAMC(1, -1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}

func TestDMS(t *testing.T) {
	t.Run("valid: values at lower bound", func(t *testing.T) {
		got, err := CalcDMS(0, 0)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: values at upper bound", func(t *testing.T) {
		got, err := CalcDMS(1, 1)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: equal values", func(t *testing.T) {
		got, err := CalcDMS(0.5, 0.5)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: complimentary values", func(t *testing.T) {
		got, err := CalcDMS(0.25, 0.75)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: arbitrary values", func(t *testing.T) {
		got, err := CalcDMS(0.1, 0.2)

		assert.Equal(t, 0.7, got)
		require.NoError(t, err)
	})

	t.Run("valid: near-equal values", func(t *testing.T) {
		got, err := CalcDMS(0.50000000000001, 0.49999999999999)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("invalid: negative abstractness", func(t *testing.T) {
		got, err := CalcDMS(-1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative instability", func(t *testing.T) {
		got, err := CalcDMS(1, -1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: abstractness is NaN", func(t *testing.T) {
		got, err := CalcDMS(math.NaN(), 0)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: instability is NaN", func(t *testing.T) {
		got, err := CalcDMS(0, math.NaN())

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}
