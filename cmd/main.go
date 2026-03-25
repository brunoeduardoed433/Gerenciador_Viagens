package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com.br/brunoeduardoed433/gerenciador_viagens/viagem"
)

func main() {

	err := viagem.CarregarLista()

	if err != nil {
		slog.Error("Erro ao carregar lista", "erro", err, "data", time.Now())
		panic(err)
	}

	iniciarMenu()

}

func iniciarMenu() {
	for {
		var opcao int
		fmt.Print("\n" + `OPÇÕES: ` + "\n\n" + `1. Cadastrar Viagem` + "\n" + `2. Adicionar Nota a uma Viagem` + "\n" + `3. Listar Tudo` + "\n" + "4. Editar Viagem" + "\n" + "5. Editar Nota" + "\n" + `6. Deletar Viagem` + "\n" + `7. Sair` + "\n")
		fmt.Print("\nDigite a opção desejada: ")
		fmt.Scanln(&opcao)
		switch opcao {
		case 1:
			viagem.CadastrarViagem()
		case 2:
			viagem.AdicionarNota()
		case 3:
			viagem.ListarTudo()
		case 4:
			viagem.EditarCidade()
		case 5:
			//EditarNota()
			fmt.Println("Em desenvolvimento")
		case 6:
			viagem.DeletarCidade()
		case 7:
			fmt.Print("Saindo...")
			return
		default:
			fmt.Println("Opção Inválida, tente novamente: ")
		}
	}
}
