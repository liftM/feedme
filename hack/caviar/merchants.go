package caviar

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type SetLocationRequest struct {
	Address Address `json:"address"`
	// AddressOrigin     string      `json:"address_origin"`
	Geolocation Geolocation `json:"geolocation"`
	// IsCurrentLocation bool        `json:"is_current_location"`
	// Source            string      `json:"source"`
}

type Address struct {
	City          string  `json:"city"`
	Instructions  *string `json:"instructions"`
	PostalCode    string  `json:"postal_code"`
	State         string  `json:"state"`
	StreetAddress string  `json:"street_address"`
}

func (s *Session) SetLocation(addr Address) error {
	payload := SetLocationRequest{
		Address: addr,
		Geolocation: Geolocation{
			Latitude:  "37.7519437",
			Longitude: "-122.420297",
		},
	}
	res, err := s.PostJSON("/api/v1/web/set_location", payload)
	log.Println(res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
	return err
}

type MerchantListing struct {
	Merchants []Merchant
}

type Merchant struct {
	DeliveryFeeIsLow              bool         `json:"deliveryFeeIsLow"`
	DeliveryFee                   string       `json:"delivery_fee"`
	DeliveryFeeText               string       `json:"delivery_fee_text"`
	Description                   string       `json:"description"`
	DistanceToMerchant            string       `json:"distance_to_merchant"`
	DistanceToMerchantText        string       `json:"distance_to_merchant_text"`
	DomID                         string       `json:"dom_id"`
	FulfillmentTypesAllowed       []string     `json:"fulfillment_types_allowed"`
	Geolocation                   *Geolocation `json:"geolocation"`
	ID                            int          `json:"id"`
	ImageSet                      []Image      `json:"image_set"`
	IsFavorite                    bool         `json:"is_favorite"`
	Labels                        [][2]string  `json:"labels"`
	MerchantGeolocation           *Geolocation `json:"merchant_geolocation"`
	Name                          string       `json:"name"`
	Neighborhood                  string       `json:"neighborhood"`
	OpenAtText                    string       `json:"open_at_text"`
	PriceCategory                 string       `json:"price_category"`
	ShouldShowAJAXETAs            bool         `json:"should_show_ajax_etas"`
	ShouldShowMerchantUnavailable bool         `json:"should_show_merchant_unavailable"`
	ShouldShowOpenNow             bool         `json:"should_show_open_now"`
	URL                           string       `json:"url"`
}

type Image struct {
	Width      int    `json:"width"`
	LowDPIURL  string `json:"lodpi"`
	HighDPIURL string `json:"hidpi"`
}

type Geolocation struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

func (s *Session) Merchants() (MerchantListing, error) {
	res, err := s.GetJSON("/san-francisco")
	if err != nil {
		return MerchantListing{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return MerchantListing{}, err
	}

	var listing MerchantListing
	err = json.Unmarshal(body, &listing)
	if err != nil {
		return MerchantListing{}, err
	}

	return listing, nil
}
