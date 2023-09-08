// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"fmt"
	"github.com/go-vela/cli/internal/output"
	"net/http"
	"strings"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Add creates a worker based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for worker configuration")
	// send API call to get a worker registration for the given worker
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela#AdminWorkerService.Register
	wr, _, err := client.Admin.Worker.RegisterWorker(c.Hostname)
	if err != nil || wr == nil {
		return fmt.Errorf("unable to retrieve worker registration details: %w", err)
	}

	logrus.Tracef("got worker registration details, adding worker %q", c.Hostname)

	// create a custom http client using the registration token received
	// from .RegisterToken as a Bearer token.
	//
	// we will make a call to the registration endpoint on the worker
	// at the given c.Address for the worker.
	workerRegistrationURL := strings.TrimSuffix(c.Address, "/")
	workerRegistrationURL = fmt.Sprintf("%s/register", workerRegistrationURL)

	// send request using client
	resp, err := client.Call("POST", workerRegistrationURL, wr, nil)
	// create a new request for the given URL (c.Address)
	//req, err := http.NewRequestWithContext(context.Background(), "POST", workerRegistrationURL, nil)
	if err != nil {
		return fmt.Errorf("unable to form request for worker registration endpoint")
	}

	// add the authorization header using the registration token
	//req.Header.Add("Authorization", "Bearer "+registerToken.GetToken())
	//
	//// add the user agent for the request
	//req.Header.Add("User-Agent", client.UserAgent)

	// create a new http client
	//httpClient := http.DefaultClient
	//httpClient.Timeout = time.Second * 15
	//
	//// perform the request to the worker registration endpoint
	//resp, err := httpClient.Do(req)
	//if err != nil {
	//	return fmt.Errorf("error on response to worker registration endpoint")
	//}
	//defer resp.Body.Close()

	// read the body response
	//bodyBytes, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return fmt.Errorf("unable to read response body for worker registration call %v", resp.StatusCode)
	//}

	// if the call was successful but didn't register successfully
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error while registering worker at %q, received: %d - %s", workerRegistrationURL, resp.StatusCode, resp.Body.Close())
	}

	out := fmt.Sprintf("worker %q registered successfully", c.Hostname)

	logrus.Tracef("worker %q registered", c.Hostname)

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the worker in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(out)
	case output.DriverJSON:
		// output the worker in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(out)
	case output.DriverSpew:
		// output the worker in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(out)
	case output.DriverYAML:
		// output the worker in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(out)
	default:
		// output the worker in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(out)
	}
}
