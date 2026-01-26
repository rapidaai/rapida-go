// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// MinUint64 returns the minimum of two uint64 numbers
func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
