package main

import (
	"os"
	"strconv"
)

type mapa struct {
	barcos    int
	jugadores int
}

type barco struct {
	tamaño         int
	horizontal     int
	cantidad_barco int
}

/*func inicializar_barco() {
	m := barco{
		tamaño:         3,
		horizontal:     1,
		cantidad_barco: 1,
	}
}
*/
/*func imprimir(m *mapa) {
	for i := 0; i < 10; i++ {
		for j := 0; i < 10; j++ {
			fmt.Print(m.tablero[i][j])
		}
	}
}	*/

func imprimir(mp [][]int) {
	for i := 0; i < len(mp); i++ {
		print("\n")
		for j := 0; j < len(mp); j++ {
			print(mp[i][j])
		}
	}
}

func main() {
	
	cant_jugadores, _ := strconv.Atoi(os.Args[1])
	cant_barcos, _ := strconv.Atoi(os.Args[2])
	x, _ := strconv.Atoi(os.Args[3])
	y, _ := strconv.Atoi(os.Args[4])

	print("cantidad jugadores = ", cant_jugadores)
	print("\ncantidad barcos = ", cant_barcos)

	//crear mapa*
	mp := make([][]int, x)
	for i := 0; i < x; i++ {
		mp[i] = make([]int, y)
	}

	imprimir(mp)

}
