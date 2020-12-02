package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem2A struct {

}

func (this *Problem2A) Solve() {
	Log.Info("Problem 2A solver beginning!")


	file, err := os.Open("source-data/input-day-02a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	validCount := 0;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			parts := strings.Split(line, ":");
			policyRaw := parts[0];
			policy := &SimplePasswordPolicy{};
			policy.Parse(policyRaw);
			password := strings.TrimSpace(parts[1]);
			if(policy.Validate(password)){
				Log.Info("Password %s was valid for policy %s", password, policy.Describe())
				validCount++;
			} else{
				Log.Info("Password %s was INVALID for policy %s", password, policy.Describe())
			}
		}
	}
	Log.Info("Finished parsing password file - found %d valid passwords", validCount);


}

type SimplePasswordPolicy struct {
	MinLetters int;
	MaxLetters int;
	LetterInQuestion uint8;
}

func (this *SimplePasswordPolicy) Parse(line string) bool {
	parts := strings.Split(line, " ");
	rangeParts := strings.Split(parts[0], "-");
	val, err := strconv.ParseInt(rangeParts[0], 10, 64);
	if(err != nil){
		return false;
	}
	this.MinLetters = int(val);
	val, err = strconv.ParseInt(rangeParts[1], 10, 64);
	if(err != nil){
		return false;
	}
	this.MaxLetters = int(val);
	this.LetterInQuestion = parts[1][0];
	return true;
}

func (this *SimplePasswordPolicy) Describe() string {
	return fmt.Sprintf("%d-%d %c", this.MinLetters, this.MaxLetters, this.LetterInQuestion);
}

func (this *SimplePasswordPolicy) Validate(password string) bool {
	count := 0;
	len := len(password);
	for i := 0; i < len; i++{
		if(password[i] == this.LetterInQuestion){
			count++;
		}
	}
	return count >= this.MinLetters && count <= this.MaxLetters;
}