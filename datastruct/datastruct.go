package datastruct

// type HelloWorldRequest struct {
// 	NAME string `json:"name"`
// }

// type HelloWorldResponse struct {
// 	STATUS  int    `json:"code"`
// 	MESSAGE string `json:"message"`
// }

// type HelloDaerahRequest struct {
// 	NAME    string `json:"name"`
// 	KELAMIN string `json:"jenis_kelamin"`
// 	ASAL    string `json:"asal_kota"`
// }

// type HelloDaerahResponse struct {
// 	STATUS  int    `json:"code"`
// 	MESSAGE string `json:"message"`
// }

type GetBookingRequest struct {
	// ROUTE_ID   int `json:"route_id"`
	// ITENARY_ID int `json:"itenary_id"`

	ID_DELIVERY int `json:"idt_delivery"`
	ROUTING_STATUS string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
	LAST_KNOWN_LOCATION string `json:"last_known_location"`
	ID_ITENARY int `json:"fk_id_itenary"`
	ID_ROUTE_SPEC int `json:"fk_id_route_spec"`

}

type GetBookingResponse struct {
	// ID_DELIVERY         int    `json:"delivery_id"`
	// ROUTE_ID            int    `json:"route_id"`
	// ITENARY_ID          int    `json:"itenary_id"`
	// ROUTING_STATUS      string `json:"routing_status"`
	// TRANSPORT_STATUS    string `json:"transport_status"`
	// LAST_KNOWN_LOCATION string `json:"last_known_location"`
	// RESPONSE_CODE       int    `json:"responseCode"`
	// RESPONSE_DESC       string `json:"responseDesc"`

	
	MESSAGE string `json:"message"`
}

type Delivery struct {
	ID_DELIVERY         int
	ROUTING_STATUS      string
	TRANSPORT_STATUS    string
	LAST_KNOWN_LOCATION string
	ID_ITENARY          int
	ID_ROUTE            int
}
