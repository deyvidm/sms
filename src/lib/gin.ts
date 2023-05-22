import { writable, type Writable } from 'svelte/store';
import type { Contact, CurrentUser, Event } from './Types';

export const currentUser = writable<CurrentUser>({username:"",contacts:[]})
export const userContacts = writable<Contact[]>([])
export const userEvents = writable<Event[]>([])

export class APIClient {
    private token = ""
    private user;

    constructor() {
    }

    public SignOut() {
        this.token = ""
        this.user = null
        currentUser.update(u => u = {username:"", contacts:[]})
        window.sessionStorage.setItem("store", "")
    }

    public async UpdateContacts(): Promise<boolean> {
        const response = await fetch("http://localhost:8080/api/contacts", {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + this.token
            },
        });

        if (!response.ok) {
            return false
        }
        const data = await response.json()
        if (data.status != "success") {
            return false
        }
        userContacts.update(u => u = data.data)
        return true
    }

    public async AddContact(first_name: string, last_name: string, phone: string): Promise<boolean> {
        const response = await fetch("http://localhost:8080/api/contacts/new", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + this.token
            },
            body: JSON.stringify({first_name, last_name, phone}),
        });

        if (!response.ok) {
            return false
        }
        const data = await response.json()
        if (data.status != "success") {
            return false
        }
        return true
    }

    public async UserLogin(username: string, password: string): Promise<boolean> {
        const response = await fetch("http://localhost:8080/api/users/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            const data = await response.json();
            if (data.status == "success") {
                this.token = data.data.token;
                this.user = data.data.user
                currentUser.update(u => u = this.user)
                userContacts.update(c => c = data.data.user.contacts)
                userEvents.update(c => c = data.data.user.events)
            }
            return true;
        } else {
            return false;
        }

    }
}

export const apiClient = new APIClient()