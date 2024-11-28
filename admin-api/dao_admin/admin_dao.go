package dao_admin

type ContainerStatsDto struct {
	Id       string `json:"ID"`
	Name     string `json:"Name"`
	CPU      string `json:"CPUPerc"`
	MemPerc  string `json:"MemPerc"`
	MemUsage string `json:"MemUsage"`
}

type ContainersStatsDto []ContainerStatsDto
