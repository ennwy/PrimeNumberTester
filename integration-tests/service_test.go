package integration_test

import (
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"math"

	"github.com/ennwy/prime_number_tester/internal/server"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

var (
	addr        = net.JoinHostPort(os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))
	url         = "http://" + addr + "/primes"
	contentType = "application/json"
)

var ErrResp = server.ErrResponse{
				Code: 422,
				Err:  "the given input is invalid",
			}

type TestCase struct {
	Data     any
	Expected server.Response
}

type ErrTestCase struct {
	Data     any
	Expected server.ErrResponse
}

func TestPrimes(t *testing.T) {
	testCase := TestCase{
		Data:     []int{5, 12, 7, 37},
		Expected: server.Response{true, false, true, true},
	}

	b, err := json.Marshal(testCase.Data)
	require.Nil(t, err)
	log.Println(string(b))

	resp, err := http.Post(url, contentType, bytes.NewReader(b))
	require.Nil(t, err)

	defer resp.Body.Close()

	var result server.Response
	b, err = io.ReadAll(resp.Body)
	require.Nil(t, err)
	json.Unmarshal(b, &result)

	require.Equal(t, testCase.Expected, result)
}

func TestForError(t *testing.T) {
	testCase := ErrTestCase{
		Data: []string{"1", "3", "7", "Nan"},
		Expected: ErrResp,
	}

	b, err := json.Marshal(testCase.Data)
	require.Nil(t, err)
	log.Println(string(b))

	resp, err := http.Post(url, contentType, bytes.NewReader(b))
	require.Nil(t, err)

	defer resp.Body.Close()

	var result server.ErrResponse
	b, err = io.ReadAll(resp.Body)
	require.Nil(t, err)
	json.Unmarshal(b, &result)

	require.Equal(t, testCase.Expected, result)
}

func TestIntOverflow(t *testing.T) {
	testCase := ErrTestCase{
		Data: []uint64{
			math.MaxUint64,
			18446744073709551557, // largest 64 bit prime that overflows int64
		},
		Expected: ErrResp,
	}

	b, err := json.Marshal(testCase.Data)
	require.Nil(t, err)
	log.Println(string(b))

	resp, err := http.Post(url, contentType, bytes.NewReader(b))
	require.Nil(t, err)

	defer resp.Body.Close()

	var result server.ErrResponse
	b, err = io.ReadAll(resp.Body)
	require.Nil(t, err)
	json.Unmarshal(b, &result)

	require.Equal(t, testCase.Expected, result)
}
