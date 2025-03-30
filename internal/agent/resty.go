package agent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"github.com/go-resty/resty/v2"
	"time"
)

type RestyAgent struct {
	client *resty.Client
}

func NewRestyAgent() *RestyAgent {
	return &RestyAgent{client: resty.New()}
}

func (r *RestyAgent) getAge(ctx context.Context, name string) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		resp, err := r.client.R().EnableTrace().Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
		if err != nil {
			return 0, err
		}

		var ageResp ageResponse
		err = json.Unmarshal(resp.Body(), &ageResp)
		if err != nil {
			return 0, err
		}

		return ageResp.Age, nil
	}
}

func (r *RestyAgent) getGender(ctx context.Context, name string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		resp, err := r.client.R().EnableTrace().Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
		if err != nil {
			return "", err
		}

		var genderResp genderResponse
		err = json.Unmarshal(resp.Body(), &genderResp)
		if err != nil {
			return "", err
		}

		return genderResp.Gender, nil
	}
}

func (r *RestyAgent) getNationalize(ctx context.Context, name string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		resp, err := r.client.R().EnableTrace().Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
		if err != nil {
			return "", err
		}

		var natResp nationalizeResponse
		err = json.Unmarshal(resp.Body(), &natResp)
		if err != nil {
			return "", err
		}

		return selectCountry(natResp.Country), nil
	}
}

// selectCountry позволяет выбрать страну с наибольшим процентом совпадения
func selectCountry(countries []country) string {
	var bestCountry country
	var bestPercent float32
	for _, c := range countries {
		if bestPercent < c.Probability {
			bestPercent = c.Probability
			bestCountry = c
		}
	}

	return bestCountry.CountryID
}

func (r *RestyAgent) EnrichInformation(ctx context.Context, user e.User) (*e.User, error) {
	if user.Name == nil || *user.Name == "" {
		return nil, errors.New("name is required to enrich")
	}

	name := *user.Name
	deadlineCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	age, err := r.getAge(deadlineCtx, name)
	if err != nil {
		return nil, err
	}
	gender, err := r.getGender(deadlineCtx, name)
	if err != nil {
		return nil, err
	}
	nationalize, err := r.getNationalize(deadlineCtx, name)
	if err != nil {
		return nil, err
	}

	user.Age, user.Gender, user.Nationality = &age, &gender, &nationalize

	return &user, nil
}
