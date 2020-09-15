package main

import (
	"fmt"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for i, i2 := range cookies {
		fmt.Printf("index:%d, cookie:%#v\n", i, i2)
	}

	c, err := r.Cookie("sessionid")
	fmt.Printf("sessionid:%v, err:%v\n", c, err)

	cookie := &http.Cookie{
		Name:   "sessionid",
		Value:  "dsfdsfsdfdsfdsfs",
		MaxAge: 3600,
	}
	http.SetCookie(w, cookie)

	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":9000", nil)
}
