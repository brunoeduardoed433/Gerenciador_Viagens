package viagem

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var proximoIdNota = 1
var viagemData ViagemData

func CadastrarViagem() {

	// Cria entrada de dados como buffer para salvar os dados temporariamente até utiliza-los.
	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("\nDigite o destino: ")
	// Ele pega toda a string até encontrar um \n e atribui a variavel destino. E o "_" ignora uma outra variavel que retorna Erro se houver.
	destino, _ := leitor.ReadString('\n')
	// Ele remove espaços e/ou \n no começo e final da string
	destino = strings.TrimSpace(destino)

	viagem := Viagem{ID: viagemData.UltimoID + 1, Destino: destino, Notas: []Nota{}}
	novaViagemList := append(viagemData.Data, viagem)
	viagemData.Data = novaViagemList

	viagemData.UltimoID = viagemData.UltimoID + 1

	SalvarDados()
}

func AdicionarNota() {
	var buscaID int
	var comentario string

	fmt.Println("Digite o ID da viagem que deseja: ")
	fmt.Scanln(&buscaID)

	viagemEncontrada := false
	for i := range viagemData.Data {

		if viagemData.Data[i].ID != buscaID {
			continue
		}

		leitor := bufio.NewReader(os.Stdin)
		fmt.Println("Adicione sua nota da viagem: ")
		comentario, _ = leitor.ReadString('\n')
		comentario = strings.TrimSpace(comentario)

		novaNota := Nota{
			ID:          proximoIdNota,
			Conteudo:    comentario,
			DataCriacao: time.Now(),
		}

		viagemData.Data[i].Notas = append(viagemData.Data[i].Notas, novaNota)
		fmt.Println("Nota adicionada com sucesso.")
		viagemEncontrada = true

		SalvarDados()

		proximoIdNota++

		return
	}

	if viagemEncontrada == false {
		fmt.Println("ID inválido! Tente novamente")
	}
}

func ListarTudo() {

	if len(viagemData.Data) == 0 {
		fmt.Println("Nenhuma viagem cadastrada.")
		return
	}

	for _, lista := range viagemData.Data {
		fmt.Printf("[ID: %d] Viagem para: %s\n\n", lista.ID, lista.Destino)
		for _, lista := range lista.Notas {
			fmt.Printf("   - [ID: %d] Nota:  %s (%02d/%02d/%04d)\n", lista.ID, lista.Conteudo, lista.DataCriacao.Day(), lista.DataCriacao.Month(), lista.DataCriacao.Year())
		}
		fmt.Println("----------------------------------------")
	}
}

func EditarCidade() {
	var pegaID int

	fmt.Println("Digite o ID da Viagem desejada: ")
	fmt.Scanln(&pegaID)
	viagemEncontrada := false

	for i := range viagemData.Data {
		if viagemData.Data[i].ID != pegaID {
			continue
		}

		leitor := bufio.NewReader(os.Stdin)
		fmt.Println("\nDigite o destino correto: ")
		novoDestino, _ := leitor.ReadString('\n')
		novoDestino = strings.TrimSpace(novoDestino)
		viagemEncontrada = true

		viagemData.Data[i] = Viagem{
			Destino: novoDestino,
			ID:      pegaID,
		}

		SalvarDados()

	}
	if viagemEncontrada {
		fmt.Println("Viagem editada com sucesso!")
	} else {
		fmt.Println("ID inválido! Tente novamente.")
	}
}

// func editarNota() {
// 	var pegaID int
// 	fmt.Println("Digite o ID da Viagem desejada: ")
// 	fmt.Scanln(&pegaID)
// 	notaEncontrada := false
// 	for _, nota := range lista {
// 		if nota. != pegaID {
// 			continue
// 		}
// 		leitor := bufio.NewReader(os.Stdin)
// 		fmt.Println("\nDigite a Nota correta para esta viagem : ")
// 		novoDestino, _ := leitor.ReadString('\n')
// 		novoDestino = strings.TrimSpace(novoDestino)
// 		notaEncontrada = true
// 		viagemData.Data[i] = Viagem{
// 			Destino: novoDestino,
// 			ID:      pegaID,
// 		}
// 	}
// 	if notaEncontrada {
// 		fmt.Println("Viagem editada com sucesso!")
// 	} else {
// 		fmt.Println("ID inválido! Tente novamente.")
// 	}
// }

func DeletarCidade() {
	var buscaID int
	var novaViagemList []Viagem
	cidadeEncontrada := false

	fmt.Println("Digite o ID da viagem que deseja deletar: ")
	fmt.Scanln(&buscaID)

	for _, viagem := range viagemData.Data {
		if viagem.ID != buscaID {
			novaViagemList = append(novaViagemList, viagem)
			continue
		}
		cidadeEncontrada = true
	}

	viagemData.Data = novaViagemList

	SalvarDados()

	if cidadeEncontrada {
		fmt.Print("\nCidade removida com sucesso\n")
	} else {
		fmt.Println("ID inválido! Tente novamente")
	}

}

func CarregarLista() error {

	_, err := os.Stat("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			arquivo, err := os.Create("data.json")

			if err != nil {
				return err
			}
			defer arquivo.Close()
		} else {
			return err
		}
	}

	conteudo, err := os.ReadFile("data.json")
	if err != nil {
		return err
	}

	if len(conteudo) > 0 {
		err = json.Unmarshal(conteudo, &viagemData)
		if err != nil {
			return err
		}
		return nil
	}
	viagemData = ViagemData{}
	SalvarDados()

	return nil
}

func SalvarDados() {

	bytesJson, err := json.Marshal(viagemData)
	fmt.Println("Salvando dados...")
	if err != nil {
		fmt.Println("Erro ao converter para Json: ", err)
		return
	}

	err = os.WriteFile("data.json", bytesJson, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo: ", err)
		return
	}
}
