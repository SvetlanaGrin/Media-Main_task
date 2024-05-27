package service

import (
	"github.com/stretchr/testify/require"
	"llcmediatelTask/internal/models"
	"strconv"
	"testing"
)

func TestServiceMoneyСhange_MoneyChange_HappyPath(t *testing.T) {
	//t.Parallel()
	cases :=[]struct{
		Amount    int
		Banknotes []int
	}{
		{
			Amount:    400,
			Banknotes: []int{100,50},
		},
		{
			Amount:    600,
			Banknotes: []int{5000,1000},
		},
		{
			Amount:    50,
			Banknotes: []int{100,50},
		},
		{
			Amount:    150,
			Banknotes: []int{150,100,50},
		},
	}
	answer:=[]struct{
		res [][]int
	}{
		{
			res:[][]int{{100, 100, 100, 100}, {50, 50, 100, 100, 100}, {50, 50, 50, 50, 100, 100}, {50, 50, 50, 50, 50, 50, 100},{50, 50, 50, 50, 50, 50, 50, 50}},
		},
		{
			[][]int(nil),
		},
		{
			res:[][]int{{50}},
		},
		{
			res:[][]int{{150}, {50, 100}, {50, 50, 50}},
		},
	}
	smc:=ServiceMoneyСhange{}
	for i,tCase:=range cases{
		t.Run(strconv.Itoa(i), func(t *testing.T){
			t.Parallel()
			smc.MoneyChange(tCase.Amount,tCase.Banknotes)
			require.Equal(t, answer[i].res ,smc.result)
		})
	} 
}

func TestServiceMoneyСhange_Exchange(t *testing.T) {
	smc:=ServiceMoneyСhange{}
	money:=models.Money{Amount:    50,
			Banknotes: []int{100,50},
		}
	res:=models.Answer{ Exchanges: [][]int{{50}}}			
	answer:=smc.Exchange(money)
	require.Equal(t, answer ,res)
	
}
