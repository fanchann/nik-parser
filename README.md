## NIK Parser Written in Go

### installation
```sh
go get github.com github.com/fanchann/nik-parser
```

example:
```go
package main

import (
	"encoding/json"
	"fmt"

	nikparser "github.com/fanchann/nik-parser"
)

func main() {
	nik := "3203012503770011"
	result, err := nikparser.ParseNIK(nik)
	if err != nil {
		fmt.Printf("%v \n", err)
		return
	}

	jsonByte, _ := json.Marshal(result)

	fmt.Println(string(jsonByte))
}
```

### author
- [fanchann](https://github.com/fanchann)