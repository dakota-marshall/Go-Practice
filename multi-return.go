package main

import (
	"errors"
	"fmt"
)

func two_returns(yes_no bool) (bool, error) {

	if yes_no {
		return true, nil
	} else {
		return false, errors.New("No no no")
	}

}

func main() {
	yesno, _ := two_returns(true)

	if yesno {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
