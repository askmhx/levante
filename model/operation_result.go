package model

const OperationResultCodeSuccess = "000000"
const OperationResultCodeFailed = "999999"

type OperationResult struct {
	Desc   string
	Code   string
	RetURL string
}
