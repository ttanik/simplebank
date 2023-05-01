package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ttanik/simplebank/token"
)

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(username, duration)
	assert.NoError(t, err)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}
func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setpuAuth     func(t *testing.T, request *http.Request, tokeMaker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setpuAuth: func(t *testing.T, request *http.Request, tokeMaker token.Maker) {
				addAuthorization(t, request, tokeMaker, authorizationTypeBearer, "user", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		}, {
			name: "NoAuthorization",
			setpuAuth: func(t *testing.T, request *http.Request, tokeMaker token.Maker) {

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		}, {
			name: "UnsupportedAuthorization",
			setpuAuth: func(t *testing.T, request *http.Request, tokeMaker token.Maker) {
				addAuthorization(t, request, tokeMaker, "CuAuthorization", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		}, {
			name: "InvalidAuthorizationFormat",
			setpuAuth: func(t *testing.T, request *http.Request, tokeMaker token.Maker) {
				addAuthorization(t, request, tokeMaker, "", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		}, {
			name: "ExpiredToken",
			setpuAuth: func(t *testing.T, request *http.Request, tokeMaker token.Maker) {
				addAuthorization(t, request, tokeMaker, authorizationTypeBearer, "user", -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)
			authPath := "/auth"
			server.router.GET(
				authPath,
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			assert.NoError(t, err)

			tc.setpuAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}
}
