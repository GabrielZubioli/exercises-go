package main

import "fmt"

func main(){
	
	var votosA, votosB, votosC int
    var votosNulos, votosBrancos int

	fmt.Print("Quantos votos o primeio candidato recebeu? ")
    fmt.Scanln(&votosA)
	fmt.Print("Quantos votos o segundo candidato recebeu? ")
    fmt.Scanln(&votosB)
	fmt.Print("Quantos votos o terçeiro candidato recebeu? ")
    fmt.Scanln(&votosC)
	fmt.Print("Quantos votos foram nulos? ")
    fmt.Scanln(&votosNulos)
	fmt.Print("Quantos votos foram em brancos? ")
    fmt.Scanln(&votosBrancos)

	totalEleitores := votosA + votosB + votosC + votosNulos + votosBrancos
    votosValidos := votosA + votosB + votosC

	percValidos := float64(votosValidos) / float64(totalEleitores) * 100
	percValidosA := float64(votosA) / float64(totalEleitores) * 100
	percValidosB := float64(votosB) / float64(totalEleitores) * 100
	percValidosC := float64(votosC) / float64(totalEleitores) * 100
	percNulos := float64(votosNulos) / float64(totalEleitores) * 100
	percBrancos := float64(votosBrancos) / float64(totalEleitores) * 100


	fmt.Print("O percentual de foi:\n")
	fmt.Print("Votos validos em relação ao total de eleitores : ", percValidos,"%\n")
	fmt.Print("Votos validos candidato A em relação ao total de eleitores : ", percValidosA,"%\n")
	fmt.Print("Votos validos candidato B em relação ao total de eleitores : ", percValidosB,"%\n")
	fmt.Print("Votos validos candidato C em relação ao total de eleitores : ", percValidosC,"%\n")
	fmt.Print("Votos nulos em relação ao total de eleitores : ", percNulos,"%\n")
	fmt.Print("Votos nem branco em relação ao total de eleitores : ", percBrancos,"%\n")

}