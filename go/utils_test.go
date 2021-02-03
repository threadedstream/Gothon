package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCostToFloat32(t *testing.T) {
	sample0 := "34r 21k"
	sample1 := "213r 45k"
	sample2 := "145r 79k"

	output0, err0 := costToFloat32(sample0)
	output1, err1 := costToFloat32(sample1)
	output2, err2 := costToFloat32(sample2)

	assert.Equal(t, err0, nil)
	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
	assert.Equal(t, output0, float32(34.21))
	assert.Equal(t, output1, float32(213.45))
	assert.Equal(t, output2, float32(145.79))
}

func TestFloat32ToCost(t *testing.T) {
	sample0 := float32(34.21)
	sample1 := float32(3.14159265359)
	sample2 := float32(52.12)
	sample3 := float32(4560)

	output0 := float32ToCost(sample0)
	output1 := float32ToCost(sample1)
	output2 := float32ToCost(sample2)
	output3 := float32ToCost(sample3)

	assert.Equal(t, output0, "34r 21k")
	assert.Equal(t, output1, "3r 14k")
	assert.Equal(t, output2, "52r 12k")
	assert.Equal(t, output3, "4560r 00k")
}
