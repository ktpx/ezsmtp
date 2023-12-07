//*******************************************
//  EZSMTP -- easy smtp mail forwarder
//  
//*******************************************

package main

import (
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"bufio"
	"os"
)

// Command line options
type Opts struct {
	Subject string
	Sender  string
	To      string
	Cc      string
	Body    string
	Server  string
	Port    string
}

var opt Opts

func init() {
	flag.StringVar(&opt.Subject, "s", "(No Subject)", "Mail `Subject`")
	flag.StringVar(&opt.To, "t", "", "`Email` Receipient`")
	flag.StringVar(&opt.Body, "m", "", "`Message` Body")
	flag.StringVar(&opt.Cc, "cc", "", "`Email` CC to Receipients`")
	flag.StringVar(&opt.Server, "h", "localhost", "`smtp` server")
	flag.StringVar(&opt.Port, "p", "25", "`smtp port` port")
	flag.StringVar(&opt.Sender, "f", "ezsmtp@localhost", "`Sender` email")
}

// Read from stdin and return string
func readStdin() string {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s = fmt.Sprintf("%s%s", s, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return s
}

func main() {

	flag.Parse()

        // TODO: add auth
	// Set up authentication information.
	// auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	
	if len(opt.To) == 0 {
		log.Fatal("No recipient specified.")
	}
	if len(opt.Body) == 0 {
	  opt.Body = readStdin()
	}
	to := []string{opt.To}
	header := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", opt.To, opt.Subject, opt.Body)
	msg := []byte(header)

	err := smtp.SendMail(opt.Server+":"+opt.Port, nil, opt.Sender, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sending mail ...", to)
}
