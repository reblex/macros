package profile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func (p *Profile) GetProfiles() []profileData {
	raw, _ := ioutil.ReadFile("profiles.json")

	var profiles []profileData

	json.Unmarshal(raw, &profiles)

	return profiles
}

func (p *Profile) Load() error {
	profiles := p.GetProfiles()

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

func (p *Profile) Save() error {
	profiles := p.GetProfiles()

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

	ioutil.WriteFile("profiles.json", data, 0644)

	return nil
}
