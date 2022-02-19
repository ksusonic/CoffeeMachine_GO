package main

import "fmt"

func Status(Water, Milk, Beans, Cups, Money *int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", *Water)
	fmt.Printf("%d ml of milk\n", *Milk)
	fmt.Printf("%d g of coffee beans\n", *Beans)
	fmt.Printf("%d disposable cups\n", *Cups)
	fmt.Printf("$%d of money\n\n", *Money)
}

func Fill(Water, Milk, Beans, Cups *int) {
	var water, milk, beans, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	*Water += water
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	*Milk += milk
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&beans)
	*Beans += beans
	fmt.Println("Write how many cups of coffee you want to add:")
	fmt.Scan(&cups)
	*Cups += cups
}

func Take(Money *int) {
	fmt.Printf("I gave you $%d\n", *Money)
	*Money = 0
}

func Buy(Water, Milk, Beans, Cups, Money *int) {
	var choice int
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var waterNeeded, milkNeeded, beansNeeded, moneyIncome = 0, 0, 0, 0
	fmt.Scan(&choice)
	if choice == 1 {
		waterNeeded = 250
		beansNeeded = 16
		moneyIncome = 4
	} else if choice == 2 {
		waterNeeded = 350
		milkNeeded = 75
		beansNeeded = 20
		moneyIncome = 7
	} else if choice == 3 {
		waterNeeded = 200
		milkNeeded = 100
		beansNeeded = 12
		moneyIncome = 6
	} else {
		return
	}
	if waterNeeded > *Water {
		fmt.Println("Sorry, not enough water!")
	} else if beansNeeded > *Beans {
		fmt.Println("Sorry, not enough beans!")
	} else if milkNeeded > *Milk {
		fmt.Println("Sorry, not enough milk!")
	} else if *Cups < 1 {
		fmt.Println("Sorry, not enough cups!")
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
		*Water -= waterNeeded
		*Milk -= milkNeeded
		*Beans -= beansNeeded
		*Money += moneyIncome
		*Cups -= 1
	}
	fmt.Println()
}

func main() {
	var water, milk, beans, cups, money = 400, 540, 120, 9, 550
	var action string
	for true {
		fmt.Println("Write action (buy, fill, take, remaining, exit): ")
		fmt.Scan(&action)
		fmt.Println()
		if action == "buy" {
			Buy(&water, &milk, &beans, &cups, &money)
		} else if action == "fill" {
			Fill(&water, &milk, &beans, &cups)
		} else if action == "take" {
			Take(&money)
		} else if action == "remaining" {
			Status(&water, &milk, &beans, &cups, &money)
		} else if action == "exit" {
			break
		}
	}
}
