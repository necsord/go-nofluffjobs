package nofluffjobs

import (
	"fmt"
	"net/http"
)

type SearchPostingRequest struct {
	Page           int            `json:"page,omitempty"`
	RawSearch      string         `json:"rawSearch,omitempty"`
	CriteriaSearch CriteriaSearch `json:"criteriaSearch,omitempty"`
}
type CriteriaSearch struct {
	Id          []string         `json:"id,omitempty"`
	Category    []string         `json:"category,omitempty"`
	City        []string         `json:"city,omitempty"`
	Company     []string         `json:"company,omitempty"`
	Employment  []string         `json:"employment,omitempty"`
	JobLanguage []string         `json:"jobLanguage,omitempty"`
	JobPosition []string         `json:"jobPosition,omitempty"`
	Keyword     []string         `json:"keyword,omitempty"`
	More        []string         `json:"more,omitempty"`
	Requirement []string         `json:"requirement,omitempty"`
	Salary      []SalaryCriteria `json:"salary,omitempty"`
	Seniority   []string         `json:"seniority,omitempty"`
}
type SalaryCriteria struct {
	Currency    string `json:"currency,omitempty"`
	GreaterThan int64  `json:"greaterThan,omitempty"`
	Period      string `json:"period,omitempty"`
}
type SearchPostingQuery struct {
	Limit          int    `url:"limit,omitempty"`
	Offset         int    `url:"offset,omitempty"`
	SalaryCurrency string `url:"salaryCurrency,omitempty"`
	SalaryPeriod   string `url:"salaryPeriod,omitempty"`
	Region         string `url:"region,omitempty"`
}

type SearchPostingResponse struct {
	SearchPostingRequest
	TotalCount int       `json:"totalCount"`
	Postings   []Posting `json:"postings"`
}
type Posting struct {
	Id                       string   `json:"id"`
	Name                     string   `json:"name"`
	Location                 Location `json:"location"`
	Posted                   int64    `json:"posted"`
	Renewed                  int64    `json:"renewed"`
	Title                    string   `json:"title"`
	Technology               string   `json:"technology"`
	Logo                     Logo     `json:"logo"`
	Category                 string   `json:"category"`
	Seniority                []string `json:"seniority"`
	Url                      string   `json:"url"`
	Regions                  []string `json:"regions"`
	Salary                   Salary   `json:"salary"`
	Flavors                  []string `json:"flavors"`
	TopInSearch              bool     `json:"topInSearch"`
	Highlighted              bool     `json:"highlighted"`
	Help4Ua                  bool     `json:"help4Ua"`
	OnlineInterviewAvailable bool     `json:"onlineInterviewAvailable"`
}
type Location struct {
	Places            []Place `json:"places"`
	FullyRemote       bool    `json:"fullyRemote"`
	CovidTimeRemotely bool    `json:"covidTimeRemotely"`
}
type Place struct {
	City string `json:"city"`
	Url  string `json:"url"`
}
type Salary struct {
	From     int64  `json:"from"`
	To       int64  `json:"to"`
	Type     string `json:"type"`
	Currency string `json:"currency"`
}
type Logo struct {
	Original               string `json:"original"`
	JobsDetails            string `json:"jobs_details"`
	JobsListing            string `json:"jobs_listing"`
	JobsDetails2x          string `json:"jobs_details_2x"`
	JobsListing2x          string `json:"jobs_listing_2x"`
	CompaniesDetails       string `json:"companies_details"`
	CompaniesListing       string `json:"companies_listing"`
	JobsDetailsWebp        string `json:"jobs_details_webp"`
	JobsListingWebp        string `json:"jobs_listing_webp"`
	CompaniesDetails2x     string `json:"companies_details_2x"`
	CompaniesListing2x     string `json:"companies_listing_2x"`
	JobsDetails2xWebp      string `json:"jobs_details_2x_webp"`
	JobsListing2xWebp      string `json:"jobs_listing_2x_webp"`
	OriginalPixelImage     string `json:"original_pixel_image"`
	CompaniesDetailsWebp   string `json:"companies_details_webp"`
	CompaniesListingWebp   string `json:"companies_listing_webp"`
	CompaniesDetails2xWebp string `json:"companies_details_2x_webp"`
	CompaniesListing2xWebp string `json:"companies_listing_2x_webp"`
}

type ErrorResponse struct {
	Message  string
	Response *http.Response
}

func (res ErrorResponse) Error() string {
	return fmt.Sprintf(
		"%v %v: %d %s",
		res.Response.Request.Method,
		res.Response.Request.URL,
		res.Response.StatusCode,
		res.Message,
	)
}
