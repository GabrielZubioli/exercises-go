package main
import "fmt"

func main() {
    var dollar, quantDollar float64
    fmt.Print("Qual a cotação do dollar? ")
    fmt.Scanln(&dollar)
    fmt.Print("Quantos dolares você tem? ")
    fmt.Scanln(&quantDollar)

    var real = dollar * quantDollar

    fmt.Printf("O valor convertido em reais é R$%.2f", real)
}
