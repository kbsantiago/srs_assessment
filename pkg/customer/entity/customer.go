package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Customer teste
type Customer struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Cpf      string        `bson:"_cpf,omitempty"`
	Nome     string        `bson:"_nome,omitempty"`
	Endereco string        `bson:"_endereco,omitempty"`
	Dividas  []Bill        `bson:"_dividas,omitempty"`
}

//Bill teste
type Bill struct {
	Valor int64     `bson:"_valor,omitempty"`
	Data  time.Time `bson:"_data,omitempty"`
}
