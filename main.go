package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	db := initDatabase()

	customerRepositoryDB := repository.NewCustomerReposotoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID}", customerHandler.GetCustomer).Methods(http.MethodGet)

	// log.Printf("Banking service started ar port %v", viper.GetInt("app.port"))
	logs.Info("Banking service started ar port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf("localhost:%v", viper.GetInt("app.port")), router)

	//Service
	// customers, err := customerService.GetCustomers()
	// customer, err := customerService.GetCustomer(2000)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)
	// fmt.Println(customer)

	// Repository
	// customers, err := customerRepository.GetAll()
	// customer, err := customerRepository.GetById(2000)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)
	// fmt.Println(customer)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable",
		viper.GetString("db.driver"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return db
}
