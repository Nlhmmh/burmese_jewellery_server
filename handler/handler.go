package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/orm_custom"
	"net/http"

	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

var _ ServerInterface = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

// Health Check
// (GET /api/health_check)
func (*Handler) GetApiHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// Get enums
// (GET /api/enums)
func (*Handler) GetApiEnums(c *gin.Context) {
	var list []*orm_custom.Enum
	if err := orm.NewQuery(qm.SQL(`
		SELECT
			n.nspname as enum_schema,
			t.typname as enum_name,
			string_agg(e.enumlabel, ',') as enum_value
		FROM pg_type t
			JOIN pg_enum e ON t.oid = e.enumtypid
			JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace
		GROUP BY enum_schema, enum_name;
		`,
	)).BindG(c, &list); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	data, err := models.ConvListFromORM(list, models.ConvEnumFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, data)
}
