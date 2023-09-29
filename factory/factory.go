package main

import (
	"fmt"
	"strconv"
)

type ILaptop interface {
	getPrice() int
	getName() string
}

type Laptop struct{
	price int
	name string
}

func (l *Laptop) getName() string {
	return l.name
}

func (l *Laptop) getPrice() int{
	return l.price
}

type AsusLaptop struct{
	Laptop
}

type MacLaptop struct{
	Laptop
}

type HPLaptop struct{
	Laptop
}

func createAsusLaptop() ILaptop {
	return &AsusLaptop{Laptop: Laptop{499990, "Asus TUF 15"}}
}

func createMacLaptop() ILaptop {
	return &MacLaptop{Laptop: Laptop{1499990, "Macbook Air M2"}}
}

func createHPLaptop() ILaptop {
	return &HPLaptop{Laptop: Laptop{399999, "Victus"}}
}

func LaptopFactory(laptopName string) (ILaptop, error) {
	if laptopName == "AsusLaptop" {
		return createAsusLaptop(), nil
	}
	if laptopName == "MacLaptop" {
		return createMacLaptop(), nil
	}
	if laptopName == "HPLaptop" {
		return createHPLaptop(), nil
	}
	return nil, nil
}

func main() {
	AsusLaptop, err := LaptopFactory("AsusLaptop")
	MacLaptop, err2 := LaptopFactory("MacLaptop")
	HPLaptop, err3 := LaptopFactory("HPLaptop")

	if err == nil{
		fmt.Println("Laptop is: " + AsusLaptop.getName() + " and costs " + strconv.Itoa(AsusLaptop.getPrice()) + "\n")
	}
	if err2 == nil{
		fmt.Println("Laptop is: " + MacLaptop.getName() + " and costs " + strconv.Itoa(MacLaptop.getPrice()) + "\n")
	}
	if err3 == nil{
		fmt.Println("Laptop is: " + HPLaptop.getName() + " and costs " + strconv.Itoa(HPLaptop.getPrice()) + "\n")
	}
	
}