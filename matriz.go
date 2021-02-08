package main
import "fmt"

func dibuja_matrices(MAATRIZ [3][3]int)(){
	for i:=0; i<len(MAATRIZ);i++ {
		fmt.Printf("%d", MAATRIZ[i])
		fmt.Printf("\n")
	}

}




func main(){

	var matriz [3][3]int

	for i:=0; i<len(matriz);i++ {
		matriz[i][0] = i
	}

	dibuja_matrices(matriz)


}