package api

import "net/http"

func (s *Server) newExpense(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte(`{"message": "new expense created"}`))
}

func (s *Server) newInvoice(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte(`{"message": "new invoice created"}`))
}

func (s *Server) getExpense(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte(`{"message": "get expense"}`))
}

func (s *Server) getInvoice(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte(`{"message": "get invoice"}`))
}
