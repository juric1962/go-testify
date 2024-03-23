package main

import (
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

/*
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
*/
