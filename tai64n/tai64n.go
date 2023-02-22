/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2022 WireGuard LLC. All Rights Reserved.
 */
// Package timestamp provides a TAI64N timestamp implementation.
package timestamp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

const (
	TimestampSize = 12
	base          = uint64(0x400000000000000a)
	whitenerMask  = uint32(0x1000000 - 1)
)

// Timestamp represents a TAI64N timestamp.
type Timestamp [TimestampSize]byte

// Now returns the current TAI64N timestamp.
func Now() Timestamp {
	return stamp(time.Now())
}

// After reports whether t1 is after t2.
func (t1 Timestamp) After(t2 Timestamp) bool {
	return bytes.Compare(t1[:], t2[:]) > 0
}

// String returns the string representation of the TAI64N timestamp.
func (t Timestamp) String() string {
	return t.ToTime().String()
}

// ToTime returns the corresponding time.Time value of the TAI64N timestamp.
func (t Timestamp) ToTime() time.Time {
	secs := binary.BigEndian.Uint64(t[:8]) - base
	nano := binary.BigEndian.Uint32(t[8:12])
	return time.Unix(int64(secs), int64(nano)).Add(time.Duration(whitenerMask-nano) * time.Nanosecond)
}

// stamp returns the TAI64N timestamp for a given time.Time value.
func stamp(t time.Time) Timestamp {
	var tai64n Timestamp
	secs := base + uint64(t.Unix())
	nano := uint32(t.Nanosecond()) &^ whitenerMask
	binary.BigEndian.PutUint64(tai64n[:8], secs)
	binary.BigEndian.PutUint32(tai64n[8:], nano)
	return tai64n
}

