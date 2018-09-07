package ikeatradfriadapter

import (
	"github.com/function61/home-automation-hub/hapitypes"
	"github.com/function61/home-automation-hub/libraries/ikeatradfri"
	"log"
)

func New(adapter *hapitypes.Adapter, config hapitypes.AdapterConfig) *hapitypes.Adapter {
	go func() {
		log.Println("ikeatradfriadapter: started")

		coapClient := ikeatradfri.NewCoapClient(
			config.TradfriUrl,
			config.TradfriUser,
			config.TradfriPsk)

		for {
			select {
			case powerMsg := <-adapter.PowerMsg:
				var responseErr error = nil

				if powerMsg.On {
					responseErr = ikeatradfri.TurnOn(powerMsg.DeviceId, coapClient)
				} else {
					responseErr = ikeatradfri.TurnOff(powerMsg.DeviceId, coapClient)
				}

				if responseErr != nil {
					log.Printf("ikeatradfriadapter: error %s", responseErr.Error())
				}
			}
		}
	}()

	return adapter
}
