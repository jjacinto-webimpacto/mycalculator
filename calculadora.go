package mycalculator

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int  {
	entradaLimpia:=strings.Split(entrada,operador);
	operador1:=parsear(entradaLimpia[0])
	operador2:=parsear(entradaLimpia[1])
	switch operador{
		case "+" :
			fmt.Println(operador1+operador2)		//realizamos la operacion aritmetica -> suma
			return operador1+operador2
		case "-" :
			fmt.Println(operador1-operador2)		//realizamos la operacion aritmetica -> resta			
			return operador1-operador2
		case "*" :
			fmt.Println(operador1*operador2)		//realizamos la operacion aritmetica -> multiplicar			
			return operador1*operador2
		case "/" :
			fmt.Println(operador1/operador2)		//realizamos la operacion aritmetica -> division	
			return operador1/operador2
		default:
			fmt.Println(operador," No esta soportado ")		//realizamos cuando ninguna validdacion se respeta
			return 0
	}	
}

func parsear(entrada string) int{
	operador,_:=strconv.Atoi(entrada)	//convertimos de string a numero
	return operador
}

func leerEntrada()string{
	scanner:=bufio.NewScanner(os.Stdin) //bufio.NewScanne libreria para el ingreso de datos
	scanner.Scan()						//permite obtener el dato ingresao
	return scanner.Text()			//leer el dato y pasarlo a la variable	
}
