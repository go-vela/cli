// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Add creates a worker based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for worker configuration")

	// send API call to get a registration token for the given worker
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela#AdminWorkerService.RegisterToken
	registerToken, _, err := client.Admin.Worker.RegisterToken(c.Hostname)
	if err != nil || registerToken == nil {
		return fmt.Errorf("unable to retrieve registration token: %w", err)
	}

	logrus.Tracef("got registration token, adding worker %q", c.Hostname)

	// create a custom http client using the registration token received
	// from .RegisterToken as a Bearer token.
	//
	// we will make a call to the registration endpoint on the worker
	// at the given c.Address for the worker.
	workerRegistrationURL := strings.TrimSuffix(c.Address, "/")
	workerRegistrationURL = fmt.Sprintf("%s/register", workerRegistrationURL)

	// create a new request for the given URL (c.Address)
	req, err := http.NewRequestWithContext(context.Background(), "POST", workerRegistrationURL, nil)
	if err != nil {
		return fmt.Errorf("unable to form request for worker registration endpoint")
	}

	// add the authorization header using the registration token
	req.Header.Add("Authorization", "Bearer "+registerToken.GetToken())

	// add the user agent for the request
	req.Header.Add("User-Agent", client.UserAgent)

	// create a new http client
	httpClient := http.DefaultClient
	httpClient.Timeout = time.Second * 15

	// perform the request to the worker registration endpoint
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error on response to worker registration endpoint")
	}
	defer resp.Body.Close()

	// read the body response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body for worker registration call")
	}

	// if the call was successful but didn't register successfully
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error while registering worker at %q, received: %d - %s", workerRegistrationURL, resp.StatusCode, string(bodyBytes))
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
		return output.JSON(out, c.Color)
	case output.DriverSpew:
		// output the worker in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(out)
	case output.DriverYAML:
		// output the worker in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(out, c.Color)
	default:
		// output the worker in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(out)
	}
}
