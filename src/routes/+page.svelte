<script>
    import { page } from "$app/stores";
    import Header from "./Header.svelte";

    /** @type {import('./$types').PageData} */
    export let data;

</script>

<svelte:head>
    <title>Logge</title>
</svelte:head>

<div class="home-page">
    {#if data.user}
        <div>
            <p>Welcome to the page sshugar {data.user}</p>
        </div>
    {/if}
</div>
