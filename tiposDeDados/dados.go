package tiposdedados

import (
	"errors"
	"fmt"
)

func Numeros() {
	/*
		Numeros inteiros (int, int8, int16, int32, int64)
		Respectivamente são representados por 1, 8, 16, 32 e 64 bits.
		O tipo int é o tipo inteiro padrão da linguagem e sua representação depende da arquitetura do sistema (32 ou 64 bits).

		Ex: int8 = 100000000000 isso dá um overflow, pois o valor máximo que um int8 pode armazenar é 127 (2^7 - 1)
	*/
	var numero int16 = 100
	fmt.Println(numero)

	var inteiro int = 200 // Quando não especificado, o tipo int utiliza por padrão a arquitetura do sistema	ou seja 32 ou 64 bits
	fmt.Println(inteiro)

	// Caso para um numero int não ocorra inferencia de tipo, o Go assume que é int mesmo sem especificar o tipo
	var numero2 = 300 // int de 32 ou 64 bits dependendo da arquitetura do sistema
	fmt.Println(numero2)

	/*
		Por padrão o tipo int suporta números negativos e positivos,
		porém para utilizar apenas números positivos, podemos utilizar o tipo uint (unsigned int)
		Caso se utilize um número negativo em uma variável do tipo uint, ocorrerá um erro de compilação
		uint segue a mesma convenção de bits do int (uint8, uint16, uint32, uint64)
	*/
	var numero3 uint = 400
	fmt.Println(numero3)

	/*
		Numeros de ponto flutuante (float32, float64)
		Representam números reais com parte decimal. O float32 utiliza 32 bits e o float64 utiliza 64 bits.
		O float64 é o tipo de ponto flutuante padrão em Go.
	*/
	var numeroFloat32 float32 = 10.5
	fmt.Println(numeroFloat32)

	var numeroFloat float64 = 10000000000000.5
	fmt.Println(numeroFloat)

	/*
		com inferência de tipo, o Go assume que é float, dependendo do tamanho do número ele
		pode assumir float32 ou float64 de acordo com a arquitetura do sistema
	*/
	numeroFloat2 := 20.99
	fmt.Println(numeroFloat2)
}

func Chars() {
	var char rune = 'A' // O tipo rune é um alias para int32 e representa um caractere Unicode e não um caractere string
	fmt.Println(char)

	var byteChar byte = 'B' // O tipo byte é um alias para uint8 e representa um caractere ASCII
	fmt.Println(byteChar)
}

func Booleanos() {
	// O tipo booleano representa valores verdadeiros ou falsos. Note que o valor padrão é false
	var variavelBooleana bool
	fmt.Println(variavelBooleana)

	var booleano bool = true
	fmt.Println(booleano)

	booleano2 := false
	fmt.Println(booleano2)
}

func ErrorTypes() {
	var erro error = nil // O tipo error é uma interface embutida em Go que representa um erro. O valor padrão é nil
	fmt.Println(erro)

	// Para a criação de erros personalizados, podemos utilizar a função errors.New() do pacote "errors"
	var customError error = errors.New("Este é um erro personalizado")
	fmt.Println(customError)
}

/*
Para todos os tipos de dados em Go existem valores inicializados padrão (zero values).
Quando uma variável é declarada sem um valor inicial, ela recebe automaticamente um valor padrão baseado no seu tipo de dado.

Os valores padrão para os tipos de dados mais comuns em Go são:

- Números inteiros (int, int8, int16, int32, int64): 0
- Números de ponto flutuante (float32, float64): 0.0
- Booleanos (bool): false
- Strings (string): "" (string vazia)
- Erros (error): nil
*/
