package main

import ("gopkg.in/gomail.v2"
// "encoding/json"
"fmt"
"bufio"
"os"
"strings"
"strconv"
)

var settings array
var email string
var pass string
var port_int int
var port_int1 int

func main() {

file, _ := os.Open("config") //имя файла, откуда читаем настройки
f := bufio.NewReader(file)

fmt.Println("\nВаши настройки:")
settings, _ := f.ReadString(0)
fmt.Print(settings)
arr := strings.Split(settings, "\n")
fmt.Print("\n")

email := strings.Split(arr[0], ":")
pass := strings.Split(arr[1], ":")
server := strings.Split(arr[2], ":")
port_str := strings.Split(arr[3], ":")
port := port_str[1]
port_int, _ := strconv.ParseInt(port, 10, 0)
port_int1 = int(port_int)

to := strings.Split(arr[4], ":")
defer file.Close()

//письмо

m := gomail.NewMessage()
m.SetHeader("From", email[1])
m.SetHeader("To", to[1], email[1]) //!!!!!!!!!!!!!!
m.SetAddressHeader("", "", "")
m.SetHeader("Subject", "Config Worked")
m.SetBody("text/html", "Config Worked!")
m.Attach("1.png")

//отправка

d := gomail.NewDialer(server[1], port_int1, email[1], pass[1])

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
panic(err)
}

}
