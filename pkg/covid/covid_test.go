package covid

import (
	"testing"
)

func TestGetter(t *testing.T) {
	covid, err := NewCovid()
	if err != nil {
		t.Error("Error creating new Covid object")
	}
	covid.SetCovidData([]CovidDataRecord{
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
	})
	if len(covid.GetCovidData()) != 9 {
		t.Error("Expected 9 but got", len(covid.GetCovidData()))
	}
}

func TestGroupByProvince(t *testing.T) {
	covid, err := NewCovid()
	if err != nil {
		t.Error("Error creating new Covid object")
	}
	covid.SetCovidData([]CovidDataRecord{
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
	})
	province := covid.GroupByProvince()
	if province["Bangkok"] != 3 {
		t.Error("Expected 3 but got", province["Bangkok"])
	}
	if province["N/A"] != 1 {
		t.Error("Expected 1 but got", province["N/A"])
	}
}

func TestGroupByAge(t *testing.T) {
	covid, err := NewCovid()
	if err != nil {
		t.Error("Error creating new Covid object")
	}
	covid.SetCovidData([]CovidDataRecord{
		{
			Age:      []*int{nil}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{50}[0],
			Province: "Bangkok",
		},
		{
			Age:      &[]int{1}[0],
			Province: "Bangkok",
		},
	})
	age := covid.GroupByAge()
	if age["0-30"] != 2 {
		t.Error("Expected 2 but got", age["0-30"])
	}
	if age["N/A"] != 1 {
		t.Error("Expected 1 but got", age["N/A"])
	}
	if age["31-60"] != 1 {
		t.Error("Expected 1 but got", age["31-60"])
	}
}
