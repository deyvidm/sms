
import PocketBase, { ListResult, type RecordFullListQueryParams } from 'pocketbase';
import type { AttendeeRecord, AttendeeResponse, AttendeeStatusOptions, ContactRecord, ContactResponse, EventRecord, EventResponse, InvitationRecord } from './pocketbase-types';
import { writable } from 'svelte/store';
import { error } from '@sveltejs/kit';

// export const pb = new PocketBase('http://YOUR-SERVER-IP-OR-URL'); // remote
export const pb = new PocketBase('http://127.0.0.1:8090'); // local

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange((auth) => {
    console.log('authStore changed', auth);
    currentUser.set(pb.authStore.model);
});

export const contacts = writable(new Array<ContactResponse>());
export const events = writable(new Array<EventResponse>());
export const attendees = writable(new Array<AttendeeResponse>());


export async function get50Contacts() {
    return pb.collection('contact').getList<ContactResponse>(1, 50, {})
}

export async function get50Events() {
    return pb.collection('event').getList<EventResponse>(1, 50, {}).then((result) => {
        events.set(result.items)
    });
}

export async function get50Attendees(event: EventResponse) {
    return pb.collection('attendee').getList<AttendeeResponse>(1, 50, {expand: "contact", filter:`event = "${event.id}"`})
}

export async function createEvent(r: EventRecord) {
    return pb.collection('event').create(r)
}

export async function createInvite(r: InvitationRecord) {
    return pb.collection('invitation').create(r)
}

export async function updateAttendee(r: AttendeeResponse){
    return pb.collection('attendee').update(r.id,{paid:r.paid})
}

export const API = {
    createContact: (r: ContactRecord) => {
        if (!pb.authStore.model?.id) {
            Promise.reject("not logged in")
        }
        return pb.collection("contact").create(r)
    }
}
