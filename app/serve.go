package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Server struct {
	host string
	port int

	appUsername string
	appPassword string

	page []byte
}

func NewServer(host string, port int, pagePath string) *Server {
	username := os.Getenv("DEMO_USERNAME")
	password := os.Getenv("DEMO_PASSWORD")

	page := []byte(defaultPage)
	if pagePath != "" {
		altPage, err := ioutil.ReadFile(pagePath)
		if err == nil {
			page = altPage
		}
	}

	return &Server{
		host:        host,
		port:        port,
		appUsername: username,
		appPassword: password,
		page:        page,
	}
}

func (s *Server) Serve() error {
	http.HandleFunc("/", s.ServeIndex)
	http.HandleFunc("/api", s.ServeAPI)

	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) ServeIndex(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(s.page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) ServeAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	if s.appUsername == "" {
		writeServerError(w, "DEMO_USERNAME environment variable not set")
		return
	}

	if s.appPassword == "" {
		writeServerError(w, "DEMO_PASSWORD environment variable not set")
		return
	}

	req, err := http.NewRequest("GET", "https://demo.secrethub.io/api/v1/basic-auth", nil)
	if err != nil {
		writeServerError(w, err.Error())
		return
	}

	req.SetBasicAuth(s.appUsername, s.appPassword)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		writeServerError(w, err.Error())
		return
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		panic(err)
	}
}

func writeServerError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, message)
}
