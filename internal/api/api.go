package api

import (
	"caltoph/internal/config"
	"caltoph/internal/health"
	"caltoph/internal/logger"
	"context"

	docs "caltoph/docs"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/oauth2"
)

// @title Caltoph API
// @version 1.0
// @BasePath /_caltoph/v1

var providers map[string]oidc_provider

type oidc_provider struct {
	provider *oidc.Provider
	config   oauth2.Config
}

type loginBody struct {
	Provider string `json:"provider" validate:"required"`
	Code     string `json:"code" validate:"required"`
}

type loginReturn struct {
	Session_id string `json:"session_id"`
}

type errormsg struct {
	Message string `json:"message"`
}

func Init() {
	logger.DebugLogger.Println("api: Initializing api")
	//Initialize oidc providers
	providers = map[string]oidc_provider{}
	for _, a := range config.Config.Oidc_provider {
		var err error
		o := oidc_provider{}
		o.provider, err = oidc.NewProvider(context.TODO(), a.Url)
		if err != nil {
			logger.FatalLogger.Panicln("Failed to initialize oidc providers")
		}
		o.config = oauth2.Config{
			ClientID:     a.Client_id,
			ClientSecret: a.Client_secret,
			Endpoint:     o.provider.Endpoint(),
		}
		providers[a.Name] = o
	}
	//Initializer gin
	if !config.Config.DevMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.GET("/_caltoph/v1/health", getHealth)
	router.GET("/_caltoph/v1/auth/oidc_servers", getOidc_Servers)
	router.POST("/_caltoph/v1/auth/login", login)
	docs.SwaggerInfo.BasePath = "/_caltoph/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	go router.Run(":8080")
	logger.DebugLogger.Println("api: Finished initializing api")
}

// getHealth godoc
// @Summary Get app health
// @Despcription Checks db and cna connection
// @Success 200
// @Failure 500
// @Router /health [get]
func getHealth(c *gin.Context) {
	logger.DebugLogger.Println("api: getHealth is called")
	health, msg := health.GetHealth()
	if health {
		c.String(200, msg)
		logger.DebugLogger.Println("api: getHealth is OK")
		return
	}
	if !health {
		c.String(500, msg)
		logger.WarningLogger.Println("api: getHealth is NOT OK")
		return
	}
}

// getOidc_Servers godoc
// @Summary Get oidc authentication providers
// @Despcription Get a list of oidc authentication provides, that can be used for authentication
// @Success 200 {array} config.Oidc_provider
// @Router /auth/oidc_servers [get]
func getOidc_Servers(c *gin.Context) {
	logger.DebugLogger.Println("api: getOidc_Servers is called")
	c.IndentedJSON(200, config.Config.Oidc_provider)
	logger.DebugLogger.Println("api: getOidc_Servers finished without error")
}

// login godoc
// @Summary Perform a login action and create a session
// @Despcription Perform a login action and create a session
// @Success 200 {object} loginReturn
// @Failure 400 {object} errormsg
// @Failure 500 {object} errormsg
// @Router /auth/login [post]
// @Param Login body loginBody true "Authorize data from oidc provider"
func login(c *gin.Context) {
	logger.DebugLogger.Println("api: login is called")

	var body loginBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.IndentedJSON(400, errormsg{Message: err.Error()})
		return
	}

	o, ok := providers[body.Provider]
	if !ok {
		c.IndentedJSON(400, errormsg{Message: "Provider does not exist"})
		return
	}
	token, err := o.config.Exchange(context.TODO(), body.Code)
	if err != nil {
		logger.ErrorLogger.Println(err)
		c.IndentedJSON(500, errormsg{Message: "Failed to obtain token"})
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		logger.ErrorLogger.Println("api: login: id_token is missing")
		c.IndentedJSON(500, errormsg{Message: "Failed to obtain token"})
	}

	verifier := o.provider.Verifier(&oidc.Config{ClientID: o.config.ClientID})
	idToken, err := verifier.Verify(context.TODO(), rawIDToken)
	if err != nil {
		logger.ErrorLogger.Println(err)
		c.IndentedJSON(500, errormsg{Message: "Failed to obtain token"})
	}

	claims := map[string]interface{}{}
	err = idToken.Claims(claims)
	if err != nil {
		logger.ErrorLogger.Println(err)
		c.IndentedJSON(500, errormsg{Message: "Failed to obtain token"})
	}

	c.String(200, "This is a legit call")
}
