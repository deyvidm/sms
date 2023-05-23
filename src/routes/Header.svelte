<script lang="ts">
    import { enhance } from "$app/forms";
    import { page } from "$app/stores";

    export let user: any;
</script>

<div class="navbar bg-primary text-primary-content">
    <a href="/" class="btn btn-ghost normal-case text-xl">SMS Thing</a>

    <div class="absolute right-2">
        {#if !user}
            <a href="/login" class:active={$page.url.pathname === "/login"}>
                <button class="btn btn-ghost">Sign in</button>
            </a>

            <!-- <a href="/register">
                <button class="btn btn-ghost">Sign up</button>
            </a> -->
        {/if}

        {#if user}
            Signed in as {user}
            <a href="/logout">
                <button class="btn btn-ghost right-0">Sign Out</button>
            </a>
            <!-- <form use:enhance method="POST" action="?/logout">
                <button class="btn btn-outline-danger"
                    >Or click here to logout.</button
                >
            </form> -->
        {/if}
    </div>
</div>
