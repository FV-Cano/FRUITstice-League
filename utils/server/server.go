package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"utils/globals"
	"utils/inventory"
	"utils/jsonFile"
)

/**
* @desc: Inicia un servidor en el puerto 8080

* @param ctx: contexto de la aplicación
 */
func ServerStart(ctx context.Context) {
	mux := http.NewServeMux()
	
	// Handlers
	mux.HandleFunc("/inventory", getInventory)
	mux.HandleFunc("/newSale", newSale)
	mux.HandleFunc("/report", genReport)
	
	newServer := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		<-ctx.Done()
		fmt.Println("Apagando servidor...")
		newServer.Shutdown(context.Background())
	}()

	fmt.Print("Servidor iniciado en http://127.0.0.1:8080\n")
	if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	fmt.Println("Servidor apagado.")
}

/**
* @desc: Función genérica para realizar solicitudes HTTP

* @param method: método de la solicitud
* @param url: URL de la solicitud
* @param requestBody: cuerpo de la solicitud
* @param responseBody: cuerpo de la respuesta
* @return: error (si hubiese)
 */
func DoRequest(method, url string, requestBody interface{}, responseBody interface{}) error {
	var reqBody []byte
	var err error
	if requestBody != nil {
		reqBody, err = json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("fallo al codificar el cuerpo de la solicitud: %w", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("fallo al crear la solicitud: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("fallo al realizar la solicitud: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("respuesta no exitosa: %d", resp.StatusCode)
	}

	if responseBody != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("fallo al leer la respuesta: %w", err)
		}
		err = json.Unmarshal(body, responseBody)
		if err != nil {
			return fmt.Errorf("fallo al decodificar la respuesta: %w", err)
		}
	}

	return nil
}

/**
* @desc: Registra una nueva venta

* @param input: datos de la venta
*/
func RequestNewSale(input string) {
	parts := strings.Split(input, ",")
	if len(parts) != 3 {
		fmt.Println("Error: Formato incorrecto.")
		return
	}

	quantity, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		fmt.Println("Error: La cantidad debe ser un número entero.")
		return
	}

	saleReq := globals.SaleRequest{
		ClientName:  strings.TrimSpace(parts[0]),
		ProductName: strings.ToLower(strings.TrimSpace(parts[1])),
		Quantity:    quantity,
	}

	DoRequest("POST", "http://localhost:8080/newSale", saleReq, nil)
}

// =================== API =====================

/**
* @desc: Request para obtener el inventario
*/
func getInventory(w http.ResponseWriter, r *http.Request) {
	// Obtener inventario
	inventory, err := jsonFile.Read()
	if err != nil {
		http.Error(w, "Error al obtener el inventario", http.StatusInternalServerError)
		return
	}

	// Convertir el inventario a JSON
	inventoryJSON, err := json.Marshal(inventory)
	if err != nil {
		http.Error(w, "Error al convertir el inventario a JSON", http.StatusInternalServerError)
		return
	}

	// Responder con el inventario
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(inventoryJSON)
}

/**
* @desc: Request para registrar una nueva venta
*/
func newSale(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud
	var saleReq globals.SaleRequest
	err := json.NewDecoder(r.Body).Decode(&saleReq)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	inventory.RegisterSale(saleReq.ClientName, saleReq.ProductName, saleReq.Quantity)

	// Responder con el inventario actualizado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Venta registrada"))
}

func genReport(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud
	var reportReq globals.ReportRequest
	err := json.NewDecoder(r.Body).Decode(&reportReq)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	inventory.CreateReport(reportReq.Type)

	w.WriteHeader(http.StatusOK)
}