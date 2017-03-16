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
	CAR_COLOR    = "color"
	CAR_ID       = "id"
	HEADER_TABLE = "header"
	ERR_PARAMS   = "err_param"
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

			if validation(1, false, keys, CarList) {
				StringSlot := keys[1]
				slot, _ := strconv.Atoi(StringSlot)
				for i := 1; i <= slot; i++ {
					id := strconv.Itoa(i)

					AddList := CarSlot{
						Slot:  id,
						CarId: "",
						Color: "",
					}
					CarList = append(CarList, AddList)
				}
				wording("", "Created parking lot with "+StringSlot+" slots")
			} else {
				wording(ERR_PARAMS, "")
			}

			break
		case "park":

			if validation(2, true, keys, CarList) {
				CarNumber := keys[1]
				CarColor := keys[2]

				var FreeSlot int
				var SlotNum string

				for k, v := range CarList {
					if v.CarId == "" {
						FreeSlot = k
						SlotNum = v.Slot
						break
					}
				}
				CarList[FreeSlot].CarId = CarNumber
				CarList[FreeSlot].Color = CarColor

				wording("", "Allocated slot number : "+SlotNum)
			} else {
				wording(ERR_PARAMS, "")
			}
			break
		case "status":
			wording(HEADER_TABLE, "")
			findData("", "", CarList)
			break
		case "leave":
			if validation(1, true, keys, CarList) {
				Id := keys[1]
				for k, v := range CarList {
					if v.Slot == Id {
						CarList[k].CarId = ""
						CarList[k].Color = ""
						break
					}
				}
			} else {
				wording(ERR_PARAMS, "")
			}
			break
		case "find":
			if validation(2, false, keys, CarList) {
				FindKey := keys[1]
				FindValue := keys[2]

				wording(HEADER_TABLE, "")

				if FindKey == CAR_ID {
					findData(FindValue, CAR_ID, CarList)

				} else {
					findData(FindValue, CAR_COLOR, CarList)
				}
			} else {
				wording(ERR_PARAMS, "")
			}
			break
		case "exit":
			return
		default:
			wording("", "keyword not found")
		}

	}

}

func wording(wtype, msg string) {

	switch wtype {
	case HEADER_TABLE:
		fmt.Println("No 					Id							Color")
		break
	case ERR_PARAMS:
		fmt.Println("Missing Parameters")
		break
	default:
		fmt.Println(msg)
	}
}

func validation(cond int, clist bool, keys []string, CarList []CarSlot) bool {

	if clist && len(CarList) == 0 {
		fmt.Println("Car Slot not created yet")
		return false
	}
	if cond == 1 {
		if len(keys) == 2 && keys[1] != "" {
			return true
		}
	} else if cond == 2 {
		if len(keys) == 3 && keys[1] != "" && keys[2] != "" {
			return true
		}
	}

	return false
}

func findData(FindValue, KeyType string, CarList []CarSlot) {

	if KeyType == CAR_ID {
		for _, v := range CarList {
			if v.CarId == FindValue {
				wording("", v.Slot+" 			"+v.CarId+"							"+v.Color)
			}
		}
	} else if KeyType == CAR_COLOR {
		for _, v := range CarList {
			if v.Color == FindValue {
				wording("", v.Slot+" 			"+v.CarId+"							"+v.Color)
			}
		}
	} else {
		for _, v := range CarList {
			wording("", v.Slot+" 			"+v.CarId+"							"+v.Color)
		}
	}

}
