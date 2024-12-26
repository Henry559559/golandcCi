package commands

import (
	"bufio"
	"fmt"
	"golandcli/expenses"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("Digita Tu Valor: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.Replace(str, "\r\n", "", 1)
	return str, nil
}

func ShowInConsole(expensesList []float32) {

	fmt.Println(contentString(expensesList))
}

func contentString(expensesList []float32) string {
	builder := strings.Builder{}

	fmt.Println("")
	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
		max, min, avg, sum := expensesDetails(expensesList)
		if i == len(expensesList)-1 {
			fmt.Println("==========================================================================")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", sum))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
		}
	}
	return builder.String()
}

func expensesDetails(expensesList []float32) (max, min, avg, sum float32) {

	if len(expensesList) == 0 {
		return
	}

	sum = expenses.Sum(expensesList...)
	min = expenses.Min(expensesList...)
	max = expenses.Max(expensesList...)
	avg = expenses.Average(expensesList...)
	return
}

func Export(fileName string, list []float32) error {
	f, err := os.Create(fileName)
	if err != nil {
		
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(list))
	if err != nil {
		return err
	}

	return w.Flush()
}
