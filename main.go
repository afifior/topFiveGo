package main

import (
	"flag"
	"fmt"
	"net/http"
)

type addHandler struct {
	cm chunksmap
}
type topFiveHandler struct {
	cm chunksmap
}

var cm chunksmap

func main() {
	flag.Parse()
	cm := new(chunksmap)
	cm.dictionary = make(map[string]int)
	mh := maxheap{heapArray: []HeapElement{}, values: make(map[string]int, 5), size: 0, maxsize: 5}
	cm.hp = &mh

	cm.addString("is ball ball ball ball eggs eggs pool pool wild daily last")
	cm.addString("is is is is is")
	fmt.Printf("%s", cm.getTopMembers())
	http.Handle("/", &topFiveHandler{cm: *cm})
	http.Handle("/add", &addHandler{cm: *cm})
	http.ListenAndServe(":3000", nil)
}
func (h *topFiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.cm.getTopMembers()))
}
func (h *addHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	newString := r.FormValue("str")
	if newString != "" {
		h.cm.addString(newString)
	}
	fmt.Fprint(w, addForm)
}

const addForm = `
<html><body>
<form method="POST" action="/add">
Add String: <input type="text" name="str">
<input type="submit" value="Add">
</form>
</html></body>
`
