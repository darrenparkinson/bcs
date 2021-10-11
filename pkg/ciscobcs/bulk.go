package ciscobcs

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BulkResults struct {
	LineCount                  int
	CountOfTypes               map[string]int
	UnrecognisedTypes          map[string]int
	Devices                    []Device
	TrackSummaries             []TrackSummary
	TrackSmupieRecommendations []TrackSmupieRecommendation
	SWEoxBulletins             []SWEOXBulletin
	HWEoxBulletins             []HWEOXBulletin
	FNBulletins                []FNBulletin
	PSIRTBulletins             []PSIRTBulletin
	Errors                     []error
}

func (s *BulkService) Retrieve(ctx context.Context, customerID string) (*BulkResults, error) {
	url := fmt.Sprintf("%s/customer/%s/bulk/alerts", s.client.BaseURL, customerID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return s.client.makeBulkRequest(ctx, req)
}

func (s *BulkService) Download(ctx context.Context, customerID string, w io.Writer) error {
	url := fmt.Sprintf("%s/customer/%s/bulk/alerts", s.client.BaseURL, customerID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return s.client.makeRequestToWriter(ctx, req, w)
}

// ParseBulkFile will take a raw jsonlines file and return the results as a BulkResults struct
// Note that it is not part of a service since no details are required just to parse a file.
func ParseBulkFile(filename string) (*BulkResults, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return scanBulk(file)
}

// makeBulkRequest provides a function specifically for the bulk data which is sent as jsonlines format
func (c *Client) makeBulkRequest(ctx context.Context, req *http.Request) (*BulkResults, error) {
	req.Header.Add("x-api-key", c.APIKey)
	rc := req.WithContext(ctx)
	res, err := c.HTTPClient.Do(rc)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var ciscobcsErr error
		switch res.StatusCode {
		case 400:
			ciscobcsErr = ErrBadRequest
		case 401:
			ciscobcsErr = ErrUnauthorized
		case 403:
			ciscobcsErr = ErrForbidden
		case 500:
			ciscobcsErr = ErrInternalError
		default:
			ciscobcsErr = ErrUnknown
		}
		return nil, ciscobcsErr
	}
	return scanBulk(res.Body)
}

func scanBulk(body io.Reader) (*BulkResults, error) {
	results := &BulkResults{}
	scanner := bufio.NewScanner(body)
	// scanner has a limit of 65k, so lets set a larger buffer for it to use
	const maxCapacity = 65536 * 2
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	// Set up our results
	lineCount := 0
	countOfTypes := make(map[string]int)
	unrecognisedTypes := make(map[string]int)
	results.Devices = []Device{}
	results.TrackSummaries = []TrackSummary{}
	results.TrackSmupieRecommendations = []TrackSmupieRecommendation{}
	results.SWEoxBulletins = []SWEOXBulletin{}
	results.HWEoxBulletins = []HWEOXBulletin{}
	results.FNBulletins = []FNBulletin{}
	results.PSIRTBulletins = []PSIRTBulletin{}

	for scanner.Scan() {
		line := scanner.Bytes()
		lineCount++
		// first we need to check the line type before we can unmarshal it
		var lineType BulkTypeChecker
		err := json.Unmarshal(line, &lineType)
		if err != nil {
			return nil, errors.New("error unmarshalling type: check input file")
		}
		// Process each type from here
		switch lineType.Type {
		case "device":
			countOfTypes[lineType.Type]++
			var v Device
			err := json.Unmarshal(line, &v)
			if err != nil {
				results.Errors = append(results.Errors, fmt.Errorf("device: %w", err))
			}
			results.Devices = append(results.Devices, v)
		case "track_summary":
			countOfTypes[lineType.Type]++
			var v TrackSummary
			err := json.Unmarshal(line, &v)
			if err != nil {
				results.Errors = append(results.Errors, fmt.Errorf("track_summary: %w", err))
			}
			results.TrackSummaries = append(results.TrackSummaries, v)
		case "track_smupie_recommendation":
			countOfTypes[lineType.Type]++
			var v TrackSmupieRecommendation
			err := json.Unmarshal(line, &v)
			if err != nil {
				results.Errors = append(results.Errors, fmt.Errorf("track_smupie_recommendation: %w", err))
			}
			results.TrackSmupieRecommendations = append(results.TrackSmupieRecommendations, v)
		case "sw_eox_bulletin":
			countOfTypes[lineType.Type]++
			var v SWEOXBulletin
			err := json.Unmarshal(line, &v)
			if err != nil {
				results.Errors = append(results.Errors, fmt.Errorf("sw_eox_bulletin: %w", err))
			}
			results.SWEoxBulletins = append(results.SWEoxBulletins, v)
		case "hw_eox_bulletin":
			countOfTypes[lineType.Type]++
			var v HWEOXBulletin
			err := json.Unmarshal(line, &v)
			if err != nil {
				results.Errors = append(results.Errors, fmt.Errorf("hw_eox_bulletin: %w", err))
			}
			results.HWEoxBulletins = append(results.HWEoxBulletins, v)
		case "fn_bulletin":
			countOfTypes[lineType.Type]++
			var v FNBulletin
			err := json.Unmarshal(line, &v)
			if err != nil {
				results.Errors = append(results.Errors, fmt.Errorf("fn_bulletin: %w", err))
			}
			results.FNBulletins = append(results.FNBulletins, v)
		case "psirt_bulletin":
			countOfTypes[lineType.Type]++
			var v PSIRTBulletin
			err := json.Unmarshal(line, &v)
			if err != nil {
				results.Errors = append(results.Errors, fmt.Errorf("psirt_bulletin: %w", err))
			}
			results.PSIRTBulletins = append(results.PSIRTBulletins, v)
		default:
			unrecognisedTypes[lineType.Type]++
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	results.LineCount = lineCount
	results.CountOfTypes = countOfTypes
	results.UnrecognisedTypes = unrecognisedTypes
	return results, nil
}
