package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	//@ request is from signin form ,
	
USERNAME, DBPASSWORD, DBHOST, DBPORT, DBNAME:=initializeSQL();
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USERNAME, DBPASSWORD, DBHOST, DBPORT, DBNAME);
 db,err:=sql.Open("mysql",dsn);
 if err!=nil{
     fmt.Print("problem while reading database. ");
    }
    defer db.Close();
    err=db.Ping();
    if err!=nil{
        fmt.Print("Problem while connecting to database . RESTART AGAIN. ")
        return;
        
    }//! CORS ERROR HANDLE IT . 
     chiRouter:=chi.NewRouter();
   
   
     chiRouter.Post("/signin",func(w http.ResponseWriter, r *http.Request) {
     CORSfix(w);
      signIn(db,w,r);
      });
      chiRouter.Post("/sendrequest",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
        sendRequest(db,w,r);
        });
        chiRouter.Post("/homehistory",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
        homeHistory(db,w,r);
      })
        chiRouter.Post("/chathistory",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
        loadChatHistory(db,w,r);
      })
        chiRouter.Post("/getnamesfromlogintable",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
        getNamesFromLoginTable(db,w,r);
      })
        chiRouter.Post("/getnamesfromsearchhistorytable",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
        getNamesFromUserSearchHistoryTable(db,w,r);
      })
        chiRouter.Post("/getnames",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
        getNamesFromLoginTable(db,w,r);
      })
        chiRouter.Post("/getcurrentuseridandflag",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
      getCurrentUserIdAndFlag(db,w,r);//@ userid ra flag return garcha. 
      })
        chiRouter.Get("/wschatbody",func(w http.ResponseWriter, r *http.Request) {
        CORSfix(w);
      chatBodySocket(db,w,r);
      })
      
      
      log.Fatal(http.ListenAndServe(":8080",chiRouter));
      //// ALWAYS CLOSE THE BODY AND OTHER THINGS FROM table or other stuffs.
   }
   
   
   








    