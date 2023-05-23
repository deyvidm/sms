import { fail, redirect } from '@sveltejs/kit';
import { apiClient } from '$lib/gin.js';

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ fetch, cookies, request }) => {
		const data = await request.formData();
		data.set("username", "user2")
		data.set("password", "hunter2")

		const body = await apiClient.post(fetch, '/users/login',
		{
			username: data.get('username'),
		    password: data.get('password')
		});

		if (body.errors) {
			return fail(401, body);
		}

		console.log("/users/login body: ")
		console.log(body)
		apiClient.setToken(body.data.token)
		const value = btoa(JSON.stringify(body.data));
		cookies.set('jwt', value, { path: '/' });

		throw redirect(307, '/');
	}
};
