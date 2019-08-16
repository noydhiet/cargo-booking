package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

type HelloDaerahRequest struct {
	NAME          string `json:"name"`
	JENIS_KELAMIN string `json:"jenis_kelamin"`
	ASAL_KOTA     string `json:"asal_kota"`
}

type HelloDaerahResponse struct {
	STATUS  int    `json:"code`
	MESSAGE string `json:"message"`
}

type GetItenaryRequest struct {
	VOYAGE_NUMBER   int    `json:"voyage_number"`
	UNLOAD_LOCATION string `json:"unload_location"`
	LOAD_LOCATION   string `json:"load_location"`
}

type GetItenaryResponse struct {
	STATUS          int    `json:"code`
	LOAD_LOCATION   string `json:"load_location`
	UNLOAD_LOCATION string `json:"unload_location"`
	VOYAGE_NUMBER   int    `json:"voyage_number"`
	LOAD_TIME       int    `json:"load_time"`
	UNLOAD_TIME     int    `json:"unload_time"`
}
