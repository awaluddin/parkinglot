package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Parking ")
	fmt.Println("---------------------")

	CarList := [][]string{}

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		keys := strings.Split(text, " ")

		switch keys[0] {
		case "create":
			if keys[1] != "" {
				slot, _ := strconv.Atoi(keys[1])
				for i := 1; i <= slot; i++ {
					id := strconv.Itoa(i)
					AddList := []string{id, "", ""}
					CarList = append(CarList, AddList)
				}
			}
			fmt.Println("parking car")
			break
		case "park":
			CarNumber := keys[1]
			CarColor := keys[2]

			FreeSlot := 0
			for k, v := range CarList {
				if v[1] == "" {
					FreeSlot = k
					break
				}
			}
			CarList[FreeSlot][1] = CarNumber
			CarList[FreeSlot][2] = CarColor

			fmt.Println("parking car")
			break
		case "status":
			fmt.Println("No 				Id							Color")
			for _, v := range CarList {
				fmt.Println(v[0] + " 			" + v[1] + "							" + v[2])
			}
			break
		case "leave":
			Id := keys[1]
			for k, v := range CarList {
				if v[0] == Id {
					CarList[k][1] = ""
					CarList[k][2] = ""
					break
				}
			}
			break
		case "find":
			FindKey := keys[1]
			FindValue := keys[2]
			FindKeyIndex := 0

			if FindKey == "id" {
				FindKeyIndex = 1
			} else {
				FindKeyIndex = 2
			}

			fmt.Println("No 				Id							Color")
			for _, v := range CarList {
				if v[FindKeyIndex] == FindValue {
					fmt.Println(v[0] + " 			" + v[1] + "							" + v[2])
				}

			}
			break
		default:
			fmt.Println("keyword not found")
		}

	}

}
