package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/restuwahyu13/go-supertest/supertest"
	. "github.com/smartystreets/goconvey/convey"
	"syreclabs.com/go/faker"
)

var router = SetupRouter()
var accessToken string
var studentId interface{}

func TestLoginHandler(t *testing.T) {

	Convey("Auth Login Handler Group", t, func() {

		Convey("Login User Account Is Not Registered", func() {
			payload := gin.H{
				"email":    "anto13@zetmail.com",
				"password": "qwerty12",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/login")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusNotFound, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "User account is not registered", response.Message)
			})
		})

		Convey("Login Failed User Account Is Not Active", func() {
			payload := gin.H{
				"email":    "carmelo_marquardt@weissnat.info",
				"password": "testing13",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/login")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusForbidden, rr.Code)
				assert.Equal(t, http.MethodPost, response.Method)
				assert.Equal(t, "User account is not active", response.Message)
			})
		})

		Convey("Login Error Username Or Password Is Wrong", func() {
			payload := gin.H{
				"email":    "eduardo.wehner@greenholtadams.net",
				"password": "testing",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/login")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusForbidden, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Username or password is wrong", response.Message)
			})
		})

		Convey("Login Success", func() {
			payload := gin.H{
				"email":    "eduardo.wehner@greenholtadams.net",
				"password": "qwerty12345",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/login")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Login successfully", response.Message)

				var token map[string]interface{}
				encoded := util.Strigify(response.Data)
				_ = json.Unmarshal(encoded, &token)

				accessToken = token["accessToken"].(string)
			})
		})
	})
}

func TestRegisterHandler(t *testing.T) {

	Convey("Auth Register Handler Group", t, func() {

		Convey("Register New Account", func() {
			payload := gin.H{
				"fullname": faker.Internet().Email(),
				"email":    faker.Internet().Email(),
				"password": "testing13",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/register")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusCreated, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Register new account successfully", response.Message)
			})
		})

	})
}

func TestForgotHandler(t *testing.T) {

	Convey("Auth Forgot Password Handler Group", t, func() {

		Convey("Forgot Password If Email Not Exist", func() {

			payload := gin.H{
				"email": "santosi131@zetmail.com",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/forgot-password")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusNotFound, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Email is not never registered", response.Message)
			})
		})

		Convey("Forgot Password If Account Is Not Active", func() {

			payload := gin.H{
				"email": "santoso13@zetmail.com",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/forgot-password")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusForbidden, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "User account is not active", response.Message)
			})
		})

		Convey("Forgot Password To Get New Password", func() {

			payload := gin.H{
				"email": "samsul1@zetmail.com",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/forgot-password")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Forgot password successfully", response.Message)
			})
		})

	})
}

func TestResendHandler(t *testing.T) {

	Convey("Auth Resend Token Handler Group", t, func() {

		Convey("Resend New Token If Email Not Exist", func() {

			payload := gin.H{
				"email": "santosi131@zetmail.com",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/resend-token")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusNotFound, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Email is not never registered", response.Message)
			})
		})

		Convey("Resend Token If Account Is Active", func() {

			payload := gin.H{
				"email": "restuwahyu13@zetmail.com",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/resend-token")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusForbidden, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "User account hash been active", response.Message)
			})
		})

		Convey("Forgot Password To Get New Password", func() {

			payload := gin.H{
				"email": "santoso13@zetmail.com",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/resend-token")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Resend new activation token successfully", response.Message)
			})
		})

	})
}

func TestResetHandler(t *testing.T) {

	Convey("Auth Reset Password Handler Group", t, func() {

		Convey("Reset Old Password To New Password", func() {
			payload := gin.H{
				"email":     "eduardo.wehner@greenholtadams.net",
				"password":  "qwerty12345",
				"cpassword": "qwerty12345",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/change-password/" + accessToken)
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Change new password successfully", response.Message)
			})
		})

	})
}

func TestCreateStudentHandler(t *testing.T) {

	Convey("Test Handler Create Student Group", t, func() {

		Convey("Create New Student Is Conflict", func() {

			payload := gin.H{
				"name": "bagus budiawan",
				"npm":  201543502292,
				"fak":  "mipa",
				"bid":  "tehnik informatika",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/student")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessToken)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusConflict, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Npm student already exist", response.Message)
			})
		})

		Convey("Create New Student", func() {

			payload := gin.H{
				"name": faker.Internet().FreeEmail(),
				"npm":  faker.RandomInt(25, 50),
				"fak":  "mipa",
				"bid":  "tehnik informatika",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/student")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessToken)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusCreated, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Create new student account successfully", response.Message)
			})
		})
	})
}

func TestResultsStudentHandler(t *testing.T) {

	Convey("Test Handler Results Student By ID Group", t, func() {

		Convey("Results All Student", func() {

			test := supertest.NewSuperTest(router, t)

			test.Get("/api/v1/student")
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessToken)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				var objects []map[string]interface{}
				encoded := util.Strigify(response.Data)
				_ = json.Unmarshal(encoded, &objects)

				studentId = objects[0]["ID"]

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "Results Students data successfully", response.Message)
			})
		})
	})
}

func TestResultStudentHandler(t *testing.T) {

	Convey("Test Handler Result Student By ID Group", t, func() {

		Convey("Result Specific Student If StudentID Is Not Exist", func() {

			ID := "00f85d71-083b-4089-9d20-bb1054df4575"

			test := supertest.NewSuperTest(router, t)

			test.Get("/api/v1/student/" + ID)
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessToken)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusNotFound, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "Student data is not exist or deleted", response.Message)
			})
		})

		Convey("Result Specific Student By ID", func() {

			ID := studentId

			test := supertest.NewSuperTest(router, t)

			test.Get("/api/v1/student/" + ID.(string))
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessToken)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)

				mapping := make(map[string]interface{})
				encode := util.Strigify(response.Data)
			_:
				json.Unmarshal(encode, &mapping)

				assert.Equal(t, ID, mapping["ID"])
				assert.Equal(t, "Result Student data successfully", response.Message)
			})
		})
	})
}

func TestDeleteStudentHandler(t *testing.T) {

	Convey("Test Handler Delete Student By ID Group", t, func() {

		Convey("Delete Specific Student If StudentID Is Not Exist", func() {

			ID := "00f85d71-083b-4089-9d20-bb1054df4575"

			test := supertest.NewSuperTest(router, t)

			test.Delete("/api/v1/student/" + ID)
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessToken)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusForbidden, rr.Code)
				assert.Equal(t, http.MethodDelete, req.Method)
				assert.Equal(t, "Student data is not exist or deleted", response.Message)
			})
		})

		Convey("Delete Specific Student By ID", func() {

			ID := studentId

			test := supertest.NewSuperTest(router, t)

			test.Delete("/api/v1/student/" + ID.(string))
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessToken)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodDelete, req.Method)
				assert.Equal(t, "Delete student data successfully", response.Message)
			})
		})
	})
}
