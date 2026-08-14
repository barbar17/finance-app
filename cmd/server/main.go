package main

import (
	"log"

	"github.com/barbar17/finance-app/internal/auth"
	"github.com/barbar17/finance-app/internal/db"
	"github.com/barbar17/finance-app/internal/module/dashboard"
	errorpage "github.com/barbar17/finance-app/internal/module/error_page"
	"github.com/barbar17/finance-app/internal/module/login"
	"github.com/barbar17/finance-app/internal/module/transaction"
	"github.com/barbar17/finance-app/internal/utils"
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

	// LoadTemplates load the html template (layout) in the web/pages folder
	utils.LoadTemplates()
	// Static load the static file including css and js
	r.Static("/static", "./web/static")

	//unguarded routes
	errorpage.ErrorPageRoutes(r)
	login.LoginRoutes(r, sessionStore)

	//AuthGuard makes sure the routes below is guarded by authentication
	r.Use(auth.AuthGuard(sessionStore))
	dashboard.DashboardRoutes(r)
	transaction.TransactionRoutes(r)

	r.Run(":8080")
}
