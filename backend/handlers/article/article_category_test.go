package article

import (
	"testing"
	"net/http/httptest"
)

func Test(t *testing.T) {

	req := httptest.NewRequest("GET", "localhost:3000/api/v1/article/get-all-categories", nil)
}