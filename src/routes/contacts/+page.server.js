
import { apiClient } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch, parent, locals, url }) {

    const { user } = await parent();
    if (!user) throw redirect(307, '/');

    const body = await apiClient.get('/contacts')
    if (body.errors) {
        return fail(401, body);
    }
    
    return {contacts: body.data}
}
