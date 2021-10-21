package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func request_header(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Header)
	for key, value := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s = %s\n", key, value[0]))
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	io.WriteString(w, "200\n")
}

func getEnv(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	cmd := exec.Command("go", "env")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	//因为结果是字节数组，需要转换成string
	cmd_res := string(output)
	//获取GOVERSION
	go_env_string := strings.Split(cmd_res, "\n")
	for i := 0; i < len(go_env_string); i++ {
		if strings.Contains(go_env_string[i], "GOVERSION") {
			fmt.Println(go_env_string[i])
			io.WriteString(w, go_env_string[i])
		}
	}

}

func main() {
	http.HandleFunc("/", request_header)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/getenv", getEnv)
	http.ListenAndServe(":8888", nil)
}
