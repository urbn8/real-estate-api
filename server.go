package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

func main() {
	log.SetLevel(log.DebugLevel)

	if err := mongo.Startup(); err != nil {
		log.Fatalf("error[%s] while startup mongodb connection", err)
	}

	if err := EnsureIndexProperty(); err != nil {
		log.Fatalf("error[%s] while ensure index on properties collection", err)
	}

	if err := seedDataIfNeeded(); err != nil {
		log.Fatalf("error[%s] while seed data", err)
	}

	category := categoryHandlers{}
	property := propertyHandlers{}
	images := imageHandlers{}
	contactInfo := contactInfoHandlers{}

	router := fasthttprouter.New()
	router.GET("/", Index)

	router.GET("/categories", category.find)
	router.POST("/categories", category.create)

	router.GET("/properties", property.find)
	router.POST("/properties", property.create)
	router.GET("/properties/:id", property.get)
	router.PUT("/properties/:id", property.update)
	router.DELETE("/properties/:id", property.remove)
	router.GET("/api/categories", NotFound)

	router.GET("/images", images.find)
	router.POST("/images", images.create)
	router.DELETE("/images/:id", images.remove)

	router.GET("/contact_info/defaults", contactInfo.find)

	if err := fasthttp.ListenAndServe(":9001", router.Handler); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("fail start server")
	}
}

func Index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	fmt.Fprint(ctx, "Welcome to REAL ESTATE API!\n")
}

func NotFound(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	log.WithFields(log.Fields{
		"path": string(ctx.RequestURI()),
	}).Infoln("not found")
	ctx.SetStatusCode(404)
}
