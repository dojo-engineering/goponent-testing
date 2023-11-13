package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Car struct {
	Make           string
	Model          string
	Id             string
	RegistrationId string
}

var cars []Car

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func carGet(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/car/")
	i := slices.IndexFunc(cars, func(c Car) bool {
		return c.Id == id
	})

	if i == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	car := cars[i]

	res, err := http.Get("https://www.example.com/example-payload")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	car.RegistrationId = string(b)

	err = json.NewEncoder(w).Encode(car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func carPost(w http.ResponseWriter, req *http.Request) {
	car := Car{}
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	car.Id = uuid.NewString()
	cars = append(cars, car)

	err = json.NewEncoder(w).Encode(car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func buildServer() http.Handler {
	server := http.NewServeMux()

	server.HandleFunc("/hello", hello)
	server.HandleFunc("/headers", headers)
	server.HandleFunc("/car", carPost)
	server.HandleFunc("/car/", carGet)
	return server
}
