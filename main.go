package main

import (
	"fmt"
	"valid/valid"
)

type Information struct {
	Name 		string		`validate:"required" bind:"Name" message:"大啊啊大"`
	Age 		int			`validate:"required" bind:"Age" alias:"年龄"`
	Address 	string		`validate:"required" bind:"Address" alias:"地址"`
	Hobby 		[]string	`validate:"required" bind:"Hobby"`
	ID 			*ID			`validate:"required" bind:"ID"`
}

type ID struct {
	IDNumber 		string	`validate:"required" bind:"IDNumber"`
	IDName			string	`validate:"required" bind:"IDName" alias:"id名称" message:"呵呵"`
	Date 			string	`validate:"required" bind:"Date"`
	Sex 			int		`validate:"oneof=1 2" bind:"Sex"`
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
	requestId := &IDRequest{
		IDNumber: "440823",
		IDName:   "林先生",
		Date:     "2019-11-11",
		Sex:      4,
	}
	request := &InformationRequest{
		Name:    "Liam",
		Age:     0,
		Address: "中国",
		Hobby:   []string{"篮球", "弹琴"},
		ID:      requestId,
	}
	v, err := valid.NewValid(info)
	if err != nil {
		fmt.Printf("错误")
	}
	ve := v.Inject(request).Valid()
	fmt.Printf("%+v", ve)
}


