package repository

import (
	"log"

	"github.com/kbsantiago/srs_assessment/pkg/customer/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

//MongoRepository mongodb repo
type MongoRepository struct {
	db     string
	server string
}

//NewMongoRepository create new repository
func NewMongoRepository(db string, server string) *MongoRepository {
	return &MongoRepository{
		db:     db,
		server: server,
	}
}

//Connect manager database connection
func (repo *MongoRepository) Connect() {
	session, err := mgo.Dial(repo.server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(repo.db)
}

func collection() *mgo.Collection {
	return db.C("customers")
}

//FindByDocument search by document value
func (repo *MongoRepository) FindByDocument(cpf string) (entity.Customer, error) {
	repo.Connect()
	var customer entity.Customer
	conditions := bson.M{"_cpf": bson.M{"$eq": cpf}}
	err := collection().Find(conditions).One(&customer)
	return customer, err
}

//Create saves company information in database
func (repo *MongoRepository) Create(customer entity.Customer) error {
	repo.Connect()
	res := collection().Insert(&customer)
	return res
}
