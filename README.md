# Go ezsmtp

## Easy smtp forwarder

Simple cli utility that connects to a local smtp server to easily send
email.  Currenly does no authentification.

```
$ ezsmtp -m "testing" -t user@example.org -f sysadm@example.org -cc admin@other.org
```

Usage

```
Usage of ./ezsmtp:
  -cc Email
        Email CC to Receipients`
  -f Sender
        Sender email (default "ezsmtp@localhost")
  -h smtp
        smtp server (default "localhost")
  -m Message
        Message Body
  -p smtp port
        smtp port port (default "25")
  -s Subject
        Mail Subject (default "(No Subject)")
  -t Email
        Email Receipient`
```





