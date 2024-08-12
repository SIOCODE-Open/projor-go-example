package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

var (
    userAccountRepository = NewUserAccountRepository()
    productRepository = NewProductRepository()
)

func HandleUserAccountCreate(w http.ResponseWriter, r *http.Request) {
    userAccount := &UserAccount{}
    if err := json.NewDecoder(r.Body).Decode(userAccount); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := userAccountRepository.Create(userAccount); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func HandleUserAccountUpdate(w http.ResponseWriter, r *http.Request) {
    userAccount := &UserAccount{}
    if err := json.NewDecoder(r.Body).Decode(userAccount); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := userAccountRepository.Update(userAccount); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleUserAccountDelete(w http.ResponseWriter, r *http.Request) {
    userAccount := &UserAccount{}
    if err := json.NewDecoder(r.Body).Decode(userAccount); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := userAccountRepository.Delete(userAccount); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleUserAccountGetByID(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/user-account/"):]
    userAccount, err := userAccountRepository.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    if err := json.NewEncoder(w).Encode(userAccount); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleUserAccountGetAll(w http.ResponseWriter, r *http.Request) {
    userAccounts, err := userAccountRepository.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(userAccounts); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleUserAccountRequest(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
    switch r.Method {
    case http.MethodOptions:
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Allow", "GET, POST, PUT, DELETE, OPTIONS")
    case http.MethodPost:
        HandleUserAccountCreate(w, r)
    case http.MethodPut:
        HandleUserAccountUpdate(w, r)
    case http.MethodDelete:
        HandleUserAccountDelete(w, r)
    case http.MethodGet:
        if r.URL.Path == "/user-account" {
            HandleUserAccountGetAll(w, r)
        } else {
            HandleUserAccountGetByID(w, r)
        }
    default:
        http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
    }
}
func HandleProductCreate(w http.ResponseWriter, r *http.Request) {
    product := &Product{}
    if err := json.NewDecoder(r.Body).Decode(product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := productRepository.Create(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func HandleProductUpdate(w http.ResponseWriter, r *http.Request) {
    product := &Product{}
    if err := json.NewDecoder(r.Body).Decode(product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := productRepository.Update(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleProductDelete(w http.ResponseWriter, r *http.Request) {
    product := &Product{}
    if err := json.NewDecoder(r.Body).Decode(product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := productRepository.Delete(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleProductGetByID(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/product/"):]
    product, err := productRepository.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    if err := json.NewEncoder(w).Encode(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleProductGetAll(w http.ResponseWriter, r *http.Request) {
    products, err := productRepository.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(products); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func HandleProductRequest(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
    switch r.Method {
    case http.MethodOptions:
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Allow", "GET, POST, PUT, DELETE, OPTIONS")
    case http.MethodPost:
        HandleProductCreate(w, r)
    case http.MethodPut:
        HandleProductUpdate(w, r)
    case http.MethodDelete:
        HandleProductDelete(w, r)
    case http.MethodGet:
        if r.URL.Path == "/product" {
            HandleProductGetAll(w, r)
        } else {
            HandleProductGetByID(w, r)
        }
    default:
        http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/user-account", HandleUserAccountRequest)
    http.HandleFunc("/product", HandleProductRequest)

    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}