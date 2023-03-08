<script lang="ts">
    import type pocketbaseEs from 'pocketbase';
    import { pb } from './pocketbase';
    function onClick() {
        alert("test");
    }
</script>

<button class="btn" on:click={onClick}>TEST</button>
