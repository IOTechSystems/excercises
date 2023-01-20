package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// Output JSON
type Output struct {
	ValueTotal int
	UUIDS      []string
}

// Structure for a single array instance of Device
type Devices struct {
	Devices []Device
}

// Structure for a single Device instance
type Device struct {
	Name      string //Device Name
	Type      string //Device Type
	Info      string //Device info contains UUID and String of information
	Value     string //Device value in base64
	TimeStamp string //UTC timestamp
}

// Used for spliting the UUID from the text
func getUUID(str string) string {
	sec1 := strings.Split(str, ",")     // Spliting with ',' to get the first string
	sec2 := strings.Split(sec1[0], " ") // Spliting with ' ' to get the fourth string
	sec3 := strings.Split(sec2[3], ":") // Spliting with ':' to get the second string which contains UUID
	return sec3[1]
}

// Function to decode Base64 to String and converting to integer type
func getTotal(str string) int {
	sDec, _ := base64.StdEncoding.DecodeString(str)
	num, _ := strconv.Atoi(string(sDec))
	return num
}

// Main Function
func main() {
	jsonFile, err := os.Open("data/data.json")

	if err != nil { // Open
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile) //Reading the JSON File

	var devices Devices //Creating the Device instance

	json.Unmarshal(byteValue, &devices) //Spliting the json

	var out Output //Creating a single instance of output.

	//Extaracting the totalValue and UUIDS from the json by the requirement
	for i := 0; i < len(devices.Devices); i++ {
		n, _ := strconv.ParseInt(devices.Devices[i].TimeStamp, 10, 64) //Parsing the Unix TimeStamp

		if n > time.Now().Unix() { // Checks the UNIX time

			out.UUIDS = append(out.UUIDS, getUUID(devices.Devices[i].Info))
			out.ValueTotal = out.ValueTotal + getTotal(devices.Devices[i].Value)
		}
	}

	outJson, err := json.MarshalIndent(out, "", "    ") //For readable format collecting the output

	_ = ioutil.WriteFile("output.json", outJson, 0644) //Writing the file in "output.json"
}
