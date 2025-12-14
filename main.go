package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	env "github.com/joho/godotenv"
)

func main() {
	wd, gErr := os.Getwd()

	if gErr != nil {
		log.Fatalln(gErr.Error())
	}

	envPath := strings.Join([]string{wd, "/", ".run.env"}, "")
	log.Println("main.go: envPath: " + envPath)

	hasError := env.Load(envPath)
	if hasError != nil {
		log.Fatalln("main.go: can't load secrets correctly", hasError.Error())
		return
	}
	log.Println("main.go: env loaded")

	prefix := "Bot "
	discord, err := discordgo.New(prefix + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println(err.Error())
	}

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		memberIds := strings.Split(os.Getenv("MEMBER_IDS"), ",")
		var remainingMember []string
		for _, v := range memberIds {
			if v != i.Member.User.ID {
				remainingMember = append(remainingMember, fmt.Sprintf("<@%s>", v))
			}
		}
		membersToMention := strings.Join(remainingMember, ",")

		switch i.ApplicationCommandData().Name {
		case "작업시작":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(`<@%s>가 작업을 시작했습니다. %s`, i.Member.User.ID, membersToMention),
				},
			})

		case "작업종료":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(`<@%s>가 작업을 종료했습니다. %s`, i.Member.User.ID, membersToMention),
				},
			})
		}
	})

	oErr := discord.Open()
	if oErr != nil {
		fmt.Println(oErr.Error())
	}
	if oErr != nil {
		fmt.Println(oErr.Error())
	}
	defer discord.Close()

	appID := discord.State.User.ID
	guildID := os.Getenv("SERVER_ID")
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "작업시작",
			Description: "작업 시작 시간 기록",
		},
		{
			Name:        "작업종료",
			Description: "작업 종료 시간 기록",
		},
	}

	for _, cmd := range commands {
		_, err := discord.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			fmt.Println("명령어 등록 실패:", err)
		}
	}

	fmt.Println("bot running")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("bot stopped")
}
