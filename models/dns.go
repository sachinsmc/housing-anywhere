package models

type DNS struct {
	SectorID int `json:"sector_id"`
}

type DNSRequest struct {
	X        string `json:"x"`
	Y        string `json:"y"`
	Z        string `json:"z"`
	Velocity string `json:"vel"`
}

type DNSResponse struct {
	Location string `json:"loc"`
}
