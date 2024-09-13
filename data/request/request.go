package request

type CreateTagsRequest struct {
	Name string `json:"name"`
}

type UpdateTagsRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
