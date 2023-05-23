
import { apiClient } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ parent, locals, url }) {
	 
    const { user } = await parent();
	if (!user) throw redirect(307, '/');

    console.log(locals)
    apiClient.setToken(locals.user)
    console.log("before GET")
    const body = await apiClient.get('/contacts')

    if (body.errors) {
        return fail(401, body);
    }

    console.log("/contacts body: ")
    console.log(body)
	return {contacts: body.data}
}
