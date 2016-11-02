package graphite

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGraphiteFunctions(t *testing.T) {
	Convey("Testing Graphite Functions", t, func() {

		Convey("formatting time range for now", func() {

			timeRange := formatTimeRange("now")
			So(timeRange, ShouldEqual, "now")

		})

		Convey("formatting time range for now-1m", func() {

			timeRange := formatTimeRange("now-1m")
			So(timeRange, ShouldEqual, "now-1min")

		})

		Convey("formatting time range for now-1M", func() {

			timeRange := formatTimeRange("now-1M")
			So(timeRange, ShouldEqual, "now-1mon")

		})

		Convey("fix interval format in query for 1m", func() {

			timeRange := formatTimeRange("aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1m'), 4)")
			So(timeRange, ShouldEqual, "aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1min'), 4)")

		})

		Convey("fix interval format in query for 1M", func() {

			timeRange := formatTimeRange("aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1M'), 4)")
			So(timeRange, ShouldEqual, "aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1mon'), 4)")

		})

	})
}
