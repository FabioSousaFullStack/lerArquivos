package main

import(
	"fmt"
	"os"
	
)

func main(){
	arquivo, err := os.Open("arquivo.txt")

	fmt.Println(arquivo)
	
	if err != nil {
		fmt.Println("Aconteceu o seguinte erro: ",err)
	}
}