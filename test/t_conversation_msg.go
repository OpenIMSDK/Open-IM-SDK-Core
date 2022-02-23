package test

import (
	"encoding/json"
	"fmt"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/sdk_params_callback"
	"open_im_sdk/pkg/server_api_params"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"time"
)

//func DotestSetConversationRecvMessageOpt() {
//	var callback BaseSuccFailedTest
//	callback.funcName = utils.GetSelfFuncName()
//	var idList []string
//	idList = append(idList, "18567155635")
//	jsontest, _ := json.Marshal(idList)
//	open_im_sdk.SetConversationRecvMessageOpt(&callback, string(jsontest), 2)
//	fmt.Println("SetConversationRecvMessageOpt", string(jsontest))
//}
//
//func DoTestGetMultipleConversation() {
//	var callback BaseSuccFailedTest
//	callback.funcName = utils.GetSelfFuncName()
//	var idList []string
//	fmt.Println("DoTestGetMultipleConversation come here")
//	idList = append(idList, "single_13977954313", "group_77215e1b13b75f3ab00cb6570e3d9618")
//	jsontest, _ := json.Marshal(idList)
//	open_im_sdk.GetMultipleConversation(string(jsontest), &callback)
//	fmt.Println("GetMultipleConversation", string(jsontest))
//}
//
//func DoTestGetConversationRecvMessageOpt() {
//	var callback BaseSuccFailedTest
//	callback.funcName = utils.GetSelfFuncName()
//	var idList []string
//	idList = append(idList, "18567155635")
//	jsontest, _ := json.Marshal(idList)
//	open_im_sdk.GetConversationRecvMessageOpt(&callback, string(jsontest))
//	fmt.Println("GetConversationRecvMessageOpt", string(jsontest))
//}

func DoTestGetHistoryMessage(userID string) {
	var testGetHistoryCallBack GetHistoryCallBack
	testGetHistoryCallBack.OperationID = utils.OperationIDGenerator()
	var params sdk_params_callback.GetHistoryMessageListParams
	params.UserID = userID
	params.Count = 10
	open_im_sdk.GetHistoryMessageList(testGetHistoryCallBack, testGetHistoryCallBack.OperationID, utils.StructToJsonString(params))
}

//func DoTestDeleteConversation(conversationID string) {
//	var testDeleteConversation DeleteConversationCallBack
//	open_im_sdk.DeleteConversation(conversationID, testDeleteConversation)
//
//}

type DeleteConversationCallBack struct {
}

func (d DeleteConversationCallBack) OnError(errCode int32, errMsg string) {
	fmt.Printf("DeleteConversationCallBack , errCode:%v,errMsg:%v\n", errCode, errMsg)
}

func (d DeleteConversationCallBack) OnSuccess(data string) {
	fmt.Printf("DeleteConversationCallBack , success,data:%v\n", data)
}

type DeleteMessageFromLocalStorageCallBack struct {
}

func (d DeleteMessageFromLocalStorageCallBack) OnError(errCode int32, errMsg string) {
	fmt.Printf("DeleteMessageFromLocalStorageCallBack , errCode:%v,errMsg:%v\n", errCode, errMsg)
}

func (d DeleteMessageFromLocalStorageCallBack) OnSuccess(data string) {
	fmt.Printf("DeleteMessageFromLocalStorageCallBack , success,data:%v\n", data)
}

type TestGetAllConversationListCallBack struct {
	OperationID string
}

func (t TestGetAllConversationListCallBack) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestGetAllConversationListCallBack ", errCode, errMsg)
}

func (t TestGetAllConversationListCallBack) OnSuccess(data string) {
	log.Info(t.OperationID, "TestGetAllConversationListCallBack ", data)
}
func DoTestGetAllConversation() {
	var test TestGetAllConversationListCallBack
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.GetAllConversationList(test, test.OperationID)

}

type TestGetConversationListSplitCallBack struct {
	OperationID string
}

func (t TestGetConversationListSplitCallBack) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestGetConversationListSplitCallBack err ", errCode, errMsg)
}

func (t TestGetConversationListSplitCallBack) OnSuccess(data string) {
	log.Info(t.OperationID, "TestGetConversationListSplitCallBack  success", data)
}
func DoTestGetConversationListSplit() {
	var test TestGetConversationListSplitCallBack
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.GetConversationListSplit(test, test.OperationID, 1, 2)

}

type TestGetOneConversationCallBack struct {
}

func (t TestGetOneConversationCallBack) OnError(errCode int32, errMsg string) {
	fmt.Printf("TestGetOneConversationCallBack , errCode:%v,errMsg:%v\n", errCode, errMsg)
}

func (t TestGetOneConversationCallBack) OnSuccess(data string) {
	fmt.Printf("TestGetOneConversationCallBack , success,data:%v\n", data)
}

//func DoTestGetOneConversation(sourceID string, sessionType int) {
//	var test TestGetOneConversationCallBack
//	//GetOneConversation(Friend_uid, SingleChatType, test)
//	open_im_sdk.GetOneConversation(sourceID, sessionType, test)
//
//}
func DoTestCreateTextMessage(text string) string {
	operationID := utils.OperationIDGenerator()
	return open_im_sdk.CreateTextMessage(operationID, text)
}

func DoTestCreateImageMessageFromFullPath() string {
	operationID := utils.OperationIDGenerator()
	return open_im_sdk.CreateImageMessageFromFullPath(operationID, "C:\\1.jpg")
	//open_im_sdk.SendMessage(&testSendMsg, operationID, s, , "", utils.StructToJsonString(o))
}

//func DoTestSetConversationDraft() {
//	var test TestSetConversationDraft
//	open_im_sdk.SetConversationDraft("single_c93bc8b171cce7b9d1befb389abfe52f", "hah", test)
//
//}
type TestSetConversationDraft struct {
}

func (t TestSetConversationDraft) OnError(errCode int32, errMsg string) {
	fmt.Printf("SetConversationDraft , OnError %v\n", errMsg)
}

func (t TestSetConversationDraft) OnSuccess(data string) {
	fmt.Printf("SetConversationDraft , OnSuccess %v\n", data)
}

type GetHistoryCallBack struct {
	OperationID string
}

func (g GetHistoryCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "GetHistoryCallBack err", errCode, errMsg)
}

func (g GetHistoryCallBack) OnSuccess(data string) {
	log.Info(g.OperationID, "get History success ", data)
}

type MsgListenerCallBak struct {
}

func (m MsgListenerCallBak) OnRecvNewMessage(msg string) {
	var mm sdk_struct.MsgStruct
	err := json.Unmarshal([]byte(msg), &mm)
	if err != nil {
		log.Error("", "Unmarshal failed", err.Error())
	} else {
		log.Info("", "recv time: ", time.Now().UnixNano(), "send_time: ", mm.SendTime, " client_msg_id: ", mm.ClientMsgID, "server_msg_id", mm.ServerMsgID)
	}

}

type TestSearchLocalMessages struct {
	OperationID string
}

func (t TestSearchLocalMessages) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "SearchLocalMessages , OnError %v\n", errMsg)
}

func (t TestSearchLocalMessages) OnSuccess(data string) {
	log.Info(t.OperationID, "SearchLocalMessages , OnSuccess %v\n", data)
}
func DoTestSearchLocalMessages() {
	var t TestSearchLocalMessages
	operationID := utils.OperationIDGenerator()
	t.OperationID = operationID
	var p sdk_params_callback.SearchLocalMessagesParams
	//p.SessionType = constant.SingleChatType
	p.SourceID = "18090680773"
	p.KeywordList = []string{}
	p.SearchTimePeriod = 24 * 60 * 60 * 10
	open_im_sdk.SearchLocalMessages(t, operationID, utils.StructToJsonString(p))
}

type TestDeleteConversation struct {
	OperationID string
}

func (t TestDeleteConversation) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestDeleteConversation , OnError %v\n", errMsg)
}

func (t TestDeleteConversation) OnSuccess(data string) {
	log.Info(t.OperationID, "TestDeleteConversation , OnSuccess %v\n", data)
}
func DoTestDeleteConversation() {
	var t TestDeleteConversation
	operationID := utils.OperationIDGenerator()
	t.OperationID = operationID
	conversationID := "single_17396220460"
	open_im_sdk.DeleteConversation(t, operationID, conversationID)
}
func (m MsgListenerCallBak) OnRecvC2CReadReceipt(data string) {
	fmt.Println("OnRecvC2CReadReceipt , ", data)
}

func (m MsgListenerCallBak) OnRecvMessageRevoked(msgId string) {
	fmt.Println("OnRecvMessageRevoked ", msgId)
}

type conversationCallBack struct {
}

func (c conversationCallBack) OnSyncServerStart() {
	panic("implement me")
}

func (c conversationCallBack) OnSyncServerFinish() {
	panic("implement me")
}

func (c conversationCallBack) OnSyncServerFailed() {
	panic("implement me")
}

func (c conversationCallBack) OnNewConversation(conversationList string) {
	log.Info("", "OnNewConversation returnList is ", conversationList)
}

func (c conversationCallBack) OnConversationChanged(conversationList string) {
	log.Info("", "OnConversationChanged returnList is", conversationList)
}

func (c conversationCallBack) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	log.Info("", "OnTotalUnreadMessageCountChanged returnTotalUnreadCount is ", totalUnreadCount)
}

type testMarkC2CMessageAsRead struct {
}

func (testMarkC2CMessageAsRead) OnSuccess(data string) {
	fmt.Println(" testMarkC2CMessageAsRead  OnSuccess", data)
}

func (testMarkC2CMessageAsRead) OnError(code int32, msg string) {
	fmt.Println("testMarkC2CMessageAsRead, OnError", code, msg)
}

//func DoTestMarkC2CMessageAsRead() {
//	var test testMarkC2CMessageAsRead
//	readid := "2021-06-23 12:25:36-7eefe8fc74afd7c6adae6d0bc76929e90074d5bc-8522589345510912161"
//	var xlist []string
//	xlist = append(xlist, readid)
//	jsonid, _ := json.Marshal(xlist)
//	open_im_sdk.MarkC2CMessageAsRead(test, Friend_uid, string(jsonid))
//}

func DoTestSendMsg(sendId, recvID string) {
	m := "mmmmmmmmtest:Gordon->sk" + sendId + ":" + recvID + ":"
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateTextMessage(m)
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "121313"
	o.Desc = "45464"
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, recvID, "", utils.StructToJsonString(o))
}

func DoTestSendImageMsg(sendId, recvID string) {
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateImageMessageFromFullPath()
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "121313"
	o.Desc = "45464"
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, recvID, "", utils.StructToJsonString(o))
}
