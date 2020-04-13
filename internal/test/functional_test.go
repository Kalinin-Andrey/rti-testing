package test

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Kalinin-Andrey/rti-testing/pkg/config"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/session"

	"github.com/Kalinin-Andrey/rti-testing/internal/app/api"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/offer"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/user"
	"github.com/Kalinin-Andrey/rti-testing/internal/test/mock"
)

type Client struct {
	client		*http.Client
	token		string
}

type App struct {
	app				*api.App
	sessionRepo		*mock.SessionRepository
	userRepo		*mock.UserRepository
	productRepo		*mock.ProductRepository
	componentRepo	*mock.ComponentRepository
	priceRepo		*mock.PriceRepository
	ruleRepo		*mock.RuleApplicabilityRepository
}

var (
	c = Client{
		client:	&http.Client{},
	}
	a		App
	passhash, _ = hex.DecodeString("3a73acfdb534ddded4c0109383ee3e5a66314113d1ff691aaf4b3ee073c8fc2edd06d48f0555ec3783f4c479994e3eee3433734c29b05f08be0e9739b956b88d8fe872bd0a0942214e94fd4001e757fa3b66a2b9925de2e800c55ef49baa4c03")
	u		= &user.User{
		ID:			1,
		Name:		"demo1",
		Passhash:	string(passhash),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
	}
	conditions = []condition.Condition{
		{
			RuleName: "technology",
			Value:    "xpon",
		},
		{
			RuleName: "internetSpeed",
			Value:    "200",
		},
	}
	expectedOffer = offer.Offer{
		Product: product.Product{
			Name: "Игровой",
			Components: []component.Component{{
				IsMain: true,
				Name:   "Интернет",
				Prices: []price.Price{{
					Cost: 765,
					Type: price.TypeCost,
				}},
			}},
		},
		TotalCost: price.Price{
			Cost: 765,
			Type: price.TypeCost,
		},
	}
)

func start() *httptest.Server {
	if a.app == nil {
		os.Chdir("../..")	//
		cfg, err := config.Get()
		if err != nil {
			log.Fatalln("Can not load the config")
		}
		cfg.Log.OutputPaths = []string{"stdout", "log/test_app.log"}

		a.app = api.New(NewCommonApp(*cfg), *cfg)

		a.sessionRepo	= a.app.SessionRepository.(*mock.SessionRepository)
		a.userRepo		= a.app.Domain.User.Repository.(*mock.UserRepository)
		a.productRepo	= a.app.Domain.Product.Repository.(*mock.ProductRepository)
		a.componentRepo	= a.app.Domain.Component.Repository.(*mock.ComponentRepository)
		a.priceRepo		= a.app.Domain.Price.Repository.(*mock.PriceRepository)
		a.ruleRepo		= a.app.Domain.RuleApplicability.Repository.(*mock.RuleApplicabilityRepository)

		a.userRepo.Response.First.Entity	= u
		a.userRepo.Response.First.Err		= nil
		a.userRepo.Response.Get.Entity		= u
		a.userRepo.Response.Get.Err			= nil
		a.userRepo.Response.Query.List		= []user.User{*u}
		a.userRepo.Response.Query.Err		= nil

		a.sessionRepo.Response.GetByUserID.Entity	= &session.Session{
			ID:			1,
			UserID:		1,
			User:		*u,
		}
		a.sessionRepo.Response.GetByUserID.Err	= nil
	}
	return httptest.NewServer(a.app.Server.Handler)
}

func TestRegister(t *testing.T) {
	caseName := "Register"
	var result		interface{}
	expectedStatus := http.StatusCreated

	ts := start()

	a.userRepo.Response.Create.Err	= nil
	reqBody := strings.NewReader(`{
	"username": "demo",
	"password": "demo"
}`)
	uri := "/api/register"

	req, _ := http.NewRequest(http.MethodPost, ts.URL + uri, reqBody)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		t.Errorf("[%s] request error: %v", caseName, err)
		return
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != expectedStatus {
		t.Errorf("[%s] expected http status %v, got %v", caseName, expectedStatus, resp.StatusCode)
	}

	err = json.Unmarshal(resBody, &result)
	if err != nil {
		t.Errorf("[%s] cant unpack json: %v", caseName, err)
		return
	}

	m, ok := result.(map[string]interface{})
	if !ok {
		t.Errorf("[%s] result %v cant assign to map[string]interface{}", caseName, result)
		return
	}

	if _, ok := m["token"]; !ok {
		t.Errorf("[%s] result %v do not contain a token", caseName, m)
		return
	}

	ts.Close()
}

func TestLogin(t *testing.T) {
	caseName := "Login"
	var result		interface{}
	expectedStatus := http.StatusOK
	ts := start()

	reqBody := strings.NewReader(`{
	"username": "demo1",
	"password": "demo1"
}`)
	uri := "/api/login"

	req, _ := http.NewRequest(http.MethodPost, ts.URL + uri, reqBody)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		t.Errorf("[%s] request error: %v", caseName, err)
		return
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != expectedStatus {
		t.Errorf("[%s] expected http status %v, got %v", caseName, expectedStatus, resp.StatusCode)
	}

	err = json.Unmarshal(resBody, &result)
	if err != nil {
		t.Errorf("[%s] cant unpack json: %v", caseName, err)
		return
	}

	m, ok := result.(map[string]interface{})
	if !ok {
		t.Errorf("[%s] result %v cant assign to map[string]interface{}", caseName, result)
		return
	}

	token, ok := m["token"]
	if !ok {
		t.Errorf("[%s] result %v do not contain a token", caseName, m)
		return
	}

	if c.token, ok = token.(string); !ok {
		t.Errorf("[%s] can not assign to string token %v", caseName, token)
		return
	}

	ts.Close()
}


func TestCalculateOffer(t *testing.T) {
	caseName :=		"CalculateOffer"
	var result		interface{}
	var expected	interface{}
	ts := start()

	uri := "/api/product/1/offer"
	reqBody := strings.NewReader(`[
		{
			"RuleName": "technology",
			"Value":    "xpon"
		},
		{
			"RuleName": "internetSpeed",
			"Value":    "200"
		}
]`)
	expectedStatus := http.StatusOK

	a.productRepo.Response.Get.Entity	= &p
	a.productRepo.Response.Get.Err		= nil
	expectedData	:= expectedOffer

	req, _ := http.NewRequest(http.MethodPost, ts.URL + uri, reqBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + c.token)
	resp, err := c.client.Do(req)
	if err != nil {
		t.Errorf("[%s] request error: %v", caseName, err)
		return
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != expectedStatus {
		t.Errorf("[%s] expected http status %v, got %v", caseName, expectedStatus, resp.StatusCode)
	}

	err = json.Unmarshal(resBody, &result)
	if err != nil {
		t.Errorf("[%s] cant unpack json: %v", caseName, err)
		return
	}

	jsonData, err := json.Marshal(expectedData)
	json.Unmarshal(jsonData, &expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("[%s] results not match\nGot: %#v\nExpected: %#v", caseName, result, expectedData)
	}

	ts.Close()
}


