package box

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/wavynote"
)

type BoxHandler struct {
	dbInfo wavynote.DataBaseInfo
}

func NewBoxHandler(dbInfo wavynote.DataBaseInfo) *BoxHandler {
	h := &BoxHandler{
		dbInfo: dbInfo,
	}
	return h
}

type testJSON struct {
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

// ShowConversation godoc
// @Summary      노트를 주고 받는 대화방 목록 조회(최대 3개)
// @Description  노트를 주고 받는 대화방 목록 조회(최대 3개)
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        id   query     string  false  "user id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/conversationlist [get]
func (h *BoxHandler) ShowConversation(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	// var reqInfo *testJSON
	// err = c.BindJSON(&reqInfo)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(reqInfo.Name)
	// fmt.Println(reqInfo.Emoji)
	// fmt.Printf("%#v", []byte(reqInfo.Emoji))

	userId := c.Query("id")
	fmt.Println("user_id:", userId)

	conversationList := []restapi.ConversationInfo{}
	conversationList = append(conversationList, restapi.ConversationInfo{
		ConverstaionId: "1afc571d-61bf-4cef-95ce-ab791f999297",
		OppNickName:    "누군가",
		NoteCount:      20,
	})
	conversationList = append(conversationList, restapi.ConversationInfo{
		ConverstaionId: "e1ce587c-c6c0-46fc-b59d-fd8316a4502a",
		OppNickName:    "또다른누군가",
		NoteCount:      2,
	})

	c.IndentedJSON(
		http.StatusOK,
		restapi.ConversationListResponse{
			Conversations: conversationList,
		},
	)
}

// ShowConversationNoteList godoc
// @Summary      특정 친구와 주고받은 노트 목록
// @Description  특정 친구와 주고받은 노트 목록
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        id  query     string  false  "conversation id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/notelist [get]
func (h *BoxHandler) ShowConversationNoteList(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	conversationId := c.Query("id")
	fmt.Println("id:", conversationId)

	conversationNoteList := []restapi.ConversationNoteInfo{}
	conversationNoteList = append(conversationNoteList, restapi.ConversationNoteInfo{
		NoteId:  "1afc571d-61bf-4cef-95ce-ab791f999297",
		FromId:  "somebody@naver.com",
		ToId:    "wavynoteadmin@gmail.com",
		Title:   "일상공유해요",
		Preview: "일상공유에 대한 본문내용입니다.",
		SendAt:  "2023-10-29 21:00:48",
	})
	conversationNoteList = append(conversationNoteList, restapi.ConversationNoteInfo{
		NoteId:  "09d05df1-2958-4a3d-b910-3b4fb079327b",
		FromId:  "wavynoteadmin@gmail.com",
		ToId:    "somebody@naver.com",
		Title:   "나의첫번째노트",
		Preview: "나의 첫 웨이비노트 본문 내용입니다.",
		SendAt:  "2023-11-01 23:20:12",
	})

	c.IndentedJSON(
		http.StatusOK,
		restapi.ConverstaionNoteListResponse{
			ConversationNotes: conversationNoteList,
		},
	)
}

// ShowConversationNote godoc
// @Summary      특정 노트 조회
// @Description  특정 노트 조회
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        id  query     string  false  "note id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/show [get]
func (h *BoxHandler) ShowConversationNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	noteId := c.Query("id")
	fmt.Println("id:", noteId)

	c.IndentedJSON(
		http.StatusOK,
		restapi.NoteInfo{
			NoteId:   "09d05df1-2958-4a3d-b910-3b4fb079327b",
			FolderId: "980e71ba-0395-49aa-833e-3ebc76b3ec88",
			FromId:   "wavynoteadmin@gmail.com",
			ToId:     "somebody@naver.com",
			SaveAt:   "2023-11-01 21:00:00",
			SendAt:   "2023-11-01 23:20:12",
			Title:    "나의첫번째노트",
			Content:  "나의 첫 웨이비노트 본문 내용입니다.",
			Keywords: []string{"일상∙생각", "마음챙김"},
		},
	)
}

// DeleteConversation godoc
// @Summary      특정 대화방 삭제
// @Description  특정 대화방 삭제
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        id  query     string  false  "conversation id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/conversation [delete]
func (h *BoxHandler) DeleteConversation(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	conversationId := c.Query("id")
	fmt.Println("id:", conversationId)

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "특정 대화방 삭제",
		},
	)
}
