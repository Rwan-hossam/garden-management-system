package main

import "fmt"

type Plant struct {
	Id               string      `json:"id"`
	Name             string      `json:"name"`
	plantType        typeOfPlant `json:"type"`
	PlantingSeason   string      `json:"planting_season"`
	CareInstructions string      `json:"care_instructions"`
}

type typeOfPlant int

const (
	Vegetable typeOfPlant = iota
	Fruit
	Herb
)

func (pt typeOfPlant) String() string {
	switch pt {
	case Vegetable:
		return "Vegetable"
	case Fruit:
		return "Fruit"
	case Herb:
		return "Herb"
	default:
		return "Unknown"
	}
}
func AddPlant(plants *[]Plant, newPlant Plant) error {

	for _, plant := range *plants {
		if plant.Id == newPlant.Id {
			return fmt.Errorf("Error: A plant with the same ID (%s) already exists", newPlant.Id)

		}
	}

	*plants = append(*plants, newPlant)
	return nil
}

func RemovePlant(plants *[]Plant, id string) string {
	for i, plant := range *plants {
		if plant.Id == id {
			*plants = append((*plants)[:i], (*plants)[i+1:]...)
			return (plant.Name + " removed successfully")
		}
	}
	return ("Plant with given ID not found")
}

func UpdatePlant(plants *[]Plant, id string, updatedName string, updatedPlantType typeOfPlant, updatedPlantingSeason string, updatedCareInstructions string) string {
	for i, plant := range *plants {
		if plant.Id == id {
			(*plants)[i].Name = updatedName
			(*plants)[i].plantType = updatedPlantType
			(*plants)[i].PlantingSeason = updatedPlantingSeason
			(*plants)[i].CareInstructions = updatedCareInstructions
			return (plant.Name + " Updated successfully ")
		}

	}
	return ("plant with given ID not found")
}

func SearchPlants(plants []Plant, searchedPlant string) []Plant {
	var results []Plant
	for _, plant := range plants {
		if plant.Name == searchedPlant || plant.PlantingSeason == searchedPlant {
			results = append(results, plant)
		}
	}
	return results
}
func GenerateHarvestReport(plants []Plant) {
	report := make(map[string]int)
	for _, plant := range plants {
		report[plant.plantType.String()]++
	}

	fmt.Println("Harvest Summary:")
	for plantType, count := range report {
		fmt.Printf("%s: %d plants\n", plantType, count)
	}
}
