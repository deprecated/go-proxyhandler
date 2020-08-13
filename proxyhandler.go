package proxyhandler

import (
	"os"
	"bufio"
	"errors"
	"sync"
	"math/rand"
)

// ProxyHandler defines the handler struct
type ProxyHandler struct {
	Path     string
	Proxies []string
	Mux sync.Mutex
}

// Create creates a new ProxyHandler
func Create(path string) (*ProxyHandler, error) {
	handler := &ProxyHandler{Path: path, Proxies: []string{}}
	
	return handler, nil
}

// Init initializes the ProxyHandler
func (h *ProxyHandler) Init() error {
	file, err := os.Open(h.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		h.Proxies = append(h.Proxies, scanner.Text())
	}

	return nil
}

// SmartRotateProxy returns the next proxy in the list and appends the old one to the end of the list
func (h *ProxyHandler) SmartRotateProxy() (string, error) {
	if (len(h.Proxies) == 0) {
		return "", errors.New("No Proxies loaded")
	}

	var p string
	h.Mux.Lock()
	p, h.Proxies = h.Proxies[0], h.Proxies[1:]
	h.Proxies = append(h.Proxies, p)
	h.Mux.Unlock()

	return p, nil
}

// RandomProxy returns a random proxy from the list
func (h *ProxyHandler) RandomProxy() (string, error) {
	if (len(h.Proxies) == 0) {
		return "", errors.New("No Proxies loaded")
	}

	i := rand.Intn(len(h.Proxies))
	p := h.Proxies[i]

	return p, nil
}

