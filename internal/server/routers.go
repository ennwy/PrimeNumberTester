package server

import (
	"net/http"
	"log"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func Primes(ctx *fiber.Ctx) error {
	ctx.Context().SetContentType(fiber.MIMEApplicationJSON)
	var numbers []int
	log.Println(ctx.Body())
	if err := json.Unmarshal(ctx.Body(), &numbers); err != nil {
		resp := ErrResponse{
			Code: http.StatusUnprocessableEntity,
			Err:  ErrInvalidInput,
		}

		b, err := json.Marshal(resp)
		ctx.Context().SetStatusCode(resp.Code)
		ctx.Context().SetBody(b)

		return err
	}

	b, err := json.Marshal(getPrimes(numbers))
	ctx.Context().SetBody(b)

	return err
}

func getPrimes(numbers []int) Response {
	primes := make(Response, 0, len(numbers))

	for _, n := range numbers {
		primes = append(primes, isPrime(n))
	}

	return primes
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n != 2 && n % 2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
