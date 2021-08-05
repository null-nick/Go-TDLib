// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// PushReceiverID Contains a globally unique push receiver identifier, which can be used to identify which account has received a push notification
type PushReceiverID struct {
	tdCommon
	ID JSONInt64 `json:"id"` // The globally unique identifier of push notification subscription
}

// MessageType return the string telegram-type of PushReceiverID
func (pushReceiverID *PushReceiverID) MessageType() string {
	return "pushReceiverId"
}

// NewPushReceiverID creates a new PushReceiverID
//
// @param iD The globally unique identifier of push notification subscription
func NewPushReceiverID(iD JSONInt64) *PushReceiverID {
	pushReceiverIDTemp := PushReceiverID{
		tdCommon: tdCommon{Type: "pushReceiverId"},
		ID:       iD,
	}

	return &pushReceiverIDTemp
}

// RegisterDevice Registers the currently used device for receiving push notifications. Returns a globally unique identifier of the push notification subscription
// @param deviceToken Device token
// @param otherUserIDs List of user identifiers of other users currently using the application
func (client *Client) RegisterDevice(deviceToken DeviceToken, otherUserIDs []int32) (*PushReceiverID, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "registerDevice",
		"device_token":   deviceToken,
		"other_user_ids": otherUserIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverID PushReceiverID
	err = json.Unmarshal(result.Raw, &pushReceiverID)
	return &pushReceiverID, err

}

// GetPushReceiverID Returns a globally unique push notification subscription identifier for identification of an account, which has received a push notification. Can be called synchronously
// @param payload JSON-encoded push notification payload
func (client *Client) GetPushReceiverID(payload string) (*PushReceiverID, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getPushReceiverId",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverID PushReceiverID
	err = json.Unmarshal(result.Raw, &pushReceiverID)
	return &pushReceiverID, err

}