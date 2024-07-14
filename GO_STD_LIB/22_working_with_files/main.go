package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func UsingWriteConvenienceFunction() {
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}
	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n",
		time.Now().Format("Mon 15:04:05"), total)
	err := os.WriteFile("output.txt", []byte(dataStr), 0666)
	if err == nil {
		fmt.Println("Output file created")
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func UsingOpenFileForWriting() {
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}
	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)

	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err == nil {
		defer file.Close()

		file.WriteString(dataStr)
	} else {
		Printfln("Error: %v", err.Error())
	}

}

func WritingJSONdata() {
	cheapProducts := []Product{}

	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}

	file, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)

	if err == nil {
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func CreatingTemporaryFile() {
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	file, err := os.CreateTemp(".", "tempfile-*.json")
	if err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func WorkingWithPath() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir component: %v", filepath.Dir(path))
	Printfln("File component: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))
}

func CreatingDirectories() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	err = os.MkdirAll(filepath.Dir(path), 0766)
	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}

func EnumeratingFiles() {
	path, err := os.Getwd()
	if err == nil {
		dirEntries, err := os.ReadDir(path)
		if err == nil {
			for _, dentry := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}

// Determining Whether a File Exists

func CheckingWhetherFileExist() {
	targetFiles := []string{"no_such_file.txt", "config.json"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		if os.IsNotExist(err) {
			Printfln("File does not exist: %v", name)
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File %v, Size: %v", info.Name(), info.Size())
		}
	}
}

func LocatingFiles() {
	path, err := os.Getwd()
	if err == nil {
		matches, err := filepath.Glob(filepath.Join(path, "*.json"))
		if err == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v", path, info.Size())
	return
}

func WalkingADirectory() {
	path, err := os.Getwd()
	if err == nil {
		err = filepath.WalkDir(path, callback)
	} else {
		Printfln("Error %v", err.Error())
	}
}

func main() {
	// for _, p := range Products {
	// 	Printfln("Product %v, Category %v, Price %.2f", p.Name, p.Category, p.Price)
	// }

	// UsingWriteConvenienceFunction()
	// UsingOpenFileForWriting()
	// WritingJSONdata()
	// CreatingTemporaryFile()
	// WorkingWithPath()
	// CreatingDirectories()
	// EnumeratingFiles()
	// CheckingWhetherFileExist()
	// LocatingFiles()
	WalkingADirectory()
}
