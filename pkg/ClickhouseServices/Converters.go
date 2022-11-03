package ClickhouseServices

import (
	DataStructures "Portal/Core/internal/structures"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

func ConvertDataToCHModel(topic, data string, key string, probes map[string]string, consType string) (interface{}, error) {
	switch topic {
	case "newcust-0":
		var chModel DataStructures.CustomerTrafficDown

		chModel, err := ConvertCustomerTrafficDownToCHModel(data)
		AddAdditionalFields(&chModel.GeneralTrafficFields, key, probes, consType)
		if err != nil {
			return nil, err
		}

		return chModel, nil
	case "newcust-1":
		var chModel DataStructures.CustomerTrafficUp

		chModel, err := ConvertCustomerTrafficUpToCHModel(data)
		AddAdditionalFields(&chModel.GeneralTrafficFields, key, probes, consType)

		if err != nil {
			return nil, err
		}

		return chModel, nil
	}

	return nil, fmt.Errorf("topic is not defined. Tried:  %s", topic)
}

func ConvertCustomerTrafficDownToCHModel(data string) (DataStructures.CustomerTrafficDown, error) {
	var buffer []DataStructures.CustomerTrafficDown

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})
	r := strings.NewReader(data)
	err := gocsv.UnmarshalWithoutHeaders(r, &buffer)

	if err != nil {
		return DataStructures.CustomerTrafficDown{}, err
	}

	return buffer[0], nil
}

func ConvertCustomerTrafficUpToCHModel(data string) (DataStructures.CustomerTrafficUp, error) {
	var buffer []DataStructures.CustomerTrafficUp

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})
	r := strings.NewReader(data)
	err := gocsv.UnmarshalWithoutHeaders(r, &buffer)

	return buffer[0], err
}

func getCountryByProb(probe string, probesToCountries map[string]string) (string, error) {
	for in := range probesToCountries {
		if in == probe { // toDo: make them all lowercase or rely on strict config?
			return probesToCountries[probe], nil
		}
	}

	return "", fmt.Errorf("cannot find a country for given probe. Probe: %s", probe)
}

func AddAdditionalFields(fields *DataStructures.GeneralTrafficFields, key string, probesToCountries map[string]string, consType string) {
	fields.Probe = strings.Split(key, "#")[0]
	country, err := getCountryByProb(fields.Probe, probesToCountries)

	if err != nil {
		log.Fatalln(err)
	}

	fields.Country = country

	i, err := strconv.ParseInt(strings.Split(key, "#")[1], 10, 64)
	if err != nil {
		log.Fatalln("Error in probe timestamp. Got: ", strings.Split(key, "#")[1])
	}

	fields.ProbeTimestamp = time.Unix(i, 0)

	consTypeUint, er := strconv.ParseUint(consType, 10, 0)

	if er != nil {
		log.Fatalln("Error in given typeConnection. Must be 0 or 1. Got : ", consType)
	}
	fields.TypeConnection = uint8(consTypeUint)
}
