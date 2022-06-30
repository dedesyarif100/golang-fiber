package main

import (
	// "bytes"
	// "context"
	"fmt"
	"go-fiber-playground/storage"
	"log"
	"os"
	// "reflect"
	// "time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// type Book struct {
// 	Author    string `json:"author"`
// 	Title     string `json:"title"`
// 	Publisher string `json:"publisher"`
// }

type Repository struct {
	DB *gorm.DB
}

// type Person struct {
//     Name string `json:"name" xml:"name" form:"name"`
//     Pass string `json:"pass" xml:"pass" form:"pass"`
// }

type SomeStruct struct {
	Name string
	Age  uint8
}

type Person struct {
    UserName		string		`query:"name" reqHeader:"name"`
    Password		string		`query:"pass" reqHeader:"pass"`
    OurProducts		[]string	`query:"products" reqHeader:"products"`
}

func (r *Repository) ContextSaveFile(c *fiber.Ctx) error {
	fmt.Println(c.MultipartForm())
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["file"]
		fmt.Println("FILES :", files)
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			if err := c.SaveFile(file, fmt.Sprintf("file/%s", file.Filename)); err != nil {
				return err
			}
		}
		return err
	}
	return nil
}

func (r *Repository) handler(c *fiber.Ctx) error {
	fmt.Println("MASUK HANDLER")

	// # BaseURL
		// fmt.Println(c.BaseURL())

	// # Bind
		// data := c.Bind(fiber.Map{
		// 	"Title": "Hello, World!",
		// })
		// fmt.Println(data)
		// nama := "DEDE SYARIFUDIN HIDAYAT"
		// return c.Render("template/tmp.html", fiber.Map{
		// 	"Name": nama,
		// })

	// # Body
		// return c.Send(c.Body())

	// # BodyParser
		// p := new(Person)
		// if err := c.BodyParser(p); err != nil {
		// 	return err
		// }
		// // fmt.Println(p)
		// log.Println(p.Name)
		// log.Println(p.Pass)
		// return nil

	// Cookie
		// cookie := new(fiber.Cookie)
		// cookie.Name = "name"
		// cookie.Value = "dede"
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// c.Cookie(cookie)
		// // c.ClearCookie()
		// return nil

	// Cookies
		// c.Cookies("name")
		// fmt.Println(c.Cookies("name2", "febby"))
		// return nil

	// Download
		// return c.Download("file/data.pdf", "report.pdf");

	// Format
		// data := []int{1, 2, 3, 4, 5}
		// c.Format(data)
		// // c.Format("TES")
		// return nil

	// # FormFile
		// file, _ := c.FormFile("file/output")
		// // if err != nil {
		// // 	log.Fatal(err)
		// // }
		// return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

	// # FormValue
		// c.FormValue("name")
		// return nil

	// # Get
		// c.Get("something", "john")
		// return nil

	// # HostName
		// fmt.Println(c.Hostname())
		// return nil

	// # IP
		// fmt.Println(c.IP())
		// return nil

	// # IPs
		// fmt.Println(c.IPs())
		// return nil

	// # Is
		// fmt.Println(c.Is("json"))
		// return nil

	// # IsFromLocal
		// fmt.Println(c.IsFromLocal())
		// return nil

	// # JSON
		// data := SomeStruct{
		// 	Name: "Grame",
		// 	Age:  20,
		// }
		// // return c.JSON(data)
		// return c.JSON(fiber.Map{
		// 	"name": data.Name,
		// 	"age": data.Age,
		// })

	// # JSONP
		// data := SomeStruct{
		// 	Name: "Grame",
		// 	Age:  20,
		// }
		// // return c.JSONP(data)
		// return c.JSONP(data, "customFunc")

	// # Links
		// c.Links(
		// 	"http://api.example.com/users?page=2", "next",
		// 	"http://api.example.com/users?page=5", "last",
		// )
		// return nil

	// # Locals
		// fmt.Println(c.Locals("user"))
		// // return nil
		// return c.Status(fiber.StatusOK).SendString("Welcome, admin!")

	// # Location
		// return nil

	// # Method
		// fmt.Println(c.Method())
		// return nil

	// # MultipartForm

	// # Next
		// fmt.Println("1st route!")
		// return c.Next()

	// # OriginalURL
		// fmt.Println(c.OriginalURL())
		// return nil

	// # Params
		// c.Params("name")
		// return nil

	// # ParamsInt
		// id, _ := c.ParamsInt("id")
		// fmt.Println(id)
		// return nil

	// # Path
		// fmt.Println(c.Path())
		// return nil

	// # Protocol
		// fmt.Println(c.Protocol())
		// return nil

	// # Query
		// fmt.Println(c.Query("id"))
		// fmt.Println(c.Query("name"))
		// return nil

	// # QueryParser
		// p := new(Person)
		// if err := c.QueryParser(p); err != nil {
        //     return err
        // }
		// log.Println(p.UserName)     // john
        // log.Println(p.Password)     // doe
        // log.Println(p.OurProducts) // [shoe, hat]
		// return nil

	// # Range

	// # Redirect
		// return c.Redirect("/user/dede")

	// # RedirectToRoute
		// return c.RedirectToRoute("user", fiber.Map{
		// 	"name": "fiber",
		// 	"queries": map[string]string{"data[0][name]": "john", "data[0][age]": "10", "test": "doe"},
		// })

	// # RedirectBack
		// return c.SendString("Home page")
		// c.Set("Content-Type", "text/html")
		// return c.SendString(`<a href="/">Back</a>`)
		// return c.RedirectBack("/")

	// # Render
		// return c.Render("index", fiber.Map{
		// 	"Title": "Hello, World!",
		// })

	// # Request
		// fmt.Println(string(c.Request().Header.Method()))
		// return nil

	// # ReqHeaderParser
		// p := new(Person)
		// if err := c.ReqHeaderParser(p); err != nil {
		// 	return err
		// }
		// log.Println(p.UserName)     // john
        // log.Println(p.Password)     // doe
        // log.Println(p.OurProducts) // [shoe, hat]
		// return nil

	// # Response
		// fmt.Println(c.Response())
		// return nil
		
	// # RestartRouting
		// c.Path("/")
		// return c.RestartRouting()

	// # Route
		// app.Get("/user/:id", r.handler)
		// route := c.Route()
		// fmt.Println(route.Method, route.Path, route.Params, route.Handlers)
		// return nil

	// # SaveFile
		// Parse the multipart form:
		// return c.Render("uploadFile", fiber.Map{})

	// # SaveFileToStorage

	// # Secure
		// fmt.Println(c.Protocol())
		// return nil

	// # Send
		// return c.Send([]byte("Hello, World!"))
		// return c.SendString("Hello, World!")
		// return c.SendStream(bytes.NewReader([]byte("DEDE SYARIF")))

	// # SendFile
		// return c.SendFile("template/404.html", false);

	// # SendStatus
		// c.SendString("Hello, World!")
		// return c.SendStatus(415)

	// # Set
		// c.Set("Content-Type", "text/plain")
		// return nil

	// # SetParserDecoder
		
	// # SetUserContext
		// ctx := context.Background()
		// c.SetUserContext(ctx)
		// return nil

	// # Stale

	// # Status
		// fmt.Println(c.Status(fiber.StatusOK))
		// fmt.Println(reflect.TypeOf(c.Status(fiber.StatusOK)))
		// return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		// return c.Status(fiber.StatusNotFound).SendFile("file/404.png")

	// # Subdomains
		// fmt.Println(c.Subdomains())
		// return nil

	// # Type
		// fmt.Println(c.Type("html"))
		// c.Type("png")
		// return nil

	// # UserContext
		// ctx := c.UserContext()
		// fmt.Println(ctx)
		// return nil

	// # Vary
		// c.Vary("Origin")     // => Vary: Origin
		// c.Vary("User-Agent") // => Vary: Origin, User-Agent
	
		// // No duplicates
		// c.Vary("Origin") // => Vary: Origin, User-Agent
	
		// c.Vary("Accept-Encoding", "Accept")
		// // => Vary: Origin, User-Agent, Accept-Encoding, Accept
		// return nil

	// # Write
		// c.Write([]byte("Hello, World!")) // => "Hello, World!"
		// fmt.Fprintf(c, "\n %s \n", "Hello, World!") // "Hello, World!Hello, World!"
		// return nil

	// # Writef
		// world := "World!"
		// c.Writef("Hello, %s", world) // => "Hello, World!"
		// fmt.Fprintf(c, "\n %s \n", "Hello, World!") // "Hello, World!Hello, World!"
		// return nil

	// # WriteString
		// c.WriteString("Hello, World!") // => "Hello, World!"
		// fmt.Fprintf(c, "\n %s \n", "Hello, World!") // "Hello, World!Hello, World!"
		// return nil

	// # XHR
		fmt.Println(c.XHR())
		return nil





		


	// # Write
		// c.Write([]byte("Hello, World!"))
		// return nil

	// return c.JSON(c.App().Stack())
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	app.Get("/upload-file", r.handler)
	app.Post("/upload-file", r.ContextSaveFile)

	app.Get("/user/:id", r.handler)
	api := app.Group("/api", r.handler)
	v1 := api.Group("/v1", r.handler)
	v1.Get("/list", r.handler)
	v1.Get("/user", r.handler)
	v1.Get("/user/:id", r.handler)
	// v1.Post("/user", r.handler)
	app.Post("/", r.handler)

	api2 := app.Group("/api2", r.handler)
	v2 := api2.Group("/v2", r.handler)
	v2.Get("/list", r.handler)
	v2.Get("/user", r.handler)

	app.Get("/user/:name", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	}).Name("user")

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("MASUK BACK")
		return nil
	})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host: 		os.Getenv("DB_HOST"),
		Port: 		os.Getenv("DB_PORT"),
		Password: 	os.Getenv("DB_PASS"),
		User: 		os.Getenv("DB_USER"),
		SSLMode: 	os.Getenv("DB_SSLMODE"),
		DBName: 	os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repository{
		DB: db,
	}

	// app := fiber.New()
	engine := html.New("./template", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

