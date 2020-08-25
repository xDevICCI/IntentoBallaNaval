package main

import (
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
	tamaño     int   //
	id         int   //
	horizontal bool  //
	comida     int   //cantidad de comida
	cabeza     bool  //comienza en false // true para decir que es cabeza
	fieldNext  *mapa //esta variable permitira ir restando ( comer) el valor de alimento y guardando
	cuerpo     bool  // false si es cabeza true si es cuerpo
	next       *gusano
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
					if j > 0 {
						mp[x+j][y].guzanito.cuerpo = true
					}
				}
				mp[x][y].guzanito.cabeza = true
			} else if mp[x][y+1].activo == false && mp[x][y+2].activo == false { //llene de forma horizontal
				for j := 0; j < 3; j++ {
					mp[x][y+j].guzanito = iniciarlizar_gusano()
					mp[x][y+j].guzanito.id = +i
					mp[x][y+j].activo = true
					if j > 0 {
						mp[x+j][y].guzanito.cuerpo = true
					}
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
				if j > 0 {
					mp[x+j][y].guzanito.cuerpo = true
				}
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
func buscar_cuerpo(mp [][]mapa, id int) (int, int, int, int) {
	var x []int
	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp); j++ {
			if mp[i][j].guzanito.id == id && mp[i][j].guzanito.cabeza == false { // si es cuerpo

				x = append(x, i)
				x = append(x, j)

			}
		}
	}
	return x[0], x[1], x[2], x[3]
}

func color(mp [][]mapa, i int, j int) { //funcion para pasar cambiar de color
	if mp[i][j].guzanito.id == 0 {
		print(ColorBlue, " ■ ")
	} else if mp[i][j].guzanito.id == 1 {
		print(ColorYellow, " ■ ")
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
		//J SE MUEVE HORIZONTAL
		//I SE MUEVE VERTICAL
		mp[i][j].guzanito.fieldNext = &mp[i][j+1]
		mp[i][j].guzanito.fieldNext.numero = mp[i][j].guzanito.fieldNext.numero - 1
		mp[i][j+1].numero = mp[i][j].guzanito.fieldNext.numero
		avanzar(mp, 1, i, j)
	case 2:
		mp[i][j].guzanito.fieldNext = &mp[i-1][j]
		mp[i][j].guzanito.fieldNext.numero = mp[i][j].guzanito.fieldNext.numero - 1
		mp[i-1][j].numero = mp[i][j].guzanito.fieldNext.numero
		avanzar(mp, 2, i, j)
	case 3:
		mp[i][j].guzanito.fieldNext = &mp[i+1][j]
		mp[i][j].guzanito.fieldNext.numero = mp[i][j].guzanito.fieldNext.numero - 1
		mp[i+1][j].numero = mp[i][j].guzanito.fieldNext.numero
		avanzar(mp, 3, i, j)
	case 4:
		mp[i][j].guzanito.fieldNext = &mp[i][j-1]
		mp[i][j].guzanito.fieldNext.numero = mp[i][j].guzanito.fieldNext.numero - 1
		mp[i][j-1].numero = mp[i][j].guzanito.fieldNext.numero
		avanzar(mp, 4, i, j)

	}
}

func avanzar(mp [][]mapa, numero int, i int, j int) {
	var aux mapa
	switch numero {
	case 1:
		//ESTE ES PARA LA DERECHA
		if mp[i][j].guzanito.fieldNext.numero == 0 {
			mp[i][j+1] = mp[i][j]
			mp[i][j] = aux

		}
	case 2:
		//ESTE ES PARA ARRIBA
		if mp[i][j].guzanito.fieldNext.numero == 0 {
			mp[i-1][j] = mp[i][j]
			mp[i][j] = aux

		}
	case 3:
		//ESTE ES PARA ABAJO
		if mp[i][j].guzanito.fieldNext.numero == 0 {
			mp[i+1][j] = mp[i][j]
			mp[i][j] = aux
		}
	case 4:
		//ESTE ES PARA LA IZQUIERDA
		if mp[i][j].guzanito.fieldNext.numero == 0 {
			mp[i][j-1] = mp[i][j]
			mp[i][j] = aux
		}
	}
}

func buscar(mp [][]mapa) { // funcion para buscar donde comer
	print("\n", len(mp))

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp); j++ {
			if mp[i][j].guzanito.cabeza == true {
				print("\n gusano cabeza en [", i, "],[", j, "]\n")
				if j == 0 { //primera columna eje J
					if i == 0 {
						if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")
							comer(mp, 1, i, j)

						} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
							comer(mp, 3, i, j)
						}
					} else if 0 < i && i < len(mp)-1 {
						if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
							comer(mp, 3, i, j)
						} else if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")
							comer(mp, 1, i, j)
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
							comer(mp, 2, i, j)
						}
					} else if i == len(mp)-1 {
						if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")
							comer(mp, 1, i, j)
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
							comer(mp, 2, i, j)
						}

					}
				}
				if j == len(mp)-1 { // ultima columna eje J
					if i == 0 {
						if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
							comer(mp, 3, i, j)
						} else if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
							comer(mp, 4, i, j)
						}
					} else if i > 0 && i < len(mp)-1 {
						if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
							comer(mp, 4, i, j)
						} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓ ")
							comer(mp, 3, i, j)
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
							comer(mp, 2, i, j)
						}

					} else if i == len(mp)-1 {
						if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
							comer(mp, 2, i, j)
						} else if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
							comer(mp, 4, i, j)
						}
					}
				}
				if i == 0 { // primera fila eje i
					if j > 0 && j < len(mp)-1 {
						if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("←")
							comer(mp, 4, i, j)
						} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {
							print("↓")
							comer(mp, 3, i, j)
						} else if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→")
							comer(mp, 1, i, j)
						}
					}
				}
				if i == len(mp)-1 { //ultima fila eje i
					if j < 0 && j < len(mp)-1 {
						if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {
							print("← ")
							comer(mp, 4, i, j)
						} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 {
							print("↑ ")
							comer(mp, 2, i, j)
						} else if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 {
							print("→ ")
							comer(mp, 1, i, j)
						}
					}
				}
				if j > 0 && j < len(mp)-1 && i > 0 && i < len(mp)-1 { // dentro de la matriz

					if mp[i][j+1].activo == false && mp[i][j+1].numero != 0 { // come al lado izq
						print("→ ", mp[i][j+1].numero)
						comer(mp, 1, i, j)

					} else if mp[i][j-1].activo == false && mp[i][j-1].numero != 0 {

						print("← ", mp[i][j-1].numero)
						comer(mp, 4, i, j)

					} else if mp[i+1][j].activo == false && mp[i+1][j].numero != 0 {

						print("↓ ", mp[i+1][j].numero)
						comer(mp, 3, i, j)

					} else if mp[i-1][j].activo == false && mp[i-1][j].numero != 0 { //arriba

						print("↑ ", mp[i-1][j].numero)
						comer(mp, 2, i, j)

					} else {
						print("no puede comer ")
					}
				} else {
					print("no hay mas camino")
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
	for {
		imprimir(mp)
		buscar(mp)
		time.Sleep(2 * time.Second)
		if 1 == 0 {
			break
		}
	}

}
