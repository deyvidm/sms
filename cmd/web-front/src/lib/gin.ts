import { writable, type Writable } from 'svelte/store';
import type { Contact, User, Event } from './gripes';
import { error, json } from '@sveltejs/kit';


export const currentUser = writable<User>()
export const userContacts = writable<Contact[]>([])
export const userEvents = writable<Event[]>([])
const base = 'http://localhost:8080/api';


export class APIClient {
    private token = "";

    constructor() {
    }

    public getToken() {
        return this.token
    }
    public setToken(token: string) {
        this.token = token
    }

    public buildReq(method, path, data) {
        let headers = {
            "access-control-allow-origin": "*",
            'Access-Control-Allow-Headers': "*",
            'Access-Control-Allow-Methods': "*",
        }
        const token = this.getToken()
        if (token.length > 0) {
            headers['Authorization'] = `Bearer ${token}`
        }

        if (data && method != "GET") {
            headers['Content-Type'] = 'application/json';
            return {
                method: method,
                headers: headers,
                body: JSON.stringify(data)
            }
        }
        return {
            method: method,
            headers: headers,
        }

    }


    private async send({method, path, data}) {
        let req = this.buildReq(method, path, data)
        const res = await fetch(`${base}${path}`, req);
        if (res.ok || res.status === 422) {
            const text = await res.text();
            return text ? JSON.parse(text) : {};
        }
        throw error(res.status);
    }

    public async get(path) {
        return this.send({ method: 'GET', path, data: {} });
    }

    public async del(path) {
        return this.send({ method: 'DELETE', path, data: {} });
    }

    public async post(path, data) {
        return this.send({ method: 'POST', path, data });
    }

    public async put(path, data) {
        return this.send({ method: 'PUT', path, data });
    }

    public async patch(path, data) {
        return this.send({ method: 'PATCH', path, data });
    }



    public SignOut() {
        this.setToken("")
    }

    public async GetInvitations(eventID: string) {
        return this.get("/events/" + eventID)
    }

    public async UserLogin(username: string | null, password: string | null): Promise<any> {
        // const response = await fetch("http://localhost:8080/api/users/login", {
        //     method: "POST",
        //     headers: {
        //         "Content-Type": "application/json",
        //     },
        //     body: JSON.stringify({ username, password }),
        // });

        // let data;
        // const body = this.post("/users/login", { username, password })
        //     .then((respData) => {
        //         console.log(respData)
        //         this.token = respData.data.token;
        //         this.user = respData.data.user
        //         data = respData.data
        //         console.log("return respData")
        //     })

        // return body
    }
}


export const apiClient = new APIClient()


// async function send({ method, path, data, token }) {
// 	const opts = { method, headers: {}, body: ""};

// 	if (data) {
// 		opts.headers['Content-Type'] = 'application/json';
// 		opts.body = JSON.stringify(data);
// 	}

// 	if (token) {
// 		opts.headers['Authorization'] = `Token ${token}`;
// 	}

// 	const res = await fetch(`${base}/${path}`, opts);
// 	if (res.ok || res.status === 422) {
// 		const text = await res.text();
// 		return text ? JSON.parse(text) : {};
// 	}

// 	throw error(res.status);
// }

// export function get(path, token) {
// 	return send({ method: 'GET', path, data: [], token });
// }

// export function del(path, token) {
// 	return send({ method: 'DELETE', path, data: [], token });
// }

// export function post(path, data, token) {
// 	return send({ method: 'POST', path, data, token });
// }

// export function put(path, data, token) {
// 	return send({ method: 'PUT', path, data, token });
// }



// export class APIClient {
//     private token = ""
//     private user;

//     constructor() {
//     }

//     public SignOut() {
//         this.token = ""
//         this.user = null
//         currentUser.update(u => u = {username:"", contacts:[]})
//         window.sessionStorage.setItem("store", "")
//     }

//     public async UpdateContacts(): Promise<boolean> {
//         const response = await fetch("http://localhost:8080/api/contacts", {
//             method: "GET",
//             headers: {
//                 "Authorization": "Bearer " + this.token
//             },
//         });

//         if (!response.ok) {
//             return false
//         }
//         const data = await response.json()
//         if (data.status != "success") {
//             return false
//         }
//         userContacts.update(u => u = data.data)
//         return true
//     }

//     public async AddContact(first_name: string, last_name: string, phone: string): Promise<boolean> {
//         const response = await fetch("http://localhost:8080/api/contacts/new", {
//             method: "POST",
//             headers: {
//                 "Content-Type": "application/json",
//                 "Authorization": "Bearer " + this.token
//             },
//             body: JSON.stringify({first_name, last_name, phone}),
//         });

//         if (!response.ok) {
//             return false
//         }
//         const data = await response.json()
//         if (data.status != "success") {
//             return false
//         }
//         return true
//     }

//     public async UserLogin(username: string, password: string): Promise<boolean> {
//         const response = await fetch("http://localhost:8080/api/users/login", {
//             method: "POST",
//             headers: {
//                 "Content-Type": "application/json",
//             },
//             body: JSON.stringify({ username, password }),
//         });

//         if (response.ok) {
//             const data = await response.json();
//             if (data.status == "success") {
//                 this.token = data.data.token;
//                 this.user = data.data.user
//                 currentUser.update(u => u = this.user)
//                 userContacts.update(c => c = data.data.user.contacts)
//                 userEvents.update(c => c = data.data.user.events)
//             }
//             return true;
//         } else {
//             return false;
//         }

//     }



// private async send({ method, path, data, token }) {
// 	const opts = { method, headers: {}, body: "" };

// 	if (data) {
// 		opts.headers['Content-Type'] = 'application/json';
// 		opts.body = JSON.stringify(data);
// 	}

// 	if (token) {
// 		opts.headers['Authorization'] = `Bearer ${this.token}`;
// 	}

// 	const res = await fetch(`${base}/${path}`, opts);
// 	if (res.ok || res.status === 422) {
// 		const text = await res.text();
// 		return text ? JSON.parse(text) : {};
// 	}

// 	throw error(res.status);
// }

// private get(path, token) {
// 	return this.send({ method: 'GET', path, data:[], token });
// }

// private del(path, token) {
// 	return this.send({ method: 'DELETE', path, data:[], token });
// }

// private post(path, data, token) {
// 	return this.send({ method: 'POST', path, data, token });
// }

// private put(path, data, token) {
// 	return this.send({ method: 'PUT', path, data, token });
// }

// }




// export const apiClient = new APIClient()