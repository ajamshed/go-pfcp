// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewReportingTriggers creates a new ReportingTriggers IE.
func NewReportingTriggers(triggers uint16) *IE {
	return newUint16ValIE(ReportingTriggers, triggers)
}

// ReportingTriggers returns ReportingTriggers in uint16 if the type of IE matches.
func (i *IE) ReportingTriggers() (uint16, error) {
	if i.Type != ReportingTriggers {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint16(i.Payload[0:2]), nil
}

// HasLIUSA reports whether reporting trigger has LIUSA bit.
func (i *IE) HasLIUSA() bool {
	switch i.Type {
	case ReportingTriggers:
		if len(i.Payload) < 1 {
			return false
		}

		u8 := uint8(i.Payload[0])
		return has8thBit(u8)
	case UsageReportTrigger:
		if len(i.Payload) < 2 {
			return false
		}
		u8 := uint8(i.Payload[1])
		return has3rdBit(u8)
	default:
		return false
	}
}

// HasDROTH reports whether reporting trigger has DROTH bit.
func (i *IE) HasDROTH() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has7thBit(u8)
	default:
		return false
	}
}

// HasSTOPT reports whether reporting trigger has STOPT bit.
func (i *IE) HasSTOPT() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has6thBit(u8)
	default:
		return false
	}
}

// HasSTART reports whether reporting trigger has START bit.
func (i *IE) HasSTART() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has5thBit(u8)
	default:
		return false
	}
}

// HasQUHTI reports whether reporting trigger has QUHTI bit.
func (i *IE) HasQUHTI() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has4thBit(u8)
	default:
		return false
	}
}

// HasTIMTH reports whether reporting trigger has TIMTH bit.
func (i *IE) HasTIMTH() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has3rdBit(u8)
	default:
		return false
	}
}

// HasVOLTH reports whether reporting trigger has VOLTH bit.
func (i *IE) HasVOLTH() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has2ndBit(u8)
	default:
		return false
	}
}

// HasPERIO reports whether reporting trigger has PERIO bit.
func (i *IE) HasPERIO() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has1stBit(u8)
	default:
		return false
	}
}

// HasEVEQU reports whether reporting trigger has EVEQU bit.
func (i *IE) HasEVEQU() bool {
	switch i.Type {
	case ReportingTriggers:
		if len(i.Payload) < 2 {
			return false
		}

		u8 := uint8(i.Payload[1])
		return has6thBit(u8)
	case UsageReportTrigger:
		if len(i.Payload) < 3 {
			return false
		}

		u8 := uint8(i.Payload[2])
		return has1stBit(u8)
	default:
		return false
	}
}

// HasEVETH reports whether reporting trigger has EVETH bit.
func (i *IE) HasEVETH() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ReportingTriggers:
		u8 := uint8(i.Payload[1])
		return has5thBit(u8)
	case UsageReportTrigger:
		u8 := uint8(i.Payload[1])
		return has8thBit(u8)
	default:
		return false
	}
}

// HasMACAR reports whether reporting trigger has MACAR bit.
func (i *IE) HasMACAR() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ReportingTriggers:
		u8 := uint8(i.Payload[1])
		return has4thBit(u8)
	case UsageReportTrigger:
		u8 := uint8(i.Payload[1])
		return has7thBit(u8)
	default:
		return false
	}
}

// HasENVCL reports whether reporting trigger has ENVCL bit.
func (i *IE) HasENVCL() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ReportingTriggers:
		u8 := uint8(i.Payload[1])
		return has3rdBit(u8)
	case UsageReportTrigger:
		u8 := uint8(i.Payload[1])
		return has6thBit(u8)
	default:
		return false
	}
}

// HasTIMQU reports whether reporting trigger has TIMQU bit.
func (i *IE) HasTIMQU() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[1])
		return has2ndBit(u8)
	default:
		return false
	}
}

// HasVOLQU reports whether reporting trigger has VOLQU bit.
func (i *IE) HasVOLQU() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ReportingTriggers, UsageReportTrigger:
		u8 := uint8(i.Payload[1])
		return has1stBit(u8)
	default:
		return false
	}
}