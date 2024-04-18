package main

import (
	"fmt"
	"os"
)

func main() {

	type Animal struct {
		name, typeAnimal, commands, birthDate string

		// Не понимаю, как вводить с консоли дату и как пользовать t.Time()
	}
	var allAnimals = map[string]Animal{
		"Whiskars": {name: "Whiskars", typeAnimal: "Cat", commands: "Sit, Pounce", birthDate: "2019.05.15"},
		"Fido":     {name: "Fido", typeAnimal: "Dog", commands: "Sit, Stay, Fetch", birthDate: "2020.01.01"},
		"Hammy":    {name: "Hammy", typeAnimal: "Hamster", commands: "Roll, Hide", birthDate: "2021.03.10"},
		"Thunder":  {name: "Thunder", typeAnimal: "Horse", commands: "Trot, Canter, Gallop", birthDate: "2015.07.21"},
	}
	//Добавление нового животного

	//Вывод списка команд животного

	//Добавление новых комманд

	//Сортировка по возрасту

	fmt.Println("Выберите номер действия:\n", "1. Добавление нового животного\n",
		"2. Список команд животного\n", "3. Добавление новых команд\n",
		"4. Сортировка животных по возрасту\n", "5. Выход")

	var numMenu uint8

	fmt.Scan(&numMenu)
	if numMenu == 1 {
		var newName, addType, addcomands, addBirthDate string
		fmt.Println("Введите имя животного:")
		fmt.Scan(&newName)
		fmt.Println("Введите вид животного:")
		fmt.Scan(&addType)
		fmt.Println("Введите выполняемые команды:")
		fmt.Scan(&addcomands)
		fmt.Println("Введите дату рождения:")
		fmt.Scan(&addBirthDate)
	} else if numMenu == 2 {
		fmt.Println("Команды какого животного хотите посмотреть?")
	} else if numMenu == 3 {
		fmt.Println("Какому животному добавить команды?")
	} else if numMenu == 4 {
		fmt.Println("Животные по дате рождения:")
	} else if numMenu == 5 {
		fmt.Println("До встречи, пирожочек!")
		os.Exit(0)
	}
	fmt.Println(allAnimals)
}
