package dbseeds

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

    "github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// XMLEnvelope has XMLCube
type XMLEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Cube    XMLCube  `xml:"Cube"`
}

// XMLCube has XMLCubeTime array
type XMLCube struct {
	XMLName   xml.Name      `xml:"Cube"`
	CubeTimes []XMLCubeTime `xml:"Cube"`
}

// XMLCubeTime has time and XMLCubeRate array
type XMLCubeTime struct {
	XMLName   xml.Name      `xml:"Cube"`
	Time      string        `xml:"time,attr"`
	CubeRates []XMLCubeRate `xml:"Cube"`
}

// XMLCubeRate has currency and rate
type XMLCubeRate struct {
	XMLName  xml.Name `xml:"Cube"`
	Currency string   `xml:"currency,attr"`
	Rate     string   `xml:"rate,attr"`
}

// Rate has model
type Rate struct {
	Time     string
	Currency string
	Rate     string
}

// ReadEnvelope reads the xml envelope
func ReadEnvelope(reader io.Reader) (XMLEnvelope, error) {
	var xmlEnvelope XMLEnvelope
	if err := xml.NewDecoder(reader).Decode(&xmlEnvelope); err != nil {
		return xmlEnvelope, err
	}
	return xmlEnvelope, nil
}

// GetRates gets the currency rate from URL, save to .xml file and returns filename
func GetRates() string {
	print("Reading from the extenal XML file Source and saving it in local server..........")
	url := "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	if response.Status != "200 OK" {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}

	xmlData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := time.Now()
	fileName := "downloads/eurofxref-hist-90d-" + t.Format("2006-01-02-15-04-05") + ".xml"
	f, err := os.Create(fileName)
	
	w := bufio.NewWriter(f)
	n, err := w.WriteString(string(xmlData))
	fmt.Printf("wrote %d bytes", n)
	w.Flush()
	defer f.Close()

	print("...Done\n")
	return fileName
}

// SaveRate saves to database
func SaveRate(rates []*Rate, db *gorm.DB) {
	// todo: save to database
	fmt.Println("Printing first item [Index=0]")
	fmt.Println(fmt.Sprintf("Time: %s Currency: %s Rate: %s", rates[0].Time, rates[0].Currency, rates[0].Rate))
	lastIndex := len(rates) - 1
	fmt.Println(fmt.Sprintf("Printing last item [Index=%d]", lastIndex))
	fmt.Println(fmt.Sprintf("Time: %s Currency: %s Rate: %s", rates[lastIndex].Time, rates[lastIndex].Currency, rates[lastIndex].Rate))
	fmt.Println("--------------Saving to Database start")
	start := time.Now()

	//defer db.DB().Close()*********************

	// db.DB() is done to get generic database object `*sql.DB` to use its functions
	//START FOR PREPARED STATEMENT WITH TRANSACTION
	txn, err := db.DB().Begin()
	if err != nil {
		log.Fatal(err)
	}

	//Delete old data if exists
	deleteStatement := `TRUNCATE table exchange_rates`
	_, err = db.DB().Exec(deleteStatement)
	if err != nil {
		panic(err)
	}

	//Prepared statement for bulk import
	stmt, err := txn.Prepare(pq.CopyIn("exchange_rates", "currency_symbol", "currency_rate", "currency_date"))

	if err != nil {
		log.Fatal(err)
	}

	for _, rate := range rates {
		_, err := stmt.Exec(rate.Currency, rate.Rate, rate.Time)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Flush all buffered data by running Exec() again
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Save to database took %s for %d rows", elapsed, len(rates)))
	fmt.Println("--------------Saving to Database end")
}

func ImportRates(db *gorm.DB) {
	fmt.Println("Start..........")
	filename := GetRates()
	//print(filename)
	// filepath.Abs appends the file name to the default working directly
	cubesFilePath, err := filepath.Abs(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open the local xml file
	file, err := os.Open(cubesFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// Read the cubes file
	xmlEnvelope, err := ReadEnvelope(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	xmlCubes := xmlEnvelope.Cube

	rates := []*Rate{}

	for i := 0; i < len(xmlCubes.CubeTimes); i++ {
		cubeTime := xmlCubes.CubeTimes[i]
		for j := 0; j < len(cubeTime.CubeRates); j++ {
			cubeRate := cubeTime.CubeRates[j]

			// get values
			r := new(Rate)
			r.Time = cubeTime.Time
			r.Currency = cubeRate.Currency
			r.Rate = cubeRate.Rate

			rates = append(rates, r)

		}
	}
	SaveRate(rates, db)
	//fmt.Println(filename)
}