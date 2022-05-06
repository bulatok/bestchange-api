package bcapi

import (
	"archive/zip"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"net/http"
	"os"
)

const (
	bcLink     = "http://api.bestchange.ru/info.zip"
	exchLink   = "https://www.bestchange.ru/"
	marketLink = "https://www.bestchange.ru/click.php?"

	zipFileName = "data.zip"

	marketsFileName = "bm_exch.dat"
	coinsFileName   = "bm_cy.dat"
	ratesFileName   = "bm_rates.dat"
)

// getZipFile makes new "GET" http request with the given http.Client
// to download .zip file with required data
//
func getZipFile(client *http.Client) error{
	req, err := http.NewRequest("GET", bcLink, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(zipFileName)
	if err != nil{
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// openFile reads .dat from data.zip file and returns body as a string
func openFile(fileName string) (string, error) {
	zipFile, err := zip.OpenReader(zipFileName)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()
	for _, v := range zipFile.File {
		if v.Name == fileName {
			r, err := v.Open()
			if err != nil {
				return "", err
			}
			newR := charmap.Windows1251.NewDecoder().Reader(r)

			data, err := io.ReadAll(newR)
			if err != nil {
				return "", err
			}
			return string(data), nil
		}
	}
	return "", fmt.Errorf("bcapi : file %s was not found", fileName)
}
