package test

import (
	"fmt"
	"joke-go/utils"
	"testing"
)

func TestAddJsonFormTag(t *testing.T) {
	rs := utils.AddJsonFormGormTag(`
		type author struct {
	Id   nulls.UUID
	Name string
	sex  string
}
	`)
	fmt.Println(rs)
}
