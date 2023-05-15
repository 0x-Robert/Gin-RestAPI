package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func index() *gin.Engine {
	r := gin.New() //gin 선언

	r.Use(gin.Logger())   //gin 내부 log, logger 미들웨어 사용 선언
	r.Use(gin.Recovery()) //gin 내부 recover, recovery 미들웨어 사용 - 패닉복구
	r.Use(CORS())         //crossdomain 미들웨어 사용 등록

	//원하는 만큼 미들웨어 추가가능
	r.GET("/bmart", aMart(), bMart) //User 지정 미들웨어

	route := r.Group("/route/v01", liteAuth())
	{
		route.POST("/post", post)
		route.PUT("/put", put)
		route.GET("/get", get)
		// nested group
		auth := route.Group("auth")
		auth.GET("/nest", nest)
	}
	return r
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 단순히 출력기능만 처리 현재는 출력예시
		fmt.Println("Authorization-word ", auth)

		c.Next() // 다음 요청 진행
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	// if err != nil {
	// 	c.JSON(204, gin.H{    // 204 or http.StatusNoContent
	// 			"message": "Fail Verified"})
	// }  else {
	// 	c.JSON(200, gin.H{    // 200 or http.StatusOK
	// 		"message": "ok",
	// 		"token":   "AccessToken",
	// 		"body":    "Body",
	// 	})
	// }

	// return
}

func post(w http.ResponseWriter, req *http.Request) {

}

// func main() {
// 	// r := gin.Default()
// 	// r.GET("/recv", func(c *gin.Context) {
// 	// 	//http-응답
// 	// 	c.JSON(200, gin.H{
// 	// 		"result": "ok",
// 	// 	})
// 	// })
// 	// // 	r.GET("/:name/:age", get)
// 	// // r.POST("/:name/:age", post)
// 	// r.Run() // listen and serve on 0.0.0.0:8080

// 	router := gin.Default()

// 	// Simple group: v1
// 	v1 := router.Group("/v1")
// 	{
// 		v1.POST("/login", temp1) //호출 Url : http://localhost:8080/v1/login
// 		v1.PUT("/submit", temp2) //호출 Url : http://localhost:8080/v1/submit
// 		v1.GET("/read", temp3) //호출 Url : http://localhost:8080/v1/read
// 	}

// 	// Simple group: v2
// 	v2 := router.Group("/v2")
// 	{
// 		v2.POST("/login", temp4) //호출 Url : http://localhost:8080/v2/login
// 		v2.DELETE("/submit", temp5) //호출 Url : http://localhost:8080/v2/submit
// 		v2.GET("/read", temp6) //호출 Url : http://localhost:8080/v2/read
// 	}

// 	router.Run(":8080")
// }

// }

func main() {
	mapi := &http.Server{
		// Addr:           config.server.port,
		Addr:           "8080",
		Handler:        index(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		return mapi.ListenAndServe()
	})

	stopSig := make(chan os.Signal) //chan 선언
	// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
	signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)
	<-stopSig //메세지 등록

	// 해당 context 타임아웃 설정, 5초후 server stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mapi.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		fmt.Println("timeout 5 seconds.")
	}
	fmt.Println("Server stop")
}
