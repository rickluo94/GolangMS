package main

import (
	docs "awesomeProject/docs"
	"awesomeProject/ent"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

// @BasePath /api/v1

type Server struct {
	db   *ent.Client
	http *gin.Engine
}

var sr Server

func initDatabase() {
	client, err := ent.Open("mysql", "nwhuualh_sa:dak@bnj3AFJ_zvj_bqx@tcp(144.48.143.151:3306)/nwhuualh_pbm?parseTime=True")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func ResponseJSON(c *gin.Context, httpCode, errCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, 400
	}

	validate := validator.New()

	check := validate.Struct(form)

	if check != nil {
		return http.StatusInternalServerError, 500
	}

	//if !check {
	//	return http.StatusBadRequest, 400
	//}

	return http.StatusOK, 200
}

// @Summary create user
// @Produce application/json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param nickname formData string true "nickname"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/create [post]
func handleCreateUser(c *gin.Context) {
	type PostParam struct {
		UserName string `form:"username" json:"username" valid:"Required; MaxSize(50)"`
		Password string `form:"password" json:"password" valid:"Required; MaxSize(50)"`
		NickName string `form:"nickname" json:"nickname" valid:"Required; MaxSize(50)"`
	}
	var form PostParam

	httpCode, errCode := BindAndValid(c, &form)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}

	usr, err := sr.db.User.
		Create().
		SetUsername(form.UserName).
		SetPassword(form.Password).
		SetNickname(form.NickName).
		Save(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "create user failed:"+err.Error(), nil)
		return
	}

	type ResponseData struct {
		UserID   uint64 `json:"UserID"`
		UserName string `json:"UserName"`
		NickName string `json:"NickName"`
	}
	var resp ResponseData
	resp.NickName = form.NickName
	resp.UserName = form.UserName
	resp.UserID = uint64(usr.ID)

	ResponseJSON(c, http.StatusOK, 200, "", resp)
}

func runHttpServer() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	sr.http = r
	docs.SwaggerInfo.BasePath = "/api/v1"

	// web api document http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	// Create
	r.POST("/user/create", handleCreateUser)

	// Listen and serve on 0.0.0.0:8080
	_ = r.Run(":8080")
}

func main() {
	//before build run if you add new api you have to command this "swag init" first;

	initDatabase()
	runHttpServer()
}
