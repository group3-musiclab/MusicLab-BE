package helper

import (
	"context"
	"musiclab-be/app/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleAPI interface {
	GetUrlAuth(state string) string
	GetToken(code string) (*oauth2.Token, error)
	CreateCalendar(token *oauth2.Token, detail CalendarDetail) error
}

type CalendarDetail struct {
	Summary  string
	Location string
	Start    string
	End      string
	Emails   []string
}

type googleAPI struct {
	conf *oauth2.Config
}

func NewGoogleApi(cfg *config.DBConfig) GoogleAPI {
	conf := &oauth2.Config{
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

	attendees := []*calendar.EventAttendee{}
	for _, e := range detail.Emails {
		a := &calendar.EventAttendee{Email: e}
		attendees = append(attendees, a)
	}

	event := &calendar.Event{
		Summary:  detail.Summary,
		Location: detail.Location,
		Start: &calendar.EventDateTime{
			DateTime: detail.Start, // Wajib format RFC3339
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			// DateTime: time.Date(2023, 02, 10, 13, 20, 10, 10, time.Local).Format(time.RFC3339), // artinya YYYY-MM-DD HH-MM-SS-NS Location
			DateTime: detail.End,
			TimeZone: "Asia/Jakarta",
		},
		Attendees: attendees,
	}

	_, err = calendarService.Events.Insert("primary", event).SendUpdates("all").Do()
	if err != nil {
		return err
	}

	return nil
}
