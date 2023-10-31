// See https://kit.svelte.dev/docs/types#app

import type { User } from "$lib/gripes";

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: User
			token: string;
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};
