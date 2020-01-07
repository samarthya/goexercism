package robotname

import (
	"math/rand"
	"strings"
	"fmt"
)

const ALPHABETS string = "ABCDEFGHIJKLMNOPQRSTUVQXYZ"

type Robot struct {
	name string
}

var GeneratedNames [] string = make([]string, 0)

func GetRandomAlphabet() string{
	words := [] rune(ALPHABETS)
	return strings.ToUpper(string(words[rand.Intn(len(words))]))
}

func GetRandomInt() int {
	return rand.Intn(10);
}

func namePresent(newName string) bool {
	for _, val := range GeneratedNames {
		if strings.Compare(val, newName) == 0 {
			return true
		}
	}
	return false
}

func GenerateRandomName() string {
	// The first time you boot them up, a random name is generated in the format of 
	// two uppercase letters followed by three digits, such as RX837 or BC811.	
	newName := ""
	
	for {
		newName = fmt.Sprintf("%s%s%d%d%d", GetRandomAlphabet(), GetRandomAlphabet(),GetRandomInt(),GetRandomInt(),GetRandomInt())

		if namePresent(newName) == false {
			GeneratedNames = append(GeneratedNames, newName)
			return newName
		}
	}
}

func (robot *Robot) Name() (string, error) {
	if robot.name == "" {
		// Generate a name
		robot.name = GenerateRandomName()
	}
	return robot.name, nil
}

func (robot *Robot) Reset() string {
	robot.name = GenerateRandomName()
	return robot.name
}