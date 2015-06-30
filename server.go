/***
    These are prototype pages for the dumaguete site
**/
package main

import (
    "log"
    "net/http"
)

//The landing page
func indexHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("in index handler") //:) 
    http.ServeFile(w, r, "static/index.html")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("in static handler")
    http.ServeFile(w, r, r.URL.Path[1:])
}

func picHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("getting pictures")
    log.Println(r.URL.Path[1:])
    http.ServeFile(w, r, r.URL.Path[1:])
}


func userHandler(w http.ResponseWriter, r *http.Request){
    log.Println("serve files")
    http.ServeFile(w, r, "static/user.html")
    /**
    * To-do
    */
}

//temporary fix now, but this enables authentication
func logInHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "static/log_in.html")
}

func registerHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "static/register.html")
}

//serve .js files, right now. Will be using jquery to do simple
//tasks as of this moment
func jsHandler(w http.ResponseWriter, r *http.Request){
    //need to serve javascript files
    http.ServeFile(w, r, r.URL.Path[1:])
}
func main() {
    //fs := http.FileServer(http.Dir("static"))
    mux := http.NewServeMux()
    //static_mux := http.NewServeMux()
    
    index_handler := http.HandlerFunc(indexHandler)
    static_handler := http.HandlerFunc(staticHandler)
    pic_handler := http.HandlerFunc(picHandler)
    js_handler := http.HandlerFunc(jsHandler)
    user_handler := http.HandlerFunc(userHandler)
    log_in_handler := http.HandlerFunc(logInHandler)
    register_handler := http.HandlerFunc(registerHandler)

    mux.Handle("/", index_handler)
    mux.Handle("/static/", static_handler)
    mux.Handle("/static/pictures/", pic_handler)
    mux.Handle("/static/js/", js_handler)
    mux.Handle("/article/", user_handler) //shows one article that is part in the
                            //data-base
    mux.Handle("/log-in/", log_in_handler)
    mux.Handle("/register/", register_handler)
    
    //http.HandleFunc("/", handler)


    log.Println("Listening...")
    http.ListenAndServe(":8080", mux)
}