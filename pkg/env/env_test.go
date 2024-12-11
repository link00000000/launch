package env

import (
	"strings"
	"testing"
)

func assertEq(t *testing.T, name, actual, expected string) {
	if expected != actual {
		t.Fatalf("assertion failed (%s == %s): expected \"%s\", got \"%s\"", name, expected, expected, actual)
	}
}

func TestReadFile(t *testing.T) {
	f := `ONE=1
TWO=2     
THREE=3 # comment
FOUR=4#comment
#comment


# comment
  # comment
  FIVE=5
SIX="6"
SEVEN="#7"
EIGHT=\\8
NINE=\"9
TEN="\'10"
ELEVEN="\"11"
TWELVE=this is twelve
THIRTEEN="this is thirteen"
FOURTEEN='this is fourteen'
FIFTEEN=" this is fifteen "
SIXTEEN= 16 
SEVENTEEN = 17
EIGHTEEN=\ 18
NINETEEN=19\ 
	TWENTY=20`

	r := strings.NewReader(f)
	env, err := Read(r)

	if err != nil {
		t.Fatalf("error while executing Read(r): %v", err)
	}

	if len(env) != 20 {
		t.Fatalf("incorrect map size: expected %d, got %d", 20, len(env))
	}

	assertEq(t, "env[\"ONE\"]", env["ONE"], "1")
	assertEq(t, "env[\"TWO\"]", env["TWO"], "2")
	assertEq(t, "env[\"THREE\"]", env["THREE"], "3")
	assertEq(t, "env[\"FOUR\"]", env["FOUR"], "4")
	assertEq(t, "env[\"FIVE\"]", env["FIVE"], "5")
	assertEq(t, "env[\"SIX\"]", env["SIX"], "6")
	assertEq(t, "env[\"SEVEN\"]", env["SEVEN"], "#7")
	assertEq(t, "env[\"EIGHT\"]", env["EIGHT"], "\\8")
	assertEq(t, "env[\"NINE\"]", env["NINE"], "\"9")
	assertEq(t, "env[\"TEN\"]", env["TEN"], "'10")
	assertEq(t, "env[\"ELEVEN\"]", env["ELEVEN"], "\"11")
	assertEq(t, "env[\"TWELVE\"]", env["TWELVE"], "this is twelve")
	assertEq(t, "env[\"THIRTEEN\"]", env["THIRTEEN"], "this is thirteen")
	assertEq(t, "env[\"FOURTEEN\"]", env["FOURTEEN"], "this is fourteen")
	assertEq(t, "env[\"FIFTEEN\"]", env["FIFTEEN"], " this is fifteen ")
	assertEq(t, "env[\"SIXTEEN\"]", env["SIXTEEN"], "16")
	assertEq(t, "env[\"SEVENTEEN\"]", env["SEVENTEEN"], "17")
	assertEq(t, "env[\"EIGHTEEN\"]", env["EIGHTEEN"], " 18")
	assertEq(t, "env[\"NINETEEN\"]", env["NINETEEN"], "19 ")
	assertEq(t, "env[\"TWENTY\"]", env["TWENTY"], "20")
}
