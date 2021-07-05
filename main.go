package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	services "github.com/gusgus-project/services"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		services.WriteResponseUpload(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("public/files/images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		services.WriteResponseUpload(w, http.StatusBadRequest, err.Error())
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		services.WriteResponseUpload(w, http.StatusBadRequest, err.Error())
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	services.WriteResponseUpload(w, http.StatusOK, tempFile.Name())
}

func newRouters() {
	// target := "https://gusts-project.herokuapp.com"
	// remote, err := url.Parse(target)
	// if err != nil {
	// 	panic(err)
	// }

	// proxy := httputil.NewSingleHostReverseProxy(remote)

	r := mux.NewRouter()

	// r.HandleFunc("/forward/{rest:.*}", handler(proxy))
	r.HandleFunc("/api/user", services.SearchUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/userJoin/{id}", services.GetJoin).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/user/{id}", services.GetUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/user", services.InsertUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/user/{id}", services.UpdateUser).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/user/{id}", services.DeleteUser).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/product", services.SearchProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/product/{id}", services.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/product", services.InsertProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/product/{id}", services.UpdateProduct).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/product/{id}", services.DeleteProduct).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/order", services.SearchOrder).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/order/{id}", services.GetOrder).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/order", services.InsertOrder).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/order/{id}", services.UpdateOrder).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/product/{id}", services.DeleteOrder).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/upload", uploadFile).Methods("POST", "OPTIONS")

	// Create the route public
	staticDir := "/public/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	fmt.Println("Start server golang ...")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		r.URL.Path = mux.Vars(r)["rest"]
// 		p.ServeHTTP(w, r)
// 	}
// }

func main() {
	newRouters()
}
