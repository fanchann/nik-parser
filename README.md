## NIK Parser Written in Go

### installation
```sh
go get github.com github.com/fanchann/nik-parser
```

example:

_nik_parser_.go
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

_nik_gen.go_

```go
package main

import (
	"fmt"

	nikparser "github.com/fanchann/nik-parser"
)

func main() {
	nikOpts := &nikparser.NIKOpts{District: "BEKASI TIMUR", PostalCode: "17111",Ttl: 20}
	niks := nikparser.NIKGen(nikOpts)
	fmt.Printf("niks: %v\n", niks)

	jsonByte, _ := json.Marshal(&niks)

	ioutil.WriteFile("nik.json", jsonByte, fs.ModePerm)
}
```

### author
- [fanchann](https://github.com/fanchann)