package main

import (
	"encoding/json"
	"fmt"
)

func Flatten_json_main() {

	fmt.Println("flatten json")

	jsonString := `{
		"labtest": {
			"subTest": {
				"loincCode": "Leukocytes[# / volume] in Urine"
			},
			"review": false
		},
		"patientId": 1652
	}`

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {

		fmt.Println("Error while unmarshalling")
	}
	flatMap := flattenJson(data)
	for key, val := range flatMap {
		fmt.Println(key, val)
	}
}

func flattenJson(input map[string]interface{}) map[string]interface{} {
	fmt.Printf("inside flatten structure, input: %+v \n", input)

	output := make(map[string]interface{})

	for key, val := range input {

		fmt.Println(key, val)
		switch val.(type) {
		case map[string]interface{}:
			flatten := flattenJson(val.(map[string]interface{}))
			for k, v := range flatten {
				output[key+"."+k] = v
			}
		case []interface{}:
			for i, va := range val.([]interface{}) {
				switch va.(type) {
				case map[string]interface{}:
					flatten := flattenJson(val.(map[string]interface{}))
					for k, v := range flatten {
						output[fmt.Sprintf("%s.%d.%s", key, i, k)] = v
					}
				default:
					output[fmt.Sprintf("%s.%d", key, i)] = va
				}
			}
		default:
			output[key] = val

		}
	}

	return output

}
