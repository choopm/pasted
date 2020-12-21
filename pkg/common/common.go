package common

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
)

// MakeFileName generates the full filepath from args:
// hostAddr, filePath, fileName
// 127.0.0.1, /data/, 7c13232c726d635ce4076f56779e42e85ec58c40.txt
func MakeFileName(dataPath, remoteAddr string) (string, string) {
	hostPort := strings.Split(remoteAddr, ":")
	hostPort = hostPort[0 : len(hostPort)-1]
	host := strings.Join(hostPort, ":")

	h := sha1.New()
	h.Write([]byte(host))
	h.Write([]byte(randStringRunes(32)))
	bs := h.Sum(nil)

	return (dataPath + "/"),
		(fmt.Sprintf("%x", bs) + ".txt")
}

func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
