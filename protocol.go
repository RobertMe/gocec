package gocec

/*
#include <libcec/cecc.h>
*/
import "C"
import (
	"fmt"
	"strings"
)

var toString *C.char

const toStringSize = 50

func init() {
	toString = (*C.char)(C.malloc(toStringSize))
}

type LogicalAddress byte

const (
	DeviceTV LogicalAddress = iota
	DeviceRecodingDevice1
	DeviceRecodingDevice2
	DeviceTuner1
	DevicePlaybackDevice1
	DeviceAudiosystem
	DeviceTuner2
	DeviceTuner3
	DevicePlaybackDevice2
	DeviceRecodingDevice3
	DeviceTuner4
	DevicePlaybackDevice3
	DeviceReserved1
	DeviceReserved2
	DeviceFreeUse
	DeviceUnregistered
	DeviceBroadcast LogicalAddress = 15
)

func (address LogicalAddress) String() string {
	C.libcec_logical_address_to_string(C.cec_logical_address(address), toString, toStringSize)

	return C.GoString(toString)
}

type PhysicalAddress [2]byte

func (address PhysicalAddress) String() string {
	builder := strings.Builder{}

	fmt.Fprintf(&builder, "%d.%d.%d.%d", (address[0] >> 4) & 0xF, address[0] & 0x0F, address[1] >> 4, address[1] & 0x0F)

	return builder.String()
}

type Opcode byte

const (
	OpcodeActiveSource              Opcode = 0x82
	OpcodeImageViewOn               Opcode = 0x04
	OpcodeTextViewOn                Opcode = 0x0D
	OpcodeInactiveSource            Opcode = 0x9D
	OpcodeRequestActiveSource       Opcode = 0x85
	OpcodeRoutingChange             Opcode = 0x80
	OpcodeRoutingInformation        Opcode = 0x81
	OpcodeSetStreamPath             Opcode = 0x86
	OpcodeStandby                   Opcode = 0x36
	OpcodeRecordOff                 Opcode = 0x0B
	OpcodeRecordOn                  Opcode = 0x09
	OpcodeRecordStatus              Opcode = 0x0A
	OpcodeRecordTvScreen            Opcode = 0x0F
	OpcodeClearAnalogueTimer        Opcode = 0x33
	OpcodeClearDigitalTimer         Opcode = 0x99
	OpcodeClearExternalTimer        Opcode = 0xA1
	OpcodeSetAnalogueTimer          Opcode = 0x34
	OpcodeSetDigitalTimer           Opcode = 0x97
	OpcodeSetExternalTimer          Opcode = 0xA2
	OpcodeSetTimerProgramTitle      Opcode = 0x67
	OpcodeTimerClearedStatus        Opcode = 0x43
	OpcodeTimerStatus               Opcode = 0x35
	OpcodeCecVersion                Opcode = 0x9E
	OpcodeGetCecVersion             Opcode = 0x9F
	OpcodeGivePhysicalAddress       Opcode = 0x83
	OpcodeGetMenuLanguage           Opcode = 0x91
	OpcodeReportPhysicalAddress     Opcode = 0x84
	OpcodeSetMenuLanguage           Opcode = 0x32
	OpcodeDeckControl               Opcode = 0x42
	OpcodeDeckStatus                Opcode = 0x1B
	OpcodeGiveDeckStatus            Opcode = 0x1A
	OpcodePlay                      Opcode = 0x41
	OpcodeGiveTunerDeviceStatus     Opcode = 0x08
	OpcodeSelectAnalogueService     Opcode = 0x92
	OpcodeSelectDigitalService      Opcode = 0x93
	OpcodeTunerDeviceStatus         Opcode = 0x07
	OpcodeTunerStepDecrement        Opcode = 0x06
	OpcodeTunerStepIncrement        Opcode = 0x05
	OpcodeDeviceVendorId            Opcode = 0x87
	OpcodeGiveDeviceVendorId        Opcode = 0x8C
	OpcodeVendorCommand             Opcode = 0x89
	OpcodeVendorCommandWithId       Opcode = 0xA0
	OpcodeVendorRemoteButtonDown    Opcode = 0x8A
	OpcodeVendorRemoteButtonUp      Opcode = 0x8B
	OpcodeSetOsdString              Opcode = 0x64
	OpcodeGiveOsdName               Opcode = 0x46
	OpcodeSetOsdName                Opcode = 0x47
	OpcodeMenuRequest               Opcode = 0x8D
	OpcodeMenuStatus                Opcode = 0x8E
	OpcodeUserControlPressed        Opcode = 0x44
	OpcodeUserControlRelease        Opcode = 0x45
	OpcodeGiveDevicePowerStatus     Opcode = 0x8F
	OpcodeReportPowerStatus         Opcode = 0x90
	OpcodeFeatureAbort              Opcode = 0x00
	OpcodeAbort                     Opcode = 0xFF
	OpcodeGiveAudioStatus           Opcode = 0x71
	OpcodeGiveSystemAudioModeStatus Opcode = 0x7D
	OpcodeReportAudioStatus         Opcode = 0x7A
	OpcodeSetSystemAudioMode        Opcode = 0x72
	OpcodeSystemAudioModeRequest    Opcode = 0x70
	OpcodeSystemAudioModeStatus     Opcode = 0x7E
	OpcodeSetAudioRate              Opcode = 0x9A
	OpcodeStartArc                  Opcode = 0xC0
	OpcodeReportArcStarted          Opcode = 0xC1
	OpcodeReportArcEnded            Opcode = 0xC2
	OpcodeRequestArcStart           Opcode = 0xC3
	OpcodeRequestArcEnd             Opcode = 0xC4
	OpcodeEndArc                    Opcode = 0xC5
	OpcodeCdc                       Opcode = 0xF8
	OpcodeNone                      Opcode = 0xFD
)

func (opcode Opcode) String() string {
	C.libcec_opcode_to_string(C.cec_opcode(opcode), toString, toStringSize)

	return C.GoString(toString)
}

type PowerStatus byte

const (
	PowerStatusOn PowerStatus = iota
	PowerStatusStandBy
	PowerStatusTransitionToOn
	PowerStatusTransitionToStandby
	PowerStatusUnknown PowerStatus = 0x99
)

func (status PowerStatus) String() string {
	C.libcec_power_status_to_string(C.cec_power_status(status), toString, toStringSize)

	return C.GoString(toString)
}
