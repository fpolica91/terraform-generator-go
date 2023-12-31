// api/main.go
package main

import (
	"encoding/json"
	"fmt"
	"generatorv/cmds"
	_ "generatorv/docs"
	"generatorv/pkgs"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// handleCreateBuckets godoc
// @Summary Create S3 buckets
// @Description Create buckets based on the provided configuration
// @Tags buckets
// @Accept json
// @Produce json
// @Param buckets body []pkgs.Bucket true "Array of Bucket Configurations"
// @Success 200 {string} string "Buckets created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /createbuckets [post]
func handleCreateBuckets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data struct {
		Buckets []pkgs.Bucket `json:"buckets"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the CreateBuckets function from cmds package
	resourcesString, err := cmds.CreateBuckets(data.Buckets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	// fmt.Fprint(w, terraformString)
	json.NewEncoder(w).Encode(struct {
		ResourcesString string `json:"resourcesString"`
	}{ResourcesString: resourcesString})

}

func handleCreateProvider(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var provider pkgs.Prov

	err = json.Unmarshal(body, &provider)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	providerString, err := cmds.CreateProvider(provider)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	variablesString, err := cmds.CreateVariables(provider.Provider)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		ProviderString  string `json:"providerString"`
		VariablesString string `json:"variablesString"`
	}{ProviderString: providerString, VariablesString: variablesString})

}

func handleCreateVirtualPrivateCloud(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data struct {
		Vpcs []pkgs.VPC `json:"vpcs"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the CreateBuckets function from cmds package
	err = cmds.CreateVirtualPrivateCloud(data.Vpcs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Buckets created successfully")
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register your handlers to the ServeMux
	mux.HandleFunc("/createprovider", handleCreateProvider)
	mux.HandleFunc("/createbuckets", handleCreateBuckets)
	mux.HandleFunc("/createvpcs", handleCreateVirtualPrivateCloud)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Define CORS policy
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	// Apply the CORS middleware to your ServeMux
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(mux)

	// Launch server with the CORS-enabled handler
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
