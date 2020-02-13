package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rahul-golang/exercise4/models"
)

//AppServiceImpl application services
type AppServiceImpl struct {
}

//AppService **
type AppService interface {
	GetFormatedDateTime(context.Context, models.DateTime) (*models.DateTime, error)
}

//NewAppServiceImpl **
func NewAppServiceImpl() AppService {

	return AppServiceImpl{}
}

//GetFormatedDateTime converts date and return
func (appService AppServiceImpl) GetFormatedDateTime(ctx context.Context,
	datetime models.DateTime) (*models.DateTime, error) {

	str1 := strings.Split(datetime.DateTime, "T")
	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, err := time.Parse(layout1, string(str1[1]))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(t.Format(layout1))
	fmt.Println(t.Format(layout2))

	datetime.DateTime = fmt.Sprint(str1[0] + "T" + fmt.Sprint(t.Format(layout2)))
	return &datetime, nil

}
