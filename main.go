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


func obtener_coordenada()
{
	var entrada_usuario int
	var coor int

	//usuario ingresará las coordenadas por teclado - INPUT CONSOLE
	for ok := true; ok; ok = res > 0 {

		fmt.Printf("Introduzca las coordenadas (x,y)")
		_, err: fmt.Scan(&entrada_usuario)

		Block{

			Try: func() {

				coor := entrada_usuario.Split("", ",") 

				if len(coor) != 2{
					fmt.Printf("Enrada Inválida, muy pocas / muchas coordenadas")
				}
				//checkea los valores de enteros  estan entre 1 y 10
				coor[0] > 9 || coor[0] < 0 || coor[1] > 9 || coor[1] < 1{
					fmt.Printf("Entrada invalida, Por favor use valores entre 1 a 10 unicamente")
				}

				return coor
	
			},
	
			Catch: func(e Exception) {
	
				fmt.Printf("Entrada Invalida. Por favor ingrese solamente valores numericos para las coordenadas", e)
	
			},
	
			Finally: func() {
	
				fmt.Printf("Finally...")
	
			},
	
		}.Do()

	}

}

//en esta funcion, leera las coordenadas el usuario atraves de una funcion ya definida
func usuario_movimiento(tablero string)
{
	var x int;
	var y int;
	res string;

	for ok := true; ok; ok = res > 0 {
		x, y = obtener_coordenada()
		res = realizar_movimiento(tablero, x, y)

		if res == "Bombazo"{
			fmt.Printf("Bombazo en " + strconv( x+1 ) )
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
			fmt.Printf("Sam bombazo en " + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1))
			verif_hundido(tablero, x, y)
			tablero[x][y] = '$'
			if verif_ganador(tablero) {
				return "ganador"
			}
		}else if res == "fallido" {
			fmt.Printf("Lo siento, " + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1)  + "no fue nada por el momento")
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
		fmt.Printf(barco + "Hundido uwu")
	}
}

func verif_ganador(tablero string){

	for i := 0; i < 9; i++ {
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
    //..
}