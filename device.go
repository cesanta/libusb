// Copyright (c) 2015 The libusb developers. All rights reserved.
// Project site: https://github.com/gotmc/libusb
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package libusb

// #cgo pkg-config: libusb-1.0
// #include <libusb.h>
import "C"

// TODO(mdr): Do I need to be hadnling the reference counts in cgo?

type speed int

const (
	speedUnknown speed = C.LIBUSB_SPEED_UNKNOWN
	speedLow     speed = C.LIBUSB_SPEED_LOW
	speedFull    speed = C.LIBUSB_SPEED_FULL
	speedHigh    speed = C.LIBUSB_SPEED_HIGH
	speedSuper   speed = C.LIBUSB_SPEED_SUPER
)

var speedCodes = map[speed]string{
	speedUnknown: "The OS doesn't report or know the device speed.",
	speedLow:     "The device is operating at low speed (1.5MBit/s)",
	speedFull:    "The device is operating at full speed (12MBit/s)",
	speedHigh:    "The device is operating at high speed (480MBit/s)",
	speedSuper:   "The device is operating at super speed (5000MBit/s)",
}

func (speed speed) String() string {
	return speedCodes[speed]
}

type supportedSpeed int

const (
	lowSpeedOperation   supportedSpeed = C.LIBUSB_LOW_SPEED_OPERATION
	fullSpeedOperation  supportedSpeed = C.LIBUSB_FULL_SPEED_OPERATION
	highSpeedOperation  supportedSpeed = C.LIBUSB_HIGH_SPEED_OPERATION
	superSpeedOperation supportedSpeed = C.LIBUSB_SUPER_SPEED_OPERATION
)

var supportedSpeeds = map[supportedSpeed]string{
	lowSpeedOperation:   "Low speed operation supported (1.5MBit/s).",
	fullSpeedOperation:  "Full speed operation supported (12MBit/s).",
	highSpeedOperation:  "High speed operation supported (480MBit/s).",
	superSpeedOperation: "Superspeed operation supported (5000MBit/s).",
}

func (speed supportedSpeed) String() string {
	return supportedSpeeds[speed]
}

type device struct {
	libusbDevice *C.libusb_device
}

func (dev *device) GetBusNumber() (uint, error) {
	busNumber, err := C.libusb_get_bus_number(dev.libusbDevice)
	if err != nil {
		return 0, err
	}
	return uint(busNumber), nil
}

func (dev *device) GetDeviceAddress() (uint, error) {
	deviceAddress, err := C.libusb_get_device_address(dev.libusbDevice)
	if err != nil {
		return 0, err
	}
	return uint(deviceAddress), nil
}

func (dev *device) GetDeviceSpeed() (speed, error) {
	deviceSpeed, err := C.libusb_get_device_speed(dev.libusbDevice)
	if err != nil {
		return 0, err
	}
	return speed(deviceSpeed), nil
}

func (dev *device) Open() (*deviceHandle, error) {
	var handle **C.libusb_device_handle
	err := C.libusb_open(dev.libusbDevice, handle)
	if err != 0 {
		return nil, ErrorCode(err)
	}
	deviceHandle := deviceHandle{
		libusbDeviceHandle: *handle,
	}

	return &deviceHandle, nil
}
