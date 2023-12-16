package utils

import (
	"log"
	"strconv"
	"time"
	"strings"
	"errors"
	"math"
	"math/rand"
	"crypto/sha256"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
)

func Init(){
	rand.Seed(time.Now().UnixNano())
}

func GetIntFromString(sid string, name string) int{

	if sid==""{
		return 0
	}
	id, _ := strconv.Atoi(sid)
	return id
}


func GetIDFromParams(c *fiber.Ctx, name string) uint64{
	sid := c.Params(name, "")
	if sid==""{
		return 0
	}
	var result uint64

	ln := len(sid)

	if ln==0{
		return 0
	}

	for i, n := range sid{
		n -= '0' // convert n to int
		if n < 0 || n > 9{
			return 0
		}
		result += uint64(n) * uint64(math.Pow10((ln - i-1)))
	}
	
	return result
}

// default will be pageLimit:20
func GetPageAndPageLimitFromMap(m map[string]string) (int, int, error){
	pageStr, ok := m["page"]
	page, err := strconv.Atoi(pageStr)
	if !ok || err!=nil{
		return 0,0, errors.New("invalid page")
	}

	pageLimitStr, ok := m["pageLimit"]
	pageLimit, err := strconv.Atoi(pageLimitStr)
	if !ok || err!=nil{
		return page, 20, nil
	}

	if pageLimit > 50 || pageLimit < 5{
		return 0,0, errors.New("invalid pageLimit")
	}

	return page, pageLimit, nil
}

// todo: add &models.User
// 0 means not access and 1 means access
// 2 means hacker trap and should report and ban user
// 3 means not found
func CheckAdminPermision(permisions string, p string) uint8{
	//  check hacker trap
	if len(permisions) < 15{
		return 0
	}

	if p[2] == '1'{
		return 2
	}


	switch p{
	case "createArticle":
		return (permisions[0]-'0')
	case "editMyArticle":
		return (permisions[1]-'0')
	case "deleteMyArticle":
		return (permisions[3]-'0')
	case "editOtherArticle":
		return (permisions[4]-'0')
	case "deleteOtherArticle":
		return (permisions[5]-'0')
	case "settings":
		return (permisions[6]-'0')
	case "work-sample":
		return (permisions[7]-'0')
	case "customer":
		return (permisions[8]-'0')
	case "deleteComment":
		return (permisions[9]-'0')
	case "plan":
		return (permisions[10]-'0')
	case "member":
		return (permisions[11]-'0')
	case "mainpage":
		return (permisions[12]-'0')
	case "article-category":
		return (permisions[13]-'0')
	}


	return 3
}

func LogError(err error){
	log.Println(err)
}


// get UTC time
func ConvertToPersianTime(t time.Time) time.Time{
	sub := t.Sub(time.Date(1970, 1, 1,0,0,0,0, time.UTC))
	sub += 3.5 * 60 * 60 * time.Second
	// 1348 10 11
	return time.Date(1348, 10, 11,0,0,0,0, time.FixedZone("Tehran", 3.5*60*60)).Add(sub)
}

// get +3.5 hours time
func ConvertToMiladiTime(t time.Time) time.Time{
	sub := t.Sub(time.Date(1348, 10, 11,0,0,0,0, time.FixedZone("Tehran", 3.5*60*60)))
	return time.Date(1970, 1, 1,0,0,0,0, time.UTC).Add(sub)
}

// t format is year-month-day
func ConvertStringToTime(t string, loc *time.Location) (time.Time, error){
	splited := strings.Split(t, "-")
	if len(splited) < 3{
		return time.Date(1000, 1, 1, 0,0,0,0, loc),errors.New("general.go: ConvertStringToTime:invalid string parameter: "+t+"\nmust splited with -")
	}
	year, err1 := strconv.Atoi(splited[0])
	month, err2 := strconv.Atoi(splited[1])
	day, err3 := strconv.Atoi(splited[2])

	var err error
	if err1!=nil{err=err1}
	if err2!=nil{err=err2}
	if err3!=nil{err=err3}

	return time.Date(year, time.Month(month), day, 0,0,0,0, loc), err	
}

func ConvertStringTimeToPersianStringTime(s string) (string, error){
	t, err := ConvertStringToTime(s, time.FixedZone("Tehran", 3.5*60*60))
	return ConvertTimeToString(ConvertToPersianTime(t)), err
}

func ConvertTimeToString(t time.Time) string{
	return strings.Split(t.String(), "T")[0]
}

func Hash(s string) string{
	h:= sha256.New()
	h.Write([]byte(s))
	return string(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}

func GetRandInt(min, max int) (int, error) {
	if max <= min {
		return 0, errors.New("invalid min and max")
	}
	return (rand.Intn(max - min) + min), nil
}