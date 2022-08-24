package main

import (
	"context"
	"fmt"
	"strconv"

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
	host     = "localhost"
	port     = 27017
	database = "gomongo"
)

func getURL() string {
	//Ideal pasar todo a un archivo .env
	//En Go existe el package os, que nos permite obtener variables de entorno por el metodo Getenv([nombre_de_la_variable])
	return fmt.Sprintf("mongodb://%s:%d/%s", host, port, database)
}
func ConnectDb() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getURL()))
	return client, err
}
func DisconnectDb(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()

	//Conexion a la base de datos
	client, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer DisconnectDb(client)

	collUser := client.Database("gomongo").Collection("users")
	collEvent := client.Database("gomongo").Collection("events")

	collEvent.InsertOne(context.TODO(), bson.D{{
		Key:   "title",
		Value: "Nashe",
	}})

	collUser.InsertOne(context.TODO(), bson.D{{
		Key:   "name",
		Value: "Nacho",
	}})
	//Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Hello world")
	})

	//Este listado de eventos se debe poder filtrar por fecha, estado, y título.
	app.Get("/event", func(c *fiber.Ctx) error {
		var events models.Events
		var filter bson.D

		if c.Query("title", "") != "" || c.Query("date", "") != "" || c.Query("state", "") != "" {

			if c.Query("title", "") != "" {
				filter = bson.D{{Key: "title", Value: c.Query("title", "")}}

				if c.Query("date", "") != "" {
					fmt.Println(filter.Map())
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
					fmt.Println(filter.Map())
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
				filter = bson.D{{Key: "state", Value: c.Query("state", "")}}

				if c.Query("title", "") != "" {
					fmt.Println(filter.Map())
					filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "state", Value: c.Query("state", "")}}

					if c.Query("date", "") != "" {
						boolValue, _ := strconv.ParseBool(c.Query("state", ""))
						filter = bson.D{{Key: "title", Value: c.Query("title", "")}, {Key: "date", Value: c.Query("date", "")}, {Key: "state", Value: boolValue}}
					}
				} else if c.Query("date", "") != "" {
					boolValue, _ := strconv.ParseBool(c.Query("state", ""))
					filter = bson.D{{Key: "date", Value: c.Query("date", "")}, {Key: "state", Value: boolValue}}
				}
			}

		} else {
			filter = bson.D{}
		}

		res, err := collEvent.Find(context.TODO(), filter)

		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(err)
		}
		fmt.Println(err)
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

		event.Id = uuid.NewString()
		event.State = true
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

		inscriptos := append(event.Inscriptos, incripcion.UserId)
		fmt.Println(inscriptos)

		filter := bson.M{"id": incripcion.EventId}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "inscriptos", Value: inscriptos}}}}
		//var eventModified models.Event
		res, err := collEvent.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			panic(err)
		}
		fmt.Println(res)
		return c.Status(fiber.StatusAccepted).JSON("Inscripto exitosamente")
	})

	app.Get("/user", func(c *fiber.Ctx) error {

		name := c.Query("name", "")

		var user models.User
		collUser.FindOne(context.TODO(), bson.M{"name": name}).Decode(&user)

		return c.Status(fiber.StatusAccepted).JSON(user)
	})
	app.Post("/user", func(c *fiber.Ctx) error {

		var user models.User

		c.BodyParser(&user)
		collUser.InsertOne(context.TODO(), user)
		return c.Status(fiber.StatusAccepted).JSON(user)
	})
	app.Listen(":3000")
}
