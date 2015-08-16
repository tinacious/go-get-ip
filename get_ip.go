package main

import (
    "bytes"
    "fmt"
    "log"
    "net"
    "os"
    "net/http"
    m "github.com/keighl/mandrill"
)

func main() {
    var buffer bytes.Buffer

    // noop favicon request otherwise duplicate emails
    http.HandleFunc("/favicon.ico", noop)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        ifaces, _ := net.Interfaces()

        for _, i := range ifaces {
            addrs, _ := i.Addrs()

            for _, addr := range addrs {
                var ip net.IP
                switch v := addr.(type) {
                case *net.IPNet:
                    ip = v.IP
                case *net.IPAddr:
                    ip = v.IP
                }

                buffer.WriteString(ip.String() + " ")
            }
        }

        // Add `x-forwarded-for` header because Heroku returns a private IP
        buffer.WriteString(" " + r.Header.Get("x-forwarded-for"))

        var str = fmt.Sprintf("The IP address is: %s ", buffer.String())
        emailResults(str)

        http.Redirect(w, r, os.Getenv("REDIRECT"), http.StatusFound)
    })

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}

func noop(http.ResponseWriter, *http.Request) {}

func emailResults(s string) {
    log.Print(s)

    client := m.ClientWithKey(os.Getenv("MANDRILL"))

    message := &m.Message{}
    message.AddRecipient(os.Getenv("EMAIL"), "", "to")
    message.FromEmail = os.Getenv("EMAIL")
    message.FromName = "Go Get IP"
    message.Subject = "IP address"
    message.Text = s

    client.MessagesSend(message)
}
