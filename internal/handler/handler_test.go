package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io"
	"llcmediatelTask/internal/models"
	"llcmediatelTask/internal/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"github.com/golang/mock/gomock"
	serviceMock "llcmediatelTask/internal/handler/mock_handler"
)

func TestHandler_valid(t *testing.T) {
	h := Handler{}
	cases := []models.Money{
		{
			Amount:    0,
			Banknotes: []int{50, 100},
		},
		{
			Amount:    50,
			Banknotes: []int{},
		},
		{
			Amount:    50,
			Banknotes: []int{50, 100},
		},
	}
	for _, tCase := range cases {
		t.Run(strconv.Itoa(tCase.Amount), func(t *testing.T) {
			t.Parallel()
			ans := h.valid(tCase)
			if tCase.Amount == 0 || len(tCase.Banknotes) == 0 {
				require.Equal(t, true, ans)
			} else {
				require.Equal(t, false, ans)
			}

		})
	}
}

func TestHandler_Exchange_HappyPath(t *testing.T) {
	
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	money:=models.Money{Amount:    50,
		Banknotes: []int{100,50},
	}
	res:=models.Answer{ Exchanges: [][]int{{50}}}	
	mockChanger := serviceMock.NewMockMoneyChanger(ctrl)
	mockChanger.EXPECT().Exchange(money).Return(res).Times(1)
	services := &service.Service{
		MoneyChanger: mockChanger,
	}
	h := &Handler{
		services: services,
	}
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/", h.Exchange)
	req := httptest.NewRequest(http.MethodGet,
		"/",
		bytes.NewBuffer(
			[]byte(`{"amount": 50,"banknotes": [100,50]}`)))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
	result := w.Result()

	defer result.Body.Close()
	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)
	var ans models.Answer
	if err := json.Unmarshal(data, &ans); err != nil {
		logrus.Fatal(err.Error())
	}
	require.Equal(t, ans,res)
}

func TestHandler_Exchange_Error(t *testing.T) {

	h := &Handler{
	}
	gin.SetMode(gin.TestMode)
	testRec:=[]struct{
		Rec  []byte
	}{
		{
			Rec: []byte(`{"amount": 50.5,"banknotes": [100,50]}`),	
			},
			{
			Rec: []byte(`{"amount": ,"banknotes": [100,50]}`),
			},
			{
			Rec: []byte(`{"amount": 50,"banknotes": [10dhfj0,50]}`),	
			},
			{
			Rec: []byte(`{"amot": 50,"banknotes": [100,50]}`),	
			},
			{
			Rec: []byte(`{"amount": 50}`),
			},
			{
			Rec: []byte(`{"amount":osdfbanknotes": [100,50]}`),	
			},
			}
			for _,tCase:=range testRec{
				t.Run(string(tCase.Rec), func(t *testing.T){
					t.Parallel()
					r := gin.Default()
					r.GET("/", h.Exchange)
					req := httptest.NewRequest(http.MethodGet,
						"/",
						bytes.NewBuffer(tCase.Rec))

					w := httptest.NewRecorder()
					r.ServeHTTP(w, req)

					require.Equal(t, w.Code, http.StatusBadRequest)

	//				require.EqualError(t, err,error("invalid character 'o' looking for beginning of value"))
					res := w.Result()
					defer res.Body.Close()
					data, _ := io.ReadAll(res.Body)
					require.Equal(t,[]byte{0x7b, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x76,
						0x61, 0x6c, 0x69, 0x64, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x20, 0x27, 0x6f, 0x27, 0x20, 0x6c, 0x6f, 0x6f, 0x6b,
						0x69, 0x6e, 0x67, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x6e, 0x69,
						0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7d},
					data)
				})
			}
}
