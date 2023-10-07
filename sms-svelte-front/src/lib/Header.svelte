<script lang="ts">
    import { currentUser, apiClient } from './gin';

    let username: string;
    let password: string;

    function signOut() {
        apiClient.SignOut()
    }  
</script>

{#if $currentUser}
    <div class="navbar bg-primary text-primary-content">
        <a href="/" class="btn btn-ghost normal-case text-xl">SMS Thing</a>
        <div class="absolute right-2">
            Signed in as {$currentUser.username}
            <a href="/">
                <button class="btn btn-ghost right-0 " on:click={signOut}>Sign Out</button>
            </a>
        </div>
    </div>
{/if}
