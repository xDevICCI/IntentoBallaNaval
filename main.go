package main

import "fmt"
import "math/rand"
import "time"
import "strconv"
import "io"
import "io/ioutil"
import "log"
import "os"



type Block struct {
    Try     func()
    Catch   func(Exception)
    Finally func()
}

 
type Exception interface{}
 

func Throw(up Exception) {

    panic(up)

}

type batalla struct{
	tablero string
	barcos int
}

 
func (tcf Block) Do() {

    if tcf.Finally != nil {
        defer tcf.Finally()

    }
    if tcf.Catch != nil {
        defer func() {
            if r := recover(); r != nil {
                tcf.Catch(r)
            }
        }()
    }
    tcf.Try()
}

func usuario_lugar_barcos(tablero string, barcos int) string{
	//permite al usuario colocar barcos y también verificar si son posiciones válidas
	
	var validado bool


	for barco := 0; i < barcos; i++ {
		validado = false
	}

}

func AI_lugar_barcos(tablero string, barcos int) string{
	//la computadora usará al azar para generar lugares de envío

	var res int
	var validado bool

	var x int
	var y int
	var o int
	var barco int

	for i := 0; i < barcos; i++ {
		validado = false
		for ok := true; ok; ok = res > 0 {
			 x  = rand.Intn(10)
			 y 	= rand.Intn(10)
			 o 	= rand.Intn(0, 1)

			 if o == 0 {
				 orientacion = "v"
			 }else{
				 orientacion = "h"
			 }

			 validado = validar(tablero, barcos[barco], x, y, orientacion)
			 //ubicar los barcos
			 fmt.Println("Maquina AI ubicando en " + barco)
			 tablero = lugar_barco(tablero, barcos[barco], barco[0], orientacion, x, y)
		}
	}

	return tablero


}


func obtener_coordenada(){
	var entrada_usuario int
	var coor int

	//usuario ingresará las coordenadas por teclado - INPUT CONSOLE
	for ok := true; ok; ok = res > 0 { //while(true)

		//aca debe recibir
		fmt.Println("Introduzca las coordenadas (x,y)")
		_, err: fmt.Scan(&entrada_usuario)

		Block{

			Try: func() {

				coor := entrada_usuario.Split(" ", ",") 

				if len(coor) != 2{
					fmt.Println("Enrada Inválida, muy pocas / muchas coordenadas")
				}
				//checkea los valores de enteros  estan entre 1 y 10
				coor[0] > 9 || coor[0] < 0 || coor[1] > 9 || coor[1] < 1
				{
					fmt.Println("Entrada invalida, Por favor use valores entre 1 a 10 unicamente")
				}

				return coor
	
			},
	
			Catch: func(e Exception) {
	
				fmt.Println("Entrada Invalida. Por favor ingrese solamente valores numericos para las coordenadas", e)
	
			},
	
			Finally: func() {
	
				fmt.Println("Finally...")
	
			},
	
		}.Do()

	}

}

/*  recibe tablero, tamaño barco, y posicion, lugar barcos */
func lugar_barco(tablero string, barco int, s int, orientacion string, x int, y int) string {

	if orientacion == "v" {
		for i := 0; i < barco; i++ {
			tablero[x+i][y] = s
		}
	}else if orientacion == "h" {
		for i := 0; i < barco; i++ {
			tablero[x][y+i] = s
		}
	}
	return tablero
}

// comprueba si el barco realmente encajará ( en las celdas tablero)
func validar(tablero string, barco int, x int, y int, orientacion string) bool {
	//verifique si el barco encajará, según el tamaño del barco, el tablero, la orientación y las coordenadas

	if orientacion == "v"  && x  + barco > 10{
		return false
	}else if orientacion == "h" && y + barco > 10 {
		return false
	}else{
		if orientacion == "v" {
			for i := 0; i < barco; i++ {
				if tablero[x+i][y] != -1 {
					return false
				}
			}
		}else if orientacion == "h" {
			for i := 0; i < barco; i++ {
				if tablero[x][y+i] != -1 {
					return false
				}
			}
		}
	}
}

//en esta funcion mira si el barco es horizontal o vertical
func v_o_h(entrada_usuario int){

	//obtiene orientacion del barco

	for ok := true; ok; ok = res > 0 {
		fmt.Println("Vertical o Horizontal (v,h)")
		_, err: fmt.Scan(&entrada_usuario)

		if entrada_usuario == "v" || entrada_usuario == "h" {
			return entrada_usuario
		}else{
			fmt.Println("Entrada invalida, por favor ingrese unicamente v o h")
		}
	}
}

//en esta funcion, leera las coordenadas el usuario atraves de una funcion ya definida
func usuario_movimiento(tablero string) {
	var x int
	var y int
	res string

	for ok := true; ok; ok = res > 0 {
		x, y = obtener_coordenada()
		res = realizar_movimiento(tablero, x, y)

		if res == "Bombazo"{
			fmt.Println("Bombazo en " + strconv( x+1 )  + "," + strconv( y+1 ) )
			verif_hundido(tablero, x, y)
			tablero[x][y] = '$'

			if verif_ganador(tablero){
				return "Ganador"
			}
		}else if res == "Fallido" {
			fmt.Println("Lo siento, " , + strconv( x + 1 ) + "," + strconv( y+1 ) + "es un Fallo")
			tablero[x][y] = "*"
		}
		if res != "Intente nuevamente" {
			return tablero
		}
	}
}



func AI_movimiento(tablero string){
	var x int
	var y int
	var res string

	for ok := true; ok; ok = res > 0 {
		x = rand.Intn(10)
		y = rand.Intn(10)

		res = realizar_movimiento(tablero, x, y)

		if res == "bombazo" {
			fmt.Println("Sam bombazo en " + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1))
			verif_hundido(tablero, x, y)
			tablero[x][y] = '$'
			if verif_ganador(tablero) {
				return "ganador"
			}
		}else if res == "fallido" {
			fmt.Println("Lo siento, " + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1)  + "no fue nada por el momento")
			tablero[x][y] = "*"
		}
		if res != "intente otra vez" {
			return tablero
		}



	}
}



func realizar_movimiento(tablero string, x int, y int) int {
	if tablero[x][y] == -1 {
		return "Oh, No! Falle uwu"
	}else if tablero[x][y] == '*' || tablero[x][y] == '$' {
		return "intente de nuevo"
	}else{
		return "Boom, HeadShot"
	}
}

//verifica que barco es alcanzado
//verifica cuantos puntos aun existen en el barco
//barco se hunde si no existen mas puntos
func verif_hundido(tablero string, x string, y string){
	if tablero[i][j] == "P" {
		barco = "Portaaviones"
	}else if tablero[i][j] == "A"{
		barco = "Acorazado"
	}else if tablero[i][j] == "S"{
		barco = "Submarino"
	}else if tablero[i][j] == "D"{
		barco = "Destructor"
	}else if tablero[i][j] == "U"{
		barco = "Uwu"
	}

	/* marca la celda como un daño y verifica si esta hundido el barco*/
	tablero[-1][barco] -=1
	if tablero[-1][barco] == 0 {
		fmt.Println(barco + "Hundido uwu")
	}
}

func verif_ganador(tablero string) bool {

	for i := 0; i < 9; i++ {//aqui falta agregar el argv
		for j := 0; j < 9; j++ {
			if tablero[i][j] != -1 && tablero[i][j] != '*' && tablero[i][j] != '$'{
				//si no es un acierto, devuelve falso
				return false
			}else{
				return true
			}
		}
	}
}

func main() {	

	cant_jugadores, err := strconv.Atoi(os.Args[1])

	cant_barcos, err2 := strconv.Atoi(os.Args[2])

	x, err := strconv.Atoi(os.Args[3])
	y, err1 := strconv.Atoi(os.Args[4])


	var tablero[] string
	var fila int

	var usuario_tablero string
	var AI_tablero string

	//supuestamente seteamos o llenamos la matriz
	for i := 0; i < x; i++ {
		fila = []
		for j := 0; j < y; j++ {
			fila = append(fila, -1)
		}
		tablero = append(tablero,fila)
	}

	usuario_tablero = deepcopy.copy(tablero)
	AI_tablero = deepcopy.copy(tablero)

	//debemos agregar los barcos o cantidad de barcos
	usuario_tablero = append(usuario_tablero, deepcopy.copy(barcos))
	AI_tablero = append(AI_tablero, deepcopy.copy(barcos))

	//POSICIONAMOS LOS BARCOS PARA CADA TABLERO
	usuario_tablero = usuario_lugar_barcos(usuario_tablero, barcos)
	AI_tablero = AI_lugar_barcos(AI_tablero, barcos)




}