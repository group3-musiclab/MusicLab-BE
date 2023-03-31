package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"musiclab-be/app/config"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleAPI interface {
	GetUrlAuth(state string) string
	GetToken(code string) (*oauth2.Token, error)
	CreateCalendar(token *oauth2.Token, detail CalendarDetail) error
	GetUserInfo(token *oauth2.Token) (GoogleCore, error)
}

type CalendarDetail struct {
	Summary             string
	Location            string
	StartTime           string
	EndTime             string
	EndDate             string
	CreatorDisplayName  string
	CreatorEmail        string
	AttendeeDisplayName string
	AttendeeEmail       string
}

type GoogleCore struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

type googleAPI struct {
	conf *oauth2.Config
}

func NewGoogleApi(cfg *config.DBConfig) GoogleAPI {
	conf := &oauth2.Config{
		RedirectURL:  cfg.GOOGLE_REDIRECT_CALLBACK,
		ClientID:     cfg.GOOGLE_CLIENT_ID,
		ClientSecret: cfg.GOOGLE_CLIENT_SECRET,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}

	return &googleAPI{conf: conf}
}

func (ga *googleAPI) GetUrlAuth(state string) string {
	return ga.conf.AuthCodeURL(state)
}

func (ga *googleAPI) GetToken(code string) (*oauth2.Token, error) {
	token, err := ga.conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (ga *googleAPI) CreateCalendar(token *oauth2.Token, detail CalendarDetail) error {
	calendarService, err := calendar.NewService(context.Background(), option.WithTokenSource(ga.conf.TokenSource(context.Background(), token)))
	if err != nil {
		return err
	}

	t1 := strings.Replace(detail.EndDate, "-", "", -1)
	t2 := strings.Replace(t1, ":", "", -1)

	endDate := t2[:15] + "Z"

	RecurrenceString := fmt.Sprintf("RRULE:FREQ=WEEKLY;UNTIL=%s", endDate)

	event := &calendar.Event{
		Summary:  detail.Summary,
		Location: detail.Location,
		Start: &calendar.EventDateTime{
			DateTime: detail.StartTime,
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			DateTime: detail.EndTime,
			TimeZone: "Asia/Jakarta",
		},
		Recurrence: []string{
			RecurrenceString,
		},
		Creator: &calendar.EventCreator{
			DisplayName: detail.CreatorDisplayName,
			Email:       detail.CreatorEmail,
		},
		Attendees: []*calendar.EventAttendee{
			{Email: detail.AttendeeEmail, DisplayName: detail.AttendeeDisplayName},
		},
	}

	_, err = calendarService.Events.Insert("primary", event).SendUpdates("all").Do()
	if err != nil {
		return err
	}

	return nil
}

func (ga *googleAPI) GetUserInfo(token *oauth2.Token) (GoogleCore, error) {

	var userGoogleCore GoogleCore

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("failed getting user info")
		return GoogleCore{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("failed reading response body")
		return GoogleCore{}, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	errjson := json.Unmarshal(contents, &userGoogleCore)
	if errjson != nil {
		log.Println("cant unmarshal json")
		return GoogleCore{}, fmt.Errorf("cant unmarshal json")
	}

	return userGoogleCore, nil
}
