package api

import (
    "TheLazyLemur/simple-expense/auth"
    "TheLazyLemur/simple-expense/service"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

func (s *Server) newExpense(w http.ResponseWriter, r *http.Request) {
    userId := auth.GetClaimsProperty(r, "id").(float64)

    w.Header().Set("Content-Type", "application/json")
    reqBody, readerr := io.ReadAll(r.Body)
    if readerr != nil {
        log.Fatal(readerr)
        return
    }

    var expenseReq createExpenseRequest
    jsonErr := json.Unmarshal(reqBody, &expenseReq)
    if jsonErr != nil {
        log.Fatal(jsonErr)
        return
    }

    expense, err := service.CreateExpense(int64(userId), expenseReq.Description, expenseReq.Amount, r.Context(), *s.store)
    if err != nil {
        log.Fatal("Failed to create expense")
        return
    }

    w.WriteHeader(http.StatusCreated)
    expenseResp := createExpenseResponse{
        Amount: expense.Amount,
    }

    pl, _ := json.Marshal(expenseResp)
    _, wErr := w.Write(pl)
    if wErr != nil {
        return
    }
}

func (s *Server) newInvoice(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Uploading an invoice") 

    _ = r.ParseMultipartForm(10 << 20)

    file, handler, fErr := r.FormFile("invoice")
    if handler.Size > 10000000.00 {
        w.WriteHeader(http.StatusBadRequest)
        _, _ = w.Write([]byte("File is too large"))
        return
    }

    if fErr != nil{
        w.WriteHeader(http.StatusBadRequest)
        _, _ = w.Write([]byte(fErr.Error()))
        return
    }
    _, err := service.CreateInvoice(file, handler, *s.store)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        _, _ = w.Write([]byte(err.Error()))
        return
    }

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

func (s *Server) newExpenseWithInvoice(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte(`{"message": "new expense with invoice created"}`))
}
