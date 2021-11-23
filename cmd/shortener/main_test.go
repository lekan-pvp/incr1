package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURLByID(t *testing.T) {
	type want struct {
		statusCode int

	}
	tests := []struct {
		name string
		request string
		want want
	}{
		{
			name: "success",
			request: "/1",
			want: want{
				statusCode: 307,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := httprouter.New()
			router.GET("/:id", GetURLByID)

			req, err := http.NewRequest(http.MethodGet, tt.request, nil)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			result := rr.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}