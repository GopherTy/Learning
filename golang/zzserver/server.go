package zzserver

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

// some constant
const (
	ReqAddr          = "https://localhost:8080"
	Addr             = ":8080"
	DefaultMaxMemory = 32 << 20 // 32 MB
	FormKey          = "data"
)

var m content

// content
type content struct {
	data map[string]bool
	mu   sync.Mutex
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	m.data = map[string]bool{}
}

// Server .
func Server() {
	s := &http.Server{
		Addr:    Addr,
		Handler: nil,
	}

	http.HandleFunc("/", result)
	log.Fatalln(s.ListenAndServeTLS("zzserver.pem", "zzserver.key"))
	log.Println("ending .......")
}

// Client .
func Client() {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := c.PostForm(ReqAddr, url.Values{"data": []string{"hello", "world", "world"}})
	handleError(err)
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	handleError(err)
}

// Clinet1 client function
func Clinet1(data []string) (rs []bool, err error) {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := c.PostForm(ReqAddr, url.Values{
		FormKey: data,
	})
	if err != nil {
		return
	}
	defer resp.Body.Close()

	tmp := make([]byte, 1024)
	n, err := resp.Body.Read(tmp)

	err = json.Unmarshal(tmp[:n], &rs)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// result .
func result(w http.ResponseWriter, r *http.Request) {
	if r.PostForm == nil {
		r.ParseMultipartForm(DefaultMaxMemory)
	}

	var rs []bool
	if vs := r.PostForm["data"]; len(vs) > 0 {
		rs = make([]bool, len(vs))

		m.mu.Lock()
		for i, v := range vs {
			if m.data[v] {
				rs[i] = true
			} else {
				m.data[v] = true
				rs[i] = false
			}
		}
		m.mu.Unlock()
	}

	data, err := json.Marshal(rs)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

func loadCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal("ReadFile: ", err)
	}
	pool.AppendCertsFromPEM(ca)
	return pool
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
