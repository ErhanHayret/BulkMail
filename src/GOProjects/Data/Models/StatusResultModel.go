package Models

type StatusResult struct{
	Message	string 	`json:"message`
	Status 	bool	`json:"status`
	Error 	error	`json:"error`
}