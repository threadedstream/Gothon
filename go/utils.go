package main

import (
	"strconv"
	"strings"
)

func sumToFloat32(sum string) (float64, error) {
	var floatSum float64
	var rubs string
	var kops string
	var err error
	split := strings.Split(sum, " ")
	rubs = split[0][0 : len(split[0])-1]
	if len(split) > 1 {
		kops = split[1][0 : len(split[0])-1]
	} else {
		kops = "0"
	}

	floatSum, err = strconv.ParseFloat(rubs+"."+kops, 32)
	if err != nil {
		return 0.0, err
	}

	return floatSum, nil
}
