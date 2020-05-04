package geo

import (
	"fmt"
	"net/http"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/fw/logger"
)

var _ Geo = (*IPStack)(nil)

// Here is IPStack's documentation: https://ipstack.com/documentation
const baseURL = "http://api.ipstack.com"

type jsonLanguage struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type jsonLocation struct {
	Capital     string         `json:"capital"`
	Languages   []jsonLanguage `json:"languages"`
	CallingCode string         `json:"calling_code"`
	IsEU        bool           `json:"is_eu"`
}

type jsonTimeZone struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	IsDaylightSaving bool   `json:"is_daylight_saving"`
}

type jsonCurrency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type jsonResponse struct {
	ContinentCode string       `json:"continent_code"`
	ContinentName string       `json:"continent_name"`
	CountryCode   string       `json:"country_code"`
	CountryName   string       `json:"country_name"`
	RegionCode    string       `json:"region_code"`
	RegionName    string       `json:"region_name"`
	City          string       `json:"city"`
	ZipCode       string       `json:"zip"`
	Longitude     float64      `json:"longitude"`
	Latitude      float64      `json:"latitude"`
	Location      jsonLocation `json:"location"`
	TimeZone      jsonTimeZone `json:"time_zone"`
	Currency      jsonCurrency `json:"currency"`
}

type IPStack struct {
	apiKey      string
	httpRequest fw.HTTPRequest
	logger      logger.Logger
}

func (I IPStack) GetLocation(ipAddress string) (Location, error) {
	url := fmt.Sprintf("%s/%s?access_key=%s", baseURL, ipAddress, I.apiKey)
	res := jsonResponse{}
	err := I.httpRequest.JSON(http.MethodGet, url, map[string]string{}, "", &res)
	if err != nil {
		I.logger.Error(err)
		return Location{}, err
	}

	var languages []Language
	for _, jsonLanguage := range res.Location.Languages {
		language := Language{
			Code: jsonLanguage.Code,
			Name: jsonLanguage.Name,
		}
		languages = append(languages, language)
	}

	return Location{
		Continent: Continent{
			Code: res.ContinentCode,
			Name: res.ContinentName,
		},
		Country: Country{
			Code: res.CountryCode,
			Name: res.CountryName,
		},
		Region: Region{
			Code: res.RegionCode,
			Name: res.RegionName,
		},
		City: res.City,
		Currency: Currency{
			Code: res.Currency.Code,
			Name: res.Currency.Name,
		},
		Languages:       languages,
		IsEuropeanUnion: res.Location.IsEU,
	}, nil
}

func NewIPStack(apiKey string, httpRequest fw.HTTPRequest, logger logger.Logger) IPStack {
	return IPStack{
		apiKey:      apiKey,
		httpRequest: httpRequest,
		logger:      logger,
	}
}
