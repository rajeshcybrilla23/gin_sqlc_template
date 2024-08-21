package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// to avoid to run in debug mode
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
