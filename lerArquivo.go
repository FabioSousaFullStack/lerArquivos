package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	arquivo, err := ioutil.ReadFile("arquivo.txt")

	fmt.Println(string(arquivo))
	
	if err != nil {
		fmt.Println("Aconteceu o seguinte erro: ",err)
	}
}