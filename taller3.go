package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type mapa struct {
	guzanito gusano
	numero   int  // dato para llenar alimento
	activo   bool // true hay gusano // False no hay gusano
}

type gusano struct {
	tamaño     int  //
	id         int  //
	horizontal bool //
	comida     int  //cantidad de comida
	cabeza     bool //comienza en false // true para decir que es cabeza

}

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	Color1            = "\u001b[35m"
	Color2            = "\u001b[36m"
	ColorReset        = "\u001b[0m"
)

func iniciarlizar_gusano() gusano {
	m := gusano{
		tamaño:     3,
		id:         0,
		horizontal: true,
	}
	return m
}
func crear_map(mp [][]mapa, cantidad_gusano int, comida int) {
	for i := 0; i < cantidad_gusano; i++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(len(mp) - 2)
		y := rand.Intn(len(mp) - 2)
		if mp[x][y].activo == false { // si la casilla esta desocupada
			if mp[x+1][y].activo == false && mp[x+2][y].activo == false && x < 5 { // si las siguientes dos casillas en direccion abajo estan desocupadas llenar
				for j := 0; j < 3; j++ { //llene de forma vertical
					mp[x+j][y].guzanito = iniciarlizar_gusano()
					mp[x+j][y].guzanito.id = +i
					mp[x+j][y].activo = true
				}
				mp[x][y].guzanito.cabeza = true
			} else if mp[x][y+1].activo == false && mp[x][y+2].activo == false { //llene de forma horizontal
				for j := 0; j < 3; j++ {
					mp[x][y+j].guzanito = iniciarlizar_gusano()
					mp[x][y+j].guzanito.id = +i
					mp[x][y+j].activo = true
				}
				mp[x][y].guzanito.cabeza = true
			}
		} else if mp[x][y].activo == true { // si esta ocupada la casilla con un gusano
			for { // randomear hasta encontrar una casilla vacia
				s1 := rand.Intn(len(mp) - 2)
				s2 := rand.Intn(len(mp) - 2)
				if mp[s1][s2+1].activo == false && mp[s1][s2+2].activo == false && mp[s1][s2].activo == false {
					x = s1
					y = s2
					break
				}
			}
			for j := 0; j < 3; j++ { //llene de forma vertical
				mp[x][y+j].guzanito = iniciarlizar_gusano()
				mp[x][y+j].guzanito.id = +i
				mp[x][y+j].activo = true

			}
			mp[x][y].guzanito.cabeza = true
		} else {
			print("n")
		}
	}
	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp); j++ {
			if mp[j][i].activo == false { // si es 0 llenar con numeros random con limite entregado por consola
				s2 := rand.Intn(comida)
				mp[i][j].numero = s2
			}
		}
	}
}

func color(mp [][]mapa, i int, j int) { //funcion para pasar cambiar de color
	if mp[i][j].guzanito.id == 0 {
		print(ColorBlue, " ■ ")
	} else if mp[i][j].guzanito.id == 1 {
		print(ColorBlack, " ■ ")
	} else if mp[i][j].guzanito.id == 2 {
		print(ColorGreen, " ■ ")
	} else if mp[i][j].guzanito.id == 3 {
		print(ColorRed, " ■ ")
	} else if mp[i][j].guzanito.id == 4 {
		print(ColorYellow, " ■ ")
	} else if mp[i][j].guzanito.id == 5 {
		print(Color1, " ■ ")
	} else if mp[i][j].guzanito.id == 0 {
		print(Color2, " ■ ")
	} else {
		print(" ■ ")
	}
}

func imprimir(mp [][]mapa) {
	for i := 0; i < len(mp); i++ {
		print("\n")
		for j := 0; j < len(mp); j++ {
			if mp[i][j].activo == true {
				color(mp, i, j)
				//print(ColorBlue, " ", mp[i][j].guzanito.id, " ")
				print(ColorReset)
			} else {
				print(" ", mp[i][j].numero, " ")
			}
		}
	}
}
func comer(mp [][]mapa, numero int, i int, j int) {
	switch numero {
	case 1:
		fmt.Println("→ ")
		mp[i][j]
	case 2:
		fmt.Println("↑ ")
	case 3:
		fmt.Println("↓ ")
	case 4:
		fmt.Println("← ")
	}
}

func buscar(mp [][]mapa) { // funcion para buscar donde comer
	print("\n", len(mp))

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp); j++ {
			if mp[i][j].guzanito.cabeza == true {
				print("\n gusano cabeza en [", i, "],[", j, "]\n")
				if j == 0 { //priera columna eje J
					if i == 0 {
						if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")

						} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
						}
					} else if 0 < i && i < len(mp)-1 {
						if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
						} else if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
						}
					} else if i == len(mp)-1 {
						if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
						}

					}
				}
				if j == len(mp)-1 { // ultima columna eje J
					if i == 0 {
						if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
						} else if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
						}
					} else if i > 0 && i < len(mp)-1 {
						if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
						} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
						}

					} else if i == len(mp)-1 {
						if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
						} else if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
						}
					}
				}
				if i == 0 { // primera fila eje i
					if j > 0 && j < len(mp)-1 {
						if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("←")
						} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓")
						} else if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→")
						}
					}
				}
				if i == len(mp)-1 { //ultima fila eje i
					if j < 0 && j < len(mp)-1 {
						if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
						} else if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")
						}
					}
				}
				if j > 0 && j < len(mp)-1 && i > 0 && i < len(mp)-1 { // dentro de la matriz

					if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 { // come al lado izq
						print("→ ", mp[i][j+1].numero)

					} else if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {

						print("← ", mp[i][j-1].numero)

					} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {

						print("↓ ", mp[i+1][j].numero)

					} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 { //arriba

						print("↑ ", mp[i-1][j].numero)

					} else {
						print("no puede comer ")
					}
				}
			}
		}
	}
}
func main() {

	gusanos, _ := strconv.Atoi(os.Args[1])
	x, _ := strconv.Atoi(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])
	comida, _ := strconv.Atoi(os.Args[4])

	print(gusanos, x, y, comida)

	mp := make([][]mapa, x)
	for i := 0; i < x; i++ {
		mp[i] = make([]mapa, y)
	}

	crear_map(mp, gusanos, comida)
	print("\n")
	imprimir(mp)
	buscar(mp)

}