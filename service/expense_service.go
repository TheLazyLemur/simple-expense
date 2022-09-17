package service

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// TODO: Remove hardcoded organisation Id
func CreateExpense(userId int64, description string, amount int64, ctx context.Context, store db.Store) (db.Expense, error) {
	if userId == 0 {
		return db.Expense{}, errors.New("userId not provided")
	}

	createExpenseParams := db.CreateExpenseParams{
		Amount: amount,
		Owner:  userId,
	}

	expense, dbErr := store.Queries.CreateExpense(context.Background(), createExpenseParams)

	return expense, dbErr
}

func CreateInvoice(file multipart.File, handler *multipart.FileHeader, store db.Store) (db.Invoice, error) {
	// Validate file is actually an image [x]
	// Save to image to disk [x]
	// Add database entry []
	arr := [5]string{"image/jpeg", "image/png", "image/jpg", "application/pdf"}
	var result bool = false

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return db.Invoice{}, err
	}
	fileType := http.DetectContentType(fileBytes)
	println(fileType)
	for _, x := range arr {
		if x == fileType {
			result = true
			break
		}
	}
	if !result {
		return db.Invoice{}, errors.New("Not a valid file type")
	}

	defer file.Close()

	log.Println(handler.Filename)

	tempFile, err := os.CreateTemp("invoices", "upload-*.png")
	if err != nil {
		return db.Invoice{}, err
	}
	defer tempFile.Close()

	_, _ = tempFile.Write(fileBytes)

	return db.Invoice{}, nil
}
