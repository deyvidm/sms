
import PocketBase, { ListResult } from 'pocketbase';
import type {  ContactsRecord,  ContactsResponse } from './pocketbase-types';
import { writable } from 'svelte/store';

// export const pb = new PocketBase('http://YOUR-SERVER-IP-OR-URL'); // remote
export const pb = new PocketBase('http://127.0.0.1:8090'); // local

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange((auth) => {
    console.log('authStore changed', auth);
    currentUser.set(pb.authStore.model);
});

export const contacts = writable(new Array<ContactsRecord>());


export async function get50Contacts() {
        const resultList = await pb.collection('contacts').getList(1, 50, {}) as ListResult<ContactsResponse>;
        contacts.set(resultList.items as ContactsRecord[])
        // contacts = resultList.items;
}

