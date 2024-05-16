package main

import (
 "fmt"
 "net/http"
 "strconv"
)

func modHandler(w http.ResponseWriter, r *http.Request) {
 operand1, err1 := strconv.ParseFloat(r.URL.Query().Get("operand1"), 64)
 operand2, err2 := strconv.ParseFloat(r.URL.Query().Get("operand2"), 64)
 if err1 != nil || err2 != nil {
 http.Error(w, "Invalid operands", http.StatusBadRequest)
 return
 }
 result := operand1 % operand2
 fmt.Fprintf(w, "%f", result)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
 operand1, err1 := strconv.ParseFloat(r.URL.Query().Get("operand1"), 64)
 operand2, err2 := strconv.ParseFloat(r.URL.Query().Get("operand2"), 64)
 if err1 != nil || err2 != nil {
 http.Error(w, "Invalid operands", http.StatusBadRequest)
 return
 }
 result := operand1 / operand2
 fmt.Fprintf(w, "%f", result)
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
 operand1, err1 := strconv.ParseFloat(r.URL.Query().Get("operand1"), 64)
 operand2, err2 := strconv.ParseFloat(r.URL.Query().Get("operand2"), 64)
 if err1 != nil || err2 != nil {
 http.Error(w, "Invalid operands", http.StatusBadRequest)
 return
 }
 result := operand1 * operand2
 fmt.Fprintf(w, "%f", result)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
 operand1, err1 := strconv.ParseFloat(r.URL.Query().Get("operand1"), 64)
 operand2, err2 := strconv.ParseFloat(r.URL.Query().Get("operand2"), 64)
 if err1 != nil || err2 != nil {
 http.Error(w, "Invalid operands", http.StatusBadRequest)
 return
 }
 result := operand1 - operand2
 fmt.Fprintf(w, "%f", result)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
 operand1, err1 := strconv.ParseFloat(r.URL.Query().Get("operand1"), 64)
 operand2, err2 := strconv.ParseFloat(r.URL.Query().Get("operand2"), 64)
 if err1 != nil || err2 != nil {
 http.Error(w, "Invalid operands", http.StatusBadRequest)
 return
 }
 result := operand1 + operand2
 fmt.Fprintf(w, "%f", result)
}

func main() {
 http.HandleFunc("/add", addHandler)
 http.HandleFunc("/sub", subHandler)
 http.HandleFunc("/mul", mulHandler)
 http.HandleFunc("/div", divHandler)
 http.HandleFunc("/mod", modHandler)
 fmt.Println("Starting server on :8080")
 http.ListenAndServe(":8080", nil)
}
