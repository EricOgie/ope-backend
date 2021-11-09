package handlers

import (
	"fmt"
	"net/http"

	"github.com/EricOgie/ope-be/logger"
)

func ManageCors(res *http.ResponseWriter) {
	logger.Info("CORS Place1")
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT")
	(*res).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	hed := (*res).Header()
	logger.Info("CORS Place2")
	fmt.Println(fmt.Sprintf("%#v", hed))

}
