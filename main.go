package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/parking_lot/parkinglot"
	parkinglotservice "github.com/parking_lot/parkinglot/service"
	parkinglotstore "github.com/parking_lot/parkinglot/store"
)

// available commands
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
	arg := os.Args

	if len(arg) == 2 {
		cmdFromFile()
	} else {
		cmd()
	}
}

func cmd() {
	p := parkinglot.GetService()

	var (
		command            string
		sum                int
		slotID             int
		registrationNumber string
		color              string
	)

loop:
	for {
		fmt.Scanf("%s", &command)

		switch command {
		case CommandExit:
			break loop
		case CommandCreateParkingLot:
			fmt.Scanf("%d", &sum)
			err := p.CreateParkingLot(sum)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Created a parking lot with %d slots\n", sum)
			}
		case CommandPark:
			fmt.Scanf("%s", &registrationNumber)
			fmt.Scanf("%s", &color)

			slot, err := p.Park(parkinglot.Car{
				RegistrationNumber: registrationNumber,
				Color:              color,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Allocated slot number: %d\n", slot)
			}
		case CommandLeave:
			fmt.Scanf("%d", &slotID)

			err := p.Leave(slotID)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Slot number %d is free\n", slotID)
			}
		case CommandRegistrationNumbersColour:
			fmt.Scanf("%s", &color)

			result, err := p.GetRegistrationNumbersByColor(color)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(strings.Join(result, ", "))
			}
		case CommandSlotNumbersColour:
			fmt.Scanf("%s", &color)

			result, err := p.GetSlotNumbersByColor(color)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(result), " ", ", ", -1), "[]"))
			}
		case CommandSlotNumberRegistrationNumber:
			fmt.Scanf("%s", &registrationNumber)

			result, err := p.GetSlotNumberByRegistrationNumber(registrationNumber)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case CommandStatus:
			result, err := p.GetStatus()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Slot No.\tRegistration No\tColour")
				for _, r := range result {
					fmt.Printf("%d\t\t%s\t\t%s\n", r.Slot, r.Car.RegistrationNumber, r.Car.Color)
				}
			}
		default:
			fmt.Println("command not found")
		}

	}
}

func cmdFromFile() {
	p := parkinglot.GetService()
	fmt.Println(p)

}
