package passing_data

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", val)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func val(w http.ResponseWriter, req *http.Request) {
	//this can retrieve params from forms or URLs (but body params have preference)
	v := req.FormValue("q")

	//If we are uploading a file we can use:
	//f, h, err := req.FormFile("q")
	//being f the file and h the header

	//to store:
	//bs, err := ioutil.ReadAll(f)
	//dst, err := os.Create(filepath.Join("./user/", h.Filename))
	//_, err = dst.Write(bs)

	fmt.Fprintln(w, "Do my search: "+v)
}

// visit this page:
// http://localhost:8080/?q=dog
