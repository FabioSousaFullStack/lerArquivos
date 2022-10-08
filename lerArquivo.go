package main

import(
	"fmt"
	"os"
	"bufio"
	
)

func main(){
	arquivo, err := os.Open("arquivo.txt")

	
	if err != nil {
		fmt.Println("Aconteceu o seguinte erro: ",err)
	}

	leitor := bufio.NewReader(arquivo)
	
	linha, err := leitor.ReadString('m')

	if err != nil {
			fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(linha)
}