package main

import (
	"io"
	"net/http"
	"strings"
)

type StringHandler struct {
	message string
}

// Creating a simple http server

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	io.WriteString(writer, sh.message)
// }

// Inspecting the request

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
// 	request *http.Request) {
// 	Printfln("Method: %v", request.Method)
// 	Printfln("URL: %v", request.URL)
// 	Printfln("HTTP Version: %v", request.Proto)
// 	Printfln("Host: %v", request.Host)
// 	for name, val := range request.Header {
// 		Printfln("Header: %v, Value: %v", name, val)
// 	}
// 	Printfln("---")
// 	io.WriteString(writer, sh.message)
// }

// Filtering Request and generating response

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
// 	request *http.Request) {
// 	if request.URL.Path == "/favicon.ico" {
// 		Printfln("Request for icon detected - returning 404")
// 		writer.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	Printfln("Request for %v", request.URL.Path)
// 	io.WriteString(writer, sh.message)
// }

// Using the response convenience functions

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
// 	request *http.Request) {
// 	Printfln("Request for %v", request.URL.Path)
// 	switch request.URL.Path {
// 	case "/favicon.ico":
// 		http.NotFound(writer, request)
// 	case "/message":
// 		io.WriteString(writer, sh.message)
// 	default:
// 		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
// 	}
// }

// Routing handler

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func UsingRoutingHandler() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func CreatingSimpleHttpServer() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})

	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func HTTPSRedirect(writer http.ResponseWriter, request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5500" + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func HandleHTTPRedirectToHTTPS() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer",
			"certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()
	err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func CreatingStaticFileRoute() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))
	// go func() {
	// 	err := http.ListenAndServeTLS(":5500", "certificate.cer",
	// 		"certificate.key", nil)
	// 	if err != nil {
	// 		Printfln("HTTPS Error: %v", err.Error())
	// 	}
	// }()
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	// for _, p := range Products {
	// 	Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	// }
	// CreatingSimpleHttpServer()
	// UsingRoutingHandler()
	// HandleHTTPRedirectToHTTPS() // Add certificate.cer and certificate.key
	CreatingStaticFileRoute()
}
