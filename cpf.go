package cpf

import (
	"github.com/mvrilo/go-cpf"
    "go.k6.io/k6/js/modules"
)

func init() {
    modules.Register("k6/x/cpf", new(CPF))
}

type CPF struct{}

func (*CPF) Cpf(formatado bool) string {
	if formatado {
		return cpf.GeneratePretty()
	}	
	return cpf.Generate()
	
}