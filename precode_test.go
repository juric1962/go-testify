package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	//"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, responseRecorder.Code, http.StatusOK)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 10
	//  req := ... // здесь нужно создать запрос к сервису
	val := fmt.Sprintf("%d", totalCount)
	req := httptest.NewRequest("GET", "/cafe?count="+val+"&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	answer := strings.Split(responseRecorder.Body.String(), ",")
	// здесь нужно добавить необходимые проверки

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, totalCount, len(answer))

}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=kazan", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), `wrong city value`)
}
