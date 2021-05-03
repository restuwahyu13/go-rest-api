package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	util "github.com/restuwahyu13/gin-rest-api/utils"
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

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/login", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, req.Method, response.Method)
			assert.Equal(t, "User account is not registered", response.Message)
		})

		Convey("Login Failed User Account Is Not Active", func() {
			payload := gin.H{
				"email": "carmelo_marquardt@weissnat.info",
				"password":  "testing13",
			}

			rr, req, err:= util.HttpTestRequest(http.MethodPost, "/api/v1/login", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, req.Method, response.Method)
			assert.Equal(t, "User account is not active", response.Message)
		})

		Convey("Login Error Username Or Password Is Wrong", func() {
			payload := gin.H{
				"email": "eduardo.wehner@greenholtadams.net",
				"password":  "testing",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/login", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, req.Method, response.Method)
			assert.Equal(t, "Username or password is wrong", response.Message)
		})

		Convey("Login Success", func() {
			payload := gin.H{
				"email": "eduardo.wehner@greenholtadams.net",
				"password":  "qwerty12345",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/login", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, req.Method, response.Method)
			assert.Equal(t, "Login successfully", response.Message)

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

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/register", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, req.Method, response.Method)
			assert.Equal(t, "Register new account successfully", response.Message)
		})

	})
}

func TestForgotHandler(t *testing.T) {

	Convey("Auth Forgot Password Handler Group", t, func() {

		Convey("Forgot Password If Email Not Exist", func() {
			payload := gin.H{
				"email": "santosi131@zetmail.com",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/forgot-password", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, "Email is not never registered", response.Message)
		})

		Convey("Forgot Password If Account Is Not Active", func() {
			payload := gin.H{
				"email": "santoso13@zetmail.com",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/forgot-password", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, "User account is not active", response.Message)
		})

		Convey("Forgot Password To Get New Password", func() {
			payload := gin.H{
				"email": "samsul1@zetmail.com",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/forgot-password", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, "Forgot password successfully", response.Message)
		})

	})
}

func TestResendHandler(t *testing.T) {

	Convey("Auth Resend Token Handler Group", t, func() {

		Convey("Resend New Token If Email Not Exist", func() {
			payload := gin.H{
				"email": "santosi131@zetmail.com",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/resend-token", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, "Email is not never registered", response.Message)
		})

		Convey("Resend Token If Account Is Active", func() {
			payload := gin.H{
				"email": "restuwahyu13@zetmail.com",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/resend-token", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, "User account hash been active", response.Message)
		})

		Convey("Forgot Password To Get New Password", func() {
			payload := gin.H{
				"email": "santoso13@zetmail.com",
			}

			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/resend-token", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, "Resend new activation token successfully", response.Message)
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
			rr, req, err := util.HttpTestRequest(http.MethodPost, "/api/v1/change-password/" +  token.(string), util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, http.MethodPost, response.Method)
			assert.Equal(t, "Change new password successfully", response.Message)
		})

	})

}

func TestCreateStudentHandler(t  *testing.T) {

	Convey("Test Handler Create Student Group", t, func() {

		Convey("Create New Student", func() {

			token := <-accessToken
			payload := gin.H{
				"name": faker.App().Name(),
				"npm": faker.RandomInt(10, 20),
				"fak": "mipa",
				"bid": "tehnik informatika",
			}

			rr, req, err := util.HttpTestRequest(http.MethodGet, "/api/v1/student", util.Strigify(payload))

			if err != nil {
				t.Error(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer " + token.(string))
			router.ServeHTTP(rr, req)

			response := util.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, rr.Code, response.StatusCode)
			assert.Equal(t, req.Method, response.Method)
			assert.Equal(t, "Create new student account successfully", response.Message)
		})

	})
}

func TestResultStudentHandler(t  *testing.T) {

		Convey("Test Handler Result Student By ID Group", t, func() {

			Convey("Result Specific Student If StudentID Is Not Exist", func() {

				token := <-accessToken
				ID := "00f85d71-083b-4089-9d20-bb1054df4575"

				rr, req, err := util.HttpTestRequest(http.MethodGet, "/api/v1/student/" + ID, nil)

				if err != nil {
					t.Error(err.Error())
				}

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer " + token.(string))
				router.ServeHTTP(rr, req)

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, rr.Code, response.StatusCode)
				assert.Equal(t, req.Method, response.Method)
				assert.Equal(t, "Student data is not exist or deleted", response.Message)
			})

			Convey("Result Specific Student By ID", func() {

				token := <-accessToken
				ID := "7a400ff9-6eca-4aba-90d8-5ad5ac21cc7d"

				rr, req, err := util.HttpTestRequest(http.MethodGet, "/api/v1/student/" + ID, nil)

				if err != nil {
					t.Error(err.Error())
				}

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer " + token.(string))
				router.ServeHTTP(rr, req)

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, rr.Code, response.StatusCode)
				assert.Equal(t, req.Method, response.Method)

				mapping := make(map[string]interface{})
				encode := util.Strigify(response.Data)
				_ : json.Unmarshal(encode, &mapping)

				assert.Equal(t, ID, mapping["ID"])
				assert.Equal(t, "Result Student data successfully", response.Message)

			})

		})
	}