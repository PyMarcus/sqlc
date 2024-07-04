package api

import (
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_sqlc.NewMockStore(ctrl)

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(acc.ID)).Times(1).Return(acc, nil)

	server := NewServer(store)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", acc.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
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
