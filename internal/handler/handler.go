package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"llcmediatelTask/internal/models"
	"llcmediatelTask/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
type MoneyChanger interface{
	Exchange(money models.Money)(models.Answer)
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/", h.Exchange)
	return router
}
func (h *Handler)Exchange(c *gin.Context){
	input:=models.Money{}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())  
		return
	}
	if h.valid(input){
		newErrorResponse(c, http.StatusBadRequest, "bad request")  
		return
	}
	
	logrus.Debug(fmt.Sprintf("request body decoded.input %+v", input))
	
	res := h.services.MoneyChanger.Exchange(input)
	if len(res.Exchanges)==0{
		logrus.Warn("banknotes are not suitable for exchange")
	}
	logrus.Debug(fmt.Sprintf("request body decoded.result %+v", res))
	
	c.JSON(http.StatusOK, res)
}

func (h *Handler)valid(input models.Money) bool{
	if input.Amount==0{
		logrus.Errorf("amount is empty")
		return true
	}
	if len(input.Banknotes)==0{
		logrus.Errorf("banknotes is empty")
		return true
	}
	return false
	
}