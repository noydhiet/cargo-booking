package transport

import (
	"cargo-booking/datastruct"
	dt "cargo-booking/datastruct"
	svc "cargo-booking/service"
	"context"
	"encoding/json"
	"log"

	//"kit/endpoint"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	GetRouteService(context.Context, dt.Route) []dt.Route
}

type aphService struct {
}

func (aphService) GetRouteService(_ context.Context, del dt.Route) []dt.Route {
	return call_ServiceGetRouteService(del)
}

func call_ServiceGetRouteService(del dt.Route) []dt.Route {
	retDel := svc.GetRoute(del)

	return retDel
}

func makeGetRouteEndpoint(aph AphService) endpoint.Endpoint {
	log.Println("Process")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.GetRouteRequest)
		paramDel := dt.Route{}
		paramDel.ORIGIN = req.ORIGIN
		paramDel.DESTINATION = req.DESTINATION
		aph.GetRouteService(ctx, paramDel)
		return datastruct.GetRouteResponse{
			ID_ROUTE_SPEC: 1,
			ORIGIN:        "Jakarta",
			DESTINATION:   "Surabaya",
			DURATION:      10,
		}, nil
	}
}

func decodeGetRouteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.GetRouteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	aph := aphService{}

	GetRouteHandler := httptransport.NewServer(
		makeGetRouteEndpoint(aph),
		decodeGetRouteRequest,
		encodeResponse,
	)
	//url path of our API service
	http.Handle("/GetRoute", GetRouteHandler)

}
