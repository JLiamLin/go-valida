package main

import (
	"fmt"
	"valid/utils"

	"valid/valid"
)

type Information struct {
	Name 		string		`bind:"Name"`
	Age 		int			`bind:"Age"`
	Address 	string		`bind:"Address"`
	Hobby 		[]string	`bind:"Hobby"`
	ID 			*ID			`bind:"ID"`
}

type ID struct {
	IDNumber 		string	`bind:"IDNumber"`
	IDName			string	`bind:"IDName"`
	Date 			string	`bind:"Date"`
	Sex 			int		`bind:"Sex"`
}

type InformationRequest struct {
	Name 		string
	Age 		int
	Address 	string
	Hobby 		[]string
	ID 			*IDRequest
}

type IDRequest struct {
	IDNumber 		string
	IDName			string
	Date 			string
	Sex 			int
}

func main()  {
	info := &Information{ID:&ID{}}
	//info := &Information{}
	//fields := GetFieldName(info)
	//fmt.Printf("%+v\n", fields)
	//_ = validator.New().Struct(info)
	requestId := &IDRequest{
		IDNumber: "440823",
		IDName:   "林先生",
		Date:     "2019-11-11",
		Sex:      1,
	}
	request := &InformationRequest{
		Name:    "Liam",
		Age:     15,
		Address: "中国",
		Hobby:   []string{"篮球", "弹琴"},
		ID:      requestId,
	}
	utils.Bind(request, info)
	fmt.Printf("%+v\n", info)
	v, err := valid.NewValid(info)
	if err != nil {
		fmt.Printf("错误")
	}
	fmt.Printf("%+v\n", v)
	//v.Inject(reuest).Registe
	a := "标签"
	b := ""
	c := fmt.Sprintf("%s - ", a, b)
	fmt.Println(c)
}


