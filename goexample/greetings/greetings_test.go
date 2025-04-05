package greetings

import (
   "testing"
   "regexp"
)

func TestHelloName(t *testing.T) {
   name := "Jim"
   want := regexp.MustCompile(`\b`+name+`\b`)

   msg, err := Hello("Jim")
   if !want.MatchString(msg) || err != nil {
      t.Fatalf(`Hello("Jim") = %q, %v, want match for %#q, nil`, msg, err, want)
   }
}

