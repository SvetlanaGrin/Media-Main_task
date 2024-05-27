package service

import (
	"llcmediatelTask/internal/models"
	"reflect"
	"slices"
)

type Service struct {
	MoneyChanger
}

func NewService() *Service {
	return &Service{
		MoneyChanger: NewServiceMoneyСhange(),
	}
}

type MoneyChanger interface{
	Exchange(money models.Money)(models.Answer)
}

type ServiceMoneyСhange struct {
	seq []int
	result [][]int
}

func NewServiceMoneyСhange() *ServiceMoneyСhange{
	return &ServiceMoneyСhange{}
}
func (cm *ServiceMoneyСhange)MoneyChange( amount int,banknotes []int){
	if amount<0{
		return
	}
	if amount==0{
		arr:=make([]int,len(cm.seq),len(cm.seq))
		copy(arr,cm.seq)
		slices.Sort(arr)
		for _,elem:=range cm.result{
			if reflect.DeepEqual(elem, arr){
				return
			}
		}
		cm.result= append(cm.result,arr)
	}
	for i:=0;i<len(banknotes);i++{
		cm.seq =append(cm.seq,banknotes[i])
		cm.MoneyChange(amount-banknotes[i],banknotes)
		cm.seq=cm.seq[:len(cm.seq)-1]
	}
	
}
func (cm *ServiceMoneyСhange)Exchange(money models.Money)(models.Answer){
	var ans models.Answer
	cm.MoneyChange(money.Amount,money.Banknotes)
	ans.Exchanges=cm.result
	return ans
}