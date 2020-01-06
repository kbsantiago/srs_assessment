package entity

import (
	"gopkg.in/mgo.v2/bson"
)

//Score teste
type Score struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Cpf          string        `json:"cpf" bson:"_cpf,omitempty"`
	Idade        string        `json:"idade" bson:"_idade,omitempty"`
	Bens         []Ownership   `json:"bens" bson:"_bes,omitempty"`
	FonteDeRenda int64         `json:"renda" bson:"_renda,omitempty"`
}

//Ownership struct
type Ownership struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Descricao string        `json:"descricao" bson:"_descricao,omitempty"`
}
