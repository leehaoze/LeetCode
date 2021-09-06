package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	version1Nums := strings.Split(version1, ".")
	version2Nums := strings.Split(version2, ".")
	swap := false

	if len(version2Nums) > len(version1Nums) {
		version1Nums, version2Nums = version2Nums, version1Nums
		swap = true
	}

	lengthOfShortVersion := len(version2Nums)
	// version1 是长的
	for i := 0; i < len(version1Nums); i++ {
		versionB := int64(0)
		if i < lengthOfShortVersion {
			versionB, _ = strconv.ParseInt(version2Nums[i], 10, 64)
		}

		versionA, _ := strconv.ParseInt(version1Nums[i], 10, 64)

		if versionA > versionB {
			if swap {
				return -1
			} else {
				return 1
			}
		} else if versionA < versionB {
			if swap {
				return 1
			} else {
				return -1
			}
		}
	}

	return 0
}


func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.0.1", "1"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
}
