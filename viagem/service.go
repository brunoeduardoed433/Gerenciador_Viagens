package viagem

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var viagemData ViagemData

func SalvarViagem(nomeViagem string) {

	viagem := Viagem{
		ID:      viagemData.ViagemUltimoID + 1,
		Destino: nomeViagem,
		Notas:   []Nota{},
	}

	novaViagemList := append(viagemData.Viagens, viagem)
	viagemData.Viagens = novaViagemList

	viagemData.ViagemUltimoID = viagemData.ViagemUltimoID + 1
	SalvarDados()
}

func EditarViagem(novoNomeViagem string, idViagem int) {
	for i := range viagemData.Viagens {
		if viagemData.Viagens[i].ID == idViagem {
			viagemData.Viagens[i].Destino = novoNomeViagem
			return
		}
		SalvarDados()
	}
}

func SalvarNota(notaViagem string, idViagem int) {
	for i := range viagemData.Viagens {
		if viagemData.Viagens[i].ID != idViagem {
			continue
		}

		novaNota := Nota{
			ID:          viagemData.NotaUltimoID + 1,
			Conteudo:    notaViagem,
			DataCriacao: time.Now(),
		}

		viagemData.Viagens[i].Notas = append(viagemData.Viagens[i].Notas, novaNota)
		viagemData.NotaUltimoID = viagemData.NotaUltimoID + 1
		SalvarDados()

		return
	}
	return
}

func EditarNota(novaNotaViagem string, idViagem int, idNota int) {

	for i := range viagemData.Viagens {
		if viagemData.Viagens[i].ID != idViagem {
			continue
		}
		for j := range viagemData.Viagens[i].Notas {
			if viagemData.Viagens[i].Notas[j].ID != idNota {
				continue
			}

			viagemData.Viagens[i].Notas[j].Conteudo = novaNotaViagem
			viagemData.Viagens[i].Notas[j].DataCriacao = time.Now()
		}
		SalvarDados()
	}
}

func DeletarViagem(viagemDeletada string, idViagem int) {
	var novaViagemList []Viagem

	for _, viagem := range viagemData.Viagens {
		if viagem.ID != idViagem {
			novaViagemList = append(novaViagemList, viagem)
			continue
		}
	}
	viagemData.Viagens = novaViagemList
	SalvarDados()
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

func VisaoGeral() {
	var viagemComNotas int

	totalViagens := len(viagemData.Viagens)
	fmt.Printf("\nTotal de Viagens: %d", totalViagens)

	for _, v := range viagemData.Viagens {
		if len(v.Notas) > 0 {
			viagemComNotas++
		}
	}

	fmt.Printf("\nTotal de Viagens que possui notas: %d", viagemComNotas)
	fmt.Println("")
	fmt.Println("Total de notas por viagem: \n")

	for _, v := range viagemData.Viagens {
		notasPorViagem := len(v.Notas)

		if len(v.Notas) == 0 {
			fmt.Printf("[ID: %d] : %s - Nenhuma nota | ", v.ID, v.Destino)
			continue
		}
		fmt.Printf("[ID: %d] : %s - %d notas | ", v.ID, v.Destino, notasPorViagem)
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
	fmt.Println("\n\nSalvando dados...")
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
