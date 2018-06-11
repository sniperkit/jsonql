package main

import (
	"fmt"

	"github.com/sniperkit/jsonql/pkg"
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
    "hobby": null,
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
    "hobby": "dancing",
    "skills": [
      "Eating",
      "Sleeping",
      "Crawling"
    ]
  }
]
`

// Must be set at build via -ldflags "-X main.VERSION=`cat VERSION`"
var VERSION string

func main() {

	parser, err := jsonql.NewStringQuery(jsonString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Query("name='elgs'"))
	//[map[skills:[Golang Java C] name:elgs gender:m age:35]] <nil>

	fmt.Println(parser.Query("name='elgs' && gender='f'"))
	//[] <nil>

	fmt.Println(parser.Query("age<10 || (name='enny' && gender='f')"))
	// [map[hobby:<nil> skills:[IC Electric design Verification] name:enny gender:f age:36] map[name:sam gender:m age:1 hobby:dancing skills:[Eating Sleeping Crawling]]] <nil>

	fmt.Println(parser.Query("age<10"))
	// [map[gender:m age:1 hobby:dancing skills:[Eating Sleeping Crawling] name:sam]] <nil>

	fmt.Println(parser.Query("1=0"))
	//[] <nil>

	fmt.Println(parser.Query("age=(2*3)^2"))
	//[map[skills:[IC Electric design Verification] name:enny gender:f age:36 hobby:<nil>]] <nil>

	fmt.Println(parser.Query("name ~= 'e.*'"))
	// [map[name:elgs gender:m age:35 skills:[Golang Java C]] map[hobby:<nil> skills:[IC Electric design Verification] name:enny gender:f age:36]] <nil>

	fmt.Println(parser.Query("name='el'+'gs'"))
	fmt.Println(parser.Query("age=30+5.0"))
	fmt.Println(parser.Query("age=40.0-5"))
	fmt.Println(parser.Query("age=70-5*7"))
	fmt.Println(parser.Query("age=70.0/2.0"))
	fmt.Println(parser.Query("age=71%36"))
	// [map[name:elgs gender:m age:35 skills:[Golang Java C]]] <nil>

	fmt.Println(parser.Query("hobby is defined"))
	// [map[name:enny gender:f age:36 hobby:<nil> skills:[IC Electric design Verification]] map[name:sam gender:m age:1 hobby:dancing skills:[Eating Sleeping Crawling]]] <nil>

	fmt.Println(parser.Query("hobby isnot defined"))
	// [map[name:sam gender:m age:1 skills:[Eating Sleeping Crawling]]] <nil>

	fmt.Println(parser.Query("hobby is null"))
	// [map[hobby:<nil> skills:[IC Electric design Verification] name:enny gender:f age:36]] <nil>

	fmt.Println(parser.Query("hobby isnot null"))
	// [map[name:sam gender:m age:1 hobby:dancing skills:[Eating Sleeping Crawling]]] <nil>
}
