package bcapi

import (
	"archive/zip"
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

// downloadZipArchive makes "GET" http request with the given http.Client
// to download .zip file with required data
func downloadZipArchive(client *http.Client) error {
	req, err := http.NewRequest(http.MethodGet, bcLink, nil)
	if err != nil {
		return wrapError(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return wrapError(err.Error())
	}
	defer resp.Body.Close()

	out, err := os.Create(zipFileName)
	if err != nil {
		return wrapError(err.Error())
	}

	defer out.Close()

	// cpu burst
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return wrapError(err.Error())
	}
	return nil
}

// openFile reads .dat from data.zip file and returns body as a string
func openFile(zipArchive *zip.ReadCloser, fileName string) (string, error) {
	for _, v := range zipArchive.File {
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
	return "", wrapError(fileName + " file not found")
}
