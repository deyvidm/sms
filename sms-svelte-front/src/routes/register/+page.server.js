import { apiClient } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ parent }) {
	const { user } = await parent();
	if (user) throw redirect(307, '/');
}

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ cookies, request }) => {
		console.log("normally I would register, but there's more interesting code to write")
		return
		const data = await request.formData();
		data.set("username", "user3")
		data.set("first_name", "randolph")
		data.set("last_name", "pickler")
		data.set("password", "hunter3")

		const body = await apiClient.post('/users/register', {
			username: data.get('username'),
			first_name: data.get('first_name'),
			last_name: data.get('last_name'),
			password: data.get('password')
		});

		if (body.errors) {
			return fail(401, body);
		}

		console.log("/users/register body: ")
		console.log(body)
		const value = btoa(JSON.stringify(body.data.user));
		cookies.set('jwt', value, { path: '/' });

		throw redirect(307, '/');
	}
};
