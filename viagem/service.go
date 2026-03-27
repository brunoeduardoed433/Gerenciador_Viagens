package viagem

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var viagemData ViagemData

func CadastrarViagem() {

	// Cria entrada de dados como buffer para salvar os dados temporariamente até utiliza-los.
	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("\nDigite o destino: ")
	// Ele pega toda a string até encontrar um \n e atribui a variavel destino. E o "_" ignora uma outra variavel que retorna Erro se houver.
	destino, _ := leitor.ReadString('\n')
	// Ele remove espaços e/ou \n no começo e final da string
	destino = strings.TrimSpace(destino)

	viagem := Viagem{ID: viagemData.ViagemUltimoID + 1, Destino: destino, Notas: []Nota{}}
	novaViagemList := append(viagemData.Viagens, viagem)
	viagemData.Viagens = novaViagemList

	viagemData.ViagemUltimoID = viagemData.ViagemUltimoID + 1

	SalvarDados()
}

func AdicionarNota() {
	var buscaID int
	var comentario string

	fmt.Println("Digite o ID da viagem que deseja: ")
	fmt.Scanln(&buscaID)

	viagemEncontrada := false
	for i := range viagemData.Viagens {

		if viagemData.Viagens[i].ID != buscaID {
			continue
		}

		leitor := bufio.NewReader(os.Stdin)
		fmt.Println("Adicione sua nota da viagem: ")
		comentario, _ = leitor.ReadString('\n')
		comentario = strings.TrimSpace(comentario)

		novaNota := Nota{
			ID:          viagemData.NotaUltimoID + 1,
			Conteudo:    comentario,
			DataCriacao: time.Now(),
		}

		viagemData.Viagens[i].Notas = append(viagemData.Viagens[i].Notas, novaNota)
		viagemData.NotaUltimoID = viagemData.NotaUltimoID + 1
		SalvarDados()
		viagemEncontrada = true

		fmt.Println("Nota adicionada com sucesso.")

		return
	}

	if viagemEncontrada == false {
		fmt.Println("ID inválido! Tente novamente")
	}
}

func ListarTudo() {

	if len(viagemData.Viagens) == 0 {
		fmt.Println("Nenhuma viagem cadastrada.")
		return
	}

	for _, lista := range viagemData.Viagens {
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

	for i := range viagemData.Viagens {
		if viagemData.Viagens[i].ID != pegaID {
			continue
		}

		leitor := bufio.NewReader(os.Stdin)
		fmt.Println("\nDigite o destino correto: ")
		novoDestino, _ := leitor.ReadString('\n')
		novoDestino = strings.TrimSpace(novoDestino)
		viagemEncontrada = true

		viagemData.Viagens[i] = Viagem{
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

func EditarNota() {
	var pegaID int

	fmt.Println("Digite o ID da nota desejada: ")
	fmt.Scanln(&pegaID)

	notaEncontrada := false

	for _, v := range viagemData.Viagens {
		if notaEncontrada {
			break
		}
		fmt.Printf("Viagem iterada %v\n", v.ID)

		for i, n := range v.Notas {
			if n.ID != pegaID {
				continue
			}

			leitor := bufio.NewReader(os.Stdin)
			fmt.Println("\nDigite a Nota correta para esta viagem : ")
			novaNota, _ := leitor.ReadString('\n')
			novaNota = strings.TrimSpace(novaNota)

			v.Notas[i] = Nota{
				ID:          pegaID,
				Conteudo:    novaNota,
				DataCriacao: time.Now(),
			}

			notaEncontrada = true
			fmt.Printf("Nota encontrada %v\n", n.ID)
			break
		}

	}

	if notaEncontrada {
		SalvarDados()
		fmt.Println("Viagem editada com sucesso!")
	} else {
		fmt.Println("ID inválido! Tente novamente.")
	}
}

func DeletarCidade() {
	var buscaID int
	var novaViagemList []Viagem
	cidadeEncontrada := false

	fmt.Println("Digite o ID da viagem que deseja deletar: ")
	fmt.Scanln(&buscaID)

	for _, viagem := range viagemData.Viagens {
		if viagem.ID != buscaID {
			novaViagemList = append(novaViagemList, viagem)
			continue
		}
		cidadeEncontrada = true
	}

	viagemData.Viagens = novaViagemList

	if cidadeEncontrada {
		fmt.Print("\nCidade removida com sucesso\n")
		SalvarDados()
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

func VisaoGeral() error {
	var viagemComNotas int
	//var NotasPorCidade int

	totalViagens := len(viagemData.Viagens)
	fmt.Printf("\nTotal de Viagens: %d", totalViagens)

	for _, v := range viagemData.Viagens {
		if len(v.Notas) > 0 {
			viagemComNotas++
		}

		// for i := range viagemData.Viagens{
		// 	if v.notas[i] == nil{

		// 	}
		// 	NotasPorCidade = 
		// 	fmt.Printf("\nQuantidade de notas por viagem: [ID: %d] : %s - %s notas", v.ID, v.Destino, v.Notas)
		// }
	}
	fmt.Printf("\nTotal de Viagens que possui notas: %d\n", viagemComNotas)

	SalvarDados()
	return nil
}
