package techcontrollers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/elysiamori/go-crud/entities"
	"github.com/elysiamori/go-crud/libraries"
	"github.com/elysiamori/go-crud/models"
)

var validation = libraries.NewValidation()
var techModel = models.NewTechModel()

/*Konsep CRUD*/
func Index(response http.ResponseWriter, request *http.Request) {

	tech, _ := techModel.FindAll()

	data := map[string]interface{}{
		"tech": tech,
	}

	temp, err := template.ParseFiles("views/tech/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/tech/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var tech entities.Tech
		tech.Name = request.Form.Get("name")
		tech.Job = request.Form.Get("job")
		tech.Programming = request.Form.Get("programming")
		tech.Date = request.Form.Get("date")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(tech)

		if vErrors != nil {
			data["tech"] = tech
			data["validation"] = vErrors
		} else {
			data["message"] = "Tech data has been saved"
			techModel.Create(tech)
		}
		/*Mengirimkan data ke database*/
		temp, _ := template.ParseFiles("views/tech/add.html")
		temp.Execute(response, data)

	}

}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		//Mengambil id yang ada pada url
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		//Mengambil data didatabase sesuai data id di tech
		var tech entities.Tech
		techModel.Find(id, &tech)

		//Menampung data variable tech
		data := map[string]interface{}{
			"tech": tech,
		}

		temp, err := template.ParseFiles("views/tech/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var tech entities.Tech
		//Syntax tambahan untuk update data
		tech.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		tech.Name = request.Form.Get("name")
		tech.Job = request.Form.Get("job")
		tech.Programming = request.Form.Get("programming")
		tech.Date = request.Form.Get("date")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(tech)

		if vErrors != nil {
			data["tech"] = tech
			data["validation"] = vErrors
		} else {
			data["message"] = "Tech data has been updated"
			techModel.Update(tech)
		}
		/*Mengirimkan data ke database*/
		temp, _ := template.ParseFiles("views/tech/edit.html")
		temp.Execute(response, data)

	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	techModel.Delete(id)

	http.Redirect(response, request, "/tech", http.StatusSeeOther)
}
