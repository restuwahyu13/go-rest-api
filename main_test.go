package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"syreclabs.com/go/faker"
)

var router = SetupRouter()
var accessToken = make(chan interface{}, 1)

func TestLoginHandler(t *testing.T) {

	Convey("Auth Login Handler Group", t, func() {

		Convey("Login User Account Is Not Registered", func() {
			payload := gin.H{
				"email": "anto13@zetmail.com",
				"password":  "qwerty12",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/login", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "User account is not registered", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 404, response.StatusCode)
		})

		Convey("Login Failed User Account Is Not Active", func() {
			payload := gin.H{
				"email": "carmelo_marquardt@weissnat.info",
				"password":  "testing13",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/login", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "User account is not active", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 403, response.StatusCode)
		})

		Convey("Login Error Username Or Password Is Wrong", func() {
			payload := gin.H{
				"email": "eduardo.wehner@greenholtadams.net",
				"password":  "testing",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/login", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Username or password is wrong", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 403, response.StatusCode)
		})

		Convey("Login Success", func() {
			payload := gin.H{
				"email": "eduardo.wehner@greenholtadams.net",
				"password":  "qwerty12345",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/login", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Login successfully", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 200, response.StatusCode)

			var token map[string]interface{}
			encoded := util.Strigify(response.Data)
			_ = json.Unmarshal(encoded, &token)
			accessToken <- token["accessToken"]
		})

	})
}

func TestRegisterHandler(t *testing.T) {

	Convey("Auth Register Handler Group", t, func() {

		Convey("Register New Account", func() {
			payload := gin.H{
				"fullname": faker.Internet().Email(),
				"email": faker.Internet().Email(),
				"password":  "testing13",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/register", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Register new account successfully", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 200, response.StatusCode)
		})

	})
}

func TestForgotHandler(t *testing.T) {

	Convey("Auth Forgot Password Handler Group", t, func() {

		Convey("Forgot Password If Email Not Exist", func() {
			payload := gin.H{
				"email": "santosi131@zetmail.com",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/forgot-password", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Email is not never registered", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 404, response.StatusCode)
		})

		Convey("Forgot Password If Account Is Not Active", func() {
			payload := gin.H{
				"email": "santoso13@zetmail.com",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/forgot-password", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "User account is not active", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 403, response.StatusCode)
		})

		Convey("Forgot Password To Get New Password", func() {
			payload := gin.H{
				"email": "samsul1@zetmail.com",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/forgot-password", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Forgot password successfully", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 200, response.StatusCode)
		})

	})
}

func TestResendHandler(t *testing.T) {

	Convey("Auth Resend Token Handler Group", t, func() {

		Convey("Resend New Token If Email Not Exist", func() {
			payload := gin.H{
				"email": "santosi131@zetmail.com",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/resend-token", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Email is not never registered", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 404, response.StatusCode)
		})

		Convey("Resend Token If Account Is Active", func() {
			payload := gin.H{
				"email": "restuwahyu13@zetmail.com",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/resend-token", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "User account hash been active", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 403, response.StatusCode)
		})

		Convey("Forgot Password To Get New Password", func() {
			payload := gin.H{
				"email": "santoso13@zetmail.com",
			}

			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/resend-token", util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Resend new activation token successfully", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 200, response.StatusCode)
		})

	})
}

func TestResetHandler(t *testing.T) {

	Convey("Auth Reset Password Handler Group", t, func() {

		Convey("Reset Old Password To New Password", func() {
			payload := gin.H{
				"email": "eduardo.wehner@greenholtadams.net",
				"password": "qwerty12345",
				"cpassword": "qwerty12345",
			}

			token := <-accessToken
			rr := util.HttpTestRequest(router, http.MethodPost, "/api/v1/change-password/" +  token.(string), util.Strigify(payload))

			response := util.Parse(rr.Body.Bytes())
			logrus.Info(response)

			assert.Equal(t, "Change new password successfully", response.Message)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, 200, response.StatusCode)
		})

	})
}