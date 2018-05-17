package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type profileData struct {
	Name     string  `json:"name"`
	Standard string  `json:"standard"`
	Age      int     `json:"age"`
	Height   float64 `json:"height"`
	Gender   string  `json:"gender"`
}

type Profile struct {
	Name string
	Data profileData
}

func (p *Profile) GetProfiles(path string) []profileData {
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var profiles []profileData

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

	ioutil.WriteFile(path, data, 0644)

	return nil
}
