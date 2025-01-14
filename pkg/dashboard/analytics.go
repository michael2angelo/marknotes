package dashboard

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/marknotes/analytics"
	"github.com/muhwyndhamhp/marknotes/pkg/models"
	pub_dashboard_analytics "github.com/muhwyndhamhp/marknotes/pub/pages/dashboards/analytics"
	"github.com/muhwyndhamhp/marknotes/template"
	"github.com/muhwyndhamhp/marknotes/utils/errs"
	"github.com/muhwyndhamhp/marknotes/utils/typeext"
)

func (fe *DashboardFrontend) Analytics(c echo.Context) error {
	slug := c.Param("slug")

	if slug == "" {
		return c.HTML(http.StatusOK, "")
	}

	var data []typeext.JSONB
	err := fe.DB.
		WithContext(c.Request().Context()).
		Model(&models.Analytics{}).
		Scopes(models.GetLatest(slug)).
		Pluck("data", &data).
		Error
	if err != nil {
		fmt.Println(errs.Wrap(err))
		return c.HTML(http.StatusOK, "")
	}

	p, err := typeext.ConvertJSONBToStruct[analytics.AnalyticsResponse](data[0])
	if err != nil {
		fmt.Println(errs.Wrap(err))
		return c.HTML(http.StatusOK, "")
	}

	resp := pub_dashboard_analytics.Analytics(&p)

	return template.AssertRender(c, http.StatusOK, resp)
}
