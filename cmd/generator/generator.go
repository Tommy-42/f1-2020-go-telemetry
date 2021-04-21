package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	f1packet "github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	header := f1packet.PacketHeader{
		PacketFormat:            2020,
		GameMajorVersion:        1,
		GameMinorVersion:        42,
		PacketVersion:           1,
		PacketID:                uint8(f1packet.CarStatusPacket),
		SessionUID:              rand.Uint64(),
		SessionTime:             1,
		FrameIdentifier:         42,
		PlayerCarIndex:          2,
		SecondaryPlayerCarIndex: 255,
	}

	carStatusData := [22]f1packet.CarStatusData{}
	carStatusData[2] = f1packet.CarStatusData{
		TractionControl:       1,
		AntiLockBrakes:        1,
		FuelMix:               2,
		FrontBrakeBias:        58,
		PitLimiterStatus:      0,
		FuelInTank:            20.62,
		FuelCapacity:          100.00,
		FuelRemainingLaps:     -0.10,
		MaxRPM:                22000,
		IdleRPM:               3000,
		MaxGears:              8,
		DrsAllowed:            0,
		DrsActivationDistance: 150,
		TyresWear: [4]uint8{
			19, 19, 19, 19,
		},
		ActualTyreCompound: 16,
		VisualTyreCompound: 16,
		TyresAgeLaps:       5,
		TyresDamage: [4]uint8{
			12, 11, 10, 11,
		},
		FrontLeftWingDamage:     0,
		FrontRightWingDamage:    5,
		RearWingDamage:          0,
		DrsFault:                0,
		EngineDamage:            0,
		GearBoxDamage:           1,
		VehicleFiaFlags:         0,
		ErsStoreEnergy:          12.000,
		ErsDeployMode:           1,
		ErsHarvestedThisLapMGUK: 11.000,
		ErsHarvestedThisLapMGUH: 10.000,
		ErsDeployedThisLap:      5.000,
	}

	carStatusPacket := f1packet.PacketCarStatusData{
		Header:        header,
		CarStatusData: carStatusData,
	}

	s, err := net.ResolveUDPAddr("udp4", "localhost:20777")
	if err != nil {
		fmt.Printf("could not resolve UDP Addr: %v\n", err)
		return
	}
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Printf("could not dial UDP: %v\n", err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	var packet bytes.Buffer
	err = binary.Write(&packet, binary.LittleEndian, carStatusPacket)
	if err != nil {
		fmt.Printf("could not convert to little endian packet: %v\n", err)
		os.Exit(1)
	}

	n, err := c.Write(packet.Bytes())
	if err != nil {
		fmt.Printf("could not write udp packet: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sent %d over udp\n", n)
}
