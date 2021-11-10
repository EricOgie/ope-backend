package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
)

func BuyStock(reqObj requestdto.UserPayRequest, claim models.Claim) responsedto.FlutterWaveResponse {
	// io.Reader
	// httpReq := http.NewRequest("POST", konstants.FLUTTERWAVE_URL, )
	payReq := reqObj.MakeFlutterPayRequest(claim)
	payBody, err := json.Marshal(payReq)
	if err != nil {
		logger.Error("JSON conversion Err: " + err.Error())
		log.Fatal(err)
	}

	client := http.Client{}
	httpReq, _ := http.NewRequest("POST", konstants.FLUTTERWAVE_URL, bytes.NewBuffer(payBody))
	// Set Header
	httpReq.Header.Set("Authorization", "Bearer FLWPUBK_TEST-db5bc2dc21efad5023ae7b13aa04cd2e-X")
	httpRes, errorr := client.Do(httpReq)

	if errorr != nil {
		logger.Error("Http Req Err: " + err.Error())
		log.Fatal(err)
	}

	// close request
	defer httpRes.Body.Close()
	resBody, exErr := ioutil.ReadAll(httpRes.Body)
	// Handle error
	if exErr != nil {
		logger.Error("http Body Etxract Err: " + exErr.Error())
		log.Fatal(err)
	}

	v, _ := json.Marshal(resBody)

	fmt.Println(fmt.Sprintf("%#v", v))
	return responsedto.FlutterWaveResponse{}

}
