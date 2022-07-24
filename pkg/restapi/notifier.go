package restapi

import (
	"fmt"
	"net/http"
	"notifications/pkg/mailer"
)

func SendAllHandler(w http.ResponseWriter, r *http.Request) {
	res := mailer.Run()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Status: %v\n", res)
}
