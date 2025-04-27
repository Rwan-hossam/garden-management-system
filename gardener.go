package main

import "fmt"

type Gardener struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Age             int      `json:"age"`
	PlotNumber      int      `json:"plot_number"`
	ContactEmail    string   `json:"contact_details"`
	PreferredPlants []string `json:"preferred_plants"`
}

func AddGardener(Gardeners *[]Gardener, newGardener Gardener) error {

	for _, plant := range *Gardeners {
		if plant.Id == newGardener.Id {
			return fmt.Errorf("Error: A Gardener with the same ID (%s) already exists", newGardener.Id)
		}
	}

	*Gardeners = append(*Gardeners, newGardener)
	return nil
}

func RemoveGardener(Gardeners *[]Gardener, id string) string {
	for i, Gardener := range *Gardeners {
		if Gardener.Id == id {
			*Gardeners = append((*Gardeners)[:i], (*Gardeners)[i+1:]...)
			return (Gardener.Name + " removed successfully")
		}
	}
	return ("Plant with given ID not found")
}

func UpdateGardener(Gardeners *[]Gardener, id string, updatedName string, updatedAge int, updatedPlotNumber int, updatedContactEmail string, updatedPreferredPlants []string) string {
	for i, Gardener := range *Gardeners {
		if Gardener.Id == id {
			(*Gardeners)[i].Name = updatedName
			(*Gardeners)[i].Age = updatedAge
			(*Gardeners)[i].PlotNumber = updatedPlotNumber
			(*Gardeners)[i].ContactEmail = updatedContactEmail
			(*Gardeners)[i].PreferredPlants = updatedPreferredPlants
			return (Gardener.Name + " Updated successfully ")
		}
	}
	return ("plant with given ID not found")

}
func ViewGardenerAssignments(Gardeners *[]Gardener, plants []Plant) {
	for _, gardener := range *Gardeners {
		fmt.Println("Gardener Name:", gardener.Name)
		fmt.Println("Gardener Email:", gardener.ContactEmail)
		fmt.Println("Preferred Plants:", gardener.PreferredPlants)
		fmt.Println("Assigned Plot:", gardener.PlotNumber)
		fmt.Println("------")
	}
}
