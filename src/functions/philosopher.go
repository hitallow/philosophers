package functions

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hitallow/philosophers/src/models"
)

// LoadPhilosophers : load a json file
func LoadPhilosophers() (*models.Philosophers, error){
	jsonFile,err := os.Open("philosophers.json")
	if err != nil {
		return nil,err
	}
	
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	
	var philosophers *models.Philosophers

	json.Unmarshal(byteValue, &philosophers)
	
	return philosophers,nil
}


// SearchPhilosopherByName : search a philosopher by name
func SearchPhilosopherByName(philosopherName string, philosophers models.Philosophers)(*models.Philosopher,error){
	for i := 0; i < len(philosophers.Philosophers); i++ {
		var atualPhilosopher models.Philosopher
	
		atualPhilosopher =  philosophers.Philosophers[i]
	
		if strings.ToLower(atualPhilosopher.Name) == strings.ToLower(philosopherName) {	
			return &atualPhilosopher,nil
		}
	}

	return nil,errors.New("philosopher not found")
}