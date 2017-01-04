# jsontime
An omittable timestamp for JSON in Go

# Usage
```
package main

import (
    "github.com/noxer/jsontime"
)

type MyStruct struct {
    CreatedAt jsontime.T `json:",omitempty"`
}

func main() {
	myStruct := &MyStruct{CreatedAt: jsontime.FromTime(time.Now())}
	j, err := json.Marshal(myStruct)
	fmt.Println(string(j), err)

	myStruct2 := &MyStruct{}
	j, err = json.Marshal(myStruct2)
	fmt.Println(string(j), err)
}
```
