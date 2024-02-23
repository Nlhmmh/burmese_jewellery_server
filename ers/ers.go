package ers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var (
	BadRequest     = &ErrResp{statusCode: http.StatusBadRequest, Code: 4, Message: "Bad Request"}
	InternalServer = &ErrResp{statusCode: http.StatusInternalServerError, Code: 5, Message: "Internal Server Error"}
	NotFound       = &ErrResp{statusCode: http.StatusNotFound, Code: 6, Message: "Not Found"}
	UnAuthorized   = &ErrResp{statusCode: http.StatusUnauthorized, Code: 7, Message: "UnAuthorized"}
	TmpRedirect    = &ErrResp{statusCode: http.StatusTemporaryRedirect, Code: 8, Message: "Redirect"}

	UnAuthorizedAdmin      = &ErrResp{statusCode: http.StatusUnauthorized, Code: 20, Message: "Unauthorized admin access"}
	UnAuthorizedAdminStaff = &ErrResp{statusCode: http.StatusUnauthorized, Code: 21, Message: "Unauthorized admin staff access"}

	UserWithEmailNotExist     = &ErrResp{statusCode: http.StatusNotFound, Code: 30, Message: "User with email does not exists"}
	UserWithEmailAlreadyExist = &ErrResp{statusCode: http.StatusConflict, Code: 31, Message: "User with email already exists"}
	PasswordWrong             = &ErrResp{statusCode: http.StatusBadRequest, Code: 32, Message: "Password is wrong"}
	DisplayNameAlreadyExist   = &ErrResp{statusCode: http.StatusConflict, Code: 33, Message: "Display name already exists"}

	ContentTitleAlreadyExist = &ErrResp{statusCode: http.StatusConflict, Code: 90, Message: "Content title already exists"}
	CartEmpty                = &ErrResp{statusCode: http.StatusBadRequest, Code: 91, Message: "Cart is empty"}
)

type ErrResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	statusCode int    `json:"-"`
	error      error  `json:"-"`
}

func (e *ErrResp) New(err error) *ErrResp {
	return &ErrResp{
		statusCode: e.statusCode,
		Code:       e.Code,
		Message:    e.Message,
		error:      err,
		Error:      err.Error(),
	}
}

func (e *ErrResp) Abort(c *gin.Context) {
	log.Error().Err(wrap(e.error)).Msg("")
	c.AbortWithStatusJSON(e.statusCode, e)
}

func (e *ErrResp) Rollback(c *gin.Context, tx *sql.Tx) {
	if err := tx.Rollback(); err != nil {
		InternalServer.Abort(c)
		return
	}
	e.Abort(c)
}

func (e *ErrResp) TmpRedirect(c *gin.Context, redirectURL string) {
	log.Error().Err(wrap(e.error)).Msg("")
	c.Redirect(http.StatusTemporaryRedirect, "/api")
}
