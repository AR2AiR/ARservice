// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()
package mvc

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Welcome []WelcomeElement

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WelcomeElement struct {
	ID               int64             `json:"id"`
	SamplingRate     interface{}       `json:"sampling_rate"`
	Timestamp        string            `json:"timestamp"`
	Location         Location          `json:"location"`
	Sensor           Sensor            `json:"sensor"`
	Sensordatavalues []Sensordatavalue `json:"sensordatavalues"`
}

type Location struct {
	ID            int64   `json:"id"`
	Latitude      string  `json:"latitude"`
	Longitude     string  `json:"longitude"`
	Altitude      string  `json:"altitude"`
	Country       Country `json:"country"`
	ExactLocation int64   `json:"exact_location"`
	Indoor        int64   `json:"indoor"`
}

type Sensor struct {
	ID         int64      `json:"id"`
	Pin        string     `json:"pin"`
	SensorType SensorType `json:"sensor_type"`
}

type SensorType struct {
	ID           int64        `json:"id"`
	Name         Name         `json:"name"`
	Manufacturer Manufacturer `json:"manufacturer"`
}

type Sensordatavalue struct {
	ID        *int64    `json:"id,omitempty"`
	Value     *Value    `json:"value"`
	ValueType ValueType `json:"value_type"`
}

type Country string
const (
	AE Country = "AE"
	Al Country = "AL"
	Ar Country = "AR"
	At Country = "AT"
	Au Country = "AU"
	Ba Country = "BA"
	Be Country = "BE"
	Bg Country = "BG"
	Bo Country = "BO"
	Br Country = "BR"
	By Country = "BY"
	CA Country = "CA"
	CD Country = "CD"
	CN Country = "CN"
	CR Country = "CR"
	Ch Country = "CH"
	Cl Country = "CL"
	Cy Country = "CY"
	Cz Country = "CZ"
	De Country = "DE"
	Dk Country = "DK"
	Do Country = "DO"
	Es Country = "ES"
	Fi Country = "FI"
	Fr Country = "FR"
	GB Country = "GB"
	Ge Country = "GE"
	Gf Country = "GF"
	Gr Country = "GR"
	Hk Country = "HK"
	Hu Country = "HU"
	ID Country = "ID"
	IL Country = "IL"
	IR Country = "IR"
	Ie Country = "IE"
	In Country = "IN"
	Is Country = "IS"
	It Country = "IT"
	Jp Country = "JP"
	Kg Country = "KG"
	Kr Country = "KR"
	Kw Country = "KW"
	Kz Country = "KZ"
	LV Country = "LV"
	La Country = "LA"
	Li Country = "LI"
	Lu Country = "LU"
	MX Country = "MX"
	Mk Country = "MK"
	My Country = "MY"
	Ng Country = "NG"
	Nl Country = "NL"
	No Country = "NO"
	Nz Country = "NZ"
	PE Country = "PE"
	Ph Country = "PH"
	Pk Country = "PK"
	Pl Country = "PL"
	Pt Country = "PT"
	Ro Country = "RO"
	Rs Country = "RS"
	Ru Country = "RU"
	SE Country = "SE"
	Sg Country = "SG"
	Si Country = "SI"
	Sk Country = "SK"
	Th Country = "TH"
	Tr Country = "TR"
	Ua Country = "UA"
	Us Country = "US"
	Vn Country = "VN"
	Xk Country = "XK"
	Za Country = "ZA"
)

type Manufacturer string
const (
	Bosch Manufacturer = "Bosch"
	DallasSemiconductor Manufacturer = "Dallas semiconductor"
	EcoCurious Manufacturer = "EcoCurious"
	Honeywell Manufacturer = "Honeywell"
	LuftdatenInfo Manufacturer = "Luftdaten.info"
	MeasurementSpecialties Manufacturer = "Measurement Specialties"
	NovaFitness Manufacturer = "Nova Fitness"
	Plantower Manufacturer = "Plantower"
	SensirionAG Manufacturer = "Sensirion AG"
	Shinyei Manufacturer = "Shinyei"
	Various Manufacturer = "various"
)

type Name string
const (
	Bme280 Name = "BME280"
	Bmp180 Name = "BMP180"
	Bmp280 Name = "BMP280"
	Dht22 Name = "DHT22"
	Ds18B20 Name = "DS18B20"
	Ds18S20 Name = "DS18S20"
	Hpm Name = "HPM"
	Htu21D Name = "HTU21D"
	Laerm Name = "Laerm"
	Pms1003 Name = "PMS1003"
	Pms3003 Name = "PMS3003"
	Pms5003 Name = "PMS5003"
	Pms7003 Name = "PMS7003"
	Ppd42NS Name = "PPD42NS"
	RadiationSBM20 Name = "Radiation SBM-20"
	Sds011 Name = "SDS011"
	Sds021 Name = "SDS021"
	Sht31 Name = "SHT31"
)

type ValueType string
const (
	CountsPerMinute ValueType = "counts_per_minute"
	DurP1 ValueType = "durP1"
	DurP2 ValueType = "durP2"
	Humidity ValueType = "humidity"
	MaxMicro ValueType = "max_micro"
	MinMicro ValueType = "min_micro"
	NoiseLAMax ValueType = "noise_LA_max"
	NoiseLAMin ValueType = "noise_LA_min"
	NoiseLAeq ValueType = "noise_LAeq"
	P0 ValueType = "P0"
	P1 ValueType = "P1"
	P2 ValueType = "P2"
	Pressure ValueType = "pressure"
	PressureAtSealevel ValueType = "pressure_at_sealevel"
	RatioP1 ValueType = "ratioP1"
	RatioP2 ValueType = "ratioP2"
	Samples ValueType = "samples"
	Temperature ValueType = "temperature"
)

type Value struct {
	Double *float64
	String *string
}

func (x *Value) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Value) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}

