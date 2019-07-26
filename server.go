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
	sc := r.FormValue("script")
	logrus.Infof("run-script {\n%s\n}", sc)

	ctrl, err := NewController()
	if err != nil {
		logrus.Errorf("new-controller: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = RunLua(sc, ctrl)
	if err != nil {
		logrus.Errorf("run-lua: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/index", http.StatusFound)
}

var appHTML = `
<html>
<head>
    <meta charset="UTF-8" name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />    
    <title>EV3</title>    
</head>
<body>    
	<div class="app">
		<h1>Enter LUA script and run</h1>

		<iframe class="null-frame" name="null-frame"></iframe>
		
		<form class="lua-form" id="lua-form" action="/runlua" method="post" target="null-frame">            		
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

.lua-script{
	font-family: courier;
	width: 800px;
	height: 400px;
}

.null-frame{
	height: 0px;
	width: 0px;
	visibility: hidden;
}
</style>
</html>
`
