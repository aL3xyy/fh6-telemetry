package main

// Just telemetry data imported from FH6 website
type Telemetry struct {
	IsRaceOn    int32
	TimestampMS uint32

	EngineMaxRpm     float32
	EngineIdleRpm    float32
	CurrentEngineRpm float32

	AccelerationX float32
	AccelerationY float32
	AccelerationZ float32

	VelocityX float32
	VelocityY float32
	VelocityZ float32

	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32

	Yaw   float32
	Pitch float32
	Roll  float32

	NormSuspensionTravelFL float32
	NormSuspensionTravelFR float32
	NormSuspensionTravelRL float32
	NormSuspensionTravelRR float32

	TireSlipRatioFL float32
	TireSlipRatioFR float32
	TireSlipRatioRL float32
	TireSlipRatioRR float32

	WheelRotationSpeedFL float32
	WheelRotationSpeedFR float32
	WheelRotationSpeedRL float32
	WheelRotationSpeedRR float32

	WheelOnRumbleStripFL int32
	WheelOnRumbleStripFR int32
	WheelOnRumbleStripRL int32
	WheelOnRumbleStripRR int32

	WheelInPuddleFL int32
	WheelInPuddleFR int32
	WheelInPuddleRL int32
	WheelInPuddleRR int32

	SurfaceRumbleFL float32
	SurfaceRumbleFR float32
	SurfaceRumbleRL float32
	SurfaceRumbleRR float32

	TireSlipAngleFL float32
	TireSlipAngleFR float32
	TireSlipAngleRL float32
	TireSlipAngleRR float32

	TireCombinedSlipFL float32
	TireCombinedSlipFR float32
	TireCombinedSlipRL float32
	TireCombinedSlipRR float32

	SuspensionTravelMetersFL float32
	SuspensionTravelMetersFR float32
	SuspensionTravelMetersRL float32
	SuspensionTravelMetersRR float32

	CarOrdinal          int32
	CarClass            int32
	CarPerformanceIndex int32
	DrivetrainType      int32
	NumCylinders        int32
	CarGroup            uint32

	SmashableVelDiff float32
	SmashableMass    float32

	PositionX float32
	PositionY float32
	PositionZ float32

	Speed  float32
	Power  float32
	Torque float32

	TireTempFL float32
	TireTempFR float32
	TireTempRL float32
	TireTempRR float32

	Boost            float32
	Fuel             float32
	DistanceTraveled float32

	BestLap         float32
	LastLap         float32
	CurrentLap      float32
	CurrentRaceTime float32

	LapNumber    uint16
	RacePosition uint8

	Accel     uint8
	Brake     uint8
	Clutch    uint8
	HandBrake uint8
	Gear      uint8

	Steer                       int8
	NormalizedDrivingLine       int8
	NormalizedAIBrakeDifference int8
}
