package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func LoadPlantsFromFile(filename string) ([]Plant, error) {
	var plants []Plant

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &plants)
	if err != nil {
		return nil, err
	}

	return plants, nil
}

func LoadGardenersFromFile(filename string) ([]Gardener, error) {
	var gardeners []Gardener

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &gardeners)
	if err != nil {
		return nil, err
	}
	return gardeners, nil
}

func SaveGardenersToFile(filename string, gardeners []Gardener) error {
	data, err := json.MarshalIndent(gardeners, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func SavePlantsToFile(filename string, plants []Plant) error {
	data, err := json.MarshalIndent(plants, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func main() {
	plants, err := LoadPlantsFromFile("plants.json")
	if err != nil {
		fmt.Println("Error loading plants:", err)
		return
	}

	gardeners, err := LoadGardenersFromFile("gardeners.json")
	if err != nil {
		fmt.Println("Error loading gardeners:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== Community Garden CLI =====")
		fmt.Println("1. View Plants")
		fmt.Println("2. View Gardeners")
		fmt.Println("3. Add Gardener")
		fmt.Println("4. Remove Gardener")
		fmt.Println("5. Add Plant")
		fmt.Println("6. Remove Plant")
		fmt.Println("7. View Harvest Report")
		fmt.Println("8. Update Gardener")
		fmt.Println("9. Update Plant")
		fmt.Println("10. Search for Plants")
		fmt.Println("11. View Gardener Assignments")

		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\n=== Plants ===")
			for _, plant := range plants {
				fmt.Println("Name:", plant.Name)
				fmt.Println("Type:", plant.plantType.String())
				fmt.Println("Season:", plant.PlantingSeason)
				fmt.Println("Instructions:", plant.CareInstructions)
				fmt.Println("------")
			}

		case 2:
			fmt.Println("\n=== Gardeners ===")
			for _, g := range gardeners {
				fmt.Println("Name:", g.Name)
				fmt.Println("Age:", g.Age)
				fmt.Println("Email:", g.ContactEmail)
				fmt.Println("Plot:", g.PlotNumber)
				fmt.Println("Preferred Plants:", g.PreferredPlants)
				fmt.Println("------")
			}

		case 3:
			var newGardener Gardener
			fmt.Print("Enter ID: ")
			fmt.Scanln(&newGardener.Id)
			fmt.Print("Enter Name: ")
			fmt.Scanln(&newGardener.Name)
			fmt.Print("Enter Age: ")
			fmt.Scanln(&newGardener.Age)
			fmt.Print("Enter Plot Number: ")
			fmt.Scanln(&newGardener.PlotNumber)
			fmt.Print("Enter Contact Email: ")
			fmt.Scanln(&newGardener.ContactEmail)
			fmt.Print("Enter Preferred Plants (comma separated): ")
			var plantsInput string
			fmt.Scanln(&plantsInput)
			newGardener.PreferredPlants = strings.Split(plantsInput, ",")
			err := AddGardener(&gardeners, newGardener)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Gardener added successfully!")
				SaveGardenersToFile("gardeners.json", gardeners)
			}

		case 4:
			fmt.Print("Enter ID of Gardener to remove: ")
			var id string
			fmt.Scanln(&id)
			msg := RemoveGardener(&gardeners, id)
			fmt.Println(msg)
			SaveGardenersToFile("gardeners.json", gardeners)

		case 5:
			var newPlant Plant
			fmt.Print("Enter Plant ID: ")
			fmt.Scanln(&newPlant.Id)
			fmt.Print("Enter Name: ")
			fmt.Scanln(&newPlant.Name)
			fmt.Println("Enter Plant Type (0 = Vegetable, 1 = Fruit, 2 = Herb): ")
			fmt.Scanln(&newPlant.plantType)

			fmt.Print("Enter Planting Season: ")
			fmt.Scanln(&newPlant.PlantingSeason)

			fmt.Print("Enter Care Instructions: ")
			fmt.Scanln(&newPlant.CareInstructions)

			err := AddPlant(&plants, newPlant)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Plant added successfully!")
				SavePlantsToFile("plants.json", plants)
			}

		case 6:
			fmt.Print("Enter Plant ID to remove: ")
			var id string
			fmt.Scanln(&id)
			msg := RemovePlant(&plants, id)
			fmt.Println(msg)
			SavePlantsToFile("plants.json", plants)

		case 7:
			GenerateHarvestReport(plants)

		case 8:
			var newgardener Gardener
			fmt.Print("Enter Gardener ID to update: ")

			fmt.Scanln(&newgardener.Id)

			fmt.Print("Enter New Name: ")
			fmt.Scanln(&newgardener.Name)

			fmt.Print("Enter New Age: ")

			fmt.Scanln(&newgardener.Age)

			fmt.Print("Enter New Plot Number: ")

			fmt.Scanln(&newgardener.PlotNumber)

			fmt.Print("Enter New Contact Email: ")
			fmt.Scanln(&newgardener.ContactEmail)

			fmt.Print("Enter New Preferred Plants (comma separated): ")
			var plantsInput string
			fmt.Scanln(&plantsInput)
			newgardener.PreferredPlants = strings.Split(plantsInput, ",")

			msg := UpdateGardener(&gardeners, newgardener.Id, newgardener.Name, newgardener.Age, newgardener.PlotNumber, newgardener.ContactEmail, newgardener.PreferredPlants)
			fmt.Println(msg)
			SaveGardenersToFile("gardeners.json", gardeners)

		case 9:
			var newplanet Plant
			fmt.Print("Enter Plant ID to update: ")
			fmt.Scanln(&newplanet.Id)

			fmt.Print("Enter New Name: ")
			fmt.Scanln(&newplanet.Name)

			fmt.Print("Enter New Plant Type (0 = Vegetable, 1 = Fruit, 2 = Herb): ")
			fmt.Scanln(&newplanet.plantType)

			fmt.Print("Enter New Planting Season: ")
			fmt.Scanln(&newplanet.PlantingSeason)

			fmt.Print("Enter New Care Instructions: ")
			fmt.Scanln(&newplanet.CareInstructions)
			msg := UpdatePlant(&plants, newplanet.Id, newplanet.Name, newplanet.plantType, newplanet.PlantingSeason, newplanet.CareInstructions)
			fmt.Println(msg)
			SavePlantsToFile("plants.json", plants)

		case 10:
			fmt.Print("Enter Plant Name or Planting Season to search: ")
			term, _ := reader.ReadString('\n')
			term = strings.TrimSpace(term)

			results := SearchPlants(plants, term)
			if len(results) == 0 {
				fmt.Println("No plants found.")
			} else {
				for _, p := range results {
					fmt.Println("Name:", p.Name)
					fmt.Println("Type:", p.plantType.String())
					fmt.Println("Season:", p.PlantingSeason)
					fmt.Println("Instructions:", p.CareInstructions)
					fmt.Println("------")
				}
			}

		case 11:
			ViewGardenerAssignments(&gardeners, plants)

		case 0:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
