package service

import (
	"im/define"
	"im/helper"
	"im/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// userID--wb.conn
var wc = make(map[string]*websocket.Conn)

func WebsocketMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}

	defer conn.Close()
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	wc[uc.Identity] = conn
	for {
		msg := new(define.MessageStruct)
		if err := conn.ReadJSON(msg); err != nil {
			log.Println("Read Error:", err)
			return
		}
		// 判断用户是否属于消息体的房间
		if _, err := models.GetUserRoomByUserIdRoomId(uc.Identity, msg.RoomIdentity); err != nil {
			log.Println("UserIdentity:", uc.Identity, "Not Exit RoomIdentity:", msg.RoomIdentity)
			return
		}

		// 保存消息
		mb := &models.MessageBasic{
			UserIdentity: uc.Identity,
			RoomIdentity: msg.RoomIdentity,
			Data:         msg.Message,
			CreatedAt:    time.Now().Unix(),
			UpdatedAt:    time.Now().Unix(),
		}
		err := models.InsertOneMessageBasic(mb)
		if err != nil {
			log.Println("[DB ERROR]:", err)
			return
		}

		// 获取在特定房间的在线用户，并发送消息
		userRooms, err := models.GetUserRoomByRoomId(msg.RoomIdentity)
		if err != nil {
			log.Println("[DB ERROR]:", err)
			return
		}
		for _, room := range userRooms {
			if cc, ok := wc[room.UserIdentity]; ok {
				if err := cc.WriteMessage(websocket.TextMessage, []byte(msg.Message)); err != nil {
					log.Println("Write Message Error:", err)
					return
				}
			}
		}

	}
}
