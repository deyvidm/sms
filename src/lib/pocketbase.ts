
import PocketBase, { ListResult } from 'pocketbase';
import type { ContactRecord, ContactResponse, EventRecord, InvitationRecord } from './pocketbase-types';
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


export async function get50Contacts() {
    return pb.collection('contact').getList<ContactResponse>(1, 50, {}).then((result)=>{
        contacts.set(result.items)
    });
}

export async function createEvent(r: EventRecord) {
    return pb.collection('event').create(r)
}

export async function createInvite(r: InvitationRecord) {
    return pb.collection('invitation').create(r)
}


export const API = {
    createContact: (r: ContactRecord) => {
        if (!pb.authStore.model?.id){
            Promise.reject("not logged in")
        }
        return pb.collection("contact").create(r)
    }
}
