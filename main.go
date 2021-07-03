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
		services.WriteResponse(w, http.StatusBadRequest, err.Error())
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
		services.WriteResponse(w, http.StatusBadRequest, err.Error())
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		services.WriteResponse(w, http.StatusBadRequest, err.Error())
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	services.WriteResponse(w, http.StatusOK, tempFile.Name())
}

func handler() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/api/user", services.SearchUser).Methods("GET")
	r.HandleFunc("/api/userJoin/{id}", services.GetJoin).Methods("GET")
	r.HandleFunc("/api/user/{id}", services.GetUser).Methods("GET")
	r.HandleFunc("/api/user", services.InsertUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", services.UpdateUser).Methods("PATCH")
	r.HandleFunc("/api/user/{id}", services.DeleteUser).Methods("DELETE")

	r.HandleFunc("/api/product", services.SearchProduct).Methods("GET")
	r.HandleFunc("/api/product/{id}", services.GetProduct).Methods("GET")
	r.HandleFunc("/api/product", services.InsertProduct).Methods("POST")
	r.HandleFunc("/api/product/{id}", services.UpdateProduct).Methods("PATCH")
	r.HandleFunc("/api/product/{id}", services.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/api/order", services.SearchOrder).Methods("GET")
	r.HandleFunc("/api/order/{id}", services.GetOrder).Methods("GET")
	r.HandleFunc("/api/order", services.InsertOrder).Methods("POST")
	r.HandleFunc("/api/order/{id}", services.UpdateOrder).Methods("PATCH")
	r.HandleFunc("/api/product/{id}", services.DeleteOrder).Methods("DELETE")

	r.HandleFunc("/api/upload", uploadFile).Methods("POST")

	// Create the route public
	staticDir := "/public/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	fmt.Println("Start server golang ...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func main() {
	handler()
}
