// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/gogap/wechat for the canonical source repository
// @license     https://github.com/gogap/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"github.com/gogap/wechat/mp"
)

const (
	// 推送到公众号URL上的事件类型
	EventTypeCardPassCheck    = "card_pass_check"     // 卡券通过审核
	EventTypeCardNotPassCheck = "card_not_pass_check" // 卡券未通过审核
	EventTypeUserGetCard      = "user_get_card"       // 领取卡券事件
	EventTypeUserDelCard      = "user_del_card"       // 删除卡券事件
	EventTypeUserViewCard     = "user_view_card"      // 进入会员卡事件推送
	EventTypeUserConsumeCard  = "user_consume_card"   // 核销事件推送
)

// 卡券通过审核，微信会把这个事件推送到开发者填写的URL
type CardPassCheckEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event  string `xml:"Event"  json:"Event"`  // 事件类型, card_pass_check
	CardId string `xml:"CardId" json:"CardId"` // 卡券ID
}

func GetCardPassCheckEvent(msg *mp.MixedMessage) *CardPassCheckEvent {
	return &CardPassCheckEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		CardId:              msg.CardId,
	}
}

// 卡券未通过审核，微信会把这个事件推送到开发者填写的URL
type CardNotPassCheckEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event  string `xml:"Event"  json:"Event"`  // 事件类型, card_not_pass_check
	CardId string `xml:"CardId" json:"CardId"` // 卡券ID
}

func GetCardNotPassCheckEvent(msg *mp.MixedMessage) *CardNotPassCheckEvent {
	return &CardNotPassCheckEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		CardId:              msg.CardId,
	}
}

// 用户在领取卡券时，微信会把这个事件推送到开发者填写的URL。
type UserGetCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event          string `xml:"Event"          json:"Event"`          // 事件类型, user_get_card
	CardId         string `xml:"CardId"         json:"CardId"`         // 卡券ID
	IsGiveByFriend int    `xml:"IsGiveByFriend" json:"IsGiveByFriend"` // 是否为转赠，1 代表是，0 代表否。
	FriendUserName string `xml:"FriendUserName" json:"FriendUserName"` // 赠送方账号（一个OpenID），"IsGiveByFriend”为1 时填写该参数。
	UserCardCode   string `xml:"UserCardCode"   json:"UserCardCode"`   // code 序列号。自定义code 及非自定义code的卡券被领取后都支持事件推送。
	OuterId        int64  `xml:"OuterId"        json:"OuterId"`        // 领取场景值，用于领取渠道数据统计。可在生成二维码接口及添加JS API 接口中自定义该字段的整型值。
}

func GetUserGetCardEvent(msg *mp.MixedMessage) *UserGetCardEvent {
	return &UserGetCardEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		CardId:              msg.CardId,
		IsGiveByFriend:      msg.IsGiveByFriend,
		FriendUserName:      msg.FriendUserName,
		UserCardCode:        msg.UserCardCode,
		OuterId:             msg.OuterId,
	}
}

// 用户在删除卡券时，微信会把这个事件推送到开发者填写的URL。
type UserDelCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event        string `xml:"Event"          json:"Event"`        // 事件类型, user_del_card
	CardId       string `xml:"CardId"         json:"CardId"`       // 卡券ID
	UserCardCode string `xml:"UserCardCode"   json:"UserCardCode"` // 商户自定义code 值。非自定code 推送为空串
}

func GetUserDelCardEvent(msg *mp.MixedMessage) *UserDelCardEvent {
	return &UserDelCardEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		CardId:              msg.CardId,
		UserCardCode:        msg.UserCardCode,
	}
}

// 用户在进入会员卡时，微信会把这个事件推送到开发者填写的URL
type UserViewCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event        string `xml:"Event"          json:"Event"`        // 事件类型, user_view_card
	CardId       string `xml:"CardId"         json:"CardId"`       // 卡券ID
	UserCardCode string `xml:"UserCardCode"   json:"UserCardCode"` // 商户自定义code 值。非自定code 推送为空串
}

func GetUserViewCardEvent(msg *mp.MixedMessage) *UserViewCardEvent {
	return &UserViewCardEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		CardId:              msg.CardId,
		UserCardCode:        msg.UserCardCode,
	}
}

// 卡券被核销时，微信会把这个事件推送到开发者填写的URL
type UserConsumeCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.CommonMessageHeader

	Event        string `xml:"Event"          json:"Event"`        // 事件类型, user_consume_card
	CardId       string `xml:"CardId"         json:"CardId"`       // 卡券ID
	UserCardCode string `xml:"UserCardCode"   json:"UserCardCode"` // 商户自定义code 值。非自定code 推送为空串
}

func GetUserConsumeCardEvent(msg *mp.MixedMessage) *UserConsumeCardEvent {
	return &UserConsumeCardEvent{
		CommonMessageHeader: msg.CommonMessageHeader,
		Event:               msg.Event,
		CardId:              msg.CardId,
		UserCardCode:        msg.UserCardCode,
	}
}
