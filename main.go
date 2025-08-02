package main

import (
	"github.com/Takumikimura1899/openapi-schema-drive-go-next/api/generated"
	"github.com/Takumikimura1899/openapi-schema-drive-go-next/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 既存のpingエンドポイント（テスト用）
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// APIハンドラーを作成
	apiHandler := handlers.NewAPIHandler()

	// 生成されたAPIルートを登録
	generated.RegisterHandlers(r, apiHandler)

	r.Run(":8080")
}
