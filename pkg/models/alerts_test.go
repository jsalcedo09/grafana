package models

import (
	"testing"

	"github.com/grafana/grafana/pkg/components/simplejson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAlertingModelTest(t *testing.T) {
	Convey("Testing Alerting model", t, func() {

		json1, _ := simplejson.NewJson([]byte(`{ "field": "value" }`))
		json2, _ := simplejson.NewJson([]byte(`{ "field": "value" }`))

		rule1 := &AlertRuleDAO{
			Expression:  json1,
			Name:        "Namn",
			Description: "Description",
		}

		rule2 := &AlertRuleDAO{
			Expression:  json2,
			Name:        "Namn",
			Description: "Description",
		}

		Convey("Testing AlertRule equals", func() {

			So(rule1.ContainsUpdates(rule2), ShouldBeFalse)
		})

		Convey("Changing the expression should contain update", func() {
			json2, _ := simplejson.NewJson([]byte(`{ "field": "newValue" }`))
			rule1.Expression = json2
			So(rule1.ContainsUpdates(rule2), ShouldBeTrue)
		})
	})
}
