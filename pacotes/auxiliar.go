package pacotes

import "fmt"

/*
*
Como na linguagem Go não temos o conceito de classes e orientação a objetos, os pacotes são uma forma de agrupar funcionalidades relacionadas.
Com isso, se criamos uma função com a primeira letra maiúscula, essa função se torna pública e pode ser acessada por outros pacotes.
Exemplo:
package pacotes

import "fmt"

func WriteMessage() {} é uma função pública que pode ser acessada por outros pacotes
=============================================================================================================================================
Caso criamos uma função com a primeira letra minúscula, essa função se torna privada e só pode ser acessada dentro do próprio pacote.
Exemplo:
package pacotes

import "fmt"

func subWriterMessage() {} é uma função privada que só pode ser acessada dentro do pacote auxiliar
*/
func WriteMessage() {
	fmt.Println("Escrevendo uma mensagem do pacote auxiliar")
	subWriterMessage()
}
