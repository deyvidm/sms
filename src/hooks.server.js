/** @type {import('@sveltejs/kit').Handle} */
export function handle({ event, resolve }) {
	const jwt = event.cookies.get('jwt');
	let user;
	try {
		if (jwt) {
			user = JSON.parse(atob(jwt))
		}
	} catch (e) {
		console.log(e)
	}
	event.locals = {user:user}

	return resolve(event);
}
