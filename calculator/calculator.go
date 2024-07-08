// pacote principal
package main

// importando pacote fmt
import "fmt"

// função prinipal
func main() {

	// declarando variaveis float e string
	var number1, number2 float64
	var operator string

	//ouput
	fmt.Println("Enter your first number:")
	//input
	fmt.Scan(&number1)
	fmt.Println("Enter your second number:")
	fmt.Scan(&number2)

	fmt.Print("Enter the operator:")
	fmt.Scan(&operator)

	//processamento
	switch operator {
	case "+":
		fmt.Print("O resultado é: ", number1+number2)
		break
	case "-":
		fmt.Print("O resultado é: ", number1-number2)
		break
	case "*":
		fmt.Print("O resultado é: ", number1*number2)
		break
	case "/":
		if number2 != 0 {
			fmt.Print("O resultado é: ", number1/number2)
		} else {
			fmt.Print("Impossível divisão por Zero!!!!!")
		}
	default:
		fmt.Print("Operação inválida!")
	}
}
