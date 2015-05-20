package user

import(
//	"html/template"
	//"io/ioutil"
	"net/http"
	"fmt"
)

func registration(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
}
	fmt.Println("POST")

	for key, values := range r.Form {   // range over map
		for _, value := range values {    // range over []string
			fmt.Println(key, value)
		}
	}
	fmt.Printf("%+v\n", r.FormValue("birthday")) // Returns nothing

}