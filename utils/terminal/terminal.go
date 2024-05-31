package terminal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"utils/globals"

	"golang.org/x/term"
)

/**
* @desc: Lee los datos ingresados en la terminal
 */
func ReadFromTerminal() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input == "" {
		return input
	}
	
	return input
}

/**
* @desc: Lee los datos ingresados por el usuario y los almacena en el slice de structs "FruitInventory"

* @param scanner: puntero a un objeto de tipo "bufio.Scanner"
* @return: struct de tipo "Fruit", booleano
 */
func FruitInput() (globals.Fruit, bool) {
	input := ReadFromTerminal()
	if input == "" {
		return globals.Fruit{}, false
	}
	
	parts := strings.Split(input, ",")
	if len(parts) != 4 {
		fmt.Println("Por favor, ingrese los datos en el formato correcto.")
		return globals.Fruit{}, false
	}

	fruitName := strings.ToLower(strings.TrimSpace(parts[0]))
	stock, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		fmt.Println("Error: El stock debe ser un número entero.")
		return globals.Fruit{}, false
	}
	
	unitPrice, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	if err != nil {
		fmt.Println("Error: El precio debe ser un número decimal.")
		return globals.Fruit{}, false
	}

	expireDate := strings.TrimSpace(parts[3])
	expireDateParsed, err := stringToTime(expireDate)
	if err != nil {
		fmt.Println("Error: La fecha de expiración debe ser en el formato 'YYYY-MM-DD'.")
		return globals.Fruit{}, false
	}

	newFruit := globals.Fruit{
		Name: 		fruitName,
		Stock: 		stock,
		UnitPrice: 	unitPrice,
		ExpDate: 	expireDateParsed,
	}
	
	return newFruit, true
}

/**
* @desc: Convierte un string en una variable de tipo "time.Time"

* @param s: string a convertir
* @return: variable de tipo "time.Time", error (si hubiese)
 */
func stringToTime(s string) (time.Time, error) {
	dateFormat := "2006-01-02"
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

/**
* @desc: Muestra el menú de opciones en la terminal
 */
func DisplayMenu() {
	width := getTerminalWidth()
	printDivider(width)
	menuTitle := "Menú de opciones"
	fmt.Println(centerText(menuTitle, width))
	fmt.Println("1. Agregar Producto")
	fmt.Println("2. Modificar Producto")
	fmt.Println("3. Eliminar Producto")
	fmt.Println("4. Ver Inventario")
	fmt.Println("5. Registrar Venta")
	fmt.Println("6. Generar Reporte")
	fmt.Println("7. Salir")
	printDivider(width)

}

/**
* @desc: Obtiene el ancho de la terminal
 */
func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return 80
	}
	return width
}

/**
* @desc: Centra un texto en la terminal
 */
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", padding)
}

/**
* @desc: Imprime un divisor en la terminal
 */
func printDivider(width int) {
	fmt.Println(strings.Repeat("=", width))
}

/**
* @desc: Limpia la terminal
 */
func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}