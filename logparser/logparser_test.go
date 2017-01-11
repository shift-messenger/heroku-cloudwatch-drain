package logparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseValidMessage(t *testing.T) {
	entry, err := Parse([]byte(`89 <45>1 2016-10-15T08:59:08.723822+00:00 host heroku web.1 - State changed from up to down`))
	assert.NoError(t, err)
	assert.Equal(t, "heroku[web.1]: State changed from up to down", entry.Message)
	assert.WithinDuration(t, time.Date(2016, 10, 15, 8, 59, 8, 723822000, time.UTC), entry.Time, time.Microsecond)
}

func TestParseHerokuRouterMessage(t *testing.T) {
	entry, err := Parse([]byte(`1119 <40>1 2012-11-30T06:45:26+00:00 host heroku router - at=info method=GET path="/api/v1/places/7285/groups/6133" host=shiftmessenger-api.herokuapp.com request_id=c8fa11d0-fdae-486d-a6cf-8adf2fdb8bb7 fwd="174.227.132.4" dyno=web.2 connect=1ms service=3375ms status=200 bytes=91021`))
	assert.NoError(t, err)
	assert.Equal(t, "heroku[router]: info GET \"/api/v1/places/7285/groups/6133\" \"shiftmessenger-api.herokuapp.com\" \"c8fa11d0-fdae-486d-a6cf-8adf2fdb8bb7\" \"174.227.132.4\" \"web.2\" 1 3375 200 91021", entry.Message)
}

func TestParseInvalidMessages(t *testing.T) {
	tests := []string{
		``,
		`89`,
		`89 <45>`,
		`89 <45>1`,
		`89 <45>1 2016-10-15T08:59:08.723822+00:00`,
		`89 <45>1 2016-10-15T08:59:08.723822+00:00 host`,
		`89 <45>1 2016-10-15T08:59:08.723822+00:00 host heroku`,
		`89 <45>1 2016-10-15T08:59:08.723822+00:00 host heroku web.1`,
		`89 <45>1 2016-10-15T08:59:08.723822+00:00 host heroku web.1 -`,
		`<45>1 2016-10-15T08:59:08.723822+00:00 host heroku web.1 - - State changed from up to down`,
	}

	for _, test := range tests {
		entry, err := Parse([]byte(test))
		assert.Error(t, err)
		assert.Nil(t, entry)
	}
}
