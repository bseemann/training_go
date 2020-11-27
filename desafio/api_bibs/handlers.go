package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome!")
}

func AttemptCapture(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", ContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	capt := Payment{}
	if err1 := json.NewDecoder(r.Body).Decode(&capt); err1 != nil {
		fmt.Fprintln(w, "error:", err1)
	}
	if strings.Contains(capt.Name,Reject) {
		resp := DeclinedResponse{"declined","An error occurred"}
		b, _ := json.Marshal(&resp)
		w.Write(b)
	} else {
		resp := SuccessResponse{"success", "000", "0"}
		b, _ := json.Marshal(&resp)
		w.Write(b)
	}
}

func Refund(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", ContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	capt := RefundPayment{}
	if err1 := json.NewDecoder(r.Body).Decode(&capt); err1 != nil {
		fmt.Fprintln(w, "error:", err1)
	}
	if strings.Contains(capt.Name,Reject) {
		resp := DeclinedResponse{"declined","An error occurred"}
		b, _ := json.Marshal(&resp)
		w.Write(b)
	} else {
		resp := SuccessResponse{"success", "001", "0"}
		b, _ := json.Marshal(&resp)
		w.Write(b)
	}
}