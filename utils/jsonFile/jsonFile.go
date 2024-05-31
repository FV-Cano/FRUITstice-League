package jsonFile

import (
	"encoding/json"
	"os"
	"utils/globals"
)

/**
* @desc: Escribe un slice de structs de tipo "Fruit" en el archivo "inventory.json"
 */
func Write() error {
	file, err := os.Create("../utils/inventory/inventory.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(globals.FruitInventory)
	if err != nil {
		return err
	}

	return nil
}

/**
* @desc: Lee el archivo "inventory.json" y lo convierte en un slice de structs de tipo "Fruit"

* @return: slice de structs de tipo "Fruit" , error (si hubiese)
 */
func Read() ([]globals.Fruit, error) {
	file, err := os.Open("../utils/inventory/inventory.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fruits []globals.Fruit
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fruits)
	if err != nil {
		return nil, err
	}

	return fruits, nil
}