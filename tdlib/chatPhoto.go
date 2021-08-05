// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// ChatPhoto Describes a chat or user profile photo
type ChatPhoto struct {
	tdCommon
	ID            JSONInt64          `json:"id"`            // Unique photo identifier
	AddedDate     int32              `json:"added_date"`    // Point in time (Unix timestamp) when the photo has been added
	Minithumbnail *Minithumbnail     `json:"minithumbnail"` // Photo minithumbnail; may be null
	Sizes         []PhotoSize        `json:"sizes"`         // Available variants of the photo in JPEG format, in different size
	Animation     *AnimatedChatPhoto `json:"animation"`     // Animated variant of the photo in MPEG4 format; may be null
}

// MessageType return the string telegram-type of ChatPhoto
func (chatPhoto *ChatPhoto) MessageType() string {
	return "chatPhoto"
}

// NewChatPhoto creates a new ChatPhoto
//
// @param iD Unique photo identifier
// @param addedDate Point in time (Unix timestamp) when the photo has been added
// @param minithumbnail Photo minithumbnail; may be null
// @param sizes Available variants of the photo in JPEG format, in different size
// @param animation Animated variant of the photo in MPEG4 format; may be null
func NewChatPhoto(iD JSONInt64, addedDate int32, minithumbnail *Minithumbnail, sizes []PhotoSize, animation *AnimatedChatPhoto) *ChatPhoto {
	chatPhotoTemp := ChatPhoto{
		tdCommon:      tdCommon{Type: "chatPhoto"},
		ID:            iD,
		AddedDate:     addedDate,
		Minithumbnail: minithumbnail,
		Sizes:         sizes,
		Animation:     animation,
	}

	return &chatPhotoTemp
}