import { apiClient, currentUser } from '$lib/gin';
import { fail, redirect } from '@sveltejs/kit';


/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, url }) {
    if (!locals) throw redirect(302, '/login');
	// let log;
    // console.log("load root page server ")
    // if (locals.user?.username){
    //     log = "logged in: " + locals.user.username
    // } else {
    //     log = "-0-"
    // }

    // console.log("Page Server", log)

	// return {

	// };
	console.log("\t2: ", apiClient.getToken())

}


// /** @type {import('./$types').Actions} */
// export const actions = {
// 	logout: async ({ cookies, locals }) => {
// 	console.log("this one")
// 		cookies.delete('jwt', { path: '/' });
// 		locals.user = null;
// 	},
// }