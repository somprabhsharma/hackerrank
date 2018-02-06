package warmup

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

// Link to problem - https://www.hackerrank.com/challenges/birthday-cake-candles/problem
func BirthdayCakeCandles() {
	inputreader := bufio.NewReader(os.Stdin)
	lineInput, _ := inputreader.ReadString('\n')
	//noOfCandles, _ := strconv.Atoi(strings.TrimSpace(lineInput))
	lineInput, _ = inputreader.ReadString('\n')
	split_arr := strings.Split(strings.TrimSpace(lineInput), " ")
	candles := make([]int, len(split_arr))
	max := 0

	// find maximum height of candle
	for index, val := range split_arr {
		candles[index], _ = strconv.Atoi(val)
		if (candles[index] > max) {
			max = candles[index]
		}
	}

	// find the number of candles that can be blown
	noOfCandlesToBeBlown := 0
	for _, val := range candles {
		if (val == max) {
			noOfCandlesToBeBlown++
		}
	}

	fmt.Printf("%d", noOfCandlesToBeBlown)
}
