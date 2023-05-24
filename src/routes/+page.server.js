import { apiClient, currentUser } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';


/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, url }) {
    if (!locals) throw redirect(302, '/login');
}


// /** @type {import('./$types').Actions} */
// export const actions = {
// 	logout: async ({ cookies, locals }) => {
// 	console.log("this one")
// 		cookies.delete('jwt', { path: '/' });
// 		locals.user = null;
// 	},
// }