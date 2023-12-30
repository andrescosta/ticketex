package enum

import "strings"

type TransactionStatus uint

const (
	Validated TransactionStatus = iota
	ErrorExtProcCC
)

func ToTransactionStatus(s string) TransactionStatus {
	switch strings.ToLower(s) {
	case "validated":
		return Validated
	case "errorextproccc":
		return ErrorExtProcCC
	}
	return ErrorExtProcCC
}
