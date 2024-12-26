package main

import (
	"flag"
	"fmt"
	"golandcli/commands"
	"log"
	"strconv"
)

func main() {
	var expenses []float32
	var export string
	flag.StringVar(&export, "export", "", "Exportar detalle del .txt")
	flag.Parse()
	for {
		input, err := commands.GetInput()
		if err != nil {
			log.Fatal(err)
		}
		if input == "stop" {
			break
		}
		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			fmt.Println("Ingresa un nuevo valor correcto; solo se permiten números.")
			continue
		}

		expenses = append(expenses, float32(expense))
	}

	if export == "" {
		fmt.Println("Por Favor Colocar Nombre Al Archivo ")
		commands.ShowInConsole(expenses)
	} else {
		commands.Export(export, expenses)
	}
}
