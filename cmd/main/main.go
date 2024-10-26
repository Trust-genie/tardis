package main

import (
	"log"
	"os"
	"tardis/internals/handlers"
	"tardis/internals/routers"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Failed to Load Secrets")
	}
}

func main() {

	port, ok := os.LookupEnv("PORT")
	if !ok || port == "" {
		log.Fatal("Error: Missing Secrets")
	}

	router := routers.NewRouter()

	//register routers
	router.GET("/", handlers.Ping)
	router.POST("/", handlers.Create)
	router.PATCH("/update", handlers.Update)
	router.DELETE("/delete", handlers.Delete)

	//start router
	router.Run(":" + port)
}

/*
func main() {
	//this main allows me to test my functionality interactively
	var (
		functions = []string{"Put", "Create", "Delete", "Insert"}

		sct               string
		inputKey          string
		inputvalue interface{}
		err               error
		reader = bufio.NewReader(os.Stdin)
	)
	log.Println("Welcome to the interactive Tester\n")
	time.Sleep(2 * time.Second)
	for {
		log.Println("Enter Selection")
		log.Println(functions)
		fmt.Scanln(&sct)
		log.Println("\nEnter Key ...")
		fmt.Scanln(&inputKey)
		log.Println("Enter Value")
		inputvalue, _ = reader.ReadString('\n')
		inputvalue = strings.TrimSpace(inputvalue)

		switch sct {
		case "Put":
			err = storage.Store.Put(inputKey, inputvalue)
			if err != nil {
				log.Printf("Error: %v", err)
			} else {
				log.Println("Success")
			}
		case "Create":
			err = storage.Store.Insert(inputKey, inputvalue)
			if err != nil {
				log.Printf("Error: %v", err)
			} else {
				log.Println("Success")
			}
		case "Get":
			value, err = storage.Store.Get(inputKey)
			if err != nil {
				log.Printf("Error: %v", err)
			} else {
				log.Panicln("%v", value)
			}
		case "Delete":
			storage.Store.Delete(inputKey)
			log.Println("Success")

		}
	}
}

*/
