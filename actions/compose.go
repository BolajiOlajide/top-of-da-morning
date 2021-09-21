package actions

import "fmt"

// ComposeTweetText returns the default text for the tweet to be tweeted 😹
func ComposeTweetText(owner string) string {
	text := fmt.Sprintf(`Top of Da Morning
Top of Da Morning
Top of Da Morning

Let's get this shit.

#TopOfDaMorning
Credit: Tiktok User (%s)`, owner)

	return text
}
