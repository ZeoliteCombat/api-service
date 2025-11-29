package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetRandomId() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(b)
}

func GetUuid() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return u.String()
}

func GetUsername() string {
	return os.Getenv("SERVICE_USERNAME")
}

func GetConnAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	return localAddr
}

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}

func GetRandom() *big.Int {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return big.NewInt(0)
	}
	return new(big.Int).SetBytes(b)
}

func GetRandomInt(min, max int64) *big.Int {
	return new(big.Int).Int64(rand.Int63n(max-min)) + big.NewInt(min)
}

func GetRandomFloat(min, max float64) float64 {
	return min + (max-min)*math.Float64frombits(rand.Int63())
}

func IsZero(v interface{}) bool {
	switch v.(type) {
	case int:
		fallthrough
	case int8:
		fallthrough
	case int16:
		fallthrough
	case int32:
		fallthrough
	case int64:
		return v.(int) == 0
	case uint:
		fallthrough
	case uint8:
		fallthrough
	case uint16:
		fallthrough
	case uint32:
		fallthrough
	case uint64:
		return v.(uint) == 0
	case float32:
		fallthrough
	case float64:
		return math.Abs(v.(float64)) < 1e-9
	case string:
		return v.(string) == ""
	case []byte:
		return len(v.([]byte)) == 0
	case time.Time:
		return v.(time.Time).UnixNano() == 0
	case *big.Int:
		return v.(*big.Int).Sign() == -1
	default:
		return false
	}
}

func Hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func GetVersion() string {
	return strings.TrimPrefix(os.Getenv("GITHUB_REF"), "refs/tags/v")
}