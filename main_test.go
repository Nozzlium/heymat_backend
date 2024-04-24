package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/nozzlium/heymat_backend/auth"
	"github.com/nozzlium/heymat_backend/budget"
	"github.com/stretchr/testify/assert"
)

var (
	DB    *sql.DB
	M     *migrate.Migrate
	app   *fiber.App
	token string
)

var budgetPlanItems = []budget.BudgetPlanRequestBody{
	{
		Title:   "Test1",
		Amount:  100000,
		Private: false,
	},
	{
		Title:   "Test2",
		Amount:  200000,
		Private: false,
	},
	{
		Title:   "Test3",
		Amount:  300000,
		Private: false,
	},
}

func initApp() {
	app = fiber.New(fiber.Config{})
	authRoute, _ := auth.Init(
		auth.Config{
			DB: DB,
		},
	)

	budgetRoute, _ := budget.Init(
		budget.Config{
			DB:             DB,
			AuthMiddleware: auth.AuthMiddleware,
		},
	)

	app.Mount("", authRoute)
	app.Mount(
		"/api/budget",
		budgetRoute,
	)
}

func initTest() {
	config := GetTestConfig()
	DB, M = InitDB(&config)
	initApp()
}

func TestMain(m *testing.M) {
	initTest()
	Truncate(M)
	Migrate(M, 0)
	m.Run()
	Truncate(M)
}

func TestRegister(
	t *testing.T,
) {
	registerBody := auth.RegisterRequestBody{
		Username: "testusername",
		Password: "testpassword",
		Email:    "testemail",
	}
	registerBodyBytes, _ := json.Marshal(
		&registerBody,
	)
	reqRegister := httptest.NewRequest(
		"POST",
		"/api/register",
		bytes.NewReader(
			registerBodyBytes,
		),
	)
	reqRegister.Header.Add(
		"content-type",
		"application/json",
	)
	respRegister, _ := app.Test(
		reqRegister,
	)
	assert.Equal(
		t,
		200,
		respRegister.StatusCode,
	)

	registerRespBytes, _ := io.ReadAll(
		respRegister.Body,
	)
	var respBodyRegister auth.RegisterResponseBody
	json.Unmarshal(
		registerRespBytes,
		&respBodyRegister,
	)
	assert.Equal(
		t,
		registerBody.Email,
		respBodyRegister.Data.Email,
	)
	assert.Equal(
		t,
		registerBody.Username,
		respBodyRegister.Data.Username,
	)
}

func TestLogin(t *testing.T) {
	loginReqBody := auth.LoginRequestBody{
		Identity: "testusername",
		Password: "testpassword",
	}
	loginReqBodyBytes, _ := json.Marshal(
		&loginReqBody,
	)
	loginReq := httptest.NewRequest(
		"POST",
		"/api/login",
		bytes.NewReader(
			loginReqBodyBytes,
		),
	)
	loginReq.Header.Add(
		"content-type",
		"application/json",
	)
	loginResp, _ := app.Test(loginReq)
	assert.Equal(
		t,
		200,
		loginResp.StatusCode,
	)

	var loginRespBody auth.LoginResponseBody
	loginRespBytes, _ := io.ReadAll(
		loginResp.Body,
	)
	json.Unmarshal(
		loginRespBytes,
		&loginRespBody,
	)
	assert.NotEqual(
		t,
		"",
		loginRespBody.Data.Token,
	)
	token = loginRespBody.Data.Token
}

func TestCreateBudgetPlan(
	t *testing.T,
) {
	for _, budgetPlanItem := range budgetPlanItems {
		tempNewBudgetBytes, _ := json.Marshal(
			&budgetPlanItem,
		)
		newBudgetRequest := httptest.NewRequest(
			"POST",
			"/api/budget",
			bytes.NewReader(
				tempNewBudgetBytes,
			),
		)
		newBudgetRequest.Header.Add(
			"content-type",
			"application/json",
		)
		newBudgetRequest.Header.Add(
			"Authorization",
			"Bearer "+token,
		)
		resp, err := app.Test(
			newBudgetRequest,
		)
		if err != nil {
			t.Error(err)
		}

		respBytes, err := io.ReadAll(
			resp.Body,
		)
		if err != nil {
			t.Error(err)
		}

		var budgetPlanResponseBody budget.BudgetPlanResponseBody
		err = json.Unmarshal(
			respBytes,
			&budgetPlanResponseBody,
		)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(
			t,
			budgetPlanItem.Title,
			budgetPlanResponseBody.Data.Title,
		)
		assert.Equal(
			t,
			budgetPlanItem.Amount,
			budgetPlanResponseBody.Data.Amount,
		)
	}
}

func TestGetBudgetPlanById(
	t *testing.T,
) {
	for i := range budgetPlanItems {
		getPlanByIdReq := httptest.NewRequest(
			"GET",
			"/api/budget/"+strconv.Itoa(
				i+1,
			),
			nil,
		)
		getPlanByIdReq.Header.Add(
			"Authorization",
			"Bearer "+token,
		)
		resp, err := app.Test(
			getPlanByIdReq,
		)
		if err != nil {
			t.Error(err)
		}

		respBytes, err := io.ReadAll(
			resp.Body,
		)
		if err != nil {
			t.Error(err)
		}

		var budgetPlanResponseBody budget.BudgetPlanResponseBody
		err = json.Unmarshal(
			respBytes,
			&budgetPlanResponseBody,
		)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(
			t,
			uint64(i+1),
			budgetPlanResponseBody.Data.ID,
		)
	}
}

func TestGetBudgetItemList(
	t *testing.T,
) {
	budgetItemListRequest := httptest.NewRequest(
		"GET",
		"/api/budget?pageSize=2&pageNo=2",
		nil,
	)
	budgetItemListRequest.Header.Add(
		"Authorization",
		"Bearer "+token,
	)
	resp, err := app.Test(
		budgetItemListRequest,
	)
	if err != nil {
		t.Error(err)
	}

	respBytes, err := io.ReadAll(
		resp.Body,
	)
	if err != nil {
		t.Error(err)
	}

	var budgetPlanItemListResponseBody budget.BudgetPlanItemListResponseBody
	err = json.Unmarshal(
		respBytes,
		&budgetPlanItemListResponseBody,
	)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(
		t,
		uint64(2),
		budgetPlanItemListResponseBody.Data.PageNo,
	)
	assert.Equal(
		t,
		uint64(1),
		budgetPlanItemListResponseBody.Data.PageSize,
	)
	item := budgetPlanItemListResponseBody.Data.BudgetPlans[0]
	expectedItem := budgetPlanItems[0]
	assert.Equal(t, uint64(1), item.ID)
	assert.Equal(
		t,
		expectedItem.Title,
		item.Title,
	)
	assert.Equal(
		t,
		expectedItem.Amount,
		item.Amount,
	)
}
