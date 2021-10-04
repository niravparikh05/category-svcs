package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/niravparikh05/category-svcs/databases"
)

const (
	CATEGORY_BY_ID = "/api/category/{id}"
	CATEGORY       = "/api/category"
)

func main() {
	fmt.Println("Portfolio Manager Services Started ...")

	fmt.Println("Reading database properties ..")
	//read properties to connect to mongo db
	dbprops, err := databases.ReadDatabaseProps(os.Getenv("PMS_CONFIG"))
	if err != nil {
		log.Fatalln("Error while reading database properties ", err.Error())
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc(CATEGORY_BY_ID, dbprops.GetCategoryById).Methods(http.MethodGet)
	router.HandleFunc(CATEGORY, dbprops.GetAllCategories).Methods(http.MethodGet)
	router.HandleFunc(CATEGORY, dbprops.CreateCategory).Methods(http.MethodPost)
	router.HandleFunc(CATEGORY_BY_ID, dbprops.UpdateCategory).Methods(http.MethodPut)
	router.HandleFunc(CATEGORY_BY_ID, dbprops.DeleteCategory).Methods(http.MethodDelete)

	fmt.Println("Listening for requests ...")
	http.ListenAndServe(":8080", router)

}
