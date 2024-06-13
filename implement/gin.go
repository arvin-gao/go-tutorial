package implement

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func config(r *gin.Engine) {
	//? setting trust proxies IP
	r.SetTrustedProxies([]string{"192.168.x.x"})
	//? Use predefined header gin.PlatformXXX
	r.TrustedPlatform = gin.PlatformGoogleAppEngine
	//? Or set your own trusted request header for another trusted proxy service
	//? Don't set it to any suspect request header, it's unsafe
	r.TrustedPlatform = "X-CDN-IP"

	//* setting max memory for forms(default = 32MiB)
	// 8 MiB
	r.MaxMultipartMemory = 8 << 20

	_ = &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func logFile(r *gin.Engine) {
	//* Console Color
	//? Disable color. When writing the logs to file.
	gin.DisableConsoleColor()
	//? Always colorize logs.
	gin.ForceConsoleColor()
	//* Setting log file for the server
	f, _ := os.Create("my-log.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//* Custom log format
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// ::1 - [Fri, 07 Dec 2018 17:04:38 JST] "GET /ping HTTP/1.1 200 122.767µs "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36" "
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	//? Define format for the log of routes
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

type Person struct {
	//? Tags: form, json, xml, binding
	User string `form:"user" json:"user" xml:"user"  binding:"required"`
	//? HTML checkbox
	CheckOption string `form:"opt[]"` // <input type="checkbox" name="opt[]" value="red" id="red">
	//? Upload ## For multi file: "[]*multipart.FileHeader"
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
	//? Form time
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
	//? Uri
	Uri string `uri:"uri" binding:"required"`
	//? Header
	Header int `header:"head"`
	//? Struct embedding
	Items Item
}
type Item struct {
	Name string `form:"name"`
}

// JSON:    c.ShouldBindJSON
// XML:     c.ShouldBindXML
// YAML:    c.ShouldBindYAML
// Uri:     c.ShouldBindUri #e.g. -> r.Get("/:uri/")
// Header:  c.ShouldBindHeader
// FORM:
//
//	Query & Post: c.ShouldBind
//	Only Query	 : c.ShouldBindQuery
//
// Custom
//
//		# binding.Form, binding.JSON, binding.XML...
//	  c.ShouldBindWith(&, binding.Form)
//
// First, Condition binding
func bindJSON(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		// pass
	}
}

func bindURI(r gin.RouterGroup) {
	var person Person
	r.GET("/:uri", func(c *gin.Context) {
		if err := c.ShouldBindUri(&person); err != nil {
			// pass
		}
	})
}

func bindHeader(c *gin.Context) {
	// > curl -H "rate:300" -H "domain:music" 127.0.0.1/
	h := struct {
		Rate   int    `header:"rate"`
		Domain string `header:"domain"`
	}{}
	if err := c.ShouldBindHeader(&h); err != nil {
		// pass
	}
	// output: {"Rate":300, "Domain":"music"}
	c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
}

func uploadFiles(c *gin.Context) {
	//  > curl -X POST http://localhost:8080/upload \
	//  	-F "file=@/Users/name/test.zip" \
	// 		-H "Content-Type: multipart/form-data"

	//? Method - 1 with bind
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		// pass
	}

	if err := c.SaveUploadedFile(person.Avatar, person.Avatar.Filename); err != nil {
		// pass
	}

	//? Method - 2 with form directly
	file, _ := c.FormFile("file")
	_ = c.SaveUploadedFile(file, "file name??")

	//? For the multi files
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		_ = c.SaveUploadedFile(file, "file name??")
	}
	//! `file.Filename` SHOULD NOT be trusted. See `Content-Disposition` on MDN and #1693
	//! MDN: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#directives
	//! #1693: https://github.com/gin-gonic/gin/issues/1693
}

func response() {
	//* Response
	// String
	// 		c.String(http.StatusBadRequest, "bad request")
	// JSON, XML, YAML
	// 		c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	// Particular JSON
	// 		c.SecureJSON(), c.JSONP(), c.AsciiJSON(), c.PureJSON(200. gin.H{"html":"<p>hi</p>"})
	// HTML
	//    c.HTML(200, "xx.html", gin.H{"title":""})
	// other
	// 		c.Status(200)
	// Redirect
	//    c.Redirect(<code>, <URL>)
	//
	//    c.Request.URL.Path = "/test2"
	//    r.HandleContext(c)
	//
	// Proto buffer
	// 		c.ProtoBuf(http.StatusOK, data)
}

func serveFile(r *gin.Engine) {
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})

	var fs http.FileSystem
	r.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("fs/file.go", fs)
	})
}

func setupMiddleware(r *gin.Engine) {
	//* Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	//? Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	//? Per r middleware, you can add as many as you desire.
	middleware := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			if c.Query("k") == "v" {
				// pass
			}
			t := time.Now()
			//? Set example variable
			c.Set("example", "12345")
			//? before request
			c.Next()
			//? after request. latency
			log.Print(time.Since(t))
			//? access the status we are sending. status
			log.Println(c.Writer.Status())
			//? Abort the process currently
			c.Abort()
		}
	}
	r.GET("/benchmark", middleware(), nil)

	//* Group middleware
	group := r.Group("/")
	group.Use(middleware())
	{
		group.POST("/xxx", nil)
	}

}

func setupCookie(c *gin.Context) {
	cookie, err := c.Cookie("username")
	if err != nil {
		c.SetCookie(
			"username",
			"user",
			3600,
			"/",
			"localhost",
			false,
			true,
		)
	}
	fmt.Println(cookie)
}

func setupHTML(r *gin.Engine) {
	{
		r.LoadHTMLGlob("templates/*")
		html := template.Must(template.ParseFiles("file1", "file2"))
		r.SetHTMLTemplate(html)
	}

	{
		r.LoadHTMLFiles("templates/index.html", "templates/second.html")
		r.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Title value",
			})
		})
	}
}

func static(r *gin.Engine) {
	r.Static("/assets", "./assets")
	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	//? undefined !?
	// r.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))
}

func testingRequest(r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
}
