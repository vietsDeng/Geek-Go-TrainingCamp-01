package test

import (
	"fmt"
	"template/internal/myapp/service"
	"testing"
)

func TestGetLevelInfos(t *testing.T)  {
	fmt.Println(service.GetLevelInfoList())
}