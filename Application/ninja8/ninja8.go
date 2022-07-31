package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

type personn struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json: "Sayings"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	// Marshall
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// UnMarshall
	s := `[{"First":"James","Last":"Bond","Age":32,
	"Sayings":["Shaken, not stirred",
	"Youth is no guarantee of innovation",
	"In his majesty's royal service"]},
	{"First":"Miss","Last":"Moneypenny","Age":27,
	"Sayings":["James, it is soo good to see you",
	"Would you like me to take care of that for you, James?",
	"I would really prefer to be a secret agent myself."]},
	{"First":"M","Last":"Hmmmm","Age":54,
	"Sayings":["Oh, James. You didn't.",
	"Dear God, what has James done now?",
	"Can someone please tell me where James Bond is?"]}]`

	fmt.Println(s)

	bs = []byte(s)
	fmt.Println("%T\n", s)
	fmt.Println("%T\n", bs)

	var people []personn

	err = json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age, v.Sayings)
	}

}
