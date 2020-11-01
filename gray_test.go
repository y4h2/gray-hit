package grayhit

import (
	"encoding/json"
	"fmt"
)

type SampleGrayPoint struct {
}

func (point SampleGrayPoint) GetName() string {
	return "test"
}
func (point SampleGrayPoint) GetValue() string {
	return "test"
}

type SampleDiv struct {
}

func (div SampleDiv) CalcIndicator() int {
	return 10
}

func ExampleABTestPolicy() {
	sampleJSON := ` {
    "layer": {
      "id": "layer1",
      "data": "something1"
    },
    "grayRules": [
      {
        "name": "source",
        "enabled": true,
        "include": [
          "A"
        ],
        "exclude": [],
        "global": false
      },
      {
        "name": "city",
        "enabled": true,
        "include": [
          "C1",
          "C2",
          "C3"
        ],
        "exclude": [],
        "global": false
      }
    ],
    "divRule": {
      "percent": 10
    }
	}`
	policy := &ABTestPolicy{}
	json.Unmarshal([]byte(sampleJSON), policy)

	sampleGrayPoint := &SampleGrayPoint{}

	sampleDiv := &SampleDiv{}

	hitGray := policy.HitGray(sampleGrayPoint)
	hitDiv, _ := policy.HitDiv(sampleDiv)

	fmt.Println(hitGray)
	fmt.Println(hitDiv)

	// Output:
	// false
	// true
}
