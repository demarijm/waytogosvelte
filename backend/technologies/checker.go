package technologies

import (
	"io/ioutil"
	"net/http"
)

func CheckTechnology(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(bodyBytes)
	technologies := []string{}

	if isWordpress(bodyString) {
		technologies = append(technologies, "WordPress")
	}

	if isWix(bodyString) {
		technologies = append(technologies, "Wix")
	}

	if isReact(bodyString) {
		technologies = append(technologies, "React")
	}

	if isTailwind(bodyString) {
		technologies = append(technologies, "Tailwind")
	}

	if isNextJS(bodyString) {
		technologies = append(technologies, "Next.js")
	}

	return technologies, nil
}
