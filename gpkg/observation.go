package gpkg

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

type Observation struct {
	Id                  int32       `json:"id"`
	SysTime             time.Time   `json:"sys_time"`
	Lat                 float64     `json:"lat"`
	Lon                 float64     `json:"lon"`
	Alt                 float64     `json:"alt"`
	Provider            string      `json:"provider"`
	GpsTime             int64       `json:"gps_time"`
	FixSatCount         int8        `json:"fix_sat_count"`
	HasRadialAccuracy   bool        `json:"has_radial_accuracy"`
	HasVerticalAccuracy bool        `json:"has_vertical_accuracy"`
	RadialAccuracy      float64     `json:"radial_accuracy"`
	VerticalAccuracy    float64     `json:"vertical_accuracy"`
	Speed               float64     `json:"speed"`
	SpeedAccuracy       float64     `json:"speed_accuracy"`
	Satellites          []Satellite `json:"satellites"`
}

func (o *Observation) getInterfacePtrs() []interface{} {

	v := reflect.Indirect(reflect.ValueOf(o))
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Addr().Interface()
	}

	return values
}

func (o *Observation) Json() []byte {
	j, err := json.Marshal(o)
	if err != nil {
		log.Fatal("could not marshal json")
	}

	return j
}

func (o *Observation) String() []string {
	var s []string
	v := reflect.Indirect(reflect.ValueOf(o))
	s = make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		s[i] = fmt.Sprintf("%v", value)
	}

	return s
}

func (o *Observation) getSQLQuery() string {
	return `
			SELECT P.id,                                                           
				P.SysTime,                                                          
				P.Lat,                                                              
				P.Lon,                                                              
				P.Alt,                                                              
				P.Provider,
				P.GPSTime,                                                          
				P.FixSatCount,                                                      
				P.HasRadialAccuracy,                                                
				P.HasVerticalAccuracy,                                              
				P.RadialAccuracy,                                                   
				P.VerticalAccuracy,                                                 
				P.Speed,                                                            
				P.SpeedAccuracy                                                    
           	FROM gps_observation_points_sat_data AS G                              
			LEFT JOIN                                                         
			gps_observation_points AS P                                       
			ON G.base_id = P.id                                               
           	ORDER BY G.base_id
			`
}
