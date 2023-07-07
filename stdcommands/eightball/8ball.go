// using code from https://github.com/xbilldozer/8ball

package eightball

import (
	"fmt"
	"math/rand"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)

var magicAnswers = [...]string{
	// Standard set of Magic 8Ball answers.
	// See https://en.wikipedia.org/wiki/Magic_8-Ball#Possible_answers

	// positive outcomes
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes, definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",

	// Neutral outcomes
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	// extra neutral outcomes
	"That depends on what that means",
	"wtf are you talking about?",
	"that's what she said",

	//Negative outcomes
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

var ShakeFailureMessage = "You can't just shake me bro, ask a question!"

var Command = &commands.YAGCommand{
	CmdCategory:         commands.CategoryFun,
	Name:                "8ball",
	Description:         "Ask the magic 8ball a question",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	Arguments: []*dcmd.ArgDef{
		{Name: "Question", Type: dcmd.String},
	},

	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		if data.SlashCommandTriggerData != nil {
			return ":question: : " + data.Args[0].Str() + "\n" + Shake(data), nil
		}
		return Shake(data), nil
	},
}

func Shake(data *dcmd.Data) string {
	if data.Args[0].Str() == "" {
		return ShakeFailureMessage
	}

	return fmt.Sprintf(":8ball: : %s",
		magicAnswers[rand.Intn(len(magicAnswers))])
}
