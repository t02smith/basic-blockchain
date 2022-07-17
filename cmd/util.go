package cmd

import (
	"fmt"
)

func AddressesAndBalancesToString(pairs map[string]int) []string {
	var output []string = []string{}

	for addr, bal := range pairs {
		output = append(output, fmt.Sprintf("%s = %d", addr, bal))
	}

	return output
}
