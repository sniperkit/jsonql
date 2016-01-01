# jsonql
SQL like JSON query library in Golang.

This library enables query against JSON using SQL like WHERE clause. Currently supported operators are: (precedences from low to high, case insensitive)

```
OR
AND
= != > < >= <= RLIKE
+ -
* / %
^
( )
```

Please note I decided not to implement `LIKE` or `NOTE LIKE` operations because `RLIKE` is enough to handle the cases `LIKE` and `NOT LIKE` handle. 

## Install
`go get -u github.com/elgs/jsonql`

## TDOD
* Implement `IS NULL` and `IS NOT NULL`

## Example
```go
package main

import (
	"fmt"
	"github.com/elgs/jsonql"
)

var jsonString = `
[
  {
    "name": "elgs",
    "gender": "m",
	"age": 35,
    "skills": [
      "Golang",
      "Java",
      "C"
    ]
  },
  {
    "name": "enny",
    "gender": "f",
    "age": 36,
	"skills": [
      "IC",
      "Electric design",
      "Verification"
    ]
  },
  {
    "name": "sam",
    "gender": "m",
	"age": 1,
    "skills": [
      "Eating",
      "Sleeping",
      "Crawling"
    ]
  }
]
`

func main() {
	parser, err := jsonql.NewStringQuery(jsonString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Query("name='elgs'"))
	//[map[skills:[Golang Java C] name:elgs gender:m age:35]] <nil>

	fmt.Println(parser.Query("name='elgs' and gender='f'"))
	//[] <nil>

	fmt.Println(parser.Query("age<10 or (name='enny' and gender='f')"))
	//[map[gender:f age:36 skills:[IC Electric design Verification] name:enny] map[age:1 skills:[Eating Sleeping Crawling] name:sam gender:m]] <nil>

	fmt.Println(parser.Query("age<10"))
	//[map[name:sam gender:m age:1 skills:[Eating Sleeping Crawling]]] <nil>

	fmt.Println(parser.Query("1=0"))
	//[] <nil>

	fmt.Println(parser.Query("age=(2*3)^2"))
	//[map[skills:[IC Electric design Verification] name:enny gender:f age:36]] <nil>
	
	fmt.Println(parser.Query("name RLIKE 'e.*'"))
	//[map[age:35 skills:[Golang Java C] name:elgs gender:m] map[skills:[IC Electric design Verification] name:enny gender:f age:36]] <nil>
	
	fmt.Println(parser.Query("name='el'+'gs'"))
	fmt.Println(parser.Query("age=30+5.0"))
	fmt.Println(parser.Query("age=40.0-5"))
	fmt.Println(parser.Query("age=70-5*7"))
	fmt.Println(parser.Query("age=70.0/2.0"))
	fmt.Println(parser.Query("age=71%36"))
	// [map[name:elgs gender:m age:35 skills:[Golang Java C]]] <nil>
}
```