package httpserver

import "testing"

func TestHttpServer_AddBalance(t *testing.T) {
	// s := New(NewConfig())
	// testCases := []struct {
	// 	name         string
	// 	user         interface{}
	// 	expectedCode interface{}
	// }{
	// 	{
	// 		name: "invalid",
	// 		user: map[string]int{
	// 			"Id":      5,
	// 			"Balance": 15,
	// 		},
	// 		expectedCode: http.StatusAccepted,
	// 	},
	// }

	// for _, tc := range testCases {
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		rec := httptest.NewRecorder()
	// 		b := &bytes.Buffer{}
	// 		json.NewEncoder(b).Encode(tc.user)
	// 		req, _ := http.NewRequest(http.MethodPost, "/getBalance", b)
	// 		s.router.ServeHTTP(rec, req)
	// 		assert.Equal(t, tc.expectedCode, rec.Code)
	// 	})
	// }
}
