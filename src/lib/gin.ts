import { writable, type Writable } from 'svelte/store';
import type { CurrentUser } from './Types';

export const currentUser = writable<CurrentUser | null>(null)

export class APIClient {
    private token = ""
    private user = null


    constructor() {
    }

    public SignOut() {
        this.token = ""
        this.user = null
        currentUser.update(u => u = null)
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
            }
            return true;
        } else {
            return false;
        }

    }
}

export const apiClient = new APIClient()