package dashboardcontroller

import (
	"net/http"
	"text/template"
)

const viewsDir = "app/views/"

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(
		viewsDir+"/_layout/header.html",
		viewsDir+"/_layout/sidebar.html",
		viewsDir+"/dashboard/index.html",
		viewsDir+"/_layout/footer.html")

	if err != nil {
		panic(err)
	}

	temp.ExecuteTemplate(w, "index", nil)
}
