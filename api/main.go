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
	err = cmds.CreateBuckets(data.Buckets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Buckets created successfully")
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
	fmt.Println(provider.Provider, "provider")
	fmt.Println(provider.ProviderSource, "provider_source")
	fmt.Println(provider.ProviderVersion, "provider_version")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = cmds.CreateProvider(provider)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = cmds.CreateVariables()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Provider created successfully")

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
	http.HandleFunc("/createprovider", handleCreateProvider)
	http.HandleFunc("/createbuckets", handleCreateBuckets)
	http.HandleFunc("/createvpcs", handleCreateVirtualPrivateCloud)
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
