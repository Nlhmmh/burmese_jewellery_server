package handler

type GetResp struct {
	Count int64 `json:"count"`
	Data  any   `json:"data"`
}
