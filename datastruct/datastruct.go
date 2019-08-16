package datastruct

type GetRouteRequest struct {
	ORIGIN      string `json:"origin"`
	DESTINATION string `json:"destination"`
}

type GetRouteResponse struct {
	ORIGIN      string `json:"origin"`
	DESTINATION string `json:"destination"`
	DURATION    int    `json:"duration"`
}


