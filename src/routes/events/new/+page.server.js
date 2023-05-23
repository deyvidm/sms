import { apiClient } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch, locals, parent }) {
    return {
        user: locals.user
    }
    // const { user } = await parent();
	// return {events: body.data}
}