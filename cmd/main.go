package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com.br/brunoeduardoed433/gerenciador_viagens/viagem"
)

func main() {

	err := viagem.CarregarLista()
	if err != nil {
		slog.Error("Erro ao carregar lista", "erro", err, "data", time.Now())
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /viagem/salvar", SalvarViagem)
	mux.HandleFunc("PUT /viagem/editar/{id}", EditarViagem)
	mux.HandleFunc("POST /viagem/{id}/nota/salvar", SalvarNota)
	mux.HandleFunc("PUT /viagem/{id}/nota/{idnota}/editar/salvar", EditarNota)
	mux.HandleFunc("DELETE /viagem/{id}/deletar/salvar", DeletarViagem)

	slog.Info("Servidor escutando em localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

// cadastrar viagem
func SalvarViagem(escritor http.ResponseWriter, requisicao *http.Request) {
	corpo, err := io.ReadAll(requisicao.Body)
	if err != nil {
		escritor.WriteHeader(http.StatusInternalServerError)
		escritor.Write([]byte("Erro ao ler corpo da requisição!"))
		return
	}

	var viagemRequest viagem.ViagemRequest
	err = json.Unmarshal(corpo, &viagemRequest)
	if err != nil {
		escritor.WriteHeader(http.StatusBadRequest)
		escritor.Write([]byte("Dados inválidos!"))
		return
	}

	viagem.SalvarViagem(viagemRequest.Destino)

	escritor.WriteHeader(http.StatusCreated)
	escritor.Write([]byte("Viagem salva com sucesso!"))
}

// EditarViagem
func EditarViagem(escritor http.ResponseWriter, requisicao *http.Request) {
	idStr := requisicao.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Printf("\nID da viagem editada: %d", id)

	corpo, err := io.ReadAll(requisicao.Body)
	if err != nil {
		escritor.WriteHeader(http.StatusInternalServerError)
		escritor.Write([]byte("Erro ao ler corpo da requisição!"))
		return
	}

	var viagemRequest viagem.ViagemRequest
	err = json.Unmarshal(corpo, &viagemRequest)
	if err != nil {
		escritor.WriteHeader(http.StatusBadRequest)
		escritor.Write([]byte("Dados inválidos!"))
		return
	}

	viagem.EditarViagem(viagemRequest.Destino, id)

	escritor.WriteHeader(http.StatusOK)
	escritor.Write([]byte("Viagem editada com sucesso!"))
}

// Cadastrar Nota
func SalvarNota(escritor http.ResponseWriter, requisicao *http.Request) {
	idStr := requisicao.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Printf("\nID da Viagem com nota: %d", id)

	corpo, err := io.ReadAll(requisicao.Body)
	if err != nil {
		escritor.WriteHeader(http.StatusInternalServerError)
		escritor.Write([]byte("Erro ao ler corpo da requisição!"))
		return
	}

	var notaRequest viagem.NotaRequest
	err = json.Unmarshal(corpo, &notaRequest)
	if err != nil {
		escritor.WriteHeader(http.StatusBadRequest)
		escritor.Write([]byte("Dados inválidos!"))
		return
	}

	viagem.SalvarNota(notaRequest.Nota, id)

	escritor.WriteHeader(http.StatusCreated)
	escritor.Write([]byte("Nota adicionada com sucesso!"))
}

// EditarNota
func EditarNota(escritor http.ResponseWriter, requisicao *http.Request) {
	idStr := requisicao.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Printf("\nID da Nota editada: %d", id)

	idStrNota := requisicao.PathValue("idnota")
	idNota, _ := strconv.Atoi(idStrNota)
	fmt.Printf("\nID da Nota editada: %d", idNota)

	corpo, err := io.ReadAll(requisicao.Body)
	if err != nil {
		escritor.WriteHeader(http.StatusInternalServerError)
		escritor.Write([]byte("Erro ao ler corpo da requisição!"))
		return
	}

	var notaRequest viagem.NotaRequest
	err = json.Unmarshal(corpo, &notaRequest)
	if err != nil {
		escritor.WriteHeader(http.StatusBadRequest)
		escritor.Write([]byte("Dados inválidos!"))
		return
	}

	viagem.EditarNota(notaRequest.Nota, id, idNota)

	escritor.WriteHeader(http.StatusOK)
	escritor.Write([]byte("Nota editada com sucesso!"))
}

// Deletar Viagem
func DeletarViagem(escritor http.ResponseWriter, requisicao *http.Request) {
	idStr := requisicao.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Printf("\nID da Nota editada: %d", id)

	corpo, err := io.ReadAll(requisicao.Body)
	if err != nil {
		escritor.WriteHeader(http.StatusInternalServerError)
		escritor.Write([]byte("Erro ao ler corpo da requisição!"))
		return
	}

	var viagemRequest viagem.ViagemRequest
	err = json.Unmarshal(corpo, &viagemRequest)
	if err != nil {
		escritor.WriteHeader(http.StatusBadRequest)
		escritor.Write([]byte("Dados inválidos!"))
		return
	}

	viagem.DeletarViagem(viagemRequest.Destino, id)

	escritor.WriteHeader(http.StatusCreated)
	escritor.Write([]byte("Viagem Deletada com sucesso!"))
}
