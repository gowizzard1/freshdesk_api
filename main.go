package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"powergen/freshdesk/appconfig"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/muhoro/log"
)

// ContextInjector context injector
type ContextInjector struct {
	ctx context.Context
	h   http.HandlerFunc
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))

}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server is running")
}
func buildConfigs(usedocker bool) appconfig.AppConfig {
	var conf appconfig.AppConfig
	if usedocker {
		conf = appconfig.AppConfig{
			Freshconn:  os.Getenv("connection_string"),
			Seq_Server: os.Getenv("seq"),
		}
	} else {
		conf = appconfig.AppConfig{
			Freshconn:      "host=localhost port=5432 user=postgres dbname=fresdesk password=root sslmode=disable",
			Seq_Server:     "http://localhost:5341",
			Seq_Server_Key: "key",
		}
	}
	appconfig.SetConfig(conf)
	return conf
}

func main() {
	usedocker := false
	buildConfigs(usedocker)
	log.BuildLogger("FRESHDESK").UseSeq(appconfig.GetConfig().Seq_Server, appconfig.GetConfig().Seq_Server_Key)
	db, err := gorm.Open("postgres", appconfig.GetConfig().Freshconn)
	if err != nil {
		log.Fatal(err.Error(), nil)
		fmt.Println(err)
	}
	db, er := gorm.Open("postgres", appconfig.GetConfig().Freshconn)
	if er != nil {
		//panic(er.Error())
	}

	ctx := context.WithValue(context.Background(), "database", db)
	router := mux.NewRouter()
	router.Handle("/test", &ContextInjector{ctx, test}).Methods("POST")
	http.ListenAndServe(":8006", router)
}
