package app

import (
	"encoding/json"
	"reflect"
	"testing"

	domain "github.com/slidebolt/sb-domain"
	managersdk "github.com/slidebolt/sb-manager-sdk"
)

func TestOnStart_PersistsSeededEntityData(t *testing.T) {
	env := managersdk.NewTestEnv(t)
	env.Start("messenger")
	env.Start("storage")

	deps := map[string]json.RawMessage{
		"messenger": env.MessengerPayload(),
	}

	app := New()
	if _, err := app.OnStart(deps); err != nil {
		t.Fatalf("OnStart: %v", err)
	}
	t.Cleanup(func() { _ = app.OnShutdown() })

	raw, err := env.Storage().Get(domain.EntityKey{
		Plugin:   PluginID,
		DeviceID: "demo_device",
		ID:       "demo_light",
	})
	if err != nil {
		t.Fatalf("get seeded entity: %v", err)
	}

	var got domain.Entity
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal seeded entity: %v", err)
	}

	if got.Plugin != PluginID || got.DeviceID != "demo_device" || got.ID != "demo_light" {
		t.Fatalf("identity = %s.%s.%s", got.Plugin, got.DeviceID, got.ID)
	}
	if got.Type != "light" || got.Name != "Demo Light" {
		t.Fatalf("entity metadata = type:%q name:%q", got.Type, got.Name)
	}

	wantCommands := []string{
		"light_turn_on",
		"light_turn_off",
		"light_set_brightness",
		"light_set_color_temp",
	}
	if !reflect.DeepEqual(got.Commands, wantCommands) {
		t.Fatalf("commands = %v, want %v", got.Commands, wantCommands)
	}

	light, ok := got.State.(domain.Light)
	if !ok {
		t.Fatalf("state type = %T, want domain.Light", got.State)
	}
	if light.Power || light.Brightness != 128 {
		t.Fatalf("state = %+v, want power=false brightness=128", light)
	}
}
