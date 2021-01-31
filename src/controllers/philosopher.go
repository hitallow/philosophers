package controllers

import (
	"math/rand"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hitallow/philosophers/src/functions"
	"github.com/hitallow/philosophers/src/models"
)

// GetPhrase : return a phrase of a philosopher
func GetPhrase(philosopherName string,context fiber.Ctx) error{
	philosophers,err := functions.LoadPhilosophers()

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
					"error":err.Error(),
			},
		)
	}

	for i := 0; i < len(philosophers.Philosophers); i++ {
		var atualPhilosopher models.Philosopher
	
		atualPhilosopher =  philosophers.Philosophers[i]
	
		if strings.ToLower(atualPhilosopher.Name) == strings.ToLower(philosopherName) {	
			return context.Status(fiber.StatusOK).JSON(fiber.Map{
				"autor":atualPhilosopher.Name,
				"phrase":atualPhilosopher.Phrases[rand.Intn(5)],
			})
		}
	}

	return context.Status(fiber.StatusNotFound).JSON(
		fiber.Map{
			"message":  "philosophers not found",
		},
	)
}

// ListPhilosophers : return a list of philosophers
func ListPhilosophers(context fiber.Ctx) error{
	
	list,err := functions.LoadPhilosophers()

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
					"error":err.Error(),
			},
		)
	}

	var philosophersNames []string

	for i := 0; i < len(list.Philosophers); i++ {
		philosophersNames = append(philosophersNames,list.Philosophers[i].Name)	
	}

	return context.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"philosophers":philosophersNames,
		},
	)

}

// ListPhilosopherPhrases : list philospher phrases
func ListPhilosopherPhrases(philosopherName string, context fiber.Ctx) error{

	philosophers,err := functions.LoadPhilosophers()


	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
					"error":err.Error(),
			},
		)
	}

	philosopher,err := functions.SearchPhilosopherByName(philosopherName,*philosophers)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message":  "philosophers not found",
			},
		)
	}

	return context.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"philosopher":philosopher.Name,
			"phrases":philosopher.Phrases,
		},
	)
	
}