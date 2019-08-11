package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/parking_lot/parkinglot"
	parkinglotservice "github.com/parking_lot/parkinglot/service"
	parkinglotstore "github.com/parking_lot/parkinglot/store"
)

// available commands known
const (
	CommandExit                         = "exit"
	CommandSlotNumberRegistrationNumber = "slot_number_for_registration_number"
	CommandSlotNumbersColour            = "slot_numbers_for_cars_with_colour"
	CommandRegistrationNumbersColour    = "registration_numbers_for_cars_with_colour"
	CommandPark                         = "park"
	CommandStatus                       = "status"
	CommandLeave                        = "leave"
	CommandCreateParkingLot             = "create_parking_lot"
)

func init() {
	plStore := parkinglotstore.New()
	plService := parkinglotservice.New(plStore)

	parkinglot.Init(plService)
}

func main() {
	var (
		command            string
		sum                int
		slotID             int
		registrationNumber string
		color              string

		isFile   bool
		counter  int
		commands []string
		line     []string
	)

	arg := os.Args

	if len(arg) == 2 {
		isFile = true
		basePath := fmt.Sprintf("%s/src/github.com/parking_lot/", os.Getenv("GOPATH"))
		pathToFile := basePath + arg[1]

		data, err := ioutil.ReadFile(pathToFile)
		if err != nil {
			panic("file not found")
		}

		body := string(data)
		commands = strings.Split(body, "\n")
	}

	p := parkinglot.GetService()

loop:
	for {
		if isFile {
			line = strings.Split(commands[counter], " ")
			command = line[0]
		} else {
			fmt.Scanf("%s", &command)
		}

		switch command {
		case CommandExit:
			break loop
		case CommandCreateParkingLot:
			if isFile {
				sum, _ = strconv.Atoi(line[1])
			} else {
				fmt.Scanf("%d", &sum)
			}

			doCreateParkingLot(p, sum)
		case CommandPark:
			if isFile {
				registrationNumber = line[1]
				color = line[2]
			} else {
				fmt.Scanf("%s", &registrationNumber)
				fmt.Scanf("%s", &color)
			}

			doPark(p, registrationNumber, color)
		case CommandLeave:
			if isFile {
				slotID, _ = strconv.Atoi(line[1])
			} else {
				fmt.Scanf("%d", &slotID)
			}

			doLeave(p, slotID)
		case CommandRegistrationNumbersColour:
			if isFile {
				color = line[1]
			} else {
				fmt.Scanf("%s", &color)
			}

			doRegistrationNumbersColour(p, color)
		case CommandSlotNumbersColour:
			if isFile {
				color = line[1]
			} else {
				fmt.Scanf("%s", &color)
			}

			doSlotNumbersColour(p, color)
		case CommandSlotNumberRegistrationNumber:
			if isFile {
				registrationNumber = line[1]
			} else {
				fmt.Scanf("%s", &registrationNumber)
			}

			doSlotNumberRegistrationNumber(p, registrationNumber)
		case CommandStatus:
			doStatus(p)
		default:
			fmt.Println("command not found")
		}

		counter++
		if isFile {
			if counter == len(commands) {
				break
			}
		}

	}
}

func doCreateParkingLot(p parkinglot.Service, sum int) {
	err := p.CreateParkingLot(sum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created a parking lot with %d slots\n", sum)
	}
}

func doPark(p parkinglot.Service, registrationNumber, color string) {
	slot, err := p.Park(parkinglot.Car{
		RegistrationNumber: registrationNumber,
		Color:              color,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Allocated slot number: %d\n", slot)
	}
}

func doLeave(p parkinglot.Service, slotID int) {
	err := p.Leave(slotID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Slot number %d is free\n", slotID)
	}
}

func doRegistrationNumbersColour(p parkinglot.Service, color string) {
	result, err := p.GetRegistrationNumbersByColor(color)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(strings.Join(result, ", "))
	}
}

func doSlotNumbersColour(p parkinglot.Service, color string) {
	result, err := p.GetSlotNumbersByColor(color)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(result), " ", ", ", -1), "[]"))
	}
}

func doStatus(p parkinglot.Service) {
	result, err := p.GetStatus()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Slot No.\tRegistration No\t\tColour")
		for _, r := range result {
			if !reflect.DeepEqual(r.Car, parkinglot.Car{}) {
				fmt.Printf("%d\t\t%s\t\t%s\n", r.Slot, r.Car.RegistrationNumber, r.Car.Color)
			}
		}
	}
}

func doSlotNumberRegistrationNumber(p parkinglot.Service, registrationNumber string) {
	result, err := p.GetSlotNumberByRegistrationNumber(registrationNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
