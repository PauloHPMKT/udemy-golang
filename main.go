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

	"github.com/PauloHPMKT/udemy-golang/funcoes"
	"github.com/PauloHPMKT/udemy-golang/pacotes"
	tiposdedados "github.com/PauloHPMKT/udemy-golang/tiposDeDados"
	"github.com/PauloHPMKT/udemy-golang/variaveis"

	/**
	Importando um pacote externo (terceiros)
	Para importar pacotes externos, você precisa primeiro instalá-los usando o comando "go get <url-do-pacote>". Ex:
	go get github.com/badoux/checkmail

	Depois de instalado, você pode importar o pacote no seu código e utilizá-lo normalmente.
	ou utilizar o comando "go mod tidy" que irá baixar todas as dependências declaradas no arquivo go.mod e go.sum
	A utilização desse comando é recomendada sempre que você adicionar ou remover pacotes do seu projeto.
	*/
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	pacotes.WriteMessage()

	// Utilizando o pacote externo checkmail para validar o formato de um e-mail
	error := checkmail.ValidateFormat("teste@mail.com")
	if error != nil {
		fmt.Println("E-mail inválido")
	}

	variaveis.Variables()
	tiposdedados.Numeros()
	tiposdedados.Chars()
	tiposdedados.Booleanos()
	tiposdedados.ErrorTypes()

	funcoes.Execute()
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
