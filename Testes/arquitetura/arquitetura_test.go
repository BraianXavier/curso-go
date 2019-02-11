package arquitetura

import (
	"runtime"
	"testing"
)

//A arquitetura da maquina é nada mais que o modelo do seu processador!

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	//...
	t.Fail() //encontrar um erro caso a maquina não utilize esse tipo de arquitetura
}
