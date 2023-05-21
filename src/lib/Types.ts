export enum InputType {
    "Text",
    "Phone"
}

export interface Contact {
    first_name: string;
    last_name: string;
    phone: string;
    id: string;
}
export interface CurrentUser {
    username: string;
    contacts: Contact[];
}