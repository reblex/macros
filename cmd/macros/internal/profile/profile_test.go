package profile

import (
	"testing"
)

func TestLoadWorkingFile(t *testing.T) {
	t.Run("Load Profile, working", func(t *testing.T) {
		var p Profile
		p.Name = "main"
		error := p.Load("cmd/macros/internal/profile/test_profiles.json")

		// fmt.Println(p.Data)

		if error != nil {
			t.Error(error.Error())
		}
	})
}

func TestSaveNewFile(t *testing.T) {
	t.Run("Save new file", func(t *testing.T) {
		pd := ProfileData{"main", "metric", 21, 190, "male", 200, 1.65, 66, 1.2, 5}
		p := Profile{"main", pd}
		error := p.Save("cmd/macros/internal/profile/test_profiles.json")

		if error != nil {
			t.Error(error.Error())
		}
	})
}
