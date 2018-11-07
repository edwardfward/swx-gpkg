package gpkg

type Motion struct {
	ObservationTime int64   `json:"observation_time"`
	AccX            float64 `json:"acceleration_x"`
	AccY            float64 `json:"acceleration_y"`
	AccZ            float64 `json:"acceleration_z"`
	LinearAccX      float64 `json:"linear_acceleration_x"`
	LinearAccY      float64 `json:"linear_acceleration_y"`
	LinearAccZ      float64 `json:"linear_acceleration_z"`
	MagX            float64 `json:"magnetic_x"`
	MagY            float64 `json:"magnetic_y"`
	MagZ            float64 `json:"magnetic_z"`
	GyroX           float64 `json:"gyro_x"`
	GyroY           float64 `json:"gyro_y"`
	GyroZ           float64 `json:"gyro_z"`
	GravityX        float64 `json:"gravity_x"`
	GravityY        float64 `json:"gravity_y"`
	GravityZ        float64 `json:"gravity_z"`
	RotVecX         float64 `json:"rotational_vector_x"`
	RotVecY         float64 `json:"rotational_vector_y"`
	RotVecZ         float64 `json:"rotational_vector_z"`
	RotVecCos       float64 `json:"rotational_vector_cos"`
	RotVecHdgAcc    float64 `json:"rot_vec_hdg_acc"`
	BaroPressure    float64 `json:"baro_pressure"`
	Humidity        float64 `json:"humidity"`
	TemperatureC    float64 `json:"temperature_c"`

	//todo add lux, prox, stationary, motion fields
	//todo add documentation and examples
}
