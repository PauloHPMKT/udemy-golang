/**
Um pacote em Go é um grupo de arquivos que obrigatoriamente estão no mesmo diretório e compartilham o mesmo nome de pacote
e eles são compilados juntos. Pacotes ajudam a organizar o código, facilitando a reutilização e a manutenção.

Por exemplo, funções, constantes e variaveis que forem declaradas em um arquivo de um pacote, podem ser utilizadas em outros arquivos
do mesmo pacote sem a necessidade de importação.

Cada programa Go começa com um pacote chamado "main". O pacote main é especial porque define um programa executável,
enquanto outros pacotes são bibliotecas que podem ser importadas por outros programas.

No caso do pacote main ele é obrigatório ter uma função main(), que é o ponto (Entrypoint) de entrada do programa.
*/

package main

/* a lista de imports é onde você declara os pacotes que serão utilizados no seu arquivo go */
import (
	"fmt"

	"github.com/PauloHPMKT/udemy-golang/pacotes"
)

func main() {
	fmt.Println("Hello, World!")
	pacotes.WriteMessage()
}

/**
IMPORTANTE:

Quando estamos trabalhando com mais de um pacote em um projeto Go, nós precisamos de um gerenciador de dependências
chamado módulo Go (Go Modules). O módulo Go é uma coleção de pacotes Go versionados juntos como uma unidade.

O Módulo nada mais é que um conjunto de pacotes que compoem um projeto. Ele é definido por um arquivo chamado
go.mod que fica na raiz do projeto.

Quando se esta usando o modulo Go é possível utilizar outra maneira de executar o programa que é usando o comando:
go build <nome-do-arquivo.go> que irá compilar o arquivo e gerar um executável com o mesmo nome do modulo.
Ex: go build main.go -> irá gerar um executável chamado "modulo-go" (nome do modulo definido no go.mod)
*/
