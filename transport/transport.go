package transport

import (
	"cargo-booking/datastruct"
	dt "cargo-booking/datastruct"

	//"cargo-booking/logging"
	"cargo-booking/service"
	"context"
	"encoding/json"
	"errors"

	//"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	// HelloWorldService(context.Context, string) string
	// HelloDaerahService(context.Context, string, string, string) string
	GetItenaryService(context.Context, dt.Tampil) []dt.Tampil
}

type aphService struct{}

var ErrEmpty = errors.New("empty string")

// // back service /Helloworld
// func (aphService) HelloWorldService(_ context.Context, name string) string {

// 	return call_ServiceHelloWorld(name)
// }
// func call_ServiceHelloWorld(name string) string {

// 	messageResponse := service.HelloWorld(name)

// 	return messageResponse

// }
// func makeHelloWorldEndpoint(aph AphService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(datastruct.HelloWorldRequest)
// 		logging.Log(fmt.Sprintf("Name Request %s", req.NAME))
// 		v := aph.HelloWorldService(ctx, req.NAME)
// 		logging.Log(fmt.Sprintf("Response Final Message %s", v))
// 		return datastruct.HelloWorldResponse{v}, nil
// 	}
// }
// func decodeHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var request datastruct.HelloWorldRequest
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		return nil, err
// 	}
// 	return request, nil
// }

// // back service /HelloDaerah
// func (aphService) HelloDaerahService(_ context.Context, name, jenis_kelamin, asal_kota string) string {

// 	return call_ServiceHelloDaerah(name, jenis_kelamin, asal_kota)
// }
// func call_ServiceHelloDaerah(name, jenis_kelamin, asal_kota string) string {

// 	messageResponse := service.HelloDaerah(name, jenis_kelamin, asal_kota)

// 	return messageResponse

// }
// func makeHelloDaerahEndpoint(aph AphService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(datastruct.HelloDaerahRequest)
// 		//logging.Log(fmt.Sprintf("Name Request %s", req.NAME))
// 		z := aph.HelloDaerahService(ctx, req.NAME, req.JENIS_KELAMIN, req.ASAL_KOTA)
// 		// logging.Log(fmt.Sprintf("Response Final Message %s", z))
// 		return datastruct.HelloDaerahResponse{200, z}, nil
// 	}
// }
// func decodeHelloDaerahRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var request datastruct.HelloDaerahRequest
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		return nil, err
// 	}
// 	return request, nil
// }

// // func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
// // 	return json.NewEncoder(w).Encode(response)
// // }

//back service /GetItenary
func (aphService) GetItenaryService(_ context.Context, tamp dt.Tampil) []dt.Tampil {

	return call_ServiceGetItenary(tamp)
}
func call_ServiceGetItenary(tamp dt.Tampil) []dt.Tampil {

	messageResponse := service.GetItenary(tamp)

	return messageResponse

}
func makeGetItenaryEndpoint(aph AphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.GetItenaryRequest)
		bismillah := dt.Tampil{}
		bismillah.VOYAGE_NUMBER = req.VOYAGE_NUMBER
		bismillah.UNLOAD_LOCATION = req.UNLOAD_LOCATION
		bismillah.LOAD_LOCATION = req.LOAD_LOCATION
		aph.GetItenaryService(ctx, bismillah)
		return datastruct.GetItenaryResponse{
			// ID_ITENARY:      1,
			// LOAD_LOCATION:   "Tanjung Priok",
			// UNLOAD_LOCATION: "Surabaya",
			// VOYAGE_NUMBER:   1234,
			// LOAD_TIME:       360,
			// UNLOAD_TIME:     360,
		}, nil
	}
}
func decodeGetItenaryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.GetItenaryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

//Register services /HelloWorld /HelloDaerah
func RegisterHttpsServicesAndStartListener() {
	aph := aphService{}

	// HelloWorldHandler := httptransport.NewServer(
	// 	makeHelloWorldEndpoint(aph),
	// 	decodeHelloWorldRequest,
	// 	encodeResponse,
	// )
	// HelloDaerahHandler := httptransport.NewServer(
	// 	makeHelloDaerahEndpoint(aph),
	// 	decodeHelloDaerahRequest,
	// 	encodeResponse,
	// )
	GetItenaryHandler := httptransport.NewServer(
		makeGetItenaryEndpoint(aph),
		decodeGetItenaryRequest,
		encodeResponse,
	)
	// http.Handle("/HelloWorld", HelloWorldHandler)
	// http.Handle("/HelloDaerah", HelloDaerahHandler)
	http.Handle("/GetItenary", GetItenaryHandler)
}
