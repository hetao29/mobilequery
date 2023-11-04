package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/xluohome/phonedata"
	"log"
	"runtime"
)

func main() {
	/**
	 * PHONE_DATA_DIR 环境变量(内容目录)
	 * https://github.com/xluohome/phonedata/blob/master/phonedata.go
	 * CMCC		//中国移动
	 * CUCC		//中国联通
	 * CTCC		//中国电信
	 * CTCC_v	//电信虚拟运营商
	 * CUCC_v	//联通虚拟运营商
	 * CMCC_v	//移动虚拟运营商
	 * CBCC		//中国广电
	 * CBCC_v	//广电虚拟运营商
	 */
	bindaddr := flag.String("b", "0.0.0.0:80", "listen port")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/query", func(c *gin.Context) {
		mobile := c.DefaultQuery("mobile", "")
		pr, err := phonedata.Find(mobile)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "pong",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "pong",
				"record":  pr,
			})
		}
	})
	log.Println("notice: bind addr:%v", *bindaddr)
	err := r.Run(*bindaddr) // listen and serve on 0.0.0.0:8080
	log.Println("error: %v", err)
}
