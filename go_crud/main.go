package main

import (
	"net/http"

	techcontroller "github.com/elysiamori/go-crud/controllers"
)

/*Web Routes*/
func main() {
	http.HandleFunc("/", techcontroller.Index)
	http.HandleFunc("/tech", techcontroller.Index)
	http.HandleFunc("/tech/index", techcontroller.Index)
	http.HandleFunc("/tech/add", techcontroller.Add)
	http.HandleFunc("/tech/edit", techcontroller.Edit)
	http.HandleFunc("/tech/delete", techcontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
