package main

import (
	"fmt"
)

type Streamer interface {
	startStream()
	notifyAll()
	subscribe(user User)
	unsubscribe(user User)
}

type User interface {
	update(string)
	getUsername() string
}

type TwitchStreamer struct {
	subscribers []User
	nickname    string
}

type TwitchUser struct {
	username string
}

func (streamer *TwitchStreamer) startStream() {
	fmt.Printf("%s, stream has been started!\n", streamer.nickname)
	streamer.notifyAll()
}

func (streamer *TwitchStreamer) notifyAll() {
	for _, user := range streamer.subscribers {
		user.update(streamer.nickname)
	}
}

func (streamer *TwitchStreamer) subscribe(user User) {
	streamer.subscribers = append(streamer.subscribers, user)
}

func (user *TwitchUser) getUsername() string {
	return user.username
}

func (streamer *TwitchStreamer) unsubscribe(userToDelete User) {
	subscribersAmount := len(streamer.subscribers)
	for i, user := range streamer.subscribers {
		if user.getUsername() == userToDelete.getUsername() {
			streamer.subscribers[subscribersAmount-1], streamer.subscribers[i] =
				streamer.subscribers[i], streamer.subscribers[subscribersAmount-1]

			streamer.subscribers = streamer.subscribers[:subscribersAmount-1]
			break
		}
	}
}

func (user *TwitchUser) update(streamerNickname string) {
	fmt.Printf("%s, %s started streaming, don't miss it!\n", user.username, streamerNickname)
}

func main() {
	Melharucos := TwitchStreamer{nickname: "melharucos"}
	Melharucos.startStream()

	fmt.Println()

	user1 := TwitchUser{"Zeus"}
	user2 := TwitchUser{"Hades"}
	user3 := TwitchUser{"Athena"}
	user4 := TwitchUser{"Cerberus"}

	Melharucos.subscribe(&user1)
	Melharucos.subscribe(&user2)
	Melharucos.subscribe(&user3)

	fmt.Println()
	Melharucos.startStream()
	fmt.Println()

	Melharucos.subscribe(&user4)

	fmt.Println()
	Melharucos.startStream()
}
