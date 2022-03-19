package errors

import (
	"reflect"

	"github.com/arikarim/go-cfa/models"
)

func RecordNotFound(data map[uint]interface{}) string {
	for key, value := range data {
		if  err := models.DB.First(value, key).Error; err != nil {
			return reflect.TypeOf(value).Elem().Name()
		}
	}
	return ""
}
