package category

import (
	"fmt"
)

type categoryType string

const (
	FEE_ONLY  categoryType = "fee-only"
	FEE_BASED categoryType = "fee-based"
	PMS       categoryType = "portfolio-management-services"
)

type Faq struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Category struct {
	Id            string       `json:"_id"`
	Type          categoryType `json:"type"`
	Description   string       `json:"description"`
	Advantages    []string     `json:"advantages"`
	Disadvantages []string     `json:"disadvantages"`
	Faq           []Faq        `json:"faq"`
}

func (category *Category) WriteToConsole() {
	fmt.Println(category.Type)
	fmt.Println(category.Description)
	fmt.Println("Advantages ------------>")
	for _, val := range category.Advantages {
		fmt.Println(val)
	}
	fmt.Println("Disadvantages ------------>")
	for _, val := range category.Disadvantages {
		fmt.Println(val)
	}
	fmt.Println("FAQ ------------>")
	for _, val := range category.Faq {
		fmt.Println(val.Question)
		fmt.Println(val.Answer)
	}
}
