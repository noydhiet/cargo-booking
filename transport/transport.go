package transport

import (
	dt "cargo-booking/datastruct"
	svc "cargo-booking/service"
	"context"
	"encoding/json"
	"log"

	//"kit/endpoint"
	"net/http"
	_ "github.com/go-sql-driver/mysql" 
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	GetBookingService(context.Context, dt.Delivery) []dt.Delivery
}

type aphService struct {
}

func (aphService) GetBookingService(_ context.Context, del dt.Delivery) []dt.Delivery {
	return call_ServiceBookingService(del)
}

func call_ServiceBookingService(del dt.Delivery) []dt.Delivery {
	retDel := svc.InsertDelivery(del)

	return retDel
}

func makeBookingEndpoint(aph AphService) endpoint.Endpoint {
	log.Println("Process")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dt.GetBookingRequest)
		paramDel := dt.Delivery{}
		paramDel.ID_DELIVERY = req.ID_DELIVERY
		paramDel.ROUTING_STATUS = req.ROUTING_STATUS
		paramDel.TRANSPORT_STATUS = req.TRANSPORT_STATUS
		paramDel.LAST_KNOWN_LOCATION = req.LAST_KNOWN_LOCATION
		paramDel.ID_ITENARY = req.ID_ITENARY
		paramDel.ID_ROUTE = req.ID_ROUTE_SPEC
		aph.GetBookingService(ctx, paramDel)
		return dt.GetBookingResponse{
			// ID_DELIVERY:         1,
			// ROUTE_ID:            1,
			// ITENARY_ID:          1,
			// ROUTING_STATUS:      "unload",
			// TRANSPORT_STATUS:    "inport",
			// LAST_KNOWN_LOCATION: "surabaya",
			// RESPONSE_CODE:       200,
			// RESPONSE_DESC:       "Success",

			MESSAGE: "Succes",
		}, nil
	}
}

func decodeBookingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request dt.GetBookingRequest
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

	GetBookingHandler := httptransport.NewServer(
		makeBookingEndpoint(aph),
		decodeBookingRequest,
		encodeResponse,
	)
	//url path of our API service
	http.Handle("/bookingcargo", GetBookingHandler)

}
