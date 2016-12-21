package main

import (
    "fmt"
    "log"
    "os"
    "flag"
    "github.com/PuerkitoBio/goquery"
)


func check(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}

func main() {

  resp, err := goquery.NewDocument("http://gobyexample.com")

  check(err)

  flag.Parse()

  var baseDir string

  cmdArg := flag.Args()

  if len(cmdArg) != 0 {
      baseDir = cmdArg[0]
  } else {
      baseDir = "learn-go-by-example"
  }

  os.Mkdir(baseDir, os.FileMode(0777))

  resp.Find("li > a").Each(func(i int, s *goquery.Selection) {
      a, _ := s.Attr("href")

      fileName := fmt.Sprintf("%02d-%v", i, a)

      filePath := fmt.Sprintf("%v/%v", baseDir, fileName)

      os.Mkdir(filePath, os.FileMode(0777))

      path := fmt.Sprintf("%v/%v.go", filePath, fileName)

      f, err := os.Create(path)

      check(err)

      defer f.Close()
  })
}