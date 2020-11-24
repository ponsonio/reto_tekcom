package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ponsonio/reto_tekcom/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMultiplyShallMultiply(t *testing.T) {

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	var jsonStr = []byte(`{"a" : 100,"b" : 10}`)

	c.Request, _ = http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonStr))

	Multiply(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	var resp MultiplyResponce
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, "1000", resp.Result)
}

func TestMultiplyShallReturnErrorBadRequest(t *testing.T) {

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	var jsonStr = []byte(`{"x" : 100,"b" : 10}`)

	c.Request, _ = http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonStr))

	Multiply(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

}

func TestMultiplyShallReturnErrorServiceError(t *testing.T) {

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	var jsonStr = []byte(`{"a" : 100,"b" : 10}`)

	c.Request, _ = http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonStr))

	serviceMock := new(mocks.Multiplicator)

	serviceMock.On("Multiply", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Service Error"))

	service = serviceMock

	Multiply(c)

	assert.EqualValues(t, http.StatusInternalServerError, response.Code)

}