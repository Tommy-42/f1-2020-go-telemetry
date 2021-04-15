# Packet Information
## Packet Types
Each packet can now carry different types of data rather than having one packet which contains everything. A header has been added to each packet as well so that versioning can be tracked and it will be easier for applications to check they are interpreting the incoming data in the correct way. Please note that all values are encoded using Little Endian format. All data is packed.

The following data types are used in the structures:

| Type | Description |
| --- | --- |
| uint8 | Unsigned 8-bit integer |
| int8 | Signed 8-bit integer | 
| uint16 | Unsigned 16-bit integer |
| int16 | Signed 16-bit integer |
| float | Floating point (32-bit) |
| uint64 | Unsigned 64-bit integer |

# Restricted data (Your Telemetry setting)
There is some data in the UDP that you may not want other players seeing if you are in a multiplayer game. This is controlled by the “Your Telemetry” setting in the Telemetry options. The options are:

* Restricted (Default) – other players viewing the UDP data will not see values for your car
* Public – all other players can see all the data for your car

Note: You can always see the data for the car you are driving regardless of the setting.

The following data items are set to zero if the player driving the car in question has their “Your Telemetry” set to “Restricted”:

Car status packet
* FuelInTank
* FuelCapacity
* FuelMix
* FuelRemainingLaps
* FrontBrakeBias
* FrontLeftWingDamage
* FrontRightWingDamage
* RearWingDamage
* EngineDamage
* GearBoxDamage
* TyresWear (All four wheels)
* TyresDamage (All four wheels)
* ErsDeployMode
* ErsStoreEnergy
* ErsDeployedThisLap
* ErsHarvestedThisLapMGUK
* ErsHarvestedThisLapMGUH
* TyresAgeLaps


# Appendices
https://forums.codemasters.com/topic/50942-f1-2020-udp-specification/?do=findComment&comment=515247


# Ressources
Based on https://forums.codemasters.com/topic/50942-f1-2020-udp-specification/