<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import { contacts, events } from './pocketbase';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import type { ContactResponse } from './pocketbase-types';
    import { createEventDispatcher } from 'svelte';
    import EventRow from './EventRow.svelte';

    function handleMessage(event) {
        console.log('event row click -- Event List');
    }
</script>

{#each $events as e}
    <EventRow event={e} on:message={handleMessage} />
{/each}
