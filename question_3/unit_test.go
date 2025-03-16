package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dstrate/7-solution-backend-challenge-3/services"
	env "github.com/joho/godotenv"
)

const (
	fakeBaconAPI = "https://baconipsum123456.com/api/?type=meat-and-filler&paras=99&format=text"
	mockResponse = `
	["Pancetta turducken chislic boudin, voluptate tongue est tenderloin cow.  
	Kevin sint mollit, beef ribs shankle irure prosciutto short ribs fatback excepteur veniam tenderloin tongue.  
	Id cow flank, alcatra rump bacon deserunt incididunt venison porchetta meatloaf irure occaecat pancetta.  
	Pig adipisicing pork chop aliqua ut landjaeger leberkas deserunt meatball ham hock.  Occaecat cow kielbasa enim.

	Lorem swine nostrud, in anim pastrami non.  
	Sint rump venison pork chop elit prosciutto fatback ut bacon shank ground round.  
	Ut beef pork burgdoggen chislic sunt t-bone incididunt.  
	Lorem landjaeger minim, drumstick laboris veniam turducken short loin officia doner bacon enim pariatur.  
	Corned beef meatloaf ullamco meatball, id salami cillum nostrud."]`
)

var expectedResult1 = map[string]int64{
	"pancetta":     2,
	"turducken":    2,
	"chislic":      2,
	"boudin":       1,
	"tongue":       2,
	"tenderloin":   2,
	"cow":          3,
	"beef":         2,
	"ribs":         1,
	"shankle":      1,
	"prosciutto":   2,
	"short ribs":   1,
	"fatback":      2,
	"flank":        1,
	"alcatra":      1,
	"rump":         2,
	"bacon":        3,
	"venison":      2,
	"porchetta":    1,
	"meatloaf":     2,
	"pig":          1,
	"pork":         3,
	"chop":         2,
	"landjaeger":   2,
	"leberkas":     1,
	"meatball":     2,
	"ham hock":     1,
	"swine":        1,
	"pastrami":     1,
	"shank":        1,
	"ground round": 1,
	"burgdoggen":   1,
	"t-bone":       1,
	"drumstick":    1,
	"short loin":   1,
	"doner":        1,
	"corned beef":  1,
	"salami":       1,
}

func mockBaconAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mockResponse))
}

func TestGetBeefSummary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockBaconAPI))
	defer server.Close()

	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	testCases := []struct {
		datasource string
		expected   map[string]int64
	}{
		{server.URL, expectedResult1},
		{fakeBaconAPI, nil},
	}
	for _, tc := range testCases {
		beef := services.GetNewBeefService()
		beef.SetBeefDataSource(tc.datasource)
		countBeefs, _ := beef.GetBeefSummary()
		if countBeefs != nil && tc.expected != nil {
			for k, v := range tc.expected {
				if countBeefs[k] != v {
					t.Errorf("GetBeefSummary(%s) = %v; want %v", k, countBeefs[k], v)
				}
			}
		} else if (countBeefs == nil && tc.expected != nil) || (countBeefs != nil && tc.expected == nil) {
			t.Errorf("GetBeefSummary = %v; want %v", countBeefs, tc.expected)
		}
	}
}
