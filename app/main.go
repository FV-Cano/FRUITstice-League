package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"utils/globals"
	"utils/inventory"
	"utils/server"
	"utils/terminal"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Inicia el servidor
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.ServerStart(ctx)
	}()

	time.Sleep(3 * time.Second)

	// Carga inicial de inventario
	fmt.Println("Bienvenido al sistema de inventario. Ingrese 1 para carga automática o 2 para carga manual.")
	option := terminal.ReadFromTerminal()
	if option == "1" {
		inventory.AutoLoad()
	} else {
		fmt.Print("Ingrese los datos de la fruta con el formato 'Nombre, Stock, Precio, Fecha de expiración (YYYY-MM-DD)' (presione 'Enter' para terminar de cargar el stock inicial): ")
		inventory.InventoryInit()
	}
	terminal.ClearTerminal()

	for {
		// Mostrar menú
		terminal.DisplayMenu()

		// Leer opción
		option := terminal.ReadFromTerminal()
		if option == "" {
			continue
		}

		// Ejecutar opción
		switch option {
		case "1":
			fmt.Println("Ingrese los datos de la fruta nueva fruta con el formato 'Nombre, Stock, Precio, Fecha de expiración (YYYY-MM-DD)': ")
			newFruit, ok := terminal.FruitInput()
			if !ok {
				fmt.Println("Error: Datos incorrectos.")
				terminal.ClearTerminal()
				continue
			}
			inventory.AddFruit(newFruit)
			terminal.ClearTerminal()
			fmt.Println("Fruta agregada correctamente.")
		case "2":
			fmt.Println("Ingrese el nombre de la fruta a modificar: ")
			fruitName := strings.ToLower(terminal.ReadFromTerminal())
			if fruitName == "" {
				fmt.Println("Error: No se ingresó un nombre.")
				continue
			}
			inventory.ModFruit(fruitName)
			terminal.ClearTerminal()
			fmt.Println("Fruta modificada correctamente.")
		case "3":
			fmt.Println("Ingrese el nombre de la fruta a eliminar: ")
			fruitName := strings.ToLower(terminal.ReadFromTerminal())
			if fruitName == "" {
				fmt.Println("Error: No se ingresó un nombre.")
				terminal.ClearTerminal()
				continue
			}
			inventory.DelFruit(fruitName)
			terminal.ClearTerminal()
			fmt.Println("Fruta eliminada correctamente.")
		case "4":
			var inventory []globals.Fruit
			server.DoRequest("GET", "http://localhost:8080/inventory", nil, &inventory)
			fmt.Println("Inventario:", inventory)
			fmt.Print("\nPresione 'Enter' para continuar...")
			terminal.ReadFromTerminal()
			terminal.ClearTerminal()
		case "5":
			fmt.Print("Ingrese los datos de la venta con el formato 'Cliente, Producto, Cantidad': ")
			input := terminal.ReadFromTerminal()
			if input == "" {
				fmt.Println("Error: No se ingresaron datos")
			}
			server.RequestNewSale(input)
			fmt.Print("\nPresione 'Enter' para continuar...")
			terminal.ReadFromTerminal()
			terminal.ClearTerminal()
		case "6":
			fmt.Println("Seleccione el tipo de reporte: \n1. Diario\n2. Semanal\n3. Mensual")
			reportType := terminal.ReadFromTerminal()
			if reportType == "" {
				fmt.Println("Error: No se ingresó un tipo de reporte.")
				continue
			}
			switch reportType {
			case "1":
				reportReq := globals.ReportRequest{Type: "daily"}
				server.DoRequest("POST", "http://localhost:8080/report", reportReq, nil)
				terminal.ClearTerminal()
			case "2":
				reportReq := globals.ReportRequest{Type: "weekly"}
				server.DoRequest("POST", "http://localhost:8080/report", reportReq, nil)
				terminal.ClearTerminal()
			case "3":
				reportReq := globals.ReportRequest{Type: "monthly"}
				server.DoRequest("POST", "http://localhost:8080/report", reportReq, nil)
				terminal.ClearTerminal()
			}
		case "7":
			cancel()
			wg.Wait()
			os.Exit(0)
		}
	}
}