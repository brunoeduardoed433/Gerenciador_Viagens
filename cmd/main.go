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

	slog.Info("Servidor escutando em localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

func EditarViagem(escritor http.ResponseWriter, requisicao *http.Request) {
	idStr := requisicao.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)
}

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