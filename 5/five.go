package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func main() {
	dat, _ := os.ReadFile("./input")
	instructions := strings.Split(string(dat), "\n")
	stacks := util.Map([]string{
		"ZJG",
		"QLRPWFVC",
		"FPMCLGR",
		"LFBWPHM",
		"GCFSVQ",
		"WHJZMQTL",
		"HFSBV",
		"FJZS",
		"MCDPFHBT",
	}, func(i string) []byte { return []byte(i) })

	for _, serializedInstruction := range instructions {
		tokens := strings.Split(serializedInstruction, " ")
		count := util.ToInt(tokens[1])
		from := util.ToInt(tokens[3]) - 1
		to := util.ToInt(tokens[5]) - 1

		curr := stacks[from][len(stacks[from])-count:]
		stacks[from] = stacks[from][:len(stacks[from])-count]
		stacks[to] = append(stacks[to], curr...)
	}

	result := string(util.Map(stacks, func(s []byte) byte { return s[len(s)-1] }))

	fmt.Println(result)

	if result != "GSLCMFBRP" {
		panic("Unexpected result")
	}
}
