package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem1A struct {

}

func (this *Problem1A) Solve() {
	Log.Info("Problem 1A solver beginning!")


	file, err := os.Open("source-data/input-day-01a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	vals := make([]int64, 0);

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			val, err := strconv.ParseInt(line, 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			vals = append(vals, val);
		}
	}

	targetVal := int64(2020);
	product := int64(0);
	len := len(vals);
	for i := 0; i < len; i++{
		for j := i + 1; j < len; j++{
			if(vals[i] + vals[j] == targetVal){
				product = vals[i] * vals[j];
			}
		}
	}

	Log.Info("Finished parsing file - magic pair is %d", product);
}
