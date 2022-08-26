package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/NachooNazar/prueba-tecnica-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// URI bd
var (
	// Obviously, this is just a test example. Do not do this in production.
	// In production, you would have the private key and public key pair generated
	// in advance. NEVER add a private key to any GitHub repo.
	host     = "localhost"
	port     = 27017
	database = "gomongo"
)

func getURL() string {
	//Ideal pasar todo a un archivo .env
	//En Go existe el package os, que nos permite obtener variables de entorno por el metodo Getenv([nombre_de_la_variable])
	return fmt.Sprintf("mongodb://%s:%d/%s", host, port, database)
}
func ConnectDb() (*mongo.Client, *mongo.Collection, *mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getURL()))
	collUser := client.Database("gomongo").Collection("users")
	collEvent := client.Database("gomongo").Collection("events")
	return client, collUser, collEvent, err
}
func DisconnectDb(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

// func createJWTToken(user models.User) (string, int64, error) {

// 	exp := time.Now().Add(time.Minute * 30).Unix()
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["user_id"] = user.Id
// 	claims["exp"] = exp
// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return "", 0, err
// 	}
// 	return t, exp, nil

// }

func main() {
	app := fiber.New()

	//Conexion a la base de datos
	client, collUser, collEvent, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer DisconnectDb(client)

	//Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	app.Post("/login", func(c *fiber.Ctx) error {

		type request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var body request
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse json",
			})

		}
		var user models.User
		collUser.FindOne(context.TODO(), bson.M{"email": body.Email}).Decode(&user)
		fmt.Println(user)
		if user.Email != body.Email || user.Password != body.Password {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Bad credentials",
			})

		}

		// token, exp, err2 := createJWTToken(user)
		// if err2 != nil {
		// 	fmt.Println(err2)
		// 	fmt.Println(exp)
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error": "could not login",
		// 		"err":   err2,
		// 	})
		// }
		// cookie := fiber.Cookie{
		// 	Name:     "jwt",
		// 	Value:    token,
		// 	Expires:  exp,
		// 	HTTPOnly: true,
		// }

		///c.Cookie(&cookie)

		// c.Set("Authorization", token)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Success": "could login",
		})
	})

	//Este listado de eventos se debe poder filtrar por fecha, estado, y título.
	app.Get("/event", func(c *fiber.Ctx) error {
		var events models.Events
		var filter bson.D

		if c.Query("title", "") != "" || c.Query("date", "") != "" || c.Query("state", "") != "" {

			if c.Query("title", "") != "" {
				filter = bson.D{{Key: "title", Value: c.Query("title", "")}}

				if c.Query("date", "") != "" {

					filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "date", Value: c.Query("date", "")}}

					if c.Query("state", "") != "" {
						boolValue, _ := strconv.ParseBool(c.Query("state", ""))
						filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "date", Value: c.Query("date", "")}, {Key: "state", Value: boolValue}}
					}
				} else if c.Query("state", "") != "" {
					boolValue, _ := strconv.ParseBool(c.Query("state", ""))
					filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "state", Value: boolValue}}
				}
			}

			if c.Query("date", "") != "" {
				filter = bson.D{{Key: "date", Value: c.Query("date", "")}}

				if c.Query("title", "") != "" {

					filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "date", Value: c.Query("date", "")}}

					if c.Query("state", "") != "" {
						boolValue, _ := strconv.ParseBool(c.Query("state", ""))
						filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "date", Value: c.Query("date", "")}, {Key: "state", Value: boolValue}}
					}
				} else if c.Query("state", "") != "" {
					boolValue, _ := strconv.ParseBool(c.Query("state", ""))
					filter = bson.D{{Key: "date", Value: c.Query("date", "")}, {Key: "state", Value: boolValue}}
				}
			}

			if c.Query("state", "") != "" {
				boolValue, _ := strconv.ParseBool(c.Query("state", ""))
				filter = bson.D{{Key: "state", Value: boolValue}}

				if c.Query("title", "") != "" {

					filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "state", Value: c.Query("state", "")}}

					if c.Query("date", "") != "" {
						boolValue, _ := strconv.ParseBool(c.Query("state", ""))
						filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "date", Value: c.Query("date", "")}, {Key: "state", Value: boolValue}}
					}
				} else if c.Query("date", "") != "" {

					filter = bson.D{{Key: "date", Value: c.Query("date", "")}, {Key: "state", Value: boolValue}}
				}
			}

		} else {
			filter = bson.D{}
		}

		res, err := collEvent.Find(context.TODO(), filter)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		for res.Next(context.TODO()) {
			var event models.Event
			res.Decode(&event)
			events = append(events, event)
		}

		// var eventts [][]string
		// for i := range events {
		// 	eventts = append(eventts, events[i].Inscriptos)
		// }
		// fmt.Println(eventts)
		return c.Status(fiber.StatusOK).JSON(events)
	})
	app.Post("/event", func(c *fiber.Ctx) error {
		var event models.Event

		if err = c.BodyParser(&event); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		actualDate := time.Now()
		var eventDate time.Time
		eventDate, err3 := time.Parse("02/01/2006", event.Date)
		if err3 != nil {
			return c.Status(fiber.StatusBadRequest).JSON("invalid date")
		}
		event.Id = uuid.NewString()
		canGo := calcDateRecent(actualDate, eventDate)
		if !canGo {
			return c.Status(fiber.StatusBadRequest).JSON("You cant go back into the time")
		}
		//event.State = true
		res, err2 := collEvent.InsertOne(context.TODO(), event)

		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err2)
		}
		return c.Status(fiber.StatusOK).JSON(res)
	})
	app.Put("/event", func(c *fiber.Ctx) error {
		type incription struct {
			EventId string `json:"eventID"`
			UserId  string `json:"userID"`
		}

		var incripcion incription

		if err := c.BodyParser(&incripcion); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		if incripcion.EventId == "" || incripcion.UserId == "" {
			return c.Status(fiber.StatusBadRequest).JSON("Bad request")
		}

		var user models.User
		collUser.FindOne(context.TODO(), bson.M{"id": incripcion.UserId}).Decode(&user)
		if len(user.Name) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON("Invalid user")
		}

		var event models.Event
		collEvent.FindOne(context.TODO(), bson.M{"id": incripcion.EventId}).Decode(&event)
		if len(event.Title) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON("Invalid event")
		}

		if !event.State {
			return c.Status(fiber.StatusBadRequest).JSON("Event closed")
		}

		//Puedo acceder a diferentes atributos -> year, month, day, hour, minute, second
		actualDate := time.Now()
		var eventDate time.Time
		eventDate, err3 := time.Parse("02/01/2006", event.Date)
		if err3 != nil {
			return c.Status(fiber.StatusBadRequest).JSON("invalid date")
		}

		canGo := calcDateRecent(actualDate, eventDate)
		if !canGo {
			return c.Status(fiber.StatusBadRequest).JSON("You cant go back into the time")
		}
		//Agrega evento al user
		myEvents := append(user.MyEvents, incripcion.EventId)
		filter := bson.M{"id": incripcion.UserId}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "myEvents", Value: myEvents}}}}
		_, err5 := collUser.UpdateOne(context.TODO(), filter, update)
		if err5 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err5)
		}
		//Agrega user al evento
		inscriptos := append(event.Inscriptos, incripcion.UserId)
		filter = bson.M{"id": incripcion.EventId}
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "inscriptos", Value: inscriptos}}}}
		_, err6 := collEvent.UpdateOne(context.TODO(), filter, update)
		if err6 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err6)
		}

		return c.Status(fiber.StatusAccepted).JSON("Inscripto exitosamente")
	})
	app.Get("/event/own", func(c *fiber.Ctx) error {
		var events models.Events

		userId := c.Query("userId")
		query := c.Query("time", "")

		fmt.Println(userId)
		var user models.User
		collUser.FindOne(context.TODO(), bson.M{"id": userId}).Decode(&user)

		for i := 0; i < len(user.MyEvents); i++ {
			filter := bson.M{"date": bson.M{"$gte": query}, "id": user.MyEvents[i]}
			var event models.Event
			res := collEvent.FindOne(context.TODO(), filter)
			res.Decode(&event)
			if len(event.Title) > 0 {
				events = append(events, event)
			}
		}

		return c.Status(fiber.StatusOK).JSON(events)
	})
	app.Put("/event/:id", func(c *fiber.Ctx) error {
		eventID := c.Params("id")
		var eventBody models.Event
		c.BodyParser(&eventBody)

		var eventRes models.Event
		collEvent.FindOne(context.TODO(), bson.D{{Key: "id", Value: eventID}}).Decode(&eventRes)
		eventBody.Id = eventID
		eventBody.Inscriptos = eventRes.Inscriptos

		if eventBody.Date == "" {
			eventBody.Date = eventRes.Date
		}
		if eventBody.Organizer == "" {
			eventBody.Organizer = eventRes.Organizer
		}
		if eventBody.Hour == "" {
			eventBody.Hour = eventRes.Hour
		}
		if eventBody.Place == "" {
			eventBody.Place = eventRes.Place
		}
		if eventBody.Title == "" {
			eventBody.Title = eventRes.Title
		}
		if eventBody.ShortDescription == "" {
			eventBody.ShortDescription = eventRes.ShortDescription
		}
		if eventBody.LargeDescription == "" {
			eventBody.LargeDescription = eventRes.LargeDescription
		}
		if !eventBody.State {
			eventBody.State = eventRes.State
		}

		res, err := collEvent.UpdateOne(context.TODO(), bson.M{"id": eventID}, bson.D{{Key: "$set", Value: eventBody}})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
		//collEvent.FindOne(context.TODO(), bson.D{{Key: "id", Value: eventID}}).Decode(&eventRes)
		return c.Status(fiber.StatusOK).JSON(res)
	})

	//RUTAS ADMIN

	app.Get("/user", func(c *fiber.Ctx) error {

		name := c.Query("name", "")
		var filter bson.M

		if name == "" {
			var users models.Users
			filter = bson.M{}
			res, err := collUser.Find(context.TODO(), filter)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(users)
			}

			if err := res.All(context.TODO(), &users); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(err)
			}

			return c.Status(fiber.StatusAccepted).JSON(users)
		} else {
			var user models.User
			filter = bson.M{"name": name}
			collUser.FindOne(context.TODO(), filter).Decode(&user)
			return c.Status(fiber.StatusAccepted).JSON(user)
		}
	})
	app.Post("/user", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user)
		validUser := validateUser(user)
		if !validUser {
			return c.Status(fiber.StatusBadRequest).JSON("Invalid user")
		}
		user.Id = uuid.NewString()
		collUser.InsertOne(context.TODO(), user)
		return c.Status(fiber.StatusAccepted).JSON(user)
	})

	app.Listen(":3000")
}

func calcDateRecent(currentDate time.Time, eventDate time.Time) bool {
	minutesCurrentDate := int(currentDate.Year())*525600 + int(currentDate.Month())*43800 + int(currentDate.Day())*1440
	minutesEventDate := int(eventDate.Year())*525600 + int(eventDate.Month())*43800 + int(eventDate.Day())*1440
	if minutesCurrentDate > minutesEventDate {
		return false
	} else {
		return true
	}
}

func validateUser(user models.User) bool {
	if user.Name != "" && user.Email != "" && user.Lastname != "" && user.Password != "" {
		return true
	}
	return false
}
