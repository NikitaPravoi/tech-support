package input

// IDInput is an input struct for requests using a path variable "id"
type IDInput struct {
	ID int64 `json:"id" example:"1" doc:"ID" path:"id"`
}
