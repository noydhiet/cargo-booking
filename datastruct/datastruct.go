package datastruct

type GetRouteRequest struct {
	ORIGIN      string `json:"origin"`
	DESTINATION string `json:"destination"`
}

type GetRouteResponse struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Duration      int    `json:"duration"`
	Id_route_spec int    `json:"id_route_spec"`
}

type Route struct {
	Id_route_spec int
	Destination   string
	Origin        string
	Duration      int
}
