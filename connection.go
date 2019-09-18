package gocec

/*
#cgo pkg-config: libcec
//#cgo CFLAGS: -Iinclude
//#cgo LDFLAGS: -lcec
#include <libcec/cecc.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Connection struct {
	connection C.libcec_connection_t;
}

type Adapter struct {
	Path, Comm string
}

func NewConnection(config *Configuration) (*Connection, error) {
	conn := &Connection{
		connection: C.libcec_initialise(&config.configuration),
	}

	return conn, nil
}

func (conn *Connection) FindAdapters() []Adapter {
	var foundDevices [10]C.cec_adapter
	count := int(C.libcec_find_adapters(conn.connection, &foundDevices[0], C.uchar(len(foundDevices)), nil))

	adapters := make([]Adapter, count)
	for i := 0; i < count; i++ {
		adapters[i] = Adapter{
			Path: C.GoStringN(&foundDevices[i].path[0], 1024),
			Comm: C.GoStringN(&foundDevices[i].comm[0], 1024),
		}
	}

	return adapters
}

func (conn *Connection) Open(adapter Adapter) error {
	result := C.libcec_open(conn.connection, C.CString(adapter.Comm), C.CEC_DEFAULT_CONNECT_TIMEOUT)

	if result < 1 {
		return errors.New("Failed to open adapter")
	}

	return nil
}

func (conn *Connection) GetAdapterAddress() (LogicalAddress, error) {
	addresses := C.libcec_get_logical_addresses(conn.connection)
	if addresses.addresses[addresses.primary] == 0 {
		return DeviceUnregistered, errors.New("unable to determine logical address of the CEC adapter")
	}
	return LogicalAddress(addresses.primary), nil
}

func (conn *Connection) ActiveDevices() []LogicalAddress {
	activeDevices := C.libcec_get_active_devices(conn.connection)

	result := make([]LogicalAddress, 0, 16)
	for address, active := range activeDevices.addresses {
		if active == 1 {
			result = append(result, LogicalAddress(address))
		}
	}

	return result
}

func (conn *Connection) GetPowerStatus(address LogicalAddress) PowerStatus {
	return PowerStatus(C.libcec_get_device_power_status(conn.connection, C.cec_logical_address(address)))
}

func (conn *Connection) GetPhysicalAddress(address LogicalAddress) PhysicalAddress {
	physical := C.libcec_get_device_physical_address(conn.connection, C.cec_logical_address(address))
	return PhysicalAddress{byte(physical >> 8), byte(physical & 0xFF)}
}

func (conn *Connection) GetVendor(address LogicalAddress) Vendor {
	vendor := C.libcec_get_device_vendor_id(conn.connection, C.cec_logical_address(address))
	return Vendor(vendor)
}

func (conn *Connection) GetOSDName(address LogicalAddress) string {
	name := C.cec_osd_name{}
	C.libcec_get_device_osd_name(conn.connection, C.cec_logical_address(address), (*C.char)(unsafe.Pointer(&name[0])))
	return C.GoString((*C.char)(unsafe.Pointer(&name[0])))
}

func (conn *Connection) Transmit(message Message) {
	var command C.cec_command

	messageLength := len(message)

	if messageLength > 0 {
		command.initiator = C.cec_logical_address(message.Source())
		command.destination = C.cec_logical_address(message.Destination())

		if messageLength > 1 {
			command.opcode_set = 1
			command.opcode = C.cec_opcode(message.Opcode())
		} else {
			command.opcode_set = 0
		}


		if parameters := message.Parameters(); parameters != nil {
			command.parameters.size = C.uint8_t(len(parameters))
			for i, val := range parameters {
				command.parameters.data[i] = C.uint8_t(val)
			}
		} else {
			command.parameters.size = 0
		}
	}

	C.libcec_transmit(conn.connection, &command)
}
