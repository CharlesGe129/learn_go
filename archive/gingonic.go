// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package archive

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/json-iterator/go"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world!",
	})
}

type Login struct {
	User string `json:"user" binding:"-"`
	Password string `json:"password" binding:"-"`
}

// swagger:parameters getBookable
type Booking struct {
	// required: true
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func Fetch(db *gorm.DB, v interface{}) ([]interface{}, error) {
	if rows, err := db.Rows(); err != nil {
		return nil, err
	} else {
		var l []interface{}
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = reflect.Indirect(val)
		}
		t := val.Type()
		defer rows.Close()
		for rows.Next() {
			fmt.Printf("row=%v\n", rows)
			v := reflect.New(t).Interface()
			if err := db.ScanRows(rows, v); err != nil {
				return nil, err
			} else {
				fmt.Println("appended")
				l = append(l, v)
			}
			fmt.Printf("l=%v\n", l)
		}
		return l, nil
	}
}

func TestMySQL() {

	var db *gorm.DB
	db, err := gorm.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	type Like struct {
		ID int `gorm:"primary_key"`
		Ip string `gorm:"type:varchar(20); not null; index:ip_idx"`
		Ua string `gorm:"type:varchar(256); no null; "`
		Title string `gorm:"type:varchar(128); not null; index: title_idx"`
		CreatedAt time.Time
	}

	if db.HasTable(&Like{}) {
		fmt.Println("Has table Like")
	} else {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}); err!= nil {
			panic(err)
		}
	}

	// add
	if err := db.Create(&Like{ID: 1, Ip: "127.0.0.1", Ua: "uauauaua", Title: "this is title 111", CreatedAt: time.Now()}); err != nil {
		fmt.Println(err)
	}
	if err := db.Create(&Like{ID: 2, Ip: "127.0.0.1", Ua: "uauauaua", Title: "this is title 222", CreatedAt: time.Now()}); err != nil {
		fmt.Println(err)
	}
	if err := db.Create(&Like{ID: 3, Ip: "127.0.0.1", Ua: "uauauaua", Title: "this is title 333", CreatedAt: time.Now()}); err != nil {
		fmt.Println(err)
	}
	if err := db.Create(&Like{ID: 4, Ip: "127.0.0.1", Ua: "uauauaua", Title: "this is title 444", CreatedAt: time.Now()}); err != nil {
		fmt.Println(err)
	}
	if err := db.Create(&Like{ID: 5, Ip: "127.0.0.1", Ua: "uauauaua", Title: "this is title 555", CreatedAt: time.Now()}); err != nil {
		fmt.Println(err)
	}

	// find
	fmt.Println("===========find==============")
	like := Like{}
	rs := db.Model(&Like{}).Where("id = ?", "1")
	rs.First(&like)
	fmt.Println(like)
	fmt.Println(like.ID)

	// count
	fmt.Println("==========count=========")
	count := -1
	db.Model(&Like{}).Count(&count)
	fmt.Printf("count=%d\n", count)

	// update
	db.Model(&like).Update("title", "new title")
	fmt.Println(like)

	// delete
	fmt.Println(rs)
	db.Where(like).Delete(Like{})
	db.Model(&Like{}).Count(&count)
	fmt.Printf("count=%d\n", count)

	l, err := Fetch(db.Table("likes"), Like{})
	if err != nil {
		panic(err)
	} else {
		for _, each := range l {
			v := reflect.Indirect(reflect.ValueOf(each))
			fmt.Println(v.FieldByName("Title"))
		}
	}
}

// User Info
//
// swagger:response UserResponse
type UserWapper struct {
	Body ResponseMessage
	ID uint32
}

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "this is v1")
		})
	}

	a := jsoniter.RawMessage(`{}`)
	fmt.Println(a)
	//
	//v2 := r.Group("/v2")
	//{
	//	v2.GET("/test", func(c *gin.Context) {
	//		c.String(http.StatusOK, "this is v2")
	//	})
	//}
	//r.GET("/ping", getting)
	//r.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	c.String(http.StatusOK, "name=%s, action=%s", name, action)
	//})
	//
	//// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	//r.GET("/welcome", func(c *gin.Context) {
	//	firstName := c.DefaultQuery("first", "Guest")
	//	lastName := c.Query("last")
	//	c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	//})
	//
	//r.POST("/form_post", func(c *gin.Context) {
	//	message := c.PostForm("message")
	//	nick := c.DefaultPostForm("nick", "anonymous")
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": "posted",
	//		"message": message,
	//		"nick": nick,
	//	})
	//})
	//
	//// Binding
	//r.POST("/loginJSON", func(c *gin.Context) {
	//	var json Login
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	//if json.User != "root" || json.Password != "123" {
	//	//	c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
	//	//	return
	//	//}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})

	// MySQL
	//TestMySQL()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	// swagger:route GET /bookable getBookable
	//
	// get a bookable date
	//
	// This will book a date
	//
	//		Responses:
	//			200: UserResponse
	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBindWith(&b, binding.Query); err == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(http.StatusOK, names)
	})

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.Run()
}