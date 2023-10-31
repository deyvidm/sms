import { apiClient } from '$lib/gin';

/** @type {import('@sveltejs/kit').Handle} */
export function handle({event, resolve }) {
	const jwt = event.cookies.get('jwt');
	event.locals = jwt ? JSON.parse(atob(jwt)) : null;
	if (event.locals) {
		apiClient.setToken(event.locals.token)
	}
	return resolve(event);
}
