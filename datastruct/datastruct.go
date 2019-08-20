package datastruct

type GetRouteRequest struct {
	ORIGIN      string `json:"origin"`
	DESTINATION string `json:"destination"`
}

type GetRouteResponse struct {
	ORIGIN        string `json:"origin"`
	DESTINATION   string `json:"destination"`
	DURATION      int    `json:"duration"`
	ID_ROUTE_SPEC int    `json:"id_route_spec"`
}

type Route struct {
	ID_ROUTE_SPEC int
	DESTINATION   string
	ORIGIN        string
	DURATION      int
}
