# Creating and Using Packages

## Understanding the Module File

* The original purpose of a module file was to enable code to be published so that it can be used in other projects and, potentially, by other developers. Module files are still used for this purpose, but Go has started to gain mainstream development, and as this has happened, the percentage of projects that are published has fallen.
* These days, the most common reason for creating a module file is that it makes it easy to install packages that have been published and has the bonus effect of allowing the use of the command

## Creating a Custom Package

* Packages make it possible to add structure to a project so that related features are grouped together.
* Create the packages/store folder and add to it a file named product.go
* Create a separate folder for a particular package
* Multiple file can be a part of same package.

```go
package store
type Product struct {
    Name, Category string
    price float64
}
```

### Using a Custom Package

```go
package main
import (
    "fmt"
    "packages/store"
)
func main() {
    product := store.Product {
        Name: "Kayak",
        Category: "Watersports",
    }
    fmt.Println("Name:", product.Name)
    fmt.Println("Category:", product.Category)
}
```

The exported features provided by the package are accessed using the package name as a prefix, like this:
```go
...
var product *store.Product = &store.Product {
...
```

### Understanding Package Access Control

* Go has an unusual approach to access control. Instead of relying on dedicated keywords, like public and private, Go examines the first letter of the names given to the features in a code file, such as types, functions, and methods. If the first letter is lowercase, then the feature can be used only within the package that defines it. Features are exported for use outside of the package by giving them an uppercase first letter.

```go
package store
type Product struct {
    Name, Category string
    price float64
}
func NewProduct(name, category string, price float64) *Product {
    return &Product{ name, category, price }
}
func (p *Product) Price() float64 {
    return p.price
}
func (p *Product) SetPrice(newPrice float64)  {
    p.price = newPrice
}
```

* The access control rules do not apply to individual function or method parameters, which means that the NewProduct function has to have an uppercase first character to be exported, but the parameter names can be lowercase.


### Using a Package Alias
* One way to deal with package name conflicts is to use an alias, which allows a package to be accessed using a different name

### Using a Dot Import
* There is a special alias, known as the dot import, that allows a package’s features to be used without using a prefix