package viagem

import "time"

type Viagem struct {
	ID      int    `json:"id"`
	Destino string `json:"destino"`
	Notas   []Nota `json:"nota"`
}

type ViagemData struct {
	ViagemUltimoID int      `json:"viagem-ultimo-id"`
	NotaUltimoID   int      `json:"nota-ultimo-id"`
	Viagens           []Viagem `json:"data"`
}

type Nota struct {
	ID          int       `json:"id"`
	Conteudo    string    `json:"conteudo"`
	DataCriacao time.Time `json:"data-cricao"`
}