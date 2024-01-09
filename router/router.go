package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController, oc controller.IOrganizationController, tec controller.ITeamController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken,
		},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.GET("/csrf", uc.CsrfToken)
	e.POST("/logout", uc.LogOut)

	// ユーザー
	u := e.Group("/users")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:jwtToken",
	}))
	u.GET("/userDetails", uc.GetLoggedInUserDetails)
	u.PUT("/updateName", uc.UpdateUserName)
	u.PUT("/assignToOrganization", uc.AssignUserToOrganization)
	u.POST("/assignToTeam", uc.AssignUserToTeam)
	u.PUT("/unassignFromTeam", uc.UnassignFromTeam)

	// 組織
	o := e.Group("/organization")
	o.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:jwtToken",
	}))
	o.GET("/created", oc.GetCreatedOrganizationsByUserId)
	o.GET("/lists", oc.ListOrganizations)
	o.POST("/create", oc.CreateOrganization)

	// チーム
	te := o.Group("/team")
	te.GET("/join", tec.GetAssignTeamByUserId)
	te.GET("/:organizationId", tec.GetTeamsByOrganizationId)
	te.POST("/:organizationId/create", tec.CreateTeam)
	te.DELETE("/:teamId", tec.DeleteTeam)

	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:jwtToken",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	// http://localhost:8080/tasks/status?taskStatus={Started, Unstarted or Completed}
	// taskStatusはcontrollerの "c.QueryParam("taskStatus")"で設定している
	t.GET("/status", tc.NarrowDownStatus)
	t.GET("/search/status", tc.FuzzySearch)
	t.GET("/by-deadlined", tc.GetTasksByDeadline)
	// t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.PUT("/:taskId/statusUpdate", tc.UpdateTaskStatus)
	t.DELETE("/:taskId", tc.DeleteTask)

	return e
}