import { apiClient } from '$lib/gin';

/** @type {import('./$types').LayoutServerLoad} */
export function load({ locals }) {
	
	console.log("load root layout server")
	return {
		user: locals.user && {
			username: locals.user.username,
		}
	};
}
