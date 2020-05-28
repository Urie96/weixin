package chatbot

import "github.com/Urie96/weixin/dao"

const JokeErrAnswer = "我现在记不起笑话了，待会儿可能会想起来"

func tellAJoke() string {
	answer, err := dao.GetRandomJokeContent()
	if err != nil {
		return JokeErrAnswer
	}
	return answer
}
