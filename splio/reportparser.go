package splio

import (
	"errors"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"time"
)

type UnsubReason string

const (
	r1 UnsubReason = "I am not interested anymore"
	r2             = "I receive too many messages"
	r3             = "I can't read the messages"
	r4             = "I never asked to subscribe"
	rz             = "Other"
)

type IndividualEmailCampaignsReportLine struct {
	SendID        string           `csv:"sendID"`
	StartTime     time.Time        `csv:"startTine"`
	CampaignID    string           `csv:"campaignID"`
	Name          string           `csv:"name"`
	OperationCode string           `csv:"operationCode"`
	Category      string           `csv:"category"`
	Subject       string           `csv:"subject"`
	Email         string           `csv:"email"`
	Status        string           `csv:"status"`
	SmtpResponse  string           `csv:"SMTP response"`
	Open          int              `csv:"open"`
	Click         int              `csv:"click"`
	UnSub         bool             `csv:"unsub"`
	UnsubReason   UnsubReason      `csv:"unsub reason"`
	Spam          bool             `csv:"spam"`
	DateBacklist  *time.Time       `csv:"date blacklist"`
	Orders        bool             `csv:"orders"`
	Turnover      *decimal.Decimal `csv:"turnover"`
	Subscriptions []List           `csv:"subscriptions"`
}

type IndividualEmailCampaignsReport struct {
	headers []string
}

func NewIndividualEmailCampaignsReport(headers []string) *IndividualEmailCampaignsReport {
	return &IndividualEmailCampaignsReport{
		headers: headers,
	}
}

func (p *IndividualEmailCampaignsReport) ParseArray(line *IndividualEmailCampaignsReportLine, source []string) error {
	if len(source) != len(p.headers) {
		return errors.New("report array does not have ")
	}

	for key, value := range source {
		switch p.headers[key] {
		case "sendID":
			line.SendID = value
		case "startTine":
		case "campaignID":
			line.CampaignID = value
		case "name":
			line.Name = value
		case "operationCode":
			line.OperationCode = value
		case "category":
			line.Category = value
		case "subject":
			line.Subject = value
		case "email":
			line.Email = value
		case "status":
			line.Status = value
		case "SMTP response":
			line.SmtpResponse = value
		case "open":
			line.Open, _ = strconv.Atoi(value)
		case "click":
			line.Click, _ = strconv.Atoi(value)
		case "unsub":
			line.UnSub = p.getBool(value)
		case "unsub reason":
			line.UnsubReason = p.getUnsubReason(value)
		case "spam":
			line.Spam = p.getBool(value)
		case "date blacklist":
		case "orders":
			line.Orders = p.getBool(value)
		case "turnover":
		case "subscriptions":
			line.Subscriptions = p.getSubscriptionsList(value)
		}
	}

	return nil
}

func (p *IndividualEmailCampaignsReport) getBool(v string) bool {
	v = strings.ToLower(v)
	if v == "true" || v == "y" || v == "1" {
		return true
	}
	return false
}

func (p *IndividualEmailCampaignsReport) getUnsubReason(v string) UnsubReason {
	return UnsubReason(v)
}

func (p *IndividualEmailCampaignsReport) getSubscriptionsList(v string) []List {
	var list []List

	v = strings.ReplaceAll(v, "+", "")
	listCode := strings.Split(v, ",")
	for _, idStr := range listCode {
		id, _ := strconv.Atoi(idStr)
		list = append(list, List{Id: id})
	}

	return list
}