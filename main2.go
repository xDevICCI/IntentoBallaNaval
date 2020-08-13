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
	tamaño         [3]int
	horizontal     bool
	cantidad_barco int
}

func inicializar_barco() barco {

	m := barco{
		horizontal:     true,
		cantidad_barco: 1,
	}

	return m
}

/*func imprimir(m *mapa) {
	for i := 0; i < 10; i++ {
		for j := 0; i < 10; j++ {
			fmt.Print(m.tablero[i][j])
		}
	}
}	*/

func insert_barcos_matriz(vector_barco []barco, matriz [][]int) {

	for i := 0; i < len(vector_barco); i++ {
		for j := 0; j < len(vector_barco); j++ {
			matriz[i][j] = vector_barco
		}
	}
}

func crear_barco(cant_jugadores int, cant_barcos int) []barco {

	var vector_barco []barco //creamos la variable para guardar los datos barco

	for i := 0; i < cant_barcos; i++ {
		vector_barco = append(vector_barco, inicializar_barco()) // agregar los datos
	}

	return vector_barco

}

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
	var vector []barco
	vector = crear_barco(cant_jugadores, cant_barcos)
	insert_barcos_matriz(vector, mp)
}
