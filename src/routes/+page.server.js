import { apiClient, currentUser } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';


/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, url }) {
	let log;

    if (locals.user?.username){
        log = "logged in: " + locals.user.username
    } else {
        log = "-0-"
    }

    console.log("Page Server", log)

	return {

	};
}


// /** @type {import('./$types').Actions} */
// export const actions = {
// 	logout: async ({ cookies, locals }) => {
// 	console.log("this one")
// 		cookies.delete('jwt', { path: '/' });
// 		locals.user = null;
// 	},
// }