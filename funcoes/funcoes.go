/*
Funções em Go podem retornar um ou múltiplos valores.
Elas também podem retornar valores de tipos diferentes, incluindo tipos personalizados e erros.

Funções que retornam múltiplos valores são úteis para retornar resultados e erros juntos,
permitindo um tratamento de erros mais claro e conciso.

Por convenção, em Go, quando uma função retorna um erro, ele é sempre o último valor retornado.
Se a função não encontrar nenhum erro, ela deve retornar nil como valor de erro.

*/

package funcoes

import "fmt"

func Execute() {
	resultado := soma(10, 20)
	fmt.Println("Resultado da soma:", resultado)

	/*
		Chamando a função que retorna múltiplos valores
		Aqui, estamos chamando a função calculosMatematicos que retorna dois valores inteiros.
		Utilizamos a sintaxe de atribuição múltipla para capturar ambos os valores retornados
		em variáveis separadas: soma e substracao.

		Caso você não queira utilizar um dos valores retornados, você pode usar o
		underscore (_) para ignorár o resultado que voce deseja.

		UM PONTO DE ATENÇÂO É A ORDEM IMPORTA NA ATRIBUIÇÃO DOS VALORES RETORNADOS.
		Ex1: _, substracao := calculosMatematicos(50, 30)
		Ex2: soma, _ := calculosMatematicos(50, 30)
	*/
	soma, substracao := calculosMatematicos(50, 30)
	fmt.Println("Resultados dos cálculos:", soma, substracao)

	resultadoMultiplicacao, err := multiplicacao(10, 0)
	if err != nil {
		fmt.Println("Erro na multiplicação:", err)
	}

	fmt.Println("Resultado da multiplicação:", resultadoMultiplicacao)

	/*
		Atribuindo uma função anônima a uma variável - type function literal
		Em Go, você pode definir funções anônimas (funções sem nome) e atribuí-las a variáveis.
		Isso permite que você crie funções de forma dinâmica e as utilize como valores de primeira classe.
	*/
	subtracao := func(a, b int) int {
		return a - b
	}

	// nesse caso da passagem de parâmetros, os tipos dos parâmetros devem ser declarados
	// Com isso a assinatura da tipagem fica "func (int, int) int"
	resultadoSubtracao := subtracao(30, 15)
	fmt.Println("Resultado da subtração:", resultadoSubtracao)
}

// tambem pode ser declarado como: func soma(a int, b int) int {}
// Essa abordagem é útil quando os parâmetros são do mesmo tipo.
func soma(a, b int) int {
	return a + b
}

/*
Função que retorna múltiplos valores

Nesse exemplo, a função calculosMatematicos retorna dois valores inteiros:
a soma e a subtração dos dois números fornecidos como argumentos.
*/
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	substracao := n1 - n2

	return soma, substracao
}

/*
Função que retorna um resultado e um erro

Por convenção em Go, quando uma função retorna um erro, ele deve ser o último valor retornado.
Se não houver erro, retornamos nil.
*/
func multiplicacao(a, b int) (int, error) {
	// Exemplo de validação que poderia gerar um erro
	calcCondition := a == 0 || b == 0
	if calcCondition {
		return 0, fmt.Errorf("não é permitido multiplicar por zero")
	}

	resultado := a * b
	return resultado, nil
}
