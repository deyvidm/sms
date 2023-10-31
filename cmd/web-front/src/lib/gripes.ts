export enum InputType {
    "Text",
    "Phone"
}

export interface Contact {
    id: string;
    first_name: string;
    last_name: string;
    phone: string;
}
export interface Event {
    id: string;
	title: string;
	capacity: number;
	start_date: string;
	end_date:   string;
}

export interface User {
    username: string;
    contacts: Contact[];
    events: Event[];
}

