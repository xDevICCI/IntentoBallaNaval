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
	horizontal bool
	numero     int
	vida       bool
}

func inicializar_barco() barco {

	m := barco{
		horizontal: true,
		id:         1,
		numero:     1,
	}
	return m
}

func insert_barcos_matriz(vector_barco []barco, matriz [][]barco, posicion int) { //funcion para insertar barcos
	//vector_barco[1] = append(vector_barco, matriz[3][5])
	for i := posicion; i < len(vector_barco); i++ {
		rand.Seed(time.Now().UnixNano())
		s1 := rand.Intn(len(matriz) - 2)
		s2 := rand.Intn(len(matriz) - 2)
		print("\n barco posicion [", s1, "][", s2, "] , en la iteracion N ", i+1)
		var aux barco
		aux.id = 1
		if matriz[s1][s2].id != 0 { // casilla ocupada
			s1 = rand.Intn(len(matriz) - 2)
			s2 = rand.Intn(len(matriz) - 2)
			if matriz[s1+1][s2].id == 0 && matriz[s1+2][s2].id == 0 && matriz[s1][s2].id == 0 {
				for k := 0; k < 3; k++ {
					matriz[s1][s2+k] = vector_barco[i]
				}
				posicion = posicion + 1
			} else {
				s1 = rand.Intn(len(matriz) - 2)
				s2 = rand.Intn(len(matriz) - 2)
				if matriz[s1][s2+1].id == 0 && matriz[s1][s2+2].id == 0 {
					for k := 0; k < 3; k++ {
						matriz[s1+k][s2] = vector_barco[i]
						matriz[s1][s2].horizontal = false
					}
				}
			}
		} else if matriz[s1][s2].id == 0 { // caso cuando las casillas esten desocupadas
			if matriz[s1][s2+1].id == 0 && matriz[s1][s2+2].id == 0 {
				for k := 0; k < 3; k++ {
					matriz[s1][s2+k] = vector_barco[i]
				}
				posicion = posicion + 1
			} else {
				print(" 2 posicion ocupada")
			}
		} else {
			print("caso que nose")
		}
	}
}

func atacar(mp [][]barco,vector_barco []) {
	print("funcion para atacar")
	for i:=0; i<len(vector_barco);i++{
		for j:=0;j<len(vector_barco);j++{
			print("asd")
		}
	}
}

func crear_barco(cant_jugadores int, cant_barcos int) []barco {

	var vector_barco []barco //creamos la variable para guardar los datos barco

	for i := 0; i < cant_barcos; i++ {
		vector_barco = append(vector_barco, inicializar_barco()) // agregar los datos
		vector_barco[i].numero = vector_barco[i].numero + i
		print(" \n atributo numero del barco ", vector_barco[i].numero)
	}

	return vector_barco

}

func imprimir(mp [][]barco) {
	print(" º")
	for k := 0; k < len(mp); k++ {
		print("  ", k, "")
	}
	for i := 0; i < len(mp); i++ {
		print("\n ", i, " ")
		for j := 0; j < len(mp); j++ {
			var aux, aux2 barco
			aux.id = 1
			aux = mp[i][j]
			if aux.id == aux2.id {
				print(" - ")
			} else {
				print(" ", aux.numero, " ")
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
	insert_barcos_matriz(vector, mp, 0)
	print("\n")
	imprimir(mp)
}
