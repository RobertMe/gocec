package gocec

/*
#include <stdio.h>
#include <libcec/cecc.h>

void setDeviceName(libcec_configuration *config, char *name)
{
	snprintf(config->strDeviceName, sizeof(config->strDeviceName), "%s", name);
}
*/
import "C"

type Configuration struct {
	configuration C.libcec_configuration
	callbacks *callbacks
}

func NewConfiguration(deviceName string, activateSource bool) (config *Configuration) {
	config = &Configuration{}
	config.configuration.clientVersion = C.LIBCEC_VERSION_CURRENT
	C.setDeviceName(&config.configuration, C.CString(deviceName))
	config.configuration.deviceTypes.types[0] = C.CEC_DEVICE_TYPE_RECORDING_DEVICE

	initialiseCallbacks(config)

	return
}

func (config *Configuration) SetDeviceName(name string) {
	C.setDeviceName(&config.configuration, C.CString(name));
}

func (config *Configuration) SetActivateSource(activate bool) {
	config.configuration.bActivateSource = convertBool(activate)
}

func (config *Configuration) SetMonitorOnly(monitor bool) {
	config.configuration.bMonitorOnly = convertBool(monitor)
}

func convertBool(value bool) C.uchar {
	if value {
		return C.uchar(1)
	} else {
		return C.uchar(0)
	}
}
