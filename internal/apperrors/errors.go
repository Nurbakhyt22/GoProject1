package apperrors

//import "net/http"
//
//type AppError struct {
//	Code    int
//	Message string
//	Err     error
//}
//
//func New(code int, message string, err error) *AppError {
//	return &AppError{
//		Code:    code,
//		Message: message,
//		Err:     err,
//	}
//}
//
//func NotFound(msg string, err error) *AppError {
//	return New(http.StatusNotFound, msg, err)
//}
//
//func BadRequest(msg string, err error) *AppError {
//	return New(http.StatusBadRequest, msg, err)
//}
//
//func Internal(msg string, err error) *AppError {
//	return New(http.StatusInternalServerError, msg, err)
//}
