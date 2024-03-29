// @title PetBan Manage API
// @version 1.0
// @description This is PetBan Test server.
// @termsOfService https://petban.net
// @contact.name API Support
// @contact.url	https://petban.net
// @contact.email 94petban@petban.net
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http

package main

import (
	_ "awesomeProject/docs"
	"awesomeProject/ent"
	"awesomeProject/ent/city"
	"awesomeProject/ent/user"
	"bytes"
	"context"
	"fmt"
	"github.com/dchest/captcha"
	_ "github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	db   *ent.Client
	http *gin.Engine
}

var sr Server

func initDatabase() {
	client, err := ent.Open("mysql", "root:1qazXSW@@tcp(127.0.0.1:3306)/GolangDev?parseTime=True")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return
	}

	sr.db = client

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
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

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	fmt.Println("file : " + file)
	fmt.Println("ext : " + ext)
	fmt.Println("id : " + id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

// @Summary Get Captcha
// @Tags Captcha
// @Produce application/json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /captcha/create [get]
func handleGetCaptcha(c *gin.Context) {
	type CaptchaResponse struct {
		CaptchaId string `json:"captchaId"`
		ImageUrl  string `json:"imageUrl"`
	}
	var resp CaptchaResponse

	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	if d.CaptchaId != "" {
		resp.CaptchaId = d.CaptchaId
		resp.ImageUrl = "/show/" + d.CaptchaId + ".png"
		ResponseJSON(c, http.StatusOK, 200, "create CaptchaID", resp)
	} else {
		ResponseJSON(c, http.StatusOK, 500, "create CaptchaID failed:", nil)
	}
}

// @Summary Get Captcha Png
// @Tags Captcha
// @Produce image/png
// @Param imageUrl path string true "imageUrl"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /captcha/show/{imageUrl} [get]
func handleGetCaptchaPng(c *gin.Context) {
	imageUrl := c.Param("imageUrl")
	logrus.Info("GetCaptchaPng : " + imageUrl)
	ServeHTTP(c.Writer, c.Request)
}

// @Summary Verify Captcha
// @Tags Captcha
// @Produce application/json
// @Param captchaId query string true "captchaId"
// @Param value query string true "value"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /captcha/verify [get]
func handleGetVerifyCaptcha(c *gin.Context) {
	type GetQuery struct {
		captchaId string `form:"captchaId" json:"captchaId"`
		value     string `form:"value" json:"value"`
	}
	var form GetQuery

	form.captchaId = c.Query("captchaId")
	form.value = c.Query("value")

	httpCode, errCode := BindAndValid(c, &form)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}

	if form.captchaId == "" || form.value == "" {
		ResponseJSON(c, http.StatusOK, 500, "captchaId or value cant be empty", nil)
	} else {
		if captcha.VerifyString(form.captchaId, form.value) {
			ResponseJSON(c, http.StatusOK, 200, "驗證成功", nil)
		} else {
			ResponseJSON(c, http.StatusOK, 200, "驗證失敗", nil)
		}
	}
}

// @Summary create user
// @Tags User
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
		Nickname string `form:"nickname" json:"nickname" valid:"Required; MaxSize(50)"`
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
		SetNickname(form.Nickname).
		Save(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "create user failed:"+err.Error(), nil)
		return
	}

	type ResponseData struct {
		UserID   uint64 `json:"UserID"`
		UserName string `json:"UserName"`
		Nickname string `json:"nickname"`
	}
	var resp ResponseData
	resp.Nickname = form.Nickname
	resp.UserName = form.UserName
	resp.UserID = uint64(usr.ID)

	ResponseJSON(c, http.StatusOK, 200, "", resp)
}

// @Summary get user
// @Tags User
// @Produce application/json
// @Param username path string true "username"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/{username} [get]
func handleGetUser(c *gin.Context) {
	userName := c.Param("username")
	if len(userName) == 0 {
		ResponseJSON(c, 200, 400, "invalid param", nil)
		return
	}

	usr, _ := sr.db.User.
		Query().
		Where(user.Username(userName)).
		First(context.Background())
	if usr == nil {
		ResponseJSON(c, http.StatusOK, 500, "user doesn't exist", nil)
		return
	}

	type ResponseData struct {
		UserID   uint64 `json:"userid"`
		UserName string `json:"username"`
		Nickname string `json:"nickname"`
	}
	var resp ResponseData
	resp.Nickname = usr.Nickname
	resp.UserName = usr.Username
	resp.UserID = uint64(usr.ID)

	ResponseJSON(c, http.StatusOK, 200, "", resp)
}

// @Summary update user
// @Tags User
// @Produce application/json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param nickname formData string true "nickname"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/update [post]
func handleUpdateUser(c *gin.Context) {
	type PostParam struct {
		UserName string `form:"username" json:"username" valid:"Required; MaxSize(50)"`
		Password string `form:"password" json:"password" valid:"Required; MaxSize(50)"`
		Nickname string `form:"nickname" json:"nickname" valid:"Required; MaxSize(50)"`
	}
	var form PostParam

	httpCode, errCode := BindAndValid(c, &form)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}

	count, _ := sr.db.User.
		Update().
		SetUsername(form.UserName).
		SetPassword(form.Password).
		SetNickname(form.Nickname).
		Where(user.Username(form.UserName)).
		Save(context.Background())

	if count == 0 {
		ResponseJSON(c, http.StatusOK, 500, "update user failed", nil)
		return
	}

	usr, _ := sr.db.User.
		Query().
		Where(user.Username(form.UserName)).
		First(context.Background())

	if usr == nil {
		ResponseJSON(c, http.StatusOK, 500, "user doesn't exist", nil)
		return
	}

	type ResponseData struct {
		UserID   uint64 `json:"userid"`
		UserName string `json:"username"`
		Nickname string `json:"nickname"`
	}
	var resp ResponseData
	resp.Nickname = form.Nickname
	resp.UserName = form.UserName
	resp.UserID = uint64(usr.ID)

	ResponseJSON(c, http.StatusOK, 200, "", resp)
}

// @Summary delete user
// @Tags User
// @Produce application/json
// @Param id formData int true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/{id} [delete]
func handleDeleteUser(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		ResponseJSON(c, 200, 400, "invalid param", nil)
		return
	}

	intID, _ := strconv.ParseInt(id, 10, 64)
	if intID < 0 {
		ResponseJSON(c, 200, 500, "delete user failed", nil)
		return
	}

	_, err := sr.db.User.
		Delete().Where(user.ID(uint64(intID))).Exec(context.Background())
	if err != nil {
		ResponseJSON(c, 200, 500, "delete user failed", nil)
		return
	}

	ResponseJSON(c, 200, 200, "delete user ok", nil)
}

// @Summary get city
// @Tags City
// @Produce application/json
// @Param ID path string true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /city/{ID} [get]
func handleGetCity(c *gin.Context) {

	id := c.Param("ID")

	if len(id) == 0 {
		ResponseJSON(c, 200, 400, "invalid param", nil)
		return
	}

	intId, _ := strconv.Atoi(id)

	if intId < 0 {
		ResponseJSON(c, 200, 500, "invalid param: number should be positive", nil)
		return
	}

	city, _ := sr.db.City.
		Query().
		Where(city.ID(intId)).
		First(context.Background())

	if city == nil {
		ResponseJSON(c, http.StatusOK, 500, "city doesn't exist", nil)
		return
	}
	type ResponseData struct {
		ID          uint64 `json:"ID"`
		Name        string `json:"Name"`
		CountryCode string `json:"CountryCode"`
		District    string `json:"District"`
		Population  uint64 `json:"Population"`
	}

	var resp ResponseData

	resp.Population = uint64(city.Population)
	resp.District = city.District
	resp.CountryCode = city.CountryCode
	resp.Name = city.Name
	resp.ID = uint64(city.ID)
	ResponseJSON(c, http.StatusOK, 200, "", resp)
}

// @Summary create city
// @Tags City
// @Produce application/json
// @Param name formData string true "name"
// @Param countrycode formData string true "countrycode"
// @Param district formData string true "district"
// @Param population formData int true "population"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /city/create [post]
func handleCreateCity(c *gin.Context) {
	type PostParam struct {
		Name        string `form:"name" json:"name"`
		CountryCode string `form:"countrycode" json:"countrycode"`
		District    string `form:"district" json:"district"`
		Population  int    `form:"population" json:"population"`
	}

	var form PostParam

	httpCode, errCode := BindAndValid(c, &form)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}

	city, err := sr.db.City.
		Create().
		SetName(form.Name).
		SetCountryCode(form.CountryCode).
		SetDistrict(form.District).
		SetPopulation(form.Population).
		Save(context.Background())

	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "create user failed:"+err.Error(), nil)
		return
	}

	type ResponseData struct {
		ID          uint64 `json:"ID"`
		Name        string `json:"Name"`
		CountryCode string `json:"CountryCode"`
		District    string `json:"District"`
		Population  uint64 `json:"Population"`
	}

	var resp ResponseData
	resp.Population = uint64(city.Population)
	resp.District = city.District
	resp.CountryCode = city.CountryCode
	resp.Name = city.Name
	resp.ID = uint64(city.ID)

	ResponseJSON(c, http.StatusOK, 200, "", resp)
}
func runHttpServer() {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	sr.http = r

	// web api document http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json")))
	//Captcha
	r.GET("/captcha/create", handleGetCaptcha)
	r.GET("/captcha/show/:imageUrl", handleGetCaptchaPng)
	r.GET("/captcha/verify", handleGetVerifyCaptcha)
	// Create
	r.POST("/user/create", handleCreateUser)
	// Read
	r.GET("/user/:username", handleGetUser)
	// Update
	r.POST("/user/update", handleUpdateUser)
	// Delete
	r.DELETE("/user/:id", handleDeleteUser)

	r.GET("/city/:ID", handleGetCity)
	r.POST("/city/create", handleCreateCity)

	// Listen and serve on 0.0.0.0:8080
	_ = r.Run(":8080")
}

func main() {
	//before build run if you add new api you have to command this "swag init" first;
	initDatabase()
	runHttpServer()
}
