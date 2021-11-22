package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lekan-pvp/incr1/internal/app/shorter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURLByID(t *testing.T) {
	type want struct {
		contentType string
		statusCode int
	}
	tests := []struct {
		name string
		request string
		id int
		long string
		shorts map[int]URLs
		want want
	}{
		{
			name: "success",
			request: "/1",
			id: 1,
			long: "http://google.com",
			want: want{
				contentType: "text/plain",
				statusCode: 307,
			},
		},
		{
			name: "not found",
			request: "/2",
			long: "http://yandex.ru",
			want: want{
				contentType: "text/plain",
				statusCode: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tt := range tests {
				short := shorter.Shorting(1)
				shorts[tt.id] = URLs{
					Long: tt.long,
					Short: short,
				}

				router := httprouter.New()
				router.GET("/:id", GetURLByID)

				req, err := http.NewRequest(http.MethodGet, tt.request, nil)
				assert.NoError(t, err)

				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				result := rr.Result()

				assert.Equal(t, tt.want.statusCode, result.StatusCode)
			}
		})
	}
}
