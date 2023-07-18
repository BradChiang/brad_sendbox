package test_tool

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// marshal 用來包裝json
// unmarshal 用來解包json到interface{}
// json key 必為 string
func JSONtest() {
	fmt.Println("JSON test start:")
	fmt.Println("marshal:")

	bolB, _ := json.Marshal(true)
	fmt.Println(bolB, ":", string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(intB, ":", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(fltB, ":", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB, ":", string(strB))

	slcB, _ := json.Marshal([]string{"apple", "pie", "mongo"})
	fmt.Println(slcB, ":", string(slcB))

	mapB, _ := json.Marshal(map[string]int{"apple": 5, "mongo": 7})
	fmt.Println(mapB, ":", string(mapB))

	res1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "pie", "mongo"},
	}
	res1B, _ := json.Marshal(res1)
	fmt.Println(res1B, ":", string(res1B))

	res2 := &response2{
		Page:   1,
		Fruits: []string{"apple", "pie", "mongo"},
	}
	res2B, _ := json.Marshal(res2)
	fmt.Println(res2B, ":", string(res2B))

	fmt.Println("unmarshal:")
	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1, "fruits": ["apple", "banana"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0], ",", res.Fruits[1])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	fmt.Println("JSON test end")
}
