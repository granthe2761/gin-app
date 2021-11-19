package handlers

import(
	"log"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"time"
)

func UploadImageHandler(c *gin.Context){
	log.Println("image posted")
	//图片名处理
	file, _ := c.FormFile("file")	
	file.Filename += fmt.Sprintf("%d",time.Now().Unix())		
	log.Println(file.Filename)
	file.Filename += ".jpg"
	
	//上传文件到指定的路径
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	dir := fmt.Sprintf("/home/go/src/github.com/granthe2761/gin-app/media/images/%d%d%d/",year,month,day)
	log.Println(dir)
	//确保文件夹已建立，否则无法保存
	err := os.MkdirAll(dir,0666)
	if err != nil{
		log.Println("创建目录错误：",err)
	}
	//保存路径处理
	dir += file.Filename
	//保存文件
	ret := c.SaveUploadedFile(file, dir)
	if ret != nil{
		log.Println(ret)
	}

	//开始计算
	
	
	//返回成功
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
