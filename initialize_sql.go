package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func initializeSQL() (string,string,string,string,string)  {
err:=godotenv.Load();
if err != nil {
fmt.Print("problem while loading .env file from forsql.go");
}
	USERNAME := os.Getenv("USERNAME_MESSENGER")
	DBPASSWORD := os.Getenv("PASSWORD_MESSENGER")
	DBHOST := os.Getenv("HOST_MESSENGER")
	DBPORT := os.Getenv("PORT_MESSENGER")
	DBNAME := os.Getenv("NAME_MESSENGER")
   
    
	return USERNAME, DBPASSWORD, DBHOST, DBPORT, DBNAME
}


