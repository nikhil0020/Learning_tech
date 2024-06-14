package main

import "io"

// The GenerateData function defines a Writer parameter, which it uses to write bytes from a string.

// I called the Close method on the PipeWriter in the goroutine that executes the GenerateData function.
// This works, but I prefer to check to see whether a Writer implements the Closer interface
// in the code that produces the data
func GenerateData(writer io.Writer) {
	data := []byte("Kayak, LifeJacket")

	writeSize := 4

	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}

		count, err := writer.Write(data[i:end])

		Printfln("Wrote %v byte(s): %v", count, string(data[i:end]))

		if err != nil {
			Printfln("Error: %v", err.Error())
		}
	}

	if closer, ok := writer.(io.Closer); ok {
		closer.Close()
	}
}

// The ConsumeData function defines a Reader parameter, which it uses to
// read bytes of data, which are then used to create a string.
func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)

	for {
		count, err := reader.Read(slice)

		if count > 0 {
			Printfln("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}

		if err == io.EOF {
			break
		}
	}

	Printfln("Read data: %v", string(data))
}
