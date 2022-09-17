package repository

import (
	"log"

	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v73/client"
)

type DAO interface {
	NewProductQuery() ProductQuery
}

type dao struct{}

var STRIPE *client.API

func stripeQb() *client.API {
	return STRIPE
}

func NewDAO(client *client.API) DAO {
	STRIPE = client
	return &dao{}
}

func NewStripe() (*client.API, error) {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}
	key := viper.GetString("stripe.key")

	// Starting a STRIPE
	client := client.New(key, nil)
	return client, nil
}

func (d *dao) NewProductQuery() ProductQuery {
	return &productQuery{}
}
