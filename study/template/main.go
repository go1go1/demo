package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认不会解析
	fmt.Println(r.Form) //
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Helllo world!") //输出到客户端
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./logins.html")
		if err != nil {
			fmt.Fprintf(w, "load template failed, %#v", err)
			return
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		fmt.Printf("username:%s\n", r.FormValue("username"))
		return
	}
}

func main() {
	//http.HandleFunc("/", sayHello)   //设置访问的路由
	//http.HandleFunc("/login", login) //

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)
	mux.HandleFunc("/login", login)

	err := http.ListenAndServe(":9000", mux) //设置监听
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
