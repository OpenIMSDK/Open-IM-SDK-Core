package open_im_sdk

import "open_im_sdk/open_im_sdk/utils"

func triggerCmdFriend() error {
	return nil

}

func triggerCmdBlackList() error {
	return nil

}

func triggerCmdFriendApplication() error {
	return nil

}

type deleteConNode struct {
	SourceID       string
	ConversationID string
	SessionType    int
}

func (u *UserRelated) triggerCmdDeleteConversationAndMessage(sourceID, conversationID string, sessionType int) error {
	c2v := cmd2Value{
		Cmd:   CmdDeleteConversation,
		Value: deleteConNode{SourceID: sourceID, ConversationID: conversationID, SessionType: sessionType},
	}

	return utils.sendCmd(u.ConversationCh, c2v, 1)
}

/*
func triggerCmdGetLoginUserInfo() error {
	c2v := cmd2Value{
		Cmd: CmdGeyLoginUserInfo,
	}
	return sendCmd(InitCh, c2v, 1)
}
*/

type updateConNode struct {
	ConId  string
	Action int //1 Delete the conversation; 2 Update the latest news in the conversation or add a conversation; 3 Put a conversation on the top;
	// 4 Cancel a conversation on the top, 5 Messages are not read and set to 0, 6 New conversations
	Args interface{}
}

func (u *UserRelated) triggerCmdNewMsgCome(msg ArrMsg) error {
	c2v := cmd2Value{
		Cmd:   CmdNewMsgCome,
		Value: msg,
	}
	utils.sdkLog("send cmd: ", u.ConversationCh)
	return utils.sendCmd(u.ConversationCh, c2v, 1)
}

func triggerCmdAcceptFriend(sendUid string) error {
	return nil

}

func triggerCmdRefuseFriend(receiveUid string) error {
	return nil
}

func (u *UserRelated) triggerCmdUpdateConversation(node updateConNode) error {
	c2v := cmd2Value{
		Cmd:   CmdUpdateConversation,
		Value: node,
	}

	return utils.sendCmd(u.ConversationCh, c2v, 1)
}

func (u *UserRelated) unInitAll() {
	c2v := cmd2Value{Cmd: CmdUnInit}
	_ = utils.sendCmd(u.ConversationCh, c2v, 1)
}

type goroutine interface {
	work(cmd cmd2Value)
	getCh() chan cmd2Value
}

func doListener(Li goroutine) {
	utils.sdkLog("doListener start.", Li.getCh())
	for {
		utils.sdkLog("doListener for.")
		select {
		case cmd := <-Li.getCh():
			if cmd.Cmd == CmdUnInit {
				utils.sdkLog("doListener goroutine.")
				return
			}
			utils.sdkLog("doListener work.")
			Li.work(cmd)
		}
	}
}
