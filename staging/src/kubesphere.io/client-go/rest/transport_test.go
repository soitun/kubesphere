/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

import (
	"context"
	"net"
	"sync"
	"testing"
)

// TestTransportForThreadSafe is meant to be run with the race detector
// to detect that the Transport generation for client-go is safe to
// call from multiple goroutines.
func TestTransportForThreadSafe(t *testing.T) {
	const (
		rootCACert = `-----BEGIN CERTIFICATE-----
MIIC4DCCAcqgAwIBAgIBATALBgkqhkiG9w0BAQswIzEhMB8GA1UEAwwYMTAuMTMu
MTI5LjEwNkAxNDIxMzU5MDU4MB4XDTE1MDExNTIxNTczN1oXDTE2MDExNTIxNTcz
OFowIzEhMB8GA1UEAwwYMTAuMTMuMTI5LjEwNkAxNDIxMzU5MDU4MIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAunDRXGwsiYWGFDlWH6kjGun+PshDGeZX
xtx9lUnL8pIRWH3wX6f13PO9sktaOWW0T0mlo6k2bMlSLlSZgG9H6og0W6gLS3vq
s4VavZ6DbXIwemZG2vbRwsvR+t4G6Nbwelm6F8RFnA1Fwt428pavmNQ/wgYzo+T1
1eS+HiN4ACnSoDSx3QRWcgBkB1g6VReofVjx63i0J+w8Q/41L9GUuLqquFxu6ZnH
60vTB55lHgFiDLjA1FkEz2dGvGh/wtnFlRvjaPC54JH2K1mPYAUXTreoeJtLJKX0
ycoiyB24+zGCniUmgIsmQWRPaOPircexCp1BOeze82BT1LCZNTVaxQIDAQABoyMw
ITAOBgNVHQ8BAf8EBAMCAKQwDwYDVR0TAQH/BAUwAwEB/zALBgkqhkiG9w0BAQsD
ggEBADMxsUuAFlsYDpF4fRCzXXwrhbtj4oQwcHpbu+rnOPHCZupiafzZpDu+rw4x
YGPnCb594bRTQn4pAu3Ac18NbLD5pV3uioAkv8oPkgr8aUhXqiv7KdDiaWm6sbAL
EHiXVBBAFvQws10HMqMoKtO8f1XDNAUkWduakR/U6yMgvOPwS7xl0eUTqyRB6zGb
K55q2dejiFWaFqB/y78txzvz6UlOZKE44g2JAVoJVM6kGaxh33q8/FmrL4kuN3ut
W+MmJCVDvd4eEqPwbp7146ZWTqpIJ8lvA6wuChtqV8lhAPka2hD/LMqY8iXNmfXD
uml0obOEy+ON91k+SWTJ3ggmF/U=
-----END CERTIFICATE-----`
		authPluginName = "auth-plugin-tr"
	)

	configurations := []struct {
		name   string
		config *Config
	}{
		{
			name: "simple config",
			config: &Config{
				Host:     "localhost:8080",
				Username: "gopher",
				Password: "g0ph3r",
			},
		},
		{
			name: "not cacheable config",
			config: &Config{
				Host:     "localhost:8080",
				Username: "gopher",
				Password: "g0ph3r",
				Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
					return nil, nil
				},
			},
		},
		{
			name: "with TLS config",
			config: &Config{
				Host:     "localhost:8080",
				Username: "gopher",
				Password: "g0ph3r",
				TLSClientConfig: TLSClientConfig{
					CAData: []byte(rootCACert),
				},
			},
		},
		{
			name: "with auth provider",
			config: &Config{
				Host:     "localhost:8080",
				Username: "gopher",
				Password: "g0ph3r",
				TLSClientConfig: TLSClientConfig{
					CAData: []byte(rootCACert),
				},
			},
		},
		{
			name: "with exec provider",
			config: &Config{
				Host:     "localhost:8080",
				Username: "gopher",
				Password: "g0ph3r",
				TLSClientConfig: TLSClientConfig{
					CAData: []byte(rootCACert),
				},
			},
		},
	}
	var wg sync.WaitGroup

	for _, tt := range configurations {
		tt := tt
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 50; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					_, err := TransportFor(tt.config)
					if err != nil {
						t.Errorf("Config: %s TransportFor() error = %v", tt.name, err)
					}
				}()
			}

		}()
	}
	wg.Wait()
}
