package uidgenerator

import (
	"testing"
)

var defCfg = Cfg{
	Alfa:      "1234567890",
	Format:    "XXX-XXXXXX-XXX",
	Validator: "[0-9]{3}-[0-9]{6}-[0-9]{3}",
}

func TestGenerator(t *testing.T) {
	simpleUID := NewGenerator(&defCfg)
	uid := simpleUID.New()
	if len(defCfg.Format) != len(uid) {
		t.Errorf("uid length (%d) != format length (%d) ", len(uid), len(defCfg.Format))
	}
	if uid == simpleUID.New() {
		t.Error("uid is permanent")
	}
}

func TestValidator(t *testing.T) {
	simpleUID := NewGenerator(&defCfg)
	if simpleUID.Validator() != defCfg.Validator {
		t.Error("validator is differ")
	}

	u1 := simpleUID.New()
	u2, err := simpleUID.Validate(u1)

	if err != nil {
		t.Errorf("Generated UID (%s) isn't validated by UID validator", u1)
	} else {
		if u1 != u2 {
			t.Error("Validator returns wrong uid")
		}
	}
}
