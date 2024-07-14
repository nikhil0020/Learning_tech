package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func EncodingJSONData() {
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var pointer *int = &ival
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}
	fmt.Print(writer.String())

	// Notice that I used the fmt.Print function to produce the output. The JSON Encoder adds a newline character after each value is encoded.
}

func EncodingSliceAndArrays() {
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)
	fmt.Print(writer.String())
}

func EncodingMap() {
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.25,
	}

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	encoder.Encode(m)

	fmt.Print(writer.String())

}

type P struct {
	Name string
	Age  int16
}

func EncodingStruct() {
	p1 := &P{"Nikhil", 21}

	var writer strings.Builder

	encoder := json.NewEncoder(&writer)

	encoder.Encode(p1)

	fmt.Print(writer.String())
}

func EffectOfPromotion() {
	Laptop := &Product{
		Name:     "Mac",
		Category: "Electronics",
		Price:    89000,
	}

	dp := DiscountedProduct{
		Laptop,
		10000,
	}

	var writer strings.Builder

	encoder := json.NewEncoder(&writer)

	encoder.Encode(&dp)

	fmt.Print(writer.String())
}

func CustomizingJSONEncodingOfStructs() {
	Laptop := &Product{
		Name:     "Mac",
		Category: "Electronics",
		Price:    89000,
	}

	dp := DiscountedProductCustomField{
		Laptop,
		10000,
	}

	var writer strings.Builder

	encoder := json.NewEncoder(&writer)

	encoder.Encode(&dp)

	fmt.Print(writer.String())
}

func OmitField() {
	Laptop := &Product{
		Name:     "Mac",
		Category: "Electronics",
		Price:    89000,
	}

	dp := DiscountedProductOmitField{
		Laptop,
		10000,
	}

	var writer strings.Builder

	encoder := json.NewEncoder(&writer)

	encoder.Encode(&dp)

	fmt.Print(writer.String())
}

func OmittingUnassignedFields() {
	var writer strings.Builder
	Laptop := &Product{
		Name:     "Mac",
		Category: "Electronics",
		Price:    89000,
	}
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProductOmitUnassigned{
		Product:  Laptop,
		Discount: 10.50,
	}
	encoder.Encode(&dp)
	dp2 := DiscountedProductOmitUnassigned{Discount: 10.50}
	encoder.Encode(&dp2)
	fmt.Print(writer.String())
}

func EncodingInterface() {
	var writer strings.Builder
	Laptop := &Product{
		Name:     "Mac",
		Category: "Electronics",
		Price:    89000,
	}
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  Laptop,
		Discount: 10.50,
	}
	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)
	fmt.Print(writer.String())
}

// Implementation of marsheler interface

func (dp *DiscountedProductOmitUnassigned) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}

		jsn, err = json.Marshal(m)
	}

	return
}

func ImplementationOfMarshalerInterface() {
	var writer strings.Builder
	Laptop := &Product{
		Name:     "Mac",
		Category: "Electronics",
		Price:    89000,
	}
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProductOmitUnassigned{
		Product:  Laptop,
		Discount: 10999,
	}
	encoder.Encode(&dp)
	dp2 := DiscountedProductOmitUnassigned{Discount: 10.50}
	encoder.Encode(&dp2)
	fmt.Print(writer.String())
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

// Decoding JSON data

func DecodingJSONData() {
	reader := strings.NewReader(`true "hello" 99.20 200`)
	vals := []interface{}{}

	decoder := json.NewDecoder(reader)

	for {
		var decoderVal interface{}

		err := decoder.Decode(&decoderVal)

		if err != nil {
			if err != io.EOF {
				Printfln("Error %v", err.Error())
			}
			break
		}

		vals = append(vals, decoderVal)
	}

	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
}

func DecodingNumbers() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}
}

func DecodingSpecificType() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	var bval bool
	var sval string
	var fpval float64
	var ival int
	vals := []interface{}{&bval, &sval, &fpval, &ival}
	decoder := json.NewDecoder(reader)
	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", ival, ival)
}

func DecodingArrays() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
}

func DecodingSpecifiedArrays() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	ints := []int{}
	mixed := []interface{}{}
	vals := []interface{}{&ints, &mixed}
	decoder := json.NewDecoder(reader)
	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", ints, ints)
	Printfln("Decoded (%T): %v", mixed, mixed)
}

func DecodingMaps() {
	reader := strings.NewReader(`{"Kayak" : 279 , "Lifejacket" : 49.95}`)

	m := map[string]interface{}{}

	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&m)

	if err != nil {
		Printfln("Error %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m, m)
		for k, v := range m {
			Printfln("Key: %v , Value: %v", k, v)
		}
	}
}

func DecodingMapWithSpecificValueType() {
	reader := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m := map[string]float64{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m, m)
		for k, v := range m {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}
}

func DecodingStruct() {
	reader := strings.NewReader(`
        {"Name":"Kayak","Category":"Watersports","Price":279}
        {"Name":"Lifejacket","Category":"Watersports" }
        {"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
    `)
	decoder := json.NewDecoder(reader)
	for {
		var val Product
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		}
	}
}

func main() {
	// EncodingJSONData()
	// EncodingSliceAndArrays()
	// EncodingMap()
	// EncodingStruct()
	// EffectOfPromotion()
	// CustomizingJSONEncodingOfStructs()
	// OmitField()
	// OmittingUnassignedFields()
	// EncodingInterface()
	// ImplementationOfMarshalerInterface()

	// Decoding

	// DecodingJSONData()
	// DecodingNumbers()
	// DecodingSpecificType()
	// DecodingArrays()
	// DecodingSpecifiedArrays()
	// DecodingMaps()
	// DecodingMapWithSpecificValueType()
	DecodingStruct()
}
