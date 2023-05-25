import { apiClient } from '$lib/gin.js';
import { fail, redirect } from '@sveltejs/kit';


export function load({ locals }) {
	console.log(locals)
	if (!locals) throw redirect(302, '/login');
}

/** @type {import('./$types').Actions} */
export const actions = {
	logout: async ({ cookies, locals }) => {
		cookies.delete('jwt', { path: '/' });
		apiClient.SignOut()
		locals = {}
		return locals
	},
}