/*
 * @Author: xiaoxu@mgtv.com
 * @Date: 2020-05-24 17:32:46
 * @Jira:
 * @Wiki:
 * @LastEditTime: 2020-08-23 00:08:05
 * @LastEditors: xiaoxu
 * @Description:
 * @FilePath: \go-admin-ui-vuef:\project\work\go\src\go-admin-demo\tools\logger.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package tools

import (
	"errors"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogger() {
	// logger.Println("--------------------------- settings.application.mode ", viper.GetString("settings.application.mode"))
	switch Mode(viper.GetString("settings.application.mode")) {
	case ModeDev, ModeTest:
		log.SetOutput(os.Stdout)
		log.SetLevel(log.TraceLevel)
	case ModeProd:
		file, err := os.OpenFile(viper.GetString("logger.dir")+"/api-"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
		if err != nil {
			log.Fatalln("log init failed")
		}

		var info os.FileInfo
		info, err = file.Stat()
		if err != nil {
			log.Fatal(err)
		}
		fileWriter := logFileWriter{file, info.Size()}
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(&fileWriter)
		// log.SetLevel(log.ErrorLevel)
		log.SetLevel(log.InfoLevel)
		log.Info("InitLogger")
	}

	log.SetReportCaller(true)
}

type logFileWriter struct {
	file *os.File
	size int64
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}
	n, e := p.file.Write(data)
	p.size += int64(n)
	//每天一个文件
	if p.file.Name() != viper.GetString("logger.dir")+"/api-"+time.Now().Format("2006-01-02")+".log" {
		p.file.Close()
		p.file, _ = os.OpenFile(viper.GetString("logger.dir")+"/api-"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
		p.size = 0
	}
	return n, e
}
