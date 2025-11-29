/*
Pacote variaveis demonstra o uso de variáveis em Go.

Existem duas maneiras principais de declarar variáveis em Go:
 1. Usando a palavra-chave 'var':
    var nomeDaVariavel tipoDaVariavel = valorInicial
 2. Usando a sintaxe curta ':=' (apenas dentro de funções):
    nomeDaVariavel := valorInicial

Go é uma linguagem de tipagem estática, o que significa que o tipo de uma variável é conhecido em tempo de compilação.
No entanto, quando você usa a sintaxe curta ':=', o compilador infere o tipo da variável com base no valor atribuído.

As variáveis declaradas fora de funções (variáveis de pacote) são inicializadas com valores zero padrão do seu tipo,
enquanto as variáveis declaradas dentro de funções (variáveis locais) devem ser inicializadas explicitamente.

Este pacote contém uma função 'Variables' que demonstra a declaração e uso de variáveis em Go.
*/
package variaveis

import "fmt"

func Variables() {
	var variavel1 string = "Olá, Mundo!" // tipo atribuído explicitamente
	variavel2 := 42                      // tipo inferido pelo compilador (implicitamente int)

	// Declaração de múltiplas variáveis
	var (
		variavel3 float64 = 3.14
		variavel4 bool    = true
		variaveis         = "Múltiplas variáveis"
	)

	// Declaração de constante (imutável) Não pode ser usada a sintaxe ':='
	const constante1 string = "Eu sou uma constante"

	// invertendo o valor das variáveis
	var variavel5 = "variavel5"
	var variavel6 = "variavel6"

	variavel5, variavel6 = variavel6, variavel5 // Essa tambem é uma forma de declarar múltiplas variáveis
	fmt.Println(variavel5, variavel6)

	// Imprimindo os valores das variáveis
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4, variaveis)
	fmt.Println(constante1)
	fmt.Println("Função Variables do pacote variaveis")
}
