package main

import (
	"fmt"
	"log"
	"math/rand"
)


func getBet(balance uint16) uint16 {
	var bet uint16 
	fmt.Println("Enter the bet balance: ")
	fmt.Scan(&bet)

	if bet > balance {
		log.Fatalf("bet can not larger than the balance.")
	} else {
		break
	}
	return  bet
}

func generateSymbolArray(symbols map[string]uint16) []string {
	symbolArr := [] string {}
	for key, value := range symbols {
		for i := 0; i < int(value) ; i++ {
			symbolArr = append(symbolArr, key)
		}
	}
	return symbolArr
}

func getRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max - min + 1) + min
	return randomNumber
}

func getSpin(reel [] string, rows, cols int) [][]string {
	result := [][] string {}
	for i := 0; i < rows; i++ {
		result = append(result, []string {})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool {}
		for row:=0; row < rows; row++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel) - 1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row]= append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func printSpin(spin [][] string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row) - 1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

func setTheBalance() uint16{
	// set the balance
	balance := uint16(200)
	return  balance
}

func checkWin(spin [][]string, multipliers map[string] uint16) []uint16 {
	lines := []uint16 {}
	for _, row := range spin {
		win := true 
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return  lines
}

func main(){

	symbols := map[string] uint16 {
		"A" : 4,
		"B" : 7,
		"C" : 12,
		"D" : 20,
	} 

	multipliers := make(map[string] uint16, 4)
	multipliers["A"] = 20
	multipliers["B"] = 10
	multipliers["C"] = 5
	multipliers["D"] = 3

	// generate symbol array
	symbolsArr := generateSymbolArray(symbols)



	// set the balance
	balance := setTheBalance()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		
		// generate a spin
		spin := getSpin(symbolsArr, 3, 3)
		printSpin(spin)

		// check if win and then update the balance
		winningLines := checkWin(spin, multipliers)
		// fmt.Println(winningLines)

		for i, multi := range winningLines {
			win := multi * bet 
			balance += win
			if multi > 0 {
				fmt.Printf("Won $%d, (%dx) on Line #%d\n", win, multi, i+1)
			}
		}
	}

	fmt.Printf("You left with $%d\n", balance)

}