package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "banana", "mango"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "banana", "mango"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "banana", "mango"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	var data map[string]interface{}

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)

	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	str2 := strs[1].(string)
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)

	str := `{"page": 1, "fruits": ["apple", "banana", "mango"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	fmt.Println(res.Fruits[1])
	fmt.Println(res.Fruits[2])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
