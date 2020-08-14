package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

type mapa struct {
	barcos    int
	jugadores int
}

type barco struct {
	id         int
	tamaño     int
	horizontal int
}

func inicializar_barco() barco {

	m := barco{
		horizontal: 0,
		id:         1,
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

func insert_barcos_matriz(vector_barco []barco, matriz [][]barco) {
	//vector_barco[1] = append(vector_barco, matriz[3][5])
	print("soy tamaño vector ", len(vector_barco), "\n")
	for i := 0; i < len(vector_barco); i++ {
		rand.Seed(time.Now().UnixNano())
		print(len(matriz))
		s1 := rand.Intn(len(matriz) - 2)
		s2 := rand.Intn(len(matriz) - 2)
		print("\n barco posicion [", s1, "][", s2, "] , en la iteracion N ", i)

		var aux barco
		aux.id = 1
		if matriz[s1][s2].id == 1 {
			print("\n encontre casilla ocupada [", s1, "][", s2, "] , en la iteracion N ", i)
		} else if matriz[s1][s2].id == 0 { // para no llenar ensima de un barco
			matriz[s1][s2] = vector_barco[i]
			matriz[s1][s2+1] = vector_barco[i]
			matriz[s1][s2+2] = vector_barco[i]
		} else {
			print("caso que nose")
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

func imprimir(mp [][]barco) {
	for i := 0; i < len(mp); i++ {

		print("\n")
		for j := 0; j < len(mp); j++ {
			var aux, aux2 barco
			aux.id = 1
			aux = mp[i][j]
			if aux.id == aux2.id {
				print(" 0 ")
			} else {
				print(" 1 ")
			}
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
	mp := make([][]barco, x)
	for i := 0; i < x; i++ {
		mp[i] = make([]barco, y)
	}

	var vector []barco
	vector = crear_barco(cant_jugadores, cant_barcos)
	insert_barcos_matriz(vector, mp)
	print("\n")
	imprimir(mp)
}
