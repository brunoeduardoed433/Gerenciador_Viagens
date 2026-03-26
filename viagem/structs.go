package viagem

import "time"

type ViagemData struct {
	UltimoID int
	Data     []Viagem
}

type ViagemNota struct {
	UltimoID int
	Note     []Nota
}

type Nota struct {
	ID          int       `json:"id"`
	Conteudo    string    `json:"conteudo"`
	DataCriacao time.Time `json:"data-cricao"`
}

type Viagem struct {
	ID      int    `json:"id"`
	Destino string `json:"destino"`
	Notas   []Nota `json:"nota"`
}
