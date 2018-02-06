package warmup

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
"math"
)

// Link to problem - https://www.hackerrank.com/challenges/mini-max-sum/problem
func MiniMaxSum() {
	inputreader := bufio.NewReader(os.Stdin)
	lineInput, _ := inputreader.ReadString('\n')
	split_arr := strings.Split(strings.TrimSpace(lineInput), " ")
	clean_arr := make([]int64, len(split_arr))
	var total int64 = 0
	var min int64 = math.MaxInt64
	var max int64 = -1

	// convert input string to array of int, find max, min and total
	for index, val := range split_arr {
		clean_arr[index], _ = strconv.ParseInt(val, 10, 64)
		if (clean_arr[index] > max) {
			max = clean_arr[index]
		}
		if (clean_arr[index] < min) {
			min = clean_arr[index]
		}
		total += clean_arr[index]
	}
	fmt.Printf("%d %d", (total-max), (total-min))
}