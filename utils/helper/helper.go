package helper

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnv(projectDirName string) {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	godotenv.Load(string(rootPath) + `/.env`)
}

// MustGetEnv get environment value
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	log.Println("ENV : ", key, value)
	if len(value) == 0 {
		return ""
	}
	return value
}

func Now() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc)
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func AdjustStructToStruct(a interface{}, b interface{}) interface{} {
	JsonStruct, _ := json.Marshal(a)
	json.Unmarshal([]byte(JsonStruct), &b)
	return b
}

func ForceInt(v interface{}) int {
	var result int
	switch v.(type) {
	case int:
		result = v.(int)
	case float64:
		result = int(v.(float64))
	case string:
		result, _ = strconv.Atoi(v.(string))
	}
	return result
}

func ForceString(v interface{}) string {
	var result string
	switch v.(type) {
	case int:
		result = strconv.Itoa(v.(int))
	case float64:
		result = strconv.Itoa(int(v.(float64)))
	case string:
		result, _ = v.(string)
	case any:
		result = v.(string)
	}
	return result
}

func ForceError(v interface{}) error {
	result := errors.New(ForceString(v))
	return result
}

func LeadingThousand(v int64) string {
	m := big.NewInt(v)
	return fmt.Sprintf("%03s", m)
}

func DecodeBase64BigInt(s string) *big.Int {
	buffer, _ := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	return big.NewInt(0).SetBytes(buffer)
}

func PrintJson(data interface{}) (res string) {
	manifestJson, _ := json.MarshalIndent(data, "", "  ")
	res = string(manifestJson)
	return
}

func IsMatchRegex(s string) (res bool) {
	rgx := "^[0-9a-zA-Z_]{6}$"
	res, _ = regexp.MatchString(rgx, s)
	return
}

func FormatToDateTime(p string) (result *time.Time, err error) {
	preResult, err := time.Parse("2006-01-02 15:04:05", p)
	if err != nil {
		return
	}

	result = &preResult
	return
}

func AdjustUrl(p string) (res string) {
	cHttps := strings.Split(p, "https://")
	cHttp := strings.Split(p, "http://")
	if len(cHttps) == 1 && len(cHttp) == 1 {
		res = fmt.Sprintf("https://%s", p)
	} else {
		res = p
	}
	return
}
