package main

import "fmt"

func dibuja_matrices(MAATRIZ [8][8]int)(){
	for i:=0; i<len(MAATRIZ);i++ {
		fmt.Printf("%d", MAATRIZ[i])
		fmt.Printf("\n")
	}

}

func pintar_tablero()(tablero2 [8][8]int){
		//Pintamos de negro las casillas impares
		for j:=0; j<len(tablero2);j=j+2 {
			for i:=1; i<len(tablero2);i=i+2 {
				tablero2[j][i] = 1
			}
		}
		//Pintamos de negro las casillas pares
		for j:=1; j<len(tablero2);j=j+2 {
			for i:=0; i<len(tablero2);i=i+2 {
				tablero2[j][i] = 1
			}
		}
	return
}

func main() {
	//Tablero
	//var tablero [8][8]int

	
	x:= pintar_tablero()

	dibuja_matrices(x)	
}