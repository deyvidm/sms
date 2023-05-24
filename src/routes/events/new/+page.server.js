import { apiClient } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch, locals, parent }) {
	console.log("---5: ", apiClient.getToken())
    
	return locals
}