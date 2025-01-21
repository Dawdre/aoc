package lib

import (
	"fmt"
	"io"
	"net/http"
)

func FetchAOCInput(url string, cookie *http.Cookie, client *http.Client) (string, error) {
	url_day := fmt.Sprintf("https://adventofcode.com/2022/day/%s/input", url)

	request, err := http.NewRequest("GET", url_day, nil)

	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	request.AddCookie(cookie)

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("can't read response body: %w", err)
	}

	return string(body), nil
}
