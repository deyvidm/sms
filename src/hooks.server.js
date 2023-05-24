import { apiClient } from '$lib/gin';

/** @type {import('@sveltejs/kit').Handle} */
export function handle({event, resolve }) {
	const jwt = event.cookies.get('jwt');
	event.locals = jwt ? JSON.parse(atob(jwt)) : null;
	if (event.locals) {
		apiClient.setToken(event.locals.token)
	}
	console.log("1:", apiClient.getToken())
	return resolve(event);
	// let data;
	// try {
	// 	if (jwt) {
	// 		data = JSON.parse(atob(jwt))
	// 	}
	// } catch (e) {
	// 	console.log(e)
	// }
	// if (!data) {
	// 	data = {
	// 		user: {},
	// 		token: "",
	// 	}
	// }
	// event.locals = {
	// 	user: data.user,
	// 	token: data.token,
	// }
	// apiClient.setToken(data.token)
	// return resolve(event);
}
