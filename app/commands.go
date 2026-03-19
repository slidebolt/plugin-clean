package app

import (
	"log"

	domain "github.com/slidebolt/sb-domain"
	messenger "github.com/slidebolt/sb-messenger-sdk"
)

func (a *App) handleCommand(addr messenger.Address, cmd any) {
	switch c := cmd.(type) {
	case domain.LightTurnOn:
		log.Printf("plugin-clean: light %s turn_on", addr.Key())
	case domain.LightTurnOff:
		log.Printf("plugin-clean: light %s turn_off transition=%v", addr.Key(), c.Transition)
	case domain.LightSetBrightness:
		log.Printf("plugin-clean: light %s set_brightness brightness=%d", addr.Key(), c.Brightness)
	case domain.LightSetColorTemp:
		log.Printf("plugin-clean: light %s set_color_temp mireds=%d", addr.Key(), c.Mireds)
	case domain.LightSetRGB:
		log.Printf("plugin-clean: light %s set_rgb r=%d g=%d b=%d", addr.Key(), c.R, c.G, c.B)
	case domain.LightSetRGBW:
		log.Printf("plugin-clean: light %s set_rgbw r=%d g=%d b=%d w=%d", addr.Key(), c.R, c.G, c.B, c.W)
	case domain.LightSetRGBWW:
		log.Printf("plugin-clean: light %s set_rgbww r=%d g=%d b=%d cw=%d ww=%d", addr.Key(), c.R, c.G, c.B, c.CW, c.WW)
	case domain.LightSetHS:
		log.Printf("plugin-clean: light %s set_hs hue=%.1f sat=%.1f", addr.Key(), c.Hue, c.Saturation)
	case domain.LightSetXY:
		log.Printf("plugin-clean: light %s set_xy x=%.4f y=%.4f", addr.Key(), c.X, c.Y)
	case domain.LightSetWhite:
		log.Printf("plugin-clean: light %s set_white white=%d", addr.Key(), c.White)
	case domain.LightSetEffect:
		log.Printf("plugin-clean: light %s set_effect effect=%s", addr.Key(), c.Effect)
	case domain.SwitchTurnOn:
		log.Printf("plugin-clean: switch %s turn_on", addr.Key())
	case domain.SwitchTurnOff:
		log.Printf("plugin-clean: switch %s turn_off", addr.Key())
	case domain.SwitchToggle:
		log.Printf("plugin-clean: switch %s toggle", addr.Key())
	case domain.FanTurnOn:
		log.Printf("plugin-clean: fan %s turn_on", addr.Key())
	case domain.FanTurnOff:
		log.Printf("plugin-clean: fan %s turn_off", addr.Key())
	case domain.FanSetSpeed:
		log.Printf("plugin-clean: fan %s set_speed percentage=%d", addr.Key(), c.Percentage)
	case domain.CoverOpen:
		log.Printf("plugin-clean: cover %s open", addr.Key())
	case domain.CoverClose:
		log.Printf("plugin-clean: cover %s close", addr.Key())
	case domain.CoverSetPosition:
		log.Printf("plugin-clean: cover %s set_position pos=%d", addr.Key(), c.Position)
	case domain.LockLock:
		log.Printf("plugin-clean: lock %s lock", addr.Key())
	case domain.LockUnlock:
		log.Printf("plugin-clean: lock %s unlock", addr.Key())
	case domain.ButtonPress:
		log.Printf("plugin-clean: button %s press", addr.Key())
	case domain.NumberSetValue:
		log.Printf("plugin-clean: number %s set_value value=%v", addr.Key(), c.Value)
	case domain.SelectOption:
		log.Printf("plugin-clean: select %s set_option option=%s", addr.Key(), c.Option)
	case domain.TextSetValue:
		log.Printf("plugin-clean: text %s set_value value=%s", addr.Key(), c.Value)
	case domain.ClimateSetMode:
		log.Printf("plugin-clean: climate %s set_mode mode=%s", addr.Key(), c.HVACMode)
	case domain.ClimateSetTemperature:
		log.Printf("plugin-clean: climate %s set_temperature temp=%v", addr.Key(), c.Temperature)
	default:
		log.Printf("plugin-clean: unknown command %T for %s", cmd, addr.Key())
	}
}
