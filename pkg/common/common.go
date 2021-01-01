package common

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
	"time"
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
	h.Write([]byte(randStringBytesMaskImprSrc(32)))
	bs := h.Sum(nil)

	return (dataPath + "/"),
		(fmt.Sprintf("%x", bs) + ".txt")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
