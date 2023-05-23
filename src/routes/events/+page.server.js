import { apiClient } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch, locals, parent }) {
    const { user } = await parent();
	if (!user) throw redirect(307, '/');

	const body = await apiClient.get(fetch, "/events")

	if (body.errors) {
        return fail(401, body);
    }

    console.log("/events body: ")
    console.log(body)
	locals.user.events = body.data 
	return {events: body.data}
}