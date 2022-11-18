package transactions

import (
	"github.com/bxcodec/faker/v3"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
	"txapp/graph/model"
)

func GetTransactions() []*model.Transaction {
	var txlist = make([]*model.Transaction, 50)

	for i := 0; i < 50; i++ {
		txlist[i] = &model.Transaction{
			ID:     uuid.NewV4().String(),
			Status: model.AllTransactionStatus[rand.Intn(len(model.AllTransactionStatus))],
			Date:   time.Now().AddDate(0, 0, rand.Intn(5)),
			Name:   faker.FirstNameFemale(),
			Type:   model.AllTransactionType[rand.Intn(len(model.AllTransactionType))],
			Amount: rand.Intn(1_000_000-1_000) + 1_000,
		}
	}

	return txlist
}
