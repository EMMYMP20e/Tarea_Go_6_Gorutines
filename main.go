package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	var input string
	var id int

	opc := 1
	cont := 0
	actuales := 0

	cPrint := make(chan int)
	cStop := make(chan int)

	for opc != 0 {
		opc = menu()
		switch opc {
		case 1:
			go Proceso(cont, cPrint, cStop)
			cPrint <- 0
			cont++
			actuales++
		case 2:
			if actuales > 0 {
				fmt.Scanln(&input)
				for i := 0; i < cont; i++ {
					cPrint <- 1
				}
				fmt.Scanln(&input)
				for i := 0; i < cont; i++ {
					cPrint <- 0
				}
			} else {
				fmt.Println("No hay procesos para mostrar")
				pausa()
			}
		case 3:
			fmt.Print("ID Proceso a eliminar: ")
			fmt.Scan(&id)
			for i := 0; i < actuales; i++ {
				cStop <- id
			}
			actuales--
			pausa()
		}
	}
}

func Proceso(id int, cPrint chan int, cStop chan int) {
	i := uint64(0)
	imprimir := false
	terminar := false
	for {
		select {
		case msg := <-cPrint:
			if msg == 1 {
				imprimir = true
			} else if msg == 0 {
				imprimir = false
			}
		case msg := <-cStop:
			if msg == id {
				terminar = true
			}
		default:

		}
		if imprimir == true {
			fmt.Printf("id %d: %d \n", id, i)
		}
		if terminar {
			break
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("Proceso ", id, " finalizado")
}

func pausa() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("[Presiona Enter para Continuar]")
	scanner.Scan()
}

func menu() int {
	var opc int
	limpiarPantalla()
	fmt.Println("1.- Agregar Proceso")
	fmt.Println("2.- Mostrar Proceso")
	fmt.Println("3.- Terminar Proceso")
	fmt.Println("0.- Salir")
	fmt.Print("Opcion: ")
	fmt.Scan(&opc)
	limpiarPantalla()
	return opc
}

func limpiarPantalla() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
