package models

type EngagementRatings struct {
	Followers int     `json:"followers"`
	Rating    float64 `json:"rating"`
	Likes     int     `json:"likes"`
	Comments  int     `json:"comments"`
}

type Response struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    *EngagementRatings `json:"data"`
}
