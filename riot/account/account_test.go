package account

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestAccountClient_GetByPUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Account
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Account{},
			doer: mock.NewJSONMockDoer(Account{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&Client{c: client}).GetByPUUID("")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestAccountClient_GetByRiotID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Account
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Account{},
			doer: mock.NewJSONMockDoer(Account{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&Client{c: client}).GetByRiotID("", "")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestAccountClient_GetMe(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Account
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Account{},
			doer: mock.NewJSONMockDoer(Account{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&Client{c: client}).GetMe("token")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestAccountClient_GetActiveShard(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *ActiveShard
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &ActiveShard{},
			doer: mock.NewJSONMockDoer(ActiveShard{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&Client{c: client}).GetActiveShard("val", "puuid")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}
