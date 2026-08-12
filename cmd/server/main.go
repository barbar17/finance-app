package main

import (
	"log"

	"github.com/barbar17/finance-app/internal/auth"
	"github.com/barbar17/finance-app/internal/dashboard"
	"github.com/barbar17/finance-app/internal/db"
	"github.com/barbar17/finance-app/internal/login"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	defer db.CloseDB()

	sessionStore := auth.NewSessionStore()

	r := gin.Default()

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// loadWebPages load the html files in the web/pages folder
	tmpl, err := loadWebPages()
	if err != nil {
		log.Fatalf("Failed to load web pages: %v", err)
	}

	//SetHTMLTemplate set the html pages and Static set the static folder for css and js
	r.SetHTMLTemplate(tmpl)
	r.Static("/static", "./web/static")

	//unguarded routes
	login.LoginRoutes(r, sessionStore)

	//AuthGuard makes sure the routes below is guarded by authentication
	r.Use(auth.AuthGuard(sessionStore))
	dashboard.DashboardRoutes(r)

	r.Run(":8080")
}
