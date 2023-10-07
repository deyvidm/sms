import { apiClient } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	const body = await apiClient.get('/contacts')
    if (body.errors) {
        return fail(401, body);
    }
    return {contacts: body.data}
}