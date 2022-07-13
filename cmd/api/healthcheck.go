package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// js := `{"status":"available","enviroment":%q,"version":%q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	data := map[string]string{
		"status":     "available",
		"enviroment": app.config.env,
		"version":    version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem abd could not process your request", http.StatusInternalServerError)
		return
	}
}
