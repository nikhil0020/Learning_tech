# Gin Framework

## Controlling log output colouring 
* To disable log’s color 
  * Run - gin.DisableConsoleColor()
* To Force log’s color 
  * Run - gin.ForceConsoleColor() , inside main function
 
## Custom HTTP configuration
* Use http.ListenAndServe()
```go
  func main() {
	router := gin.Default()
	http.ListenAndServe(":8080", router)
}    
```

#### OR

```go
func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
```

