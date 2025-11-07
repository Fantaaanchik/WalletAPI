package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mod/config"
	"io"
	"net/http"
)

// AuthMiddleware проверяет X-UserId и X-Digest
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetHeader("X-UserID")
		digest := ctx.GetHeader("X-Digest")

		if userID == "" || digest == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization headers"})
			ctx.Abort()
			return
		}
		// читаем тело запроса
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "enable no read request body"})
			ctx.Abort()
			return
		}
		//восстанавливаем тело, чтобы Gin мог его повторно использовать
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		//вычисляем HMAC-SHA1

		h := hmac.New(sha1.New, []byte(config.SecretKey))

		expected := hex.EncodeToString(h.Sum(nil))

		fmt.Println(expected)
		if digest != expected {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid digest"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
