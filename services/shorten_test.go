package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	model "github.com/harisaginting/gwyn/models"
	httpModel "github.com/harisaginting/gwyn/models/http"
	mockedRepo "github.com/harisaginting/gwyn/repositories/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ShortenServiceTestSuite struct {
	suite.Suite
	service Shorten
}

// DUMMY DATA
var shortenServiceDummyData = []model.Shorten{
	{
		ID:        1,
		Shortcode: "abc",
		Url:       "google.com",
	},
}

// this function executes before the test suite begins execution
func (s *ShortenServiceTestSuite) SetupSuite() {
	mockedRepo := new(mockedRepo.MockShortenRepository)
	mockedRepo.On("FindAll", mock.Anything).Return(shortenServiceDummyData, nil)
	s.service.repo = mockedRepo
	fmt.Println(">>> From SetupSuite", s.service.repo)
}

// this function executes after all tests executed
func (s *ShortenServiceTestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// this function executes before each test case
func (s *ShortenServiceTestSuite) SetupTest() {
	fmt.Println("-- From SetupTest")
}

// this function executes after each test case
func (s *ShortenServiceTestSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func TestShortenService(t *testing.T) {
	suite.Run(t, new(ShortenServiceTestSuite))
}

func (s *ShortenServiceTestSuite) TestList() {
	type args struct {
		ctx context.Context
		res *httpModel.ResponseList
	}
	tests := []struct {
		name    string
		service *Shorten
		args    args
		wantErr bool
	}{
		{
			name:    "test get list",
			service: &s.service,
			args: args{
				ctx: context.Background(),
				res: &httpModel.ResponseList{},
			},
		},
	}
	for _, tt := range tests {
		tt.service.List(tt.args.ctx, tt.args.res)
		s.Greater(len(tt.args.res.Items), 0, "is data not zero")
	}
}

func TestShorten_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		req httpModel.RequestCreate
	}
	tests := []struct {
		name       string
		service    *Shorten
		args       args
		wantRes    httpModel.ResponseCreate
		wantStatus int
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotStatus, err := tt.service.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shorten.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Shorten.Create() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("Shorten.Create() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestShorten_Status(t *testing.T) {
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name       string
		service    *Shorten
		args       args
		wantRes    model.Shorten
		wantStatus int
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotStatus, err := tt.service.Status(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shorten.Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Shorten.Status() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("Shorten.Status() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestShorten_Execute(t *testing.T) {
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name       string
		service    *Shorten
		args       args
		wantRes    model.Shorten
		wantStatus int
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotStatus, err := tt.service.Execute(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shorten.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Shorten.Execute() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("Shorten.Execute() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
