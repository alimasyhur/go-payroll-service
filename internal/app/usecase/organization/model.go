package organization

type CreateUpdateOrganizationRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type GetAllOrganizationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type OrganizationResponse struct {
	ID           uint                `json:"id"`
	Name         string              `json:"name"`
	Location     string              `json:"location"`
	Subdisctrict *SubdisctrictDetail `json:"subdistrict,omitempty"`
}

type SubdisctrictDetail struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
