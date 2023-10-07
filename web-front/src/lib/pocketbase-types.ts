/**
* This file was @generated using pocketbase-typegen
*/

export enum Collections {
	Attendee = "attendee",
	Contact = "contact",
	Event = "event",
	Invitation = "invitation",
	Response = "response",
	Tag = "tag",
	Users = "users",
}

// Alias types for improved usability
export type IsoDateString = string
export type RecordIdString = string
export type HTMLString = string

// System fields
export type BaseSystemFields<T = never> = {
	id: RecordIdString
	created: IsoDateString
	updated: IsoDateString
	collectionId: string
	collectionName: Collections
	expand?: T
}

export type AuthSystemFields<T = never> = {
	email: string
	emailVisibility: boolean
	username: string
	verified: boolean
} & BaseSystemFields<T>

// Record types for each collection

export enum AttendeeStatusOptions {
	"sending-invite" = "sending-invite",
	"invited" = "invited",
	"accepted" = "accepted",
	"declined" = "declined",
	"waitlist" = "waitlist",
	"uninvited" = "uninvited",
}
export type AttendeeRecord = {
	event: RecordIdString
	contact: RecordIdString
	status: AttendeeStatusOptions
	paid?: boolean
}

export enum ContactStatusOptions {
	"Pending" = "Pending",
	"Verified" = "Verified",
}
export type ContactRecord = {
	first_name: string
	last_name: string
	phone: string
	owner: RecordIdString
	status: ContactStatusOptions
}

export enum EventStatusOptions {
	"active" = "active",
	"completed" = "completed",
	"cancelled" = "cancelled",
}
export type EventRecord = {
	organizer: RecordIdString
	title: string
	description: string
	capacity: number
	start_date: IsoDateString
	end_date: IsoDateString
	send_invite_date: IsoDateString
	status: EventStatusOptions
}

export type InvitationRecord = {
	event: RecordIdString
	sent_date: IsoDateString
	title: string
	body: string
}

export type ResponseRecord = {
	sender: RecordIdString
	invitation: RecordIdString
	body: string
}

export type TagRecord = {
	owner: RecordIdString
	target: RecordIdString
	value: string
}

export type UsersRecord = {
	name?: string
	enabled?: boolean
}

// Response types include system fields and match responses from the PocketBase API
export type AttendeeResponse<Texpand = unknown> = AttendeeRecord & BaseSystemFields<Texpand>
export type ContactResponse<Texpand = unknown> = ContactRecord & BaseSystemFields<Texpand>
export type EventResponse<Texpand = unknown> = EventRecord & BaseSystemFields<Texpand>
export type InvitationResponse<Texpand = unknown> = InvitationRecord & BaseSystemFields<Texpand>
export type ResponseResponse<Texpand = unknown> = ResponseRecord & BaseSystemFields<Texpand>
export type TagResponse<Texpand = unknown> = TagRecord & BaseSystemFields<Texpand>
export type UsersResponse = UsersRecord & AuthSystemFields

// Types containing all Records and Responses, useful for creating typing helper functions

export type CollectionRecords = {
	attendee: AttendeeRecord
	contact: ContactRecord
	event: EventRecord
	invitation: InvitationRecord
	response: ResponseRecord
	tag: TagRecord
	users: UsersRecord
}

export type CollectionResponses = {
	attendee: AttendeeResponse
	contact: ContactResponse
	event: EventResponse
	invitation: InvitationResponse
	response: ResponseResponse
	tag: TagResponse
	users: UsersResponse
}