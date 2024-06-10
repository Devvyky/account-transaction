package api

import (
	"database/sql"
	"net/http"

	db "github.com/devvyky/account-transaction/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	PURCHASE             = "PURCHASE"
	INSTALLMENT_PURCHASE = "INSTALLMENT PURCHASE"
	WITHDRAWAL           = "WITHDRAWAL"
	PAYMENT              = "PAYMENT"
)

type createTransactionParams struct {
	AccountID       int64   `json:"account_id" binding:"required"`
	OperationTypeID int64   `json:"operation_type_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required,gt=0"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validateAccount(ctx, req.AccountID) {
		return
	}
	operationTypeDesc, ok := server.validateOperationType(ctx, req.OperationTypeID)
	if !ok {
		return
	}
	if operationTypeDesc != PAYMENT {
		req.Amount = -req.Amount
	}

	arg := db.CreateTransactionParams{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
	}

	txn, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, txn)
}

func (server *Server) validateAccount(ctx *gin.Context, accountId int64) bool {
	_, err := server.store.GetAccount(ctx, accountId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	return true
}

func (server *Server) validateOperationType(ctx *gin.Context, operationTypeId int64) (string, bool) {
	operationType, err := server.store.GetOperationType(ctx, operationTypeId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "operation type not found"})
			return "", false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return "", false
	}
	return operationType.Description, true
}
