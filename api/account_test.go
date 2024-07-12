package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_sqlc "github.com/PyMarcus/go_sqlc/db/mock"
	db "github.com/PyMarcus/go_sqlc/db/sqlc"
	"github.com/PyMarcus/go_sqlc/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	acc := randomAccount()

	testCases := []struct {
		name          string
		accID         int64
		buildStubs    func(store *mock_sqlc.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		// case 1
		{
			name:  "OK",
			accID: acc.ID,
			buildStubs: func(store *mock_sqlc.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(acc.ID)).Times(1).Return(acc, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},

		// case 2
		{
			name:  "Not Found",
			accID: acc.ID,
			buildStubs: func(store *mock_sqlc.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(acc.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_sqlc.NewMockStore(ctrl)
			tc.buildStubs(store)
			config := util.Config{}
			server, err := NewServer(config, store)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", acc.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})

	}
}

func randomAccount() db.Account {
	return db.Account{
		ID:        util.RandomInt(1, 100),
		Owner:     util.RandomOwner(),
		Balance:   util.RandomMoney(),
		Currency:  util.RandomCurrency(),
		CreatedAt: time.Now(),
	}
}
