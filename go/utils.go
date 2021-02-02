package main

import (
	"fmt"
	"strconv"
	"strings"
)

func costToFloat32(sum string) (float32, error) {
	var floatSum float64
	var rubs string
	var kops string
	var err error
	split := strings.Split(sum, " ")
	rubs = split[0][0 : len(split[0])-1]
	if len(split) > 1 {
		kops = split[1][0 : len(split[1])-1]
	} else {
		kops = "0"
	}

	floatSum, err = strconv.ParseFloat(rubs+"."+kops, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(floatSum), nil
}

func float32ToCost(sum float32) string {
	var rubs string
	var kops string

	sumStr := fmt.Sprintf("%.2f", sum)
	split := strings.Split(sumStr, ".")
	rubs = split[0] + "r"
	kops = split[1] + "k"
	return rubs + " " + kops
}
