package timekit

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestIsNowBetweenTimes(t *testing.T) {
	const start = "00:00:00"
	const end = "23:59:00"
	const startEndMsg = "Start: (" + start + ") and End: (" + end + ")"

	err := IsNowBetweenTimes(start, end)
	if err != nil {
		t.Error(errors.Wrap(err, startEndMsg))
		t.FailNow()
	}

	t.Log("Now (" + NowToHM() + ") is between times. " + startEndMsg)
}

func TestIsBetweenTimesBool(t *testing.T) {
	const start = "00:00:00"
	const end = "23:59:00"
	const startEndMsg = "Start: (" + start + ") and End: (" + end + ")"

	boolean := IsBetweenTimesBool(start, end)
	if boolean == false {
		t.Error(errors.New("start and end are not between times when they should've: " + startEndMsg))
		t.FailNow()
	}

	t.Log("Now (" + NowToHM() + ") is between times. " + startEndMsg)
}

// TestTimeToStr expects that a sub 10 int, for example: "1", will return "01"
func TestTimeToStr(t *testing.T) {
	require.Equal(t, "01", TimeToStr(1))
	require.Equal(t, "09", TimeToStr(9))

	require.NotEqual(t, "010", TimeToStr(10))
}

func TestCurrentDateAddTime(t *testing.T) {
	const oneDay = 24 * time.Hour
	expected := time.Now().UTC().Add(oneDay).Format("2006-01-02")
	actual := CurrentDateAddTime(oneDay, true)
	require.Equal(t, expected, actual)

	t.Log("expected (" + expected + "). Actual (" + actual + ").")
}
