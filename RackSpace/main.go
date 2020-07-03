package main

import (
	"fmt"
	"os"
)

type product struct {
	name  string
	price float32
}

var data = make(map[string]product)
var total, CH1, AP1, CF1, MK1, OM1 float32

func main() {
	data["CH1"] = product{"Chai", 3.11}
	data["AP1"] = product{"Apples", 6.00}
	data["CF1"] = product{"Coffee", 11.23}
	data["MK1"] = product{"Milk", 4.75}
	data["OM1"] = product{"Oatmeal", 3.69}
	fmt.Println("Welcome To Super Market")
	fmt.Println("Human! Buy Following Items One By One!")
	fmt.Println("CH1-->", data["CH1"])
	fmt.Println("AP1-->", data["AP1"])
	fmt.Println("CF1-->", data["CF1"])
	fmt.Println("MK1-->", data["MK1"])
	fmt.Println("OM1-->", data["OM1"])
	fmt.Println("Exit. Exit")
	items()
}
func BOGO() {
	if CF1 >= 2 {
		quantity := int(CF1 / 2)
		total -= float32(quantity) * data["CF1"].price
	}
}
func APOM() {
	if AP1 >= 1 && OM1 >= AP1 {
		total -= (AP1 / 2) * data["AP1"].price
	} else if AP1 >= 1 && OM1 < AP1 {
		total -= (OM1 / 2) * data["AP1"].price
	}
}
func APPL() {
	if AP1 >= 3 {
		total -= 1.5 * AP1
	}
}
func CHMK() {
	if MK1 >= 1 && CH1 >= 1 {
		total -= MK1 * data["MK1"].price
	}
}
func exit() {
	BOGO()
	APOM()
	APPL()
	CHMK()
	fmt.Println("Discounted Price:", total)
	os.Exit(1)
}
func items() {
choice:
	var a string
	fmt.Scanf("%v", &a)
	switch a {
	case "CH1":
		CH1++
		total += data["CH1"].price
		goto choice
	case "AP1":
		AP1++
		total += data["AP1"].price
		goto choice
	case "CF1":
		CF1++
		total += data["CF1"].price
		goto choice
	case "MK1":
		MK1++
		total += data["MK1"].price
		goto choice
	case "OM1":
		OM1++
		total += data["OM1"].price
		goto choice
	case "Exit":
		fmt.Println("Actual Price:", total)
		exit()
	}
}
