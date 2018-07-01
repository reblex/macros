package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type ProfileData struct {
	Name            string  `json:"name"`            // Name of Profile
	Standard        string  `json:"standard"`        // metric/imperial
	Age             int     `json:"age"`             // age in years
	Height          float64 `json:"height"`          // height in given standard
	Gender          string  `json:"gender"`          // gender
	CalorieConstant int     `json:"calorieConstant"` // Increase/Decrease in calories(bulk/cut)
	ActivityFactor  float64 `json:"activityFactor"`  // Activity factor, multiplied with BMR for total calories
	FatPercentage   int     `json:"fatPercentage"`   // Percentage of protein is added as fat
	ProteinFactor   float64 `json:"proteinFactor"`   // Multiplied by bodyweight in pounds to calculate amount of daily protein
	MealsPerDay     int     `json:"mealsPerDay"`     //Amount of meals per day, given same size meals
}

type Profile struct {
	Name string
	Data ProfileData
}

func getProjectPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/reblex/macros/"
}

func (p *Profile) GetProfiles(path string) []ProfileData {
	path = getProjectPath() + path

	raw, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var profiles []ProfileData

	json.Unmarshal(raw, &profiles)

	return profiles
}

func (p *Profile) Load(path string) error {
	profiles := p.GetProfiles(path)

	profileExists := false
	for i := range profiles {
		if profiles[i].Name == p.Name {
			p.Data = profiles[i]
			profileExists = true
		}
	}

	if !profileExists {
		return errors.New("Profile Not Found")
	}

	return nil
}

func (p *Profile) GetSlot(path string) (int, error) {
	profiles := p.GetProfiles(path)
	var slot int

	for i := range profiles {
		if profiles[i].Name == p.Name {
			slot = i
			break
		}
	}

	return slot, nil
}

func (p *Profile) SaveToSlot(path string, slot int) error {
	profiles := p.GetProfiles(path)

	if slot < 0 || slot > (len(profiles)-1) {
		return errors.New("Save slot out of bounds")
	}

	profiles[slot] = p.Data

	data, _ := json.MarshalIndent(profiles, "", "    ")

	ioutil.WriteFile(getProjectPath()+path, data, 0644)

	return nil
}

func (p *Profile) Save(path string) error {
	profiles := p.GetProfiles(path)

	profileExists := false
	for i := range profiles {
		if profiles[i].Name == p.Name {
			profiles[i] = p.Data
			profileExists = true
		}
	}

	if !profileExists {
		profiles = append(profiles, p.Data)
	}

	data, _ := json.MarshalIndent(profiles, "", "    ")

	ioutil.WriteFile(getProjectPath()+path, data, 0644)

	return nil
}
