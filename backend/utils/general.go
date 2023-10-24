package utils

import (
	"log"
	"strconv"
	"time"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetIntFromParams(c *fiber.Ctx, name string) uint64{
	sid := c.Params(name, "")
	if sid==""{
		return 0
	}
	id, _ := strconv.Atoi(sid)
	return uint64(id) // todo: id is int
}

// todo: add &models.User
// 0 means not access and 1 means access
// 2 means hacker trap and should report and ban user
// 3 means not found
func CheckAdminPermision(permisions string, p string) uint8{
	//  check hacker trap
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
	case "plan":
		return (permisions[10]-'0')
	case "member":
		return (permisions[11]-'0')
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
func ConvertStringToTime(t string, loc *time.Location) time.Time{
	splited := strings.Split(t, "-")
	year, _ := strconv.Atoi(splited[0])
	month, _ := strconv.Atoi(splited[1])
	day, _ := strconv.Atoi(splited[2])
	return time.Date(year, time.Month(month), day, 0,0,0,0, loc)
	
}