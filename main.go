package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	input := c.Query("input")
	fmt.Println(input)

	// tpl := template.New("code")
	tpl, err := template.ParseFiles("./tmp/main.go.tpl")
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	f, err := os.Create("./tmp/main.go")
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	if err := tpl.Execute(f, input); err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	f.Close()

	comand := "cd tmp; pwd; go run main.go"
	cmd := exec.Command("/bin/bash", "-c", comand)
	// cmd := exec.Command("cmd", "/C", comand)
	// cmd.Stdin = strings.NewReader("some input")
	// var out bytes.Buffer
	// cmd.Stdout = &out

	// // 执行 cd tmp; go run .
	// if err := cmd.Run(); err != nil {
	// 	c.String(http.StatusOK, err.Error())
	// 	return
	// }

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, fmt.Sprintf("something wrong: %s", err.Error()))
		return
	}
	fmt.Println(string(out))
	c.String(http.StatusOK, string(out))
}

func update(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "upload.html", nil)
	} else if c.Request.Method == http.MethodPost {
		c.JSON(http.StatusOK, gin.H{"name": "luojie"})
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./view/*")
	r.GET("/", get)
	r.GET("/update", update)
	r.POST("/update", update)
	r.Run()
}
