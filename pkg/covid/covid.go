package covid

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type CovidDataRecord struct {
	Age      *int   `json:"Age,omitempty"`
	Province string `json:"Province,omitempty"`
	// ProvinceId *int   `json:"ProvinceId,omitempty"`
	// ProvinceEn string `json:"ProvinceEn,omitempty"`
}

type Covid struct {
	Data []CovidDataRecord `json:"Data"`
}

type ICovid interface {
	GetCovidData() []CovidDataRecord
	FetchCovidData() error
	GroupByProvince() map[string]int
	GroupByAge() map[string]int
}

func (c *Covid) GetCovidData() []CovidDataRecord {
	return c.Data
}

func NewCovid() (*Covid, error) {
	data := &Covid{}
	err := data.FetchCovidData()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Covid) SetCovidData(payload []CovidDataRecord) {
	c.Data = payload
}

func (c *Covid) FetchCovidData() error {
	res, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		return errors.New("Error Fetch data from API " + err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error reading response body " + err.Error())
	}

	if err = json.Unmarshal(body, &c); err != nil {
		return errors.New("Error unmarshalling JSON " + err.Error())
	}
	return nil
}

func (c Covid) GroupByProvince() map[string]int {
	province := make(map[string]int)
	for _, data := range c.Data {
		if data.Province == "" {
			province["N/A"]++
		} else {
			province[data.Province]++
		}
	}
	return province
}

func (c Covid) GroupByAge() map[string]int {
	age := make(map[string]int)
	for _, data := range c.Data {
		if data.Age == nil {
			age["N/A"]++
		} else if *data.Age < 30 {
			age["0-30"]++
		} else if *data.Age >= 30 && *data.Age <= 60 {
			age["31-60"]++
		} else if *data.Age > 60 {
			age["61+"]++
		}
	}
	return age
}
