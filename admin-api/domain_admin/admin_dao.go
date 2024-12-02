package domain_admin

type ServicesResponse struct {
	Services []Service `json:"services"`
}

type Service struct {
	Name       string   `json:"name"`
	Containers []string `json:"containers"`
}
