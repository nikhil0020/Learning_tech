package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func processData(reader io.Reader) {
	b := make([]byte, 2)

	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}

		if err == io.EOF {
			break
		}
	}
}

func processDataReaderWriter(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)

	for {
		count, err := reader.Read(b)

		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
			writer.Write(b[0:count])
		}

		if err == io.EOF {
			break
		}
	}
}

func ReadingDataFromString() {
	r := strings.NewReader("Kayak")
	processData(r)
}

func ReadingDataFromStringAndWritingInStringBuilding() {
	r := strings.NewReader("This is Nikhil Singh")
	var builder strings.Builder

	processDataReaderWriter(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}

// Using The Utility Functions for Readers and Writers

// Using the Copy function achieves the same result as the previous example but more concisely.
func processDataUtilityFunction(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)

	if err != nil {
		Printfln("Error: %v", err.Error())
	}

	Printfln("Read %v bytes", count)
}

func ReadWriteWithUtilityFunction() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processDataUtilityFunction(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}

// Using the specialized Readers and Writers

// Using Pipes

// Pipes are used to connect code that consumes data through a Reader
// and code that produces code through a Writer. Add a file named data.go

// The io.Pipe function returns a PipeReader and a PipeWriter.
// The PipeReader and PipeWriter structs implement the Closer interface,
func UsingPipes() {
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)
}

// Concatenating Multiple Readers

func MultipleReaders() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")

	concatReader := io.MultiReader(r1, r2, r3)

	ConsumeData(concatReader)
}

// Combining multiple writer

func MultipleWriter() {
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder

	combinedWriter := io.MultiWriter(&w1, &w2, &w3)

	GenerateData(combinedWriter)

	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())
}

// Echoing Reads to a Writer

func EchoingReadsToWriter() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")

	concatReader := io.MultiReader(r1, r2, r3)

	var writer strings.Builder

	teeReader := io.TeeReader(concatReader, &writer)

	ConsumeData(teeReader)

	Printfln("Echo data: %v", writer.String())
}

// Limiting Read Data

func LimitingReadData() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)

	limited := io.LimitReader(concatReader, 5)

	ConsumeData(limited)
}

// Buffering data

// To see how data is processed without buffer

func ProcessingDataWithoutBuffer() {
	text := "It was a boat. A small boat."

	var reader io.Reader = NewCustomReader(strings.NewReader(text))

	var writer strings.Builder

	slice := make([]byte, 5)

	for {
		count, err := reader.Read(slice)

		if count > 0 {
			writer.Write(slice[0:count])
		}

		if err != nil {
			break
		}
	}

	Printfln("Read Data: %v", writer.String())
}

// Using bufio

func UsingBufferedReader() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)
	reader = bufio.NewReader(reader)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writer.String())
}

// Additional Buffered Reader methods

func AdditionalBufferedReaderMethod() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)
	buffered := bufio.NewReader(reader)

	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}

		if err != nil {
			break
		}
	}
	Printfln("Read Data: %v", writer.String())

}

// Custom Writer

// Performing without buffer

func WritingDataWithoutBuffer() {
	text := "It was a boat. A small boat."

	var builder strings.Builder

	var writer = NewCustomWriter(&builder)

	for i := 0; true; {
		end := i + 5

		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			break
		}

		writer.Write([]byte(text[i:end]))
		i = end
	}

	Printfln("Written data: %v", builder.String())
}

// Using Bufio

// The buffered Writer keeps data in a buffer and passes it on to the underlying
// Writer only when the buffer is full or when the Flush method is called.

func UsingBufferedWriter() {

	text := "It was a boat. A small boat."
	var builder strings.Builder
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			writer.Flush()
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())
}

// Formatting and Scanning from reader and writer

// Scanning values from a reader

func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func ScanningFromReader() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func ScanningSingle() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		_, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}
}

// Writing Formatted String to a writer

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func WritingFormattedStringToWriter() {
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writer.String())
}

// Using a replacer with a writer

func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}

func UsingReplacerWithWriter() {
	text := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	var writer strings.Builder
	writeReplaced(&writer, text, subs...)
	fmt.Println(writer.String())
}

func main() {
	// ReadingDataFromString()
	// ReadingDataFromStringAndWritingInStringBuilding()
	// ReadWriteWithUtilityFunction()
	// UsingPipes()
	// MultipleReaders()
	// MultipleWriter()
	// EchoingReadsToWriter()
	// LimitingReadData()
	// ProcessingDataWithoutBuffer()
	// UsingBufferedReader()
	// AdditionalBufferedReaderMethod()
	// WritingDataWithoutBuffer()
	// UsingBufferedWriter()
	// ScanningFromReader()
	// ScanningSingle()
	// WritingFormattedStringToWriter()
	UsingReplacerWithWriter()
}
