package app

import (
	"log"
	"net/http"
	"time"
)

func (app *Application) Serve() error {
	srv := http.Server{
		Addr:         app.Config.Port,
		Handler:      app.Config.Handler,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("server running on %s\n", app.Config.Port)
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
