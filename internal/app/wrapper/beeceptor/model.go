package beeceptor

type GetOrganizationSubdistrictResponse struct {
	Success bool                           `json:"success"`
	Message string                         `json:"message"`
	Data    GetOrganizationSubdistrictData `json:"data"`
}

type GetOrganizationSubdistrictData struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
