package inventory

import (
	"fmt"
	"sort"
	"time"
	"utils/globals"
	"utils/jsonFile"
	"utils/slice"
	"utils/terminal"

	"github.com/jung-kurt/gofpdf"
)

/**
* @desc: Lee continuamente las entradas del usuario hasta un salto de línea y las almacena en un archivo "inventory.json"
 */
func InventoryInit() {
	for {
		newFruit, ok := terminal.FruitInput()
		if !ok {
			break
		}

		globals.FruitInventory = append(globals.FruitInventory, newFruit)
	}
	jsonFile.Write()
}

/**
* @desc: Agrega una nueva fruta al inventario

* @param newFruit: nueva fruta a agregar
 */
func AddFruit(newFruit globals.Fruit) {
	globals.FruitInventory = append(globals.FruitInventory, newFruit)
	jsonFile.Write()
}

/**
* @desc: Carga automáticamente el inventario con datos predefinidos
 */
func AutoLoad() {
	globals.FruitInventory = []globals.Fruit{
		{
			Name:       "batnana",
			Stock:      20,
			UnitPrice:  19.65,
			ExpDate:    time.Date(2024, 07, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:       "straw-bat-berry",
			Stock:      15,
			UnitPrice:  12.50,
			ExpDate:    time.Date(2024, 06, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:       "supergrape",
			Stock:      25,
			UnitPrice:  8.99,
			ExpDate:    time.Date(2024, 06, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:       "aquamelon",
			Stock:      12,
			UnitPrice:  15.25,
			ExpDate:    time.Date(2024, 07, 05, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:       "green kiwi-lantern",
			Stock:      22,
			UnitPrice:  9.45,
			ExpDate:    time.Date(2024, 07, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:       "cybapple",
			Stock:      17,
			UnitPrice:  14.20,
			ExpDate:    time.Date(2024, 06, 03, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:       "plum woman",
			Stock:      14,
			UnitPrice:  16.30,
			ExpDate:    time.Date(2024, 06, 03, 0, 0, 0, 0, time.UTC),
		},
	}
	jsonFile.Write()
}

/**
* @desc: Modifica una fruta en el inventario

* @param name: nombre de la fruta a modificar
* @param newFruit: nueva fruta con los datos actualizados
 */
func ModFruit(name string) {
	for i, fruit := range globals.FruitInventory {
		if fruit.Name == name {
			fmt.Printf("Ingrese los nuevos parámetros de la fruta con el formato 'Nombre, Stock, Precio, Fecha de expiración (YYYY-MM-DD)': ")
			newFruit, ok := terminal.FruitInput()
			if !ok {
				fmt.Println("Error: Datos incorrectos.")
				return
			}
			globals.FruitInventory[i] = newFruit
			jsonFile.Write()
			return
		}
	}
	fmt.Println("Error: No se encontró la fruta.")
}

/**
* @desc: Elimina una fruta del inventario

* @param name: nombre de la fruta a eliminar
 */
func DelFruit(name string) {
	for i, fruit := range globals.FruitInventory {
		if fruit.Name == name {
			globals.FruitInventory = append(globals.FruitInventory[:i], globals.FruitInventory[i+1:]...)
			jsonFile.Write()
			return
		}
	}
	fmt.Println("Error: No se encontró la fruta.")
}

/**
* @desc: Registra una nueva venta

* @param clientName: nombre del cliente
* @param productName: nombre del producto
* @param quantity: cantidad de productos a vender
*/
func RegisterSale(clientName string, productName string, quantity int) {
	for i, fruit := range globals.FruitInventory {
		if fruit.Name == productName {
			if fruit.Stock < quantity {
				fmt.Println("Error: No hay suficiente stock.")
				return
			}
			globals.FruitInventory[i].Stock -= quantity
			sale := globals.Sale{
				ClientName: clientName,
				ProductName: productName,
				Quantity: quantity,
				Total: fruit.UnitPrice * float64(quantity),
				SaleDate: time.Now(),
			}
			globals.Sales = append(globals.Sales, sale)
			jsonFile.Write()
			fmt.Println("Venta registrada correctamente.")
			return
		}
	}
	fmt.Println("Error: No se encontró la fruta.")
}

/**
* @desc: Filtra según el tipo de reporte y obtiene los datos necesarios para armarlo

* @param reportType: tipo de reporte
* @return: slice de strings, slice de structs de tipo "Fruit", float64
*/
func GetDataForReport(reportType string) ([]string, []globals.Fruit, float64){
	var filteredSales []globals.Sale
	switch reportType {
	case "daily":
		filteredSales = slice.Filter(globals.Sales, isSaleDaily)
	case "weekly":
		filteredSales = slice.Filter(globals.Sales, isSaleWeekly)
	case "monthly":
		filteredSales = slice.Filter(globals.Sales, isSaleMonthly)
	}

	mostSales := mostSelledProducts(filteredSales)
	lowStock := slice.Filter(globals.FruitInventory, isLowStock)
	totalIncome := getTotalIncome(filteredSales)

	return mostSales, lowStock, totalIncome
}

/**
* @desc: Determina si una venta fue realizada el día de hoy
*/
func isSaleDaily(sale globals.Sale) bool {
	return sale.SaleDate.Format("2006-01-02") == globals.Today.Format("2006-01-02")
}

/**
* @desc: Determina si una venta fue realizada en la última semana (7 días atrás) y hoy
*/
func isSaleWeekly(sale globals.Sale) bool {
	return sale.SaleDate.After(globals.Week) && sale.SaleDate.Before(globals.Today.AddDate(0, 0, 1))
}

/**
* @desc: Determina si una venta fue realizada en el último mes (30 días atrás) y hoy
*/
func isSaleMonthly(sale globals.Sale) bool {
	return sale.SaleDate.After(globals.Month) && sale.SaleDate.Before(globals.Today.AddDate(0, 0, 1))
}

/**
* @desc: Obtiene los 3 productos más vendidos, sin repetir

* @param sales: slice de structs de tipo "Sale"
* @return: slice de strings
*/
func mostSelledProducts(sales []globals.Sale) []string {
	sort.Sort(slice.ByPrice(sales))

	var products []string
	for _, sale := range sales {
		if len(products) == 3 {
			break
		}
		if !slice.Contains(products, sale.ProductName) {
			products = append(products, sale.ProductName)
		}
	}

	return products 		
}

/**
* @desc: Determina si una fruta tiene un stock menor al límite establecido

* @param fruit: struct de tipo "Fruit"
* @return: booleano
*/
func isLowStock(fruit globals.Fruit) bool {
	return fruit.Stock < globals.LowStock
}

/**
* @desc: Obtiene el ingreso total de las ventas según el tipo de reporte

* @param sales: slice de structs de tipo "Sale"
* @return: float64
*/
func getTotalIncome(sales []globals.Sale) float64 {
	var total float64
	for _, sale := range sales {
		total += sale.Total
	}
	return total
}

/**
* @desc: Crea un reporte en formato PDF según el tipo de reporte

* @param reportType: tipo de reporte
* @return: error (si hubiese)
*/
func CreateReport(reportType string) error {
	mostSales, lowStock, totalIncome := GetDataForReport(reportType)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Título
	pdf.Cell(40, 10, "Reporte de ventas")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Productos mas vendidos:")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	for _, product := range mostSales {
		pdf.Cell(40, 10, product)
	}
	pdf.Ln(10)
	
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Productos con bajo stock:")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	for _, fruit := range lowStock {
		pdf.Cell(40, 10, fruit.Name)
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Ingresos totales:")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("$%.2f", totalIncome))

	return pdf.OutputFileAndClose("report.pdf")
}