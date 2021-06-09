package main

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/routers"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "app/diplom/pkg/configs"
)

// init function
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found, looking for OS environment")
	}
}

func main() {

	conn, pool, err := db.Connect(config.New().DatabaseURL)
	if err != nil {
		log.Println(err)
	}
	defer pool.Close()
	defer conn.Release()
	var name string
	name = "hj"
	fmt.Println("hjj" + name)
	//handleReq()

	router := routers.SetupRouter()
	log.Println("Service starting on port " + config.New().APIPort)
	//f, _ := os.Create(config.New().FileLogName)
	//gin.DefaultWriter = io.MultiWriter(f)

	srv := &http.Server{
		Addr:    config.New().APIPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

}

//func handleReq(){
//	http.HandleFunc("/", home_page)
//	http.ListenAndServe("localhost:8080", nil)
//}
//
//func home_page(w http.ResponseWriter, r *http.Request){
////	fmt.Fprintf(w, "Hello ghelli")
//
//	bob := models.User{
//		ID:       0,
//		Username: "vr",
//		Password: "vgrt",
//		FIO:      "rteg",
//		RoleID:   0,
//	}
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("frt")
//	fmt.Println(wd)
//	tmpl, _ := template.ParseFiles(wd + "/pages/home_page.html")
//	tmpl.Execute(w, bob)
//
//}
