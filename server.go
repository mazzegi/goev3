package goev3

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(bind string) (*Server, error) {
	s := &Server{
		httpServer: &http.Server{
			Addr: bind,
		},
		router: mux.NewRouter().StrictSlash(true),
	}
	s.router.HandleFunc("/index", s.handleGETApp).Methods("GET")
	s.router.HandleFunc("/runlua", s.handlePOSTRunLua).Methods("POST")
	s.httpServer.Handler = s.router
	return s, nil
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) handleGETApp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(appHTML))
}

func (s *Server) handlePOSTRunLua(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("runlua: %s; q=%v", r.URL, r.URL.Query())
}

var appHTML = `
<html>
<head>
    <meta charset="UTF-8" name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
    <meta http-equiv="refresh" content="120" >
    <title>Gonzo</title>    
</head>
<body>    
	<div class="app">
		<h1>Enter LUA script and run</h1>
		
		<form class="lua-form" id="lua-form" action="/runlua" method="post">            		
		<button class="lua-run-button">Run</button>
        </form>
		<textarea class="lua-script" name="script" form="lua-form"></textarea>
	</div>
</body>

<style>
.app{
	margin: auto;
    margin-top: 100px;
    max-width: 600px;    
}
</style>
</html>
`
