package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CarSlot struct {
	Slot  string
	CarId string
	Color string
}

const (
	CAR_COLOR = "color"
	CAR_ID    = "id"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("       Parking ")
	fmt.Println("---------------------")

	CarList := []CarSlot{}

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

					AddList := CarSlot{
						Slot:  id,
						CarId: "",
						Color: "",
					}
					CarList = append(CarList, AddList)
				}
			}
			break
		case "park":
			CarNumber := keys[1]
			CarColor := keys[2]

			FreeSlot := 0
			for k, v := range CarList {
				if v.CarId == "" {
					FreeSlot = k
					break
				}
			}
			CarList[FreeSlot].CarId = CarNumber
			CarList[FreeSlot].Color = CarColor

			fmt.Println("parking car")
			break
		case "status":
			fmt.Println("No 				Id							Color")
			for _, v := range CarList {
				fmt.Println(v.Slot + " 			" + v.CarId + "							" + v.Color)
			}
			break
		case "leave":
			Id := keys[1]
			for k, v := range CarList {
				if v.Slot == Id {
					CarList[k].CarId = ""
					CarList[k].Color = ""
					break
				}
			}
			break
		case "find":
			FindKey := keys[1]
			FindValue := keys[2]

			fmt.Println("No 				Id							Color")

			if FindKey == CAR_ID {
				for _, v := range CarList {
					if v.CarId == FindValue {
						fmt.Println(v.Slot + " 			" + v.CarId + "							" + v.Color)
					}
				}
			} else {
				for _, v := range CarList {
					if v.Color == FindValue {
						fmt.Println(v.Slot + " 			" + v.CarId + "							" + v.Color)
					}
				}
			}

			break
		default:
			fmt.Println("keyword not found")
		}

	}

}
