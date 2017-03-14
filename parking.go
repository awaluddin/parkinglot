package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Parking ")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		keys := strings.Split(text, " ")

		switch keys[0] {
		case "create":
			fmt.Println("parking car")
			break
		case "park":
			fmt.Println("parking car")
			break
		case "status":
			fmt.Println("list cars")
			break
		default:
			fmt.Println("keyword not found")
		}

	}

}
