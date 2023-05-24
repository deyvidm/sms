<script>
    import { page } from "$app/stores";
    import { apiClient } from "$lib/gin";
    import Header from "./Header.svelte";

    /** @type {import('./$types').PageData} */
    export let data;
    let dump = data.user ? data.user.username : "i'm null"
    apiClient.setToken(data.token)

</script>

<svelte:head>
    <title>Logge</title>
</svelte:head>

<div class="home-page">
    {#if data.user}
        <div>
            <p>Welcome to the page sshugar {dump}</p>
            <br>
            <p>Token:  {apiClient.getToken()}</p>
            <br>
            <p>locals:  {data.token}</p>
        </div>
    {/if}
</div>
